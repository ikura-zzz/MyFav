import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getFireBaseAuth } from './CommonAuthCheck';

export const Banner = () => {
  const [auth, setAuth] = useState(getFireBaseAuth());
  const [authButton, setAuthButton] = useState(<></>);
  const history = useHistory();
  const onClickSignup = () => {
    history.push('/react/signup');
  };

  const onClickSignout = async () => {
    await getFireBaseAuth().signOut();
    history.push('/react');
  };

  const backHome = () => {
    history.push('/react');
  };
  useEffect(() => {
    setAuth(getFireBaseAuth());
    if (auth.currentUser) {
      setAuthButton(
        <button
          className="bg-black hover:bg-text-gray-800 text-white ml-4 py-2 px-3"
          data-testid="signout"
          onClick={onClickSignout}
        >
          サインアウト
        </button>,
      );
    } else {
      setAuthButton(
        <button
          className="bg-black hover:bg-text-gray-800 text-white ml-4 py-2 px-3"
          data-testid="signup"
          onClick={onClickSignup}
        >
          サインアップ
        </button>,
      );
    }
  }, []);

  return (
    <nav className="font-sans bg-white text-center flex justify-between my-4 mx-auto container overflow-hidden">
      <img
        src="http://localhost/img/icon.png"
        alt="logo"
        onClick={backHome}
        className="h-10 sm:h-10 rounded-full"
      />
      <ul className="text-sm text-gray-700 list-none p-0 flex items-center">
        <li className="inline-block py-2 px-3 text-gray-900 hover:text-gray-700 no-underline">
          MyFavの使い方
        </li>
        <li>{authButton}</li>
      </ul>
    </nav>
  );
};
