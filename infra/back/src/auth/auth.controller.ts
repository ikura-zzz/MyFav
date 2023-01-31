import { Controller, Post, Req, Res } from '@nestjs/common';
import { AuthService } from './auth.service';
import { Request, Response } from 'express';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post()
  postAuth(@Req() req: Request, @Res() res: Response) {
    this.authService.postAuth(req, res).send();
  }
}
