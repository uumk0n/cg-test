import { Controller, Post, Body } from '@nestjs/common';
import { WeatherService } from './weather.service';
import { ApiResponse, ApiBody, ApiTags } from '@nestjs/swagger';
import { CurrentWeatherDto } from '../DTO/current-weather.dto';

@ApiTags('Weather')
@Controller('weather')
export class WeatherController {
    constructor(private readonly weatherService: WeatherService) {}

    @ApiBody({ type: CurrentWeatherDto })
    @ApiResponse({ status: 200, description: 'Weather information retrieved successfully' })
    @ApiResponse({ status: 400, description: 'Bad request' })
    @Post('current-weather')
    async currentWeather(@Body() data: CurrentWeatherDto): Promise<any> {
        return this.weatherService.getCurrentWeather(data);
    }
}
