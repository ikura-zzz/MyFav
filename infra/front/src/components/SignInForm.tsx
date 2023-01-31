import React, { ChangeEvent, KeyboardEvent, useState } from "react";
import { AuthRequest } from "../types/AuthRequest";
import axios from "../axiosWrapper";
import { useHistory } from "react-router-dom";
import { HttpStatusCode } from "axios";

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

  // サーバーからのレスポンスが、認証済みかどうか
  const isAuthed = (data: any): boolean => {
    return data.key !== undefined;
  };

  // サーバーへ認証依頼を飛ばす
  const auth = async (req: AuthRequest) => {
    const res = await axios.post("/auth", req);
    if (res.status === HttpStatusCode.Ok) {
      if (isAuthed(res.data)) {
        setUserId("");
        setUserPasswd("");
        history.push("/react/list");
      } else {
        if (res.data.msg === "unauthrized") {
          setErrmsg("ログイン認証に失敗しました。");
        }
      }
    } else {
      setErrmsg("予期しないエラーが発生しました。");
    }
  };

  // サインインボタンが押された時の処理
  const onClickSignin = async () => {
    if (userId === "" || userPasswd === "") {
      setErrmsg("ユーザーID、パスワードを入力してください。");
      return;
    }
    await auth({ userId, userPasswd });
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
