import React from 'react';
export type AuthButtonProps = {
  testId: string;
  onClick: () => void;
  buttonBody: string;
};

export const CommonButton = (props: AuthButtonProps) => {
  return (
    <button
      className="bg-black hover:bg-text-gray-800 text-white ml-4 py-2 px-3"
      data-testid={props.testId}
      onClick={props.onClick}
    >
      {props.buttonBody}
    </button>
  );
};
