import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getFireBaseAuth } from './CommonAuthCheck';
import './main.css';

const auth = getFireBaseAuth();
type Fav = {
  userid: number;
  favid: string;
  icon: string;
  title: string;
  category: string;
  publisher: string;
  overview: string;
  impre: string;
  timing: string;
  stars: string;
  openclose: string;
  datetime: string;
};

export const FavList = () => {
  const history = useHistory();
  const [favlist, setFavlist] = useState<Fav[]>([]);
  auth.onAuthStateChanged((user) => {
    if (!user) {
      history.push('/react');
    }
  });

  return (
    <>
      <div className="fa-2x font-sans rounded w-full max-w-md mx-auto px-8 appearance-none label-floating font-bold container">
        お気に入りリスト
      </div>
      <section className="bg-white py-4 font-sans">
        <div className="relative flex rounded mb-4 appearance-none container m-auto">
          <a
            href="https://twitter.com/share?ref_src=twsrc%5Etfw"
            className="twitter-share-button"
            data-text="わたしのお気に入りリスト"
            data-url="http://localhost/list?name={{.username}}"
            data-show-count="false"
          >
            Tweet
          </a>
          <script async src="https://platform.twitter.com/widgets.js"></script>
          <button
            className="ml-4 bg-black text-white px-2 fa-xs"
            id="copy"
            value="http://localhost/list?name=ikura"
          >
            リンクをコピー
          </button>
        </div>
        <div className="tabs container m-auto max-w-xl flex items-baseline justify-start border-gray-300 mb-10">
          <input id="all" type="radio" name="tab_item"></input>
          <label
            className="tab_item text-gray-700 text-base font-bold tracking-wide py-4 px-3"
            htmlFor="all"
          >
            すべて
          </label>
          <input id="already" type="radio" name="tab_item"></input>
          <label
            className="tab_item text-gray-700 text-base font-bold tracking-wide py-4 px-3"
            htmlFor="already"
          >
            前から
          </label>
          <input id="now" type="radio" name="tab_item"></input>
          <label
            className="tab_item text-gray-700 text-base font-bold tracking-wide py-4 px-3"
            htmlFor="now"
          >
            今好き
          </label>
          <input id="wish" type="radio" name="tab_item"></input>
          <label
            className="tab_item text-gray-700 text-base font-bold tracking-wide py-4 px-3"
            htmlFor="wish"
          >
            興味ある
          </label>
        </div>
        <ul
          className="font-sans list-none p-0 container m-auto text-gray-900"
          data-testid="physicallist"
        >
          {favlist.map((fav) => {
            return (
              <ul className="" data-testid="physicallist" key={fav.favid}>
                <a
                  id="fav"
                  href={'/fav?favid=' + fav.favid}
                  style={{ display: 'block' }}
                >
                  <input type="hidden" id="timing" value={fav.timing} />
                  <li className="">
                    <div className="">
                      <img
                        src="http://localhost/icons/apple-touch-icon.png"
                        className="rounded-full h-10 w-10 bg-gray-300 m-auto"
                        alt="favicon"
                      />
                    </div>
                    <div className="">
                      <span className="font-bold ">{fav.title}</span>
                      <div className="">
                        <p className="">{fav.category}</p>
                        <p className="">{fav.overview}</p>
                        <label htmlFor="stars" className="">
                          {fav.stars}
                        </label>
                      </div>
                    </div>
                  </li>
                </a>
              </ul>
            );
          })}
        </ul>
      </section>
      <div className="crtfav rounded-full h-10 w-10 bg-black m-auto">
        <img
          src="http://localhost/img/plus.png"
          alt="plusbuton"
          style={{ backgroundColor: 'gray' }}
          onClick={() => {
            history.push('/react/fav');
          }}
          className="rounded-full m-auto"
        />
      </div>
    </>
  );
};
