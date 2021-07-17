package utils

// SelectUserPass ユーザー名でパスワードハッシュを取得するときのSQL
const SelectUserPass string = "SELECT passhash from appusers where username=?"

// SelectUserID ユーザーIDを取得するときのSQL
const SelectUserID string = "SELECT userid from appusers where username=?"

// SelectUsernameCnt ユーザー名で、同一ユーザー名のレコード数を取得するときのSQL
const SelectUsernameCnt string = "SELECT count(*) from appusers where username=?"

// UserInsertSQL ユーザーを追加するときのSQL
const UserInsertSQL string = "INSERT INTO appusers (username, passhash,upddate) values(?,?,?)"

// FavInsertSQL Favを追加するときのSQL
const FavInsertSQL string = "INSERT INTO favs (userid, title, category, publisher, overview, impression, timing, stars, openclose, upddate) values(?,?,?,?,?,?,?,?,?,?)"

// CmnErrmsg DBアクセス系の共通エラー
const CmnErrmsg string = "予期せぬエラーが発生しました。"
