package codekata


import (
"fmt"
"flag"
)

func Flagoperation(){
  Logfile:=flag.String("logfile","default.log","Log filename to write")
  Loglevel:=flag.String("loglevel","INFO","Log level")
  Concurrency:=flag.Int("concurrency","2","Number of threads")
  fmt.Printf("Done\n")
  flag.Parse()
}
