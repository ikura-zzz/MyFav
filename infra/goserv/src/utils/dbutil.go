package utils

// SelectUserPass ユーザー名でパスワードハッシュを取得するときのSQL
const SelectUserPass string = "SELECT passhash from appusers where username=?"

// SelectUsernameCnt ユーザー名で、同一ユーザー名のレコード数を取得するときのSQL
const SelectUsernameCnt string = "SELECT count(*) from appusers where username=?"

// UserInsertSQL ユーザーを追加するときのSQL
const UserInsertSQL string = "INSERT INTO appusers (username, passhash,upddate) values(?,?,?)"

const CmnErrmsg string = "予期せぬエラーが発生しました。"
