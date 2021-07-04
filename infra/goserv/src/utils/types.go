package utils

import (
	"fmt"
	"time"
)

// Shokuzai はshokuzai構造体を表す。
type Shokuzai struct {
	Seqnum     int
	Name       string
	Category   string
	Quantity   float64
	Weight     int
	Uselimit   time.Time
	Comment    string
	Registdate time.Time
}

// ShokuzaiTypeLength 食材型の要素の数
const ShokuzaiTypeLength int = 8

// TimeFormat タイム型を文字列にするときのフォーマット
const TimeFormat = "2006-01-02"

// ShokuzaiToStrSlice 引数で渡されたShokuzai型をstringのSliceにして返す
func ShokuzaiToStrSlice(shokuzai Shokuzai) []string {
	var strSlice []string

	strSlice = append(strSlice, fmt.Sprint(shokuzai.Seqnum))
	strSlice = append(strSlice, shokuzai.Name)
	strSlice = append(strSlice, shokuzai.Category)
	strSlice = append(strSlice, fmt.Sprint(shokuzai.Quantity))
	strSlice = append(strSlice, fmt.Sprint(shokuzai.Weight))
	strSlice = append(strSlice, shokuzai.Uselimit.Format(TimeFormat))
	strSlice = append(strSlice, fmt.Sprint(shokuzai.Comment))
	strSlice = append(strSlice, shokuzai.Registdate.Format(TimeFormat))

	return strSlice
}
