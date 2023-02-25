import React, { ChangeEvent, KeyboardEvent, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { AuthRequest } from '../types/AuthRequest';
import { signInWithEmailAndPassword } from 'firebase/auth';
import { getFireBaseAuth } from './CommonAuthCheck';
import { CommonLabel, InputBox } from './InputBox';
import { CommonButton } from './CommonButton';

const auth = getFireBaseAuth();

export const SignInForm = () => {
  const [userid, setUserid] = useState('');
  const [userPasswd, setUserPasswd] = useState('');
  const [errmsg, setErrmsg] = useState('');
  const history = useHistory();

  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      if (user) {
        history.push('/react/list');
      }
    });
  });

  // ユーザーID入力イベント時の処理
  const onChangeUserId = (event: ChangeEvent<HTMLInputElement>) => {
    const { target } = event;
    if (!(target instanceof HTMLInputElement)) {
      return;
    }
    setUserid(target.value);
  };

  // パスワード入力イベント時の処理
  const onChangeUserPasswd = (event: ChangeEvent<HTMLInputElement>) => {
    const { target } = event;
    if (!(target instanceof HTMLInputElement)) {
      return;
    }
    setUserPasswd(target.value);
  };

  // Enterキー押下時はSignin処理を呼び出す
  const onKeyDownSignin = (event: KeyboardEvent<HTMLElement>) => {
    if (event.key === 'Enter') {
      onClickSignin();
    }
  };

  // サーバーへ認証依頼を飛ばす
  const requestAuth = async (req: AuthRequest): Promise<boolean> => {
    await signInWithEmailAndPassword(auth, userid, userPasswd)
      .then(async () => {
        return true;
      })
      .catch((error) => {
        setErrmsg(error.code + error.errmsg);
        return false;
      });
    return false;
  };

  // サインインボタンが押された時の処理
  const onClickSignin = async () => {
    if (userid === '' || userPasswd === '') {
      setErrmsg('ユーザーID、パスワードを入力してください。');
      return;
    }
    if (await requestAuth({ userId: userid, userPasswd })) {
      history.push('/react/list');
    }
  };
  return (
    <>
      <form className="font-sans text-sm rounded w-full max-w-md mx-auto my-8 px-8 pt-6 pb-8">
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <InputBox
            type="text"
            id="userid"
            placeHolder="メールアドレス"
            onChange={onChangeUserId}
            onKeyDown={onKeyDownSignin}
            value={userid}
            testId="username"
          />
          <CommonLabel htmlFor="userid" labelBody="メールアドレス" />
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <InputBox
            type="password"
            id="password"
            placeHolder="パスワード"
            onChange={onChangeUserPasswd}
            onKeyDown={onKeyDownSignin}
            value={userPasswd}
            testId="userpasswd"
          />

          <CommonLabel htmlFor="password" labelBody="パスワード" />
        </div>
        <div
          data-testid="errormsg"
          className="flex items-center justify-between"
        >
          {errmsg}
        </div>
        <div>
          <CommonButton
            onClick={onClickSignin}
            testId="signinbutton"
            buttonBody="サインイン"
          />
        </div>
      </form>
    </>
  );
};
