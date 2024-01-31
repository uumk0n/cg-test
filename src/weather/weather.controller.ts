import { Controller, Post, Body } from '@nestjs/common';
import { WeatherService } from './weather.service';
import { ApiResponse, ApiBody, ApiTags } from '@nestjs/swagger';
import { RegisterDto } from '../DTO/register.dto';
import { LoginDto } from '../DTO/login.dto';
import { CurrentWeatherDto } from '../DTO/current-weather.dto';

@ApiTags('Weather')
@Controller('weather')
export class WeatherController {
    constructor(private readonly weatherService: WeatherService) {}

    @ApiBody({ type: RegisterDto })
    @ApiResponse({ status: 201, description: 'User registered successfully' })
    @ApiResponse({ status: 400, description: 'Bad request' })
    @Post('register')
    async register(@Body() data: RegisterDto): Promise<{ fio: string; apiToken: string }> {
        return this.weatherService.register(data);
    }

    @ApiBody({ type: LoginDto })
    @ApiResponse({ status: 201, description: 'User logged in successfully' })
    @ApiResponse({ status: 400, description: 'Bad request' })
    @Post('login')
    async login(@Body() data: LoginDto): Promise<{ fio: string; apiToken: string }> {
        return this.weatherService.login(data);
    }

    @ApiBody({ type: CurrentWeatherDto })
    @ApiResponse({ status: 200, description: 'Weather information retrieved successfully' })
    @ApiResponse({ status: 400, description: 'Bad request' })
    @Post('current-weather')
    async currentWeather(@Body() data: CurrentWeatherDto): Promise<any> {
        return this.weatherService.getCurrentWeather(data);
    }
}
