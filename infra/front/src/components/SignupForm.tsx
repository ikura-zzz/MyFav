import React, { ChangeEvent, KeyboardEvent, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { SignupRequest } from '../types/SignupRequest';
import { createUserWithEmailAndPassword } from 'firebase/auth';
import { getFireBaseAuth } from './CommonAuthCheck';

const auth = getFireBaseAuth();

export const SignupForm = () => {
  const [userId, setUserId] = useState('');
  const [userPasswd, setUserPasswd] = useState('');
  const [retypePasswd, setRetypePasswd] = useState('');
  const [errmsg, setErrmsg] = useState('');
  const history = useHistory();

  // ユーザーID入力イベント時の処理
  const onChangeUserId = (event: ChangeEvent<HTMLInputElement>) => {
    const { target } = event;
    if (!(target instanceof HTMLInputElement)) {
      return;
    }
    setUserId(target.value);
  };

  // パスワード入力イベント時の処理
  const onChangeUserPasswd = (event: ChangeEvent<HTMLInputElement>) => {
    const { target } = event;
    if (!(target instanceof HTMLInputElement)) {
      return;
    }
    setUserPasswd(target.value);
  };

  // 再入力パスワード入力イベント時の処理
  const onChangeRetypePasswd = (event: ChangeEvent<HTMLInputElement>) => {
    const { target } = event;
    if (!(target instanceof HTMLInputElement)) {
      return;
    }
    setRetypePasswd(target.value);
  };

  // サーバーへ認証依頼を飛ばす
  const signup = async (req: SignupRequest) => {
    createUserWithEmailAndPassword(auth, req.userId, req.userPasswd).then(
      (userCredential) => {
        const user = userCredential.user;
        console.log('ユーザー作成成功' + user.uid + user.email);
        setErrmsg('');
        setUserId('');
        setUserPasswd('');
        setRetypePasswd('');
        history.push('/react');
      },
    );
    // // const res = await await axios.post('/signup', req);
    // if (res.status) {
    //   history.push('/react');
    // }
  };
  // Enterキー押下時はSignin処理を呼び出す
  const onKeyDownSignup = (event: KeyboardEvent<HTMLElement>) => {
    if (event.key === 'Enter') {
      onClickSignup();
    }
  };
  const isValidMailAdress = (userId: string): boolean => {
    const reg =
      /^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$/;
    return reg.test(userId);
  };
  // サインインボタンが押された時の処理
  const onClickSignup = async () => {
    if (userId === '' || userPasswd === '' || retypePasswd === '') {
      setErrmsg('ユーザーID、パスワード、再入力パスワードを入力してください。');
      return;
    }
    if (!isValidMailAdress(userId)) {
      setErrmsg('メールアドレスの形式に合致しません。');
      return;
    }
    if (userPasswd !== retypePasswd) {
      setErrmsg('パスワードが一致しません');
      return;
    }
    await signup({ userId, userPasswd, retypePasswd });
  };

  return (
    <>
      <form>
        <div>
          <input
            type="text"
            id="new_userid"
            placeholder="メールアドレス"
            onChange={onChangeUserId}
            onKeyDown={onKeyDownSignup}
            value={userId}
            data-testid="username"
          />
        </div>
        <div>
          <input
            type="password"
            id="new_password"
            placeholder="パスワード"
            onChange={onChangeUserPasswd}
            onKeyDown={onKeyDownSignup}
            value={userPasswd}
            data-testid="userpasswd"
          />
        </div>
        <div>
          <input
            type="password"
            id="retype_new_password"
            placeholder="パスワード再入力"
            onChange={onChangeRetypePasswd}
            onKeyDown={onKeyDownSignup}
            value={retypePasswd}
            data-testid="userpasswd"
          />
        </div>
        <div data-testid="errormsg">{errmsg}</div>
        <div>
          <input
            type="button"
            value="サインアップ"
            onClick={onClickSignup}
            data-testid="signupbutton"
          />
        </div>
      </form>
    </>
  );
};
