package logmanager

import (
	"fmt"
	"io"
	"log"
	"os"
)

const logfile string = "/var/log/myfav/myfav.log"

func Outlog(msg string) {
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	log.SetOutput(io.Writer(file))
	log.Println(msg)
}
