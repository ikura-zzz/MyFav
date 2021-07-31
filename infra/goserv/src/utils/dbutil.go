package utils

// SelectUserPass ユーザー名でパスワードハッシュを取得するときのSQL
const SelectUserPass string = "SELECT passhash FROM appusers WHERE username=?"

// SelectUserID ユーザーIDを取得するときのSQL
const SelectUserID string = "SELECT userid FROM appusers WHERE username=?"

// SelectUsernameCnt ユーザー名で、同一ユーザー名のレコード数を取得するときのSQL
const SelectUsernameCnt string = "SELECT count(*) FROM appusers WHERE username=?"

// SelectFavsByUserid FavをユーザーIDで抽出する
const SelectFavsByUserid string = "SELECT f.favid,f.title,f.category,f.publisher,f.overview,f.impression,f.timing,f.stars,f.openclose,i.image FROM favs AS f JOIN images AS i ON f.favid = i.favid WHERE f.userid=? ORDER BY category, publisher, stars DESC, upddate"

// SelectFavsByUserOpen FavをユーザーIDで抽出する
const SelectFavsByUseridAndOpen string = "SELECT f.favid,f.title,f.category,f.publisher,f.overview,f.impression,f.timing,f.stars,f.openclose,i.image FROM favs AS f JOIN images AS i ON f.favid = i.favid WHERE f.userid=? AND f.openclose=1 ORDER BY category, publisher, stars DESC, upddate"

// UserInsertSQL ユーザーを追加するときのSQL
const UserInsertSQL string = "INSERT INTO appusers (username, passhash,upddate) values(?,?,?)"

// FavInsertSQL Favを追加するときのSQL
const FavInsertSQL string = "INSERT INTO favs (userid, title, category, publisher, overview, impression, timing, stars, openclose, upddate) values(?,?,?,?,?,?,?,?,?,?)"

// ImageInsertSQL iconを追加するときのSQL
const ImageInsertSQL string = "INSERT INTO images (favid,image) value(?,?)"

// FavUpdateSQL Favを更新するときのSQL
const FavUpdateSQL string = "UPDATE favs SET title=?, category=?, publisher=?, overview=?, impression=?, timing=?, stars=?, openclose=?, upddate=? WHERE userid=? AND favid=?"

// ImageUpdateSQL iconを追加するときのSQL
const ImageUpdateSQL string = "UPDATE images SET image=? WHERE favid=?"

// UserNameUpdateSQL ユーザー名を変更するときのSQL
const UserNameUpdateSQL string = "UPDATE appusers SET username=?, upddate=? WHERE userid=?"

// UserPasswordUpdateSQL ユーザーのパスワードを変更するときのSQL
const UserPasswordUpdateSQL string = "UPDATE appusers set passhash=?, upddate=? WHERE usreid=?"

// FavDeleteSQL Favを追加するときのSQL
const FavDeleteSQL string = "DELETE FROM favs WHERE userid=? AND favid=?"

// ImageDeleteSQL iconを追加するときのSQL
const ImageDeleteSQL string = "DELETE FROM images WHERE favid=?"

// UserDeleteSQL ユーザーを削除するときのSQL
const UserDeleteSQL string = "DELETE FROM appusers WHERE userid=?"

// CmnErrmsg DBアクセス系の共通エラー
const CmnErrmsg string = "予期せぬエラーが発生しました。"
