import { HttpStatus, Injectable } from '@nestjs/common';
import { Request, Response } from 'express';
import { AuthRequest } from '../types/AuthRequest';
import { AuthResponse } from '../types/AuthResponse';
import { unauthMessage } from '../types/CommonMessages';

@Injectable()
export class AuthService {
  postAuth(req: Request, res: Response) {
    const resBody = new AuthResponse();
    const reqBody: AuthRequest = req.body;
    if (
      reqBody.userId === 'testuser2' &&
      reqBody.userPasswd === 'testpassword2'
    ) {
      res.status(HttpStatus.OK);
      resBody.key = 123456;
      resBody.message = '';
      res.json(resBody);
    } else {
      res.status(HttpStatus.OK);
      resBody.message = unauthMessage;
      res.json(resBody);
    }
    return res;
  }
}
