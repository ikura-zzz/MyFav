import React, { ChangeEvent, KeyboardEvent } from 'react';

export type InputBoxProps = {
  type: string;
  id: string;
  placeHolder: string;
  value: string;
  testId: string;
  onChange: (event: ChangeEvent<HTMLInputElement>) => void;
  onKeyDown: (event: KeyboardEvent<HTMLElement>) => void;
};

export const InputBox = (props: InputBoxProps) => {
  return (
    <input
      type={props.type}
      id={props.id}
      className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
      placeholder={props.placeHolder}
      onChange={props.onChange}
      onKeyDown={props.onKeyDown}
      value={props.value}
      required
      data-testid={props.testId}
    />
  );
};

export type CommonLabelProps = {
  htmlFor: string;
  labelBody: string;
};

export const CommonLabel = (props: CommonLabelProps) => {
  return (
    <label
      className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
      htmlFor={props.htmlFor}
    >
      {props.labelBody}
    </label>
  );
};
