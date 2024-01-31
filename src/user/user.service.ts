import { Injectable, NotFoundException, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import * as bcrypt from 'bcrypt';
import { v4 as uuidv4 } from 'uuid';
import { User } from '../entity/user.entity';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User) private readonly userRepository: Repository<User>,
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
}