// app.module.ts

import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { WeatherController } from '../weather/weather.controller';
import { WeatherService } from '../weather/weather.service';
import { User } from '../entity/user.entity';
import { Actions } from '../entity/actions.entity';
import { UserController } from 'src/user/user.controller';
import { UserService } from 'src/user/user.service';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'db',
      port: 3306,
      username: 'root',
      password: 'example',
      database: 'mydatabase',
      entities: [User, Actions],
      synchronize: false,
    }),
    TypeOrmModule.forFeature([User, Actions]),
  ],
  controllers: [WeatherController, UserController],
  providers: [WeatherService, UserService],
})
export class AppModule {}
