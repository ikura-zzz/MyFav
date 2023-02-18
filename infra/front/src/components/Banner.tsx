import React from 'react';
import { useHistory } from 'react-router-dom';
import './Banner.css';

export const Banner = () => {
  const history = useHistory();
  const onClickSignup = () => {
    history.push('/react/signup');
  };

  const backHome = () => {
    history.push('/react');
  };
  return (
    <nav className="nav">
      <img
        src="http://localhost/img/icon.png"
        alt="logo"
        onClick={backHome}
        className="logo"
      />
      <ul className="bannerList">
        <li className="howtoLink">MyFavの使い方</li>
        <button
          className="signupButton"
          data-testid="signup"
          onClick={onClickSignup}
        >
          サインアップ
        </button>
      </ul>
    </nav>
  );
};
