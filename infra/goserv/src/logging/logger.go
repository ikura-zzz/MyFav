package logging

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const logfile string = "/var/log/myfab/myfav.log"
const High string = "H"
const Mid string = "M"
const Low string = "L"

var LogFlag bool = false

func Log(msg string, priority string) {
	now := time.Now()
	year, month, day := now.Date()
	log := fmt.Sprintf("%d/%s/%d", year, month.String(), day) + " " + now.String() + " " + priority + " " + msg

	for {
		if !LogFlag {
			break
		}
		time.Sleep(30 * time.Millisecond)
	}

	LogFlag = true
	f, _ := os.Open(logfile)
	defer f.Close()

	fw := bufio.NewWriter(f)
	fw.Write([]byte(log))
	LogFlag = false
}
