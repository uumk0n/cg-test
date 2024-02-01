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
      host: process.env.DBHOST,
      port: parseInt(process.env.DBPORT, 10),
      username: process.env.DBUSERNAME,
      password: process.env.DBPASSWORD,
      database: process.env.DBDATABASE,
      entities: [User, Actions],
      synchronize: false,
    }),
    TypeOrmModule.forFeature([User, Actions]),
  ],
  controllers: [WeatherController, UserController],
  providers: [WeatherService, UserService],
})
export class AppModule {}
