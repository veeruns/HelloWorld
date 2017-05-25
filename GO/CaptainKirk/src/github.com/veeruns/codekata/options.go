package codekata


import (
"fmt"
"flag"
)
var Logfile string
var Loglevel string
var Concurrency int
var Timeout int
func Flagoperation(){

  flag.StringVar(&Logfile,"logfile","default.log","Log filename to write")

  flag.StringVar(&Loglevel,"loglevel","INFO","Log level")

  flag.IntVar(&Concurrency,"concurrency",2,"Number of threads")

  flag.IntVar(&Timeout,"timeout",2000,"Default timeout for any operations")
  fmt.Printf("Done\n")
  flag.Parse()
}
