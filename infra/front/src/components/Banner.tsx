import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getFireBaseAuth } from './CommonAuthCheck';
import { CommonButton } from './CommonButton';

export const Banner = () => {
  const [authButton, setAuthButton] = useState(<></>);
  const [currentPath, setCurrentPath] = useState('/react');
  const history = useHistory();

  const backHome = () => {
    setCurrentPath('/react');
    history.push(currentPath);
  };

  useEffect(() => {
    const auth = getFireBaseAuth();
    history.push(currentPath);
    const onClickSignup = () => {
      setCurrentPath('/react/signup');
    };

    const OnClickSignin = () => {
      setCurrentPath('/react/list');
    };

    const onClickSignout = async () => {
      await getFireBaseAuth().signOut();
      setCurrentPath('/react');
    };

    auth.onAuthStateChanged((user) => {
      if (user) {
        setAuthButton(
          <CommonButton
            testId="signout"
            onClick={onClickSignout}
            buttonBody="サインアウト"
          />,
        );
      } else if (history.location.pathname === '/react/signup') {
        setAuthButton(
          <CommonButton
            testId="signup"
            onClick={OnClickSignin}
            buttonBody="サインイン"
          />,
        );
      } else {
        setAuthButton(
          <CommonButton
            testId="signup"
            onClick={onClickSignup}
            buttonBody="サインアップ"
          />,
        );
      }
    });
  }, [currentPath, history]);

  return (
    <nav className="font-sans bg-white text-center flex justify-between my-4 mx-auto container overflow-hidden">
      <img
        src="http://localhost/img/icon.png"
        alt="logo"
        onClick={backHome}
        className="h-10 sm:h-10 rounded-full"
      />
      <ul className="text-sm text-gray-700 list-none p-0 flex items-center">
        <HowToLink />
        <li>{authButton}</li>
      </ul>
    </nav>
  );
};

const HowToLink = () => {
  return (
    <li className="inline-block py-2 px-3 text-gray-900 hover:text-gray-700 no-underline">
      MyFavの使い方
    </li>
  );
};
