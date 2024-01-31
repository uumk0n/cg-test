import { Controller, Post, Body } from '@nestjs/common';
import { ApiResponse, ApiBody, ApiTags } from '@nestjs/swagger';
import { RegisterDto } from '../DTO/register.dto';
import { LoginDto } from '../DTO/login.dto';
import { UserService } from './user.service';

@ApiTags('User')
@Controller('User')
export class UserController {
    constructor(private readonly UserService: UserService) {}

@ApiBody({ type: RegisterDto })
    @ApiResponse({ status: 201, description: 'User registered successfully' })
    @ApiResponse({ status: 400, description: 'Bad request' })
    @Post('register')
    async register(@Body() data: RegisterDto): Promise<{ fio: string; apiToken: string }> {
        return this.UserService.register(data);
    }

    @ApiBody({ type: LoginDto })
    @ApiResponse({ status: 201, description: 'User logged in successfully' })
    @ApiResponse({ status: 400, description: 'Bad request' })
    @Post('login')
    async login(@Body() data: LoginDto): Promise<{ fio: string; apiToken: string }> {
        return this.UserService.login(data);
    }
}