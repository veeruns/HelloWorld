package codekata


import (
"fmt"
"flag"
)
var Logfile string
var Loglevel string
var Concurrency int
func Flagoperation(){

  flag.StringVar(&Logfile,"logfile","default.log","Log filename to write")

  flag.StringVar(&Loglevel,"loglevel","INFO","Log level")
  
  flag.IntVar(&Concurrency,"concurrency",2,"Number of threads")
  fmt.Printf("Done\n")
  flag.Parse()
}
