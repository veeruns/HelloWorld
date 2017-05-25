package helloworld

import (
  "fmt"
  "math/rand"
)

func getrandomnumber(i int) (chan int) {
  ch := make(chan int,i)

go func() {
  for j:=1;j<11;j++ {
    rn := rand.Intn(j)
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
