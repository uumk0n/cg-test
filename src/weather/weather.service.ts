import { Injectable, NotFoundException, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import axios from 'axios';
import { User } from '../entity/user.entity';
import { Actions } from '../entity/actions.entity';


@Injectable()
export class WeatherService {
  constructor(
    @InjectRepository(User) private readonly userRepository: Repository<User>,
    @InjectRepository(Actions) private readonly actionsRepository: Repository<Actions>,
  ) { }

  async getCurrentWeather(data: { apiToken: string; city: string; language?: string }): Promise<any> {
    const user = await this.userRepository.findOne({ where: { apiToken: data.apiToken } });

    if (!user) {
      throw new NotFoundException('User not found');
    }

    if (!user.apiToken) {
      throw new BadRequestException('Invalid token');
    }

    
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
