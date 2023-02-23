import React, { ChangeEvent, KeyboardEvent, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { AuthRequest } from '../types/AuthRequest';
import { signInWithEmailAndPassword } from 'firebase/auth';
import { getFireBaseAuth } from './CommonAuthCheck';

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
  const onClickAuth = async (req: AuthRequest) => {
    await signInWithEmailAndPassword(auth, userid, userPasswd)
      .then(async () => {
        history.push('/react/list');
      })
      .catch((error) => {
        setErrmsg(error.code + error.errmsg);
      });
  };

  // サインインボタンが押された時の処理
  const onClickSignin = async () => {
    if (userid === '' || userPasswd === '') {
      setErrmsg('ユーザーID、パスワードを入力してください。');
      return;
    }
    await onClickAuth({ userId: userid, userPasswd });
  };
  return (
    <>
      <form className="font-sans text-sm rounded w-full max-w-md mx-auto my-8 px-8 pt-6 pb-8">
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <input
            type="text"
            id="userid"
            className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
            placeholder="メールアドレス"
            onChange={onChangeUserId}
            onKeyDown={onKeyDownSignin}
            value={userid}
            required
            data-testid="username"
          />
          <label
            className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
            htmlFor="userid"
          >
            メールアドレス
          </label>
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <input
            type="password"
            id="password"
            className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
            placeholder="パスワード"
            onChange={onChangeUserPasswd}
            onKeyDown={onKeyDownSignin}
            value={userPasswd}
            required
            data-testid="userpasswd"
          />
          <label
            className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
            htmlFor="password"
          >
            パスワード
          </label>
        </div>
        <div
          data-testid="errormsg"
          className="flex items-center justify-between"
        >
          {errmsg}
        </div>
        <div>
          <input
            type="button"
            className="bg-black hover:bg-black text-white py-2 px-4"
            value="サインイン"
            onClick={onClickSignin}
            data-testid="signinbutton"
          />
        </div>
      </form>
    </>
  );
};
