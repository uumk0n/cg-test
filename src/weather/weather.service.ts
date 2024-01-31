import { Injectable, NotFoundException, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import * as bcrypt from 'bcrypt';
import axios from 'axios';
import * as dotenv from 'dotenv';
import { v4 as uuidv4 } from 'uuid';
import { User } from '../entity/user.entity';
import { Actions } from '../entity/actions.entity';


@Injectable()
export class WeatherService {
  constructor(
    @InjectRepository(User) private readonly userRepository: Repository<User>,
    @InjectRepository(Actions) private readonly actionsRepository: Repository<Actions>,
  ) { }

  async register(data: { login: string; password: string; fio: string }): Promise<{ fio: string; apiToken: string }> {
    if (data.password.length < 6 || !/[.,!_]/.test(data.password)) {
      throw new BadRequestException('Password must be at least 6 characters long and contain one of the characters: . , ! _');
    }

    const apiToken = await this.generateApiToken();

    const hashedPassword = await bcrypt.hash(data.password, 10);

    const newUser = this.userRepository.create({
      login: data.login,
      hashedPassword: hashedPassword,
      fio: data.fio,
      apiToken: apiToken,
    });

    await this.userRepository.save(newUser);

    return { fio: newUser.fio, apiToken: newUser.apiToken };
  }

  private async generateApiToken(): Promise<string> {
    return uuidv4();
  }

  async login(data: { login: string; password: string }): Promise<{ fio: string; apiToken: string }> {
    const user = await this.userRepository.findOne({ where: { login: data.login } });

    if (!user) {
      throw new NotFoundException('User not found');
    }

    const isPasswordValid = await bcrypt.compare(data.password, user.hashedPassword);

    if (!isPasswordValid) {
      throw new BadRequestException('Invalid password');
    }

    return { fio: user.fio, apiToken: user.apiToken };
  }

  async getCurrentWeather(data: { apiToken: string; city: string; language?: string }): Promise<any> {
    const user = await this.userRepository.findOne({ where: { apiToken: data.apiToken } });

    if (!user) {
      throw new NotFoundException('User not found');
    }

    if (!user.apiToken) {
      throw new BadRequestException('Invalid token');
    }

    dotenv.config();
    let weatherResponse;
    try {
      const url = `http://api.weatherapi.com/v1/current.json?key=${process.env.WEATHER_API_KEY}&q=${data.city}&lang=${data.language || 'ru'}`;
      weatherResponse = await axios.get(url);
    } catch (error) {
      console.error(error);
      await this.saveAction(user.id, Date.now(), error.response?.status || null, null);
      throw new BadRequestException('Error fetching weather data');
    }
    await this.saveAction(user.id, Date.now(), weatherResponse.status, weatherResponse.data.current.temp_c);

    return weatherResponse.data;
  }

  private async saveAction(userId: number, actionTime: number, requestResult: number | null, tempC: number | null): Promise<void> {
    const newAction = this.actionsRepository.create({
      user: { id: userId },
      actionTime: actionTime,
      requestResult: requestResult,
      tempC: tempC,
    });

    await this.actionsRepository.save(newAction);
  }
}
