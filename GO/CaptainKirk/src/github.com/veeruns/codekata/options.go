package codekata


import (
"fmt"
"flag"
)

func Flagoperation(){
  var *Logfile string = flag.String("logfile","default.log","Log filename to write")
  var *Loglevel string =flag.String("loglevel","INFO","Log level")
  var *Concurrency Int =flag.Int("concurrency",2,"Number of threads")
  fmt.Printf("Done\n")
  flag.Parse()
}
