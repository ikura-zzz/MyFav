import React from "react";
import { useHistory } from "react-router-dom";

export const FavList = () => {
  const history = useHistory();
  return (
    <>
      <div>お気に入りリスト</div>
      <section>
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
      </section>
      <div>
        <img
          src="http://localhost/img/plus.png"
          alt="plusbuton"
          style={{ backgroundColor: "gray" }}
          onClick={() => {
            history.push("/react/fav");
          }}
        />
      </div>
    </>
  );
};
