package helloworld

import (
  "fmt"
  "math/rand"
)

func getrandomnumber(i int) (chan int) {
  ch := make(chan int,10)

go func() {  for i:=0;i<10;i++ {
    rn := rand.Intn(5)
    ch <- rn
  }

close(ch)
}()
return ch
}

func PrintRandom(){
ch := getrandomnumber(10)
fmt.Printf("Length of channel %d\n",len(ch))
for num:= range ch{
    fmt.Printf("%d\n",num)

}
}
