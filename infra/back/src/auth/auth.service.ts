import { HttpStatus, Injectable } from '@nestjs/common';
import { Request, Response } from 'express';
import { authRequest } from '../types/authRequests';

@Injectable()
export class AuthService {
  postAuth(req: Request, res: Response) {
    const reqBody: authRequest = req.body;
    if (
      reqBody.userId === 'testuser2' &&
      reqBody.userPasswd === 'testpassword2'
    ) {
      res.status(HttpStatus.OK);
      res.json({ key: '123456', msg: '' });
    } else {
      res.status(HttpStatus.OK);
      res.json({ msg: 'unauthrized' });
    }
    return res;
  }
}
