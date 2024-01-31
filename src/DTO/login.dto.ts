// login.dto.ts

import { ApiProperty } from '@nestjs/swagger';

export class LoginDto {
    @ApiProperty({ description: 'User login' })
    login: string;

    @ApiProperty({ description: 'User password' })
    password: string;
}
