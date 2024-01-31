// register.dto.ts

import { ApiProperty } from '@nestjs/swagger';
import { IsString, Length, Matches } from 'class-validator';

export class RegisterDto {
    @ApiProperty({ description: 'User login' })
    @IsString()
    login: string;

    @ApiProperty({
        description: 'User password (must be at least 6 characters long and contain one of the characters: . , ! _)',
        minLength: 6,
        maxLength: 255,
        example: 'password123!', 
    })
    @IsString()
    @Length(6)
    @Matches(/[.,!_]/, { message: 'Password must contain one of the characters: . , ! _' })
    password: string;

    @ApiProperty({ description: 'Full name of the user' })
    @IsString()
    fio: string;
}
