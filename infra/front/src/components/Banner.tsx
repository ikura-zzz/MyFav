import React from "react";
import { useHistory } from "react-router-dom";

export const Banner = () => {
  const history = useHistory();
  const onClickSignup = () => {
    history.push("/react/signup");
  };

  const backHome = () => {
    history.push("/react");
  };
  return (
    <>
      <img src="http://localhost/img/icon.png" alt="logo" onClick={backHome} />
      <ul>
        <li>MyFavの使い方</li>
        <li>
          <button data-testid="signup" onClick={onClickSignup}>
            サインアップ
          </button>
        </li>
      </ul>
    </>
  );
};
