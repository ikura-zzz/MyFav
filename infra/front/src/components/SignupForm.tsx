import React, { ChangeEvent, KeyboardEvent, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { SignupRequest } from '../types/SignupRequest';
import { createUserWithEmailAndPassword } from 'firebase/auth';
import { getFireBaseAuth } from './CommonAuthCheck';
import { CommonLabel, InputBox } from './InputBox';
import { CommonButton } from './CommonButton';

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
      <form className="font-sans text-sm rounded w-full max-w-md mx-auto my-8 px-8 pt-6 pb-8">
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <InputBox
            type="text"
            id="new_userid"
            placeHolder="メールアドレス"
            onChange={onChangeUserId}
            onKeyDown={onKeyDownSignup}
            value={userId}
            testId="username"
          />
          <CommonLabel htmlFor="new_userid" labelBody="メールアドレス" />
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <InputBox
            type="password"
            id="new_password"
            placeHolder="パスワード"
            onChange={onChangeUserPasswd}
            onKeyDown={onKeyDownSignup}
            value={userPasswd}
            testId="userpasswd"
          />
          <CommonLabel htmlFor="new_password" labelBody="パスワード" />
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <InputBox
            type="password"
            id="retype_new_password"
            placeHolder="パスワード再入力"
            onChange={onChangeRetypePasswd}
            onKeyDown={onKeyDownSignup}
            value={retypePasswd}
            testId="retypeuserpasswd"
          />
          <CommonLabel
            htmlFor="retype_new_password"
            labelBody="パスワード再入力"
          />
        </div>
        <div data-testid="errormsg">{errmsg}</div>
        <div className="flex items-center justify-between">
          <CommonButton
            buttonBody="サインアップ"
            onClick={onClickSignup}
            testId="signupbutton"
          />
        </div>
      </form>
    </>
  );
};
