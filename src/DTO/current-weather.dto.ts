// current-weather.dto.ts

import { ApiProperty } from '@nestjs/swagger';

export class CurrentWeatherDto {
    @ApiProperty({ description: 'User API token' })
    apiToken: string;

    @ApiProperty({ description: 'City name' })
    city: string;

    @ApiProperty({ description: 'Language preference (optional)' })
    language?: string;
}
