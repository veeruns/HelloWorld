package codekata

import (
  "testing"
  "fmt"
  "flag"
)

func TestFlagoperation(t *testing.T){
        fs := flag.NewFlagSet("unittest", flag.ExitOnError)
        fs.String("logfile", "default.log", "Some usage msg")
         fs.Parse([]string{"-logfile", "test.log"})
         flag.CommandLine=fs
         Flagoperation()
        fmt.Printf("Cmd: %s, DefaultVal: %s, ActualVal: %s, Options %s\n", "logfile", flag.Lookup("logfile").DefValue, flag.Lookup("logfile").Value)

}
