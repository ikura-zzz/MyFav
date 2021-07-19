package utils

// SelectUserPass ユーザー名でパスワードハッシュを取得するときのSQL
const SelectUserPass string = "SELECT passhash from appusers where username=?"

// SelectUserID ユーザーIDを取得するときのSQL
const SelectUserID string = "SELECT userid from appusers where username=?"

// SelectUsernameCnt ユーザー名で、同一ユーザー名のレコード数を取得するときのSQL
const SelectUsernameCnt string = "SELECT count(*) from appusers where username=?"

// SelectFavsByUserid FavをユーザーIDで抽出する
const SelectFavsByUserid string = "SELECT f.favid,f.title,f.category,f.publisher,f.overview,f.impression,f.timing,f.stars,f.openclose,i.image from favs AS f JOIN images AS i ON f.favid = i.favid where f.userid=?"

// SelectFavsByUseridAlready FavをユーザーIDで抽出する
const SelectFavsByUseridAlready string = "SELECT f.favid,f.title,f.category,f.publisher,f.overview,f.impression,f.timing,f.stars,f.openclose,i.image from favs AS f JOIN images AS i ON f.favid = i.favid where f.timing=1 AND f.userid=?"

// SelectFavsByUseridNow FavをユーザーIDで抽出する
const SelectFavsByUseridNow string = "SELECT f.favid,f.title,f.category,f.publisher,f.overview,f.impression,f.timing,f.stars,f.openclose,i.image from favs AS f JOIN images AS i ON f.favid = i.favid where f.timing=2 AND f.userid=?"

// SelectFavsByUseridWish FavをユーザーIDで抽出する
const SelectFavsByUseridWish string = "SELECT f.favid,f.title,f.category,f.publisher,f.overview,f.impression,f.timing,f.stars,f.openclose,i.image from favs AS f JOIN images AS i ON f.favid = i.favid where f.timing=3 AND f.userid=?"

// UserInsertSQL ユーザーを追加するときのSQL
const UserInsertSQL string = "INSERT INTO appusers (username, passhash,upddate) values(?,?,?)"

// FavInsertSQL Favを追加するときのSQL
const FavInsertSQL string = "INSERT INTO favs (userid, title, category, publisher, overview, impression, timing, stars, openclose, upddate) values(?,?,?,?,?,?,?,?,?,?)"

// ImageInsertSQL iconを追加するときのSQL
const ImageInsertSQL string = "INSERT INTO images (favid,image) value(?,?)"

// FavUpdateSQL Favを追加するときのSQL
const FavUpdateSQL string = "UPDATE favs SET title=?, category=?, publisher=?, overview=?, impression=?, timing=?, stars=?, openclose=?, upddate=? where userid=?, favid=?"

// ImageUpdateSQL iconを追加するときのSQL
const ImageUpdateSQL string = "UPDATE images SET image=? where favid=?"

// CmnErrmsg DBアクセス系の共通エラー
const CmnErrmsg string = "予期せぬエラーが発生しました。"
