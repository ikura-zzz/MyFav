package logmanager

import (
	"fmt"
	"log"
	"os"
)

const logfile string = "/var/log/myfav/myfav.log"

func Outlog(msg string) {
	stderr := os.Stderr
	defer func() {
		os.Stderr = stderr
	}()

	file, err := os.Open(logfile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	os.Stderr = file

	log.Println(msg)
}
