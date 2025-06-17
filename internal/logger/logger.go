package logger

import (
	"io"
	"log"
	"os"
	"time"
	"strconv"
	"strings"
)

var (
	logFile *os.File
)

func InitLogger() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		_ = os.Mkdir("logs", 0755)
	}

	var err error
	logFile, err = os.OpenFile("logs/backup.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	log.SetFlags(0)
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetPrefix("")
}

func CloseLogger() {
	if logFile != nil {
		_ = logFile.Close()
	}
}


func Info(msg string) {
	log.Printf("[%s] \033[34mINFO\033[0m  %s", timestamp(), msg)
}

func Warn(msg string) {
	log.Printf("[%s] \033[33mWARN\033[0m  %s", timestamp(), msg)
}

func Error(msg string) {
	log.Printf("[%s] \033[31mERROR\033[0m %s", timestamp(), msg)
}

func Success(msg string) {
	log.Printf("[%s] \033[32mOK\033[0m    %s", timestamp(), msg)
}

func timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func PrettyList(list []string) string {
	return "[" + strings.Join(list, ", ") + "]"
}

func PrettyInt(n int) string {
	return strconv.Itoa(n)
}
