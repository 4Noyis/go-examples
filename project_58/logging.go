package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standart logger")

	/*Loggers can be configured with flags to set their output format.
	By default, the standard logger has the log.Ldate and log.Ltime flags set, and these are collected in log.LstdFlags. We can change its flags*/

	/*It also supports emitting the file name and line from which the log function is called.*/
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	myLog := log.New(os.Stdout, "my:", log.LstdFlags)
	myLog.Println("from mylog")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)

}
