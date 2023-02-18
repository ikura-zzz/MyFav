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
      <div>新しいお気に入り</div>
      <form id="newfav" encType="multipart/form-data">
        <div>
          アイコン
          <input
            type="file"
            name="file"
            id="file"
            accept="image/*"
            data-testid="icon"
            onChange={onChangeIcon}
          />
          <p>
            <canvas id="canvas" width="0" height="0"></canvas>
          </p>
          <img src="" id="iconimg" alt="icon" />
          <input type="hidden" id="icon" name="icon" />
        </div>
        <div>
          <input
            name="title"
            id="title"
            type="text"
            placeholder="タイトル"
            required
            data-testid="favtitle"
            onChange={onChangeTitle}
            value={title}
          />
          <label htmlFor="title">タイトル</label>
        </div>
        <div>
          <input
            name="category"
            id="category"
            type="text"
            placeholder="カテゴリ"
            data-testid="favcategory"
            onChange={onChangeCategory}
            value={category}
          />
          <label htmlFor="category">カテゴリ</label>
        </div>
        <div>
          <input
            name="publisher"
            id="publisher"
            type="text"
            placeholder="著者・製作者・提供者など"
            data-testid="publisher"
            onChange={onChangePublisher}
            value={publisher}
          />
          <label htmlFor="publisher">著者・製作者・提供者など</label>
        </div>
        <div>
          <input
            name="overview"
            id="overview"
            type="text"
            placeholder="概要"
            data-testid="favoverview"
            onChange={onChangeOverview}
            value={overview}
          />
          <label htmlFor="overview">概要</label>
        </div>
        <div>
          <textarea
            name="impre"
            id="impre"
            placeholder="フリーコメント"
            data-testid="favimpression"
            onChange={onChangeImpression}
            value={impression}
          ></textarea>
          <label htmlFor="impre"></label>
        </div>
        <div>
          <div>ステータス</div>
          <div data-testid="favtiming">
            <input
              type="radio"
              id={common.already}
              name="timing"
              value={common.already}
              checked={timing === common.already}
              onClick={onClickRadioTiming}
            />
            <label htmlFor="already">前から好き/好きだった</label>
            <input
              type="radio"
              id={common.now}
              name="timing"
              value={common.now}
              checked={timing === common.now}
              onClick={onClickRadioTiming}
            />
            <label htmlFor="now">今好き</label>
            <input
              type="radio"
              id={common.wish}
              name="timing"
              value={common.wish}
              checked={timing === common.wish}
              onClick={onClickRadioTiming}
            />
            <label htmlFor="wish">興味ある</label>
          </div>
        </div>
        <div>
          <div>お気に入り度</div>
          <div data-testid="favstars">
            <input
              type="radio"
              id="one"
              name="stars"
              value={common.star1}
              checked={stars === common.star1}
              onClick={onClickRadioStars}
            />
            <label htmlFor="one">1</label>
            <input
              type="radio"
              id="two"
              name="stars"
              value={common.star2}
              checked={stars === common.star2}
              onClick={onClickRadioStars}
            />
            <label htmlFor="two">2</label>
            <input
              type="radio"
              id="three"
              name="stars"
              value={common.star3}
              checked={stars === common.star3}
              onClick={onClickRadioStars}
            />
            <label htmlFor="three">3</label>
            <input
              type="radio"
              id="four"
              name="stars"
              value={common.star4}
              checked={stars === common.star4}
              onClick={onClickRadioStars}
            />
            <label htmlFor="four">4</label>
            <input
              type="radio"
              id="five"
              name="stars"
              value={common.star5}
              checked={stars === common.star5}
              onClick={onClickRadioStars}
            />
            <label htmlFor="five">5</label>
          </div>
        </div>
        <div>
          <div>公開設定</div>
          <div data-testid="favopenclose">
            <input
              type="radio"
              id="open"
              name="openclose"
              value="open"
              checked={openClose}
              onClick={onClickRadioOpenClose}
            />
            <label htmlFor="open">公開</label>
            <input
              type="radio"
              id="close"
              name="openclose"
              value="close"
              checked={!openClose}
              onClick={onClickRadioOpenClose}
            />
            <label htmlFor="close">非公開</label>
          </div>
        </div>
        <div>
          <input
            type="button"
            value="登録"
            id="create"
            data-testid="favcreatebutton"
            onClick={onClickSubmit}
          />
          <input
            type="button"
            value="キャンセル"
            id="cancel"
            data-testid="favcancelbutton"
            onClick={onClickCancel}
          />
        </div>
      </form>

      <div id="js-popup">
        <div>
          <div id="js-close-btn">
            <i></i>
          </div>
          <input
            id="scal"
            type="range"
            value=""
            min="1"
            max="100"
            data-testid="favscal"
          />
          <br />
          <canvas id="cvs" width="500" height="500"></canvas>
          <button id="crop" data-testid="favcropimg">
            切り抜き
          </button>
        </div>
        <div id="js-black-bg"></div>
      </div>
    </div>
  );
};
