import React, { ChangeEvent, KeyboardEvent, useState } from "react";
import { AuthRequest } from "../types/AuthRequest";
import axios from "../axiosWrapper";
import { useHistory } from "react-router-dom";

export const SignInForm = () => {
  const [userId, setUserId] = useState("");
  const [userPasswd, setUserPasswd] = useState("");
  const [errmsg, setErrmsg] = useState("");
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

  // Enterキー押下時はSignin処理を呼び出す
  const onKeyDownSignin = (event: KeyboardEvent<HTMLElement>) => {
    if (event.key === "Enter") {
      onClickSignin();
    }
  };

  // サーバーへ認証依頼を飛ばす
  const auth = async (req: AuthRequest) => {
    // const res = await axios.post("/signin", req);
    // console.log(res);
    history.push("/react/list");
  };

  // サインインボタンが押された時の処理
  const onClickSignin = async () => {
    if (userId === "" || userPasswd === "") {
      setErrmsg("ユーザーID、パスワードを入力してください。");
      return;
    }
    await auth({ userId, userPasswd });
    setErrmsg("");
    setUserId("");
    setUserPasswd("");
  };
  return (
    <>
      <form>
        <div>
          <input
            type="text"
            placeholder="ユーザーID"
            onChange={onChangeUserId}
            onKeyDown={onKeyDownSignin}
            value={userId}
            required
            data-testid="username"
          />
        </div>
        <div>
          <input
            type="password"
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
