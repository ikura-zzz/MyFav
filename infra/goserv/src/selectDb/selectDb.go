package selectdb

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

// GetShokuzaiTd lsdbコマンドの実行結果（食材テーブルのselect結果をCSVに整形したもの）からHTMLテーブルを作って、String型で返却する
func GetShokuzaiTd() string {

	tdcsv, err := exec.Command("lsdb", "-s").Output()
	if err != nil {
		return "error!"
	}

	return getShokuzaiTdHTML(tdcsv)
}

// 食材（料理の材料）のみを照会
func GetShokuzaiTD_Food() string {
	tdcsv, err := exec.Command("lsdb", "-s", "food").Output()
	if err != nil {
		return "error!"
	}

	return getShokuzaiTdHTML(tdcsv)
}

// 調味料だけを照会
func GetShokuzaiTD_Seas() string {
	tdcsv, err := exec.Command("lsdb", "-s", "seas").Output()
	if err != nil {
		return "error!"
	}

	return getShokuzaiTdHTML(tdcsv)
}

// 粉類だけを照会
func GetShokuzaiTD_Powd() string {
	tdcsv, err := exec.Command("lsdb", "-s", "pwod").Output()
	if err != nil {
		return "error!"
	}

	return getShokuzaiTdHTML(tdcsv)
}

// お菓子・デザートだけを照会
func GetShokuzaiTD_Swet() string {
	tdcsv, err := exec.Command("lsdb", "-s", "swet").Output()
	if err != nil {
		return "error!"
	}

	return getShokuzaiTdHTML(tdcsv)
}

// 冷凍食品だけを照会
func GetShokuzaiTD_Froz() string {
	tdcsv, err := exec.Command("lsdb", "-s", "froz").Output()
	if err != nil {
		return "error!"
	}

	return getShokuzaiTdHTML(tdcsv)
}

// 冷凍食品だけを照会
func GetShokuzaiTD_Frut() string {
	tdcsv, err := exec.Command("lsdb", "-s", "frut").Output()
	if err != nil {
		return "error!"
	}

	return getShokuzaiTdHTML(tdcsv)
}

// lsdbの実行結果をHTMLに変換する
func getShokuzaiTdHTML(tdcsv []byte) string {
	if len(tdcsv) <= 1 {
		return "データがありません"
	}
	var tdhtml string
	tdcsvSlice := strings.Split(string(tdcsv), ",")
	clmCnt := 0
	for i := 0; i < len(tdcsvSlice); i++ {
		var nextTag string
		nextTag, clmCnt = getNextTag(tdcsvSlice[i], clmCnt)
		tdhtml = tdhtml + nextTag
	}

	return tdhtml
}

func getNextTag(record string, cnt int) (string, int) {
	var tag string
	if record == "record" {
		tag = "<tr>\n<form method=\"GET\" action=\"registShokuzai\" class=\"shokuzaiList\">\n"
		cnt = 0
		return tag, cnt
	}
	switch cnt {
	case 0:
		tag = "<input type=\"hidden\" class=\"seqnum\" id=\"seqnum\" name=\"seqnum\" value=\"" + record + "\">\n"
	case 1:
		tag = "<td><input type=\"text\" class=\"name\" id=\"name\" name=\"name\" value=\"" + record + "\"></td>\n"
	case 3:
		tag = "<td><input type=\"number\" class=\"quantity\" id=\"quantity\" name=\"quantity\" step=\"0.1\" value=\"" + record + "\"></td>\n"
	case 4:
		tag = "<td><input type=\"number\" class=\"weight\" id=\"weight\" name=\"weight\" step=\"1\" value=\"" + record + "\"></td>\n"
	case 5:
		tag = "<td><input type=\"date\" class=\"uselimit\" id=\"uselimit\" name=\"uselimit\" value=\"" + record + "\"></td>\n"
	case 6:
		tag = "<td><input type=\"text\" class=\"comment\" id=\"comment\" name=\"comment\" value=\"" + record + "\"></td>\n"
	case 7:
		tag = "<input type=\"hidden\" class=\"registdate\" id=\"registdate\" name=\"registdate\" value=\"" + record + "\">\n"
	default:
		tag = ""
	}
	r := regexp.MustCompile(`end`)
	if r.MatchString(record) {
		tag = "<td><input type=\"submit\" class=\"submit\" value=\"変更\"></td>\n</form></tr>\n"
		return tag, cnt + 1
	}
	return tag, cnt + 1
}

// ChkIdentifier ログインできるかをチェックする。
func ChkIdentifier(user string, pass string) error {
	_, err := exec.Command("lsdb", "-u", user, pass).CombinedOutput()
	if err != nil {
		return errors.New("password unmatched")
	}
	return nil
}
