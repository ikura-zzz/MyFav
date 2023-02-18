import React, { ChangeEvent, KeyboardEvent, useState } from 'react';
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
    // await signOut(auth);
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
      <form>
        <div>
          <input
            type="text"
            id="userid"
            placeholder="メールアドレス"
            onChange={onChangeUserId}
            onKeyDown={onKeyDownSignin}
            value={userid}
            required
            data-testid="username"
          />
        </div>
        <div>
          <input
            type="password"
            id="password"
            placeholder="パスワード"
            onChange={onChangeUserPasswd}
            onKeyDown={onKeyDownSignin}
            value={userPasswd}
            required
            data-testid="userpasswd"
          />
        </div>
        <div data-testid="errormsg">{errmsg}</div>
        <div>
          <input
            type="button"
            value="サインイン"
            onClick={onClickSignin}
            data-testid="signinbutton"
          />
        </div>
      </form>
    </>
  );
};
