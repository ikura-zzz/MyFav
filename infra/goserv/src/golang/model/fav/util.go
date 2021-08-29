package fav

import "strconv"

func timingconv(timing string) string {
	if timing == "already" {
		return "1"
	}
	if timing == "now" {
		return "2"
	}
	if timing == "wish" {
		return "3"
	}

	return "2"
}

func starsconv(stars string) string {
	num, err := strconv.Atoi(stars)
	if err != nil {
		return "3"
	}
	if num < 1 || 5 < num {
		return "3"
	}
	return stars
}
func opclconv(opcl string) string {
	if opcl == "open" {
		return "1"
	}
	if opcl == "close" {
		return "0"
	}

	return "0"
}
