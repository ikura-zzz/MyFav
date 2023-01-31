import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';

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

  return (
    <>
      <div>お気に入りリスト</div>
      <section>
        <div>
          Tweet
          <button
            className=""
            id="copy"
            value="http://localhost/list?name=ikura"
          >
            リンクをコピー
          </button>
        </div>
        <div>
          <input type="radio" name="tab_item"></input>
          <label>すべて</label>
          <input type="radio" name="tab_item"></input>
          <label>前から</label>
          <input type="radio" name="tab_item"></input>
          <label>今好き</label>
          <input type="radio" name="tab_item"></input>
          <label>興味ある</label>
        </div>
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
      </section>
      <div>
        <img
          src="http://localhost/img/plus.png"
          alt="plusbuton"
          style={{ backgroundColor: 'gray' }}
          onClick={() => {
            history.push('/react/fav');
          }}
        />
      </div>
    </>
  );
};
