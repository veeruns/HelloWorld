package codekata


import (
"fmt"
"flag"
)
var Options struct {
  Logfile string
  Loglevel string
  Concurrency int
  Timeout int
}
/*
var Logfile string
var Loglevel string
var Concurrency int
var Timeout int
*/
func Flagoperation(){


  flag.StringVar(&Options.Logfile,"logfile","default.log","Log filename to write")

  flag.StringVar(&Options.Loglevel,"loglevel","INFO","Log level")

  flag.IntVar(&Options.Concurrency,"concurrency",2,"Number of threads")

  flag.IntVar(&Options.Timeout,"timeout",2000,"Default timeout for any operations")
  fmt.Printf("Done\n")
  flag.Parse()
}
