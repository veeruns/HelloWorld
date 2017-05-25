package codekata


import (
"fmt"
"flag"
)

func Flagoperation(){
  var Logfile string
  flag.StringVar(&Logfile,"logfile","default.log","Log filename to write")
  var Loglevel string
  flag.StringVar(&Loglevel,"loglevel","INFO","Log level")
  var Concurrency int
  flag.IntVar(&Concurrency,"concurrency",2,"Number of threads")
  fmt.Printf("Done\n")
  flag.Parse()
}
