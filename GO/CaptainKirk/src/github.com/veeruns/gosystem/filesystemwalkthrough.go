package gosystem

import (
  "fmt"
  "os"
)

const (
  MAXDIRS=20
)

func Walkfilesystem (dir string){
      f,err := os.Open(dir)
      if err != nil {
        panic(err)
      }
      files,err:= f.Readdir(MAXDIRS)
      if err != nil {
        panic(err)
        }
       for _,fi := range files {
          fmt.Printf("%s/%s\n",dir,fi.Name())
          if fi.IsDir() {
            Walkfilesystem(fmt.Sprintf("%s/%s",fi.Name()))
          }
         }
       }
