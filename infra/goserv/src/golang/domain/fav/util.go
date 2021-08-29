package fav

func Timingconv(timing string) string {
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

func Starsconv(stars string) string {
	if stars == "1" {
		return "1"
	}
	if stars == "2" {
		return "2"
	}
	if stars == "3" {
		return "3"
	}
	if stars == "4" {
		return "4"
	}
	if stars == "5" {
		return "5"
	}

	return "3"
}
func Opclconv(opcl string) string {
	if opcl == "open" {
		return "1"
	}
	if opcl == "close" {
		return "0"
	}

	return "0"
}
