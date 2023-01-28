import React, { ChangeEvent, useState } from "react";
import { AuthRequest } from "../types/AuthRequest";
import axios from "../axiosWrapper";

export const SignInForm = () => {
  const [userId, setUserId] = useState("");
  const [userPasswd, setUserPasswd] = useState("");

  const onChangeUserId = (event: ChangeEvent<HTMLInputElement>) => {
    const { target } = event;
    if (!(target instanceof HTMLInputElement)) {
      return;
    }
    setUserId(target.value);
  };
  const onChangeUserPasswd = (event: ChangeEvent<HTMLInputElement>) => {
    const { target } = event;
    if (!(target instanceof HTMLInputElement)) {
      return;
    }
    setUserPasswd(target.value);
  };
  const authchk = async (auth: AuthRequest) => {
    axios.post("/auth", auth);
  };
  const signin = async () => {
    if (userId === "" || userPasswd === "") {
      return;
    }
    await authchk({ userId, userPasswd });
  };
  return (
    <>
      <form>
        <div>
          <input
            type="text"
            placeholder="ユーザーID"
            onChange={onChangeUserId}
            value={userId}
          />
        </div>
        <div>
          <input
            type="password"
            placeholder="パスワード"
            onChange={onChangeUserPasswd}
            value={userPasswd}
          />
        </div>
        <div>
          <input type="button" value="ログイン" onClick={signin} />
        </div>
      </form>
    </>
  );
};
