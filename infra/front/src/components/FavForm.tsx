import React, { ChangeEvent, useState, MouseEvent } from 'react';
import { CreateFavRequest } from '../types/CreateFavRequest';
import { useHistory } from 'react-router-dom';
import * as common from '../types/CommonKeyValue';

export const FavForm = () => {
  const [title, setTitle] = useState<string>('');
  const [category, setCategory] = useState<string>('');
  const [publisher, setPublisher] = useState<string>('');
  const [overview, setOverview] = useState<string>('');
  const [impression, setImpression] = useState<string>('');
  const [timing, setTiming] = useState<string>('now');
  const [stars, setStars] = useState<number>(3);
  const [openClose, setOpenClose] = useState<boolean>(false);
  const [icon, setIcon] = useState<string>('');

  const history = useHistory();

  const onChangeIcon = (event: ChangeEvent<HTMLInputElement>) => {
    setIcon(event.target.value);
  };

  const onChangeTitle = (event: ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  const onChangeCategory = (event: ChangeEvent<HTMLInputElement>) => {
    setCategory(event.target.value);
  };

  const onChangePublisher = (event: ChangeEvent<HTMLInputElement>) => {
    setPublisher(event.target.value);
  };

  const onChangeOverview = (event: ChangeEvent<HTMLInputElement>) => {
    setOverview(event.target.value);
  };

  const onChangeImpression = (event: ChangeEvent<HTMLTextAreaElement>) => {
    setImpression(event.target.value);
  };

  const onClickRadioTiming = (event: MouseEvent<HTMLInputElement>) => {
    setTiming(event.currentTarget.value);
  };

  const onClickRadioStars = (event: MouseEvent<HTMLInputElement>) => {
    setStars(Number(event.currentTarget.value));
  };

  function onClickRadioOpenClose(event: MouseEvent<HTMLInputElement>) {
    setOpenClose(event.currentTarget.value === 'open');
  }

  const onClickSubmit = () => {
    const fav: CreateFavRequest = {
      icon: icon,
      title: title,
      category: category,
      publisher: publisher,
      overview: overview,
      impression: impression,
      timing: timing,
      stars: stars,
      openclose: openClose,
    };
    console.log(fav);
  };

  function onClickCancel(event: MouseEvent<HTMLInputElement>): void {
    history.push('/react/list');
  }

  return (
    <div>
      <div className="fa-2x font-sans rounded w-full max-w-md mx-auto px-8 appearance-none font-bold">
        新しいお気に入り
      </div>
      <form
        className="font-sans text-sm rounded w-full max-w-md mx-auto my-8 px-8 pt-6 pb-8"
        id="newfav"
        encType="multipart/form-data"
      >
        <div className="relative border shadow rounded mb-4 appearance-none">
          アイコン
          <input
            type="file"
            name="file"
            id="file"
            accept="image/*"
            data-testid="icon"
            onChange={onChangeIcon}
          />
          <p className="canvas">
            <canvas
              className="canvas"
              id="canvas"
              width="0"
              height="0"
            ></canvas>
          </p>
          <img src="" id="iconimg" alt="icon" />
          <input type="hidden" id="icon" name="icon" />
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <input
            name="title"
            id="title"
            type="text"
            className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
            placeholder="タイトル"
            required
            data-testid="favtitle"
            onChange={onChangeTitle}
            value={title}
          />
          <label
            className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
            htmlFor="title"
          >
            タイトル
          </label>
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <input
            className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
            name="category"
            id="category"
            type="text"
            placeholder="カテゴリ"
            data-testid="favcategory"
            onChange={onChangeCategory}
            value={category}
          />
          <label
            className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
            htmlFor="category"
          >
            カテゴリ
          </label>
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <input
            className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
            name="publisher"
            id="publisher"
            type="text"
            placeholder="著者・製作者・提供者など"
            data-testid="publisher"
            onChange={onChangePublisher}
            value={publisher}
          />
          <label
            className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
            htmlFor="publisher"
          >
            著者・製作者・提供者など
          </label>
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <input
            className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
            name="overview"
            id="overview"
            type="text"
            placeholder="概要"
            data-testid="favoverview"
            onChange={onChangeOverview}
            value={overview}
          />
          <label
            className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
            htmlFor="overview"
          >
            概要
          </label>
        </div>
        <div className="relative border rounded mb-4 shadow appearance-none label-floating">
          <textarea
            className="w-full py-2 px-3 text-gray-700 leading-normal rounded"
            name="impre"
            id="impre"
            placeholder="フリーコメント"
            data-testid="favimpression"
            onChange={onChangeImpression}
            value={impression}
          ></textarea>
          <label
            className="absolute block text-gray-700 top-0 left-0 w-full px-3 py-2 leading-normal"
            htmlFor="impre"
          ></label>
        </div>
        <div className="favradio border shadow flex items-center mb-4 justify-between">
          <div className="radio_label">ステータス</div>
          <div className="radio_body" data-testid="favtiming">
            <input
              type="radio"
              id={common.already}
              name="timing"
              value={common.already}
              checked={timing === common.already}
              onClick={onClickRadioTiming}
            />
            <label
              className="tab_item radio_when text-base tracking-wide px-3"
              htmlFor="already"
            >
              前から好き/好きだった
            </label>
            <input
              type="radio"
              id={common.now}
              name="timing"
              value={common.now}
              checked={timing === common.now}
              onClick={onClickRadioTiming}
            />
            <label
              className="tab_item radio_when text-base tracking-wide px-3"
              htmlFor="now"
            >
              今好き
            </label>
            <input
              type="radio"
              id={common.wish}
              name="timing"
              value={common.wish}
              checked={timing === common.wish}
              onClick={onClickRadioTiming}
            />
            <label
              className="tab_item radio_when text-base tracking-wide px-3"
              htmlFor="wish"
            >
              興味ある
            </label>
          </div>
        </div>
        <div className="favradio border shadow flex items-center mb-4 justify-between">
          <div className="radio_label">お気に入り度</div>
          <div className="radio_body" data-testid="favstars">
            <input
              type="radio"
              id="one"
              name="stars"
              value={common.star1}
              checked={stars === common.star1}
              onClick={onClickRadioStars}
            />
            <label
              className="tab_item radio_star text-base tracking-wide px-5"
              htmlFor="one"
            >
              1
            </label>
            <input
              type="radio"
              id="two"
              name="stars"
              value={common.star2}
              checked={stars === common.star2}
              onClick={onClickRadioStars}
            />
            <label
              className="tab_item radio_star text-base tracking-wide px-5"
              htmlFor="two"
            >
              2
            </label>
            <input
              type="radio"
              id="three"
              name="stars"
              value={common.star3}
              checked={stars === common.star3}
              onClick={onClickRadioStars}
            />
            <label
              className="tab_item radio_star text-base tracking-wide px-5"
              htmlFor="three"
            >
              3
            </label>
            <input
              type="radio"
              id="four"
              name="stars"
              value={common.star4}
              checked={stars === common.star4}
              onClick={onClickRadioStars}
            />
            <label
              className="tab_item radio_star text-base tracking-wide px-5"
              htmlFor="four"
            >
              4
            </label>
            <input
              type="radio"
              id="five"
              name="stars"
              value={common.star5}
              checked={stars === common.star5}
              onClick={onClickRadioStars}
            />
            <label
              className="tab_item radio_star text-base tracking-wide px-5"
              htmlFor="five"
            >
              5
            </label>
          </div>
        </div>
        <div className="favradio border shadow flex items-center mb-4 justify-between">
          <div className="radio_label">公開設定</div>
          <div className="radio_body" data-testid="favopenclose">
            <input
              type="radio"
              id="open"
              name="openclose"
              value="open"
              checked={openClose}
              onClick={onClickRadioOpenClose}
            />
            <label
              className="tab_item radio_opcl text-base tracking-wide px-10"
              htmlFor="open"
            >
              公開
            </label>
            <input
              type="radio"
              id="close"
              name="openclose"
              value="close"
              checked={!openClose}
              onClick={onClickRadioOpenClose}
            />
            <label
              className="tab_item radio_opcl text-base tracking-wide px-10"
              htmlFor="close"
            >
              非公開
            </label>
          </div>
        </div>
        <div className="flex items-center justify-between">
          <input
            className="bg-black hover:bg-black text-white py-2 px-4"
            type="button"
            value="登録"
            id="create"
            data-testid="favcreatebutton"
            onClick={onClickSubmit}
          />
          <input
            className="bg-white inline-block align-baseline text-gray-500 hover:text-gray-700"
            type="button"
            value="キャンセル"
            id="cancel"
            data-testid="favcancelbutton"
            onClick={onClickCancel}
          />
        </div>
      </form>

      <div className="popup" id="js-popup">
        <div className="popup-inner">
          <div className="close-btn" id="js-close-btn">
            <i className="fas fa-times"></i>
          </div>
          <input
            id="scal"
            className="rangebar"
            type="range"
            value=""
            min="1"
            max="100"
            data-testid="favscal"
          />
          <br />
          <canvas className="canvas" id="cvs" width="500" height="500"></canvas>
          <button
            className="bg-black hover:bg-black text-white py-2 px-4"
            id="crop"
            data-testid="favcropimg"
          >
            切り抜き
          </button>
        </div>
        <div className="black-background" id="js-black-bg"></div>
        <script src="http://localhost/js/fav.js"></script>
      </div>
    </div>
  );
};
