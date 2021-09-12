package main
import (
    //"fmt"
    "errors"
    "log"
)
func main() {

   log.SetFlags(log.Ldate |log.Ltime | log.Lmicroseconds )

   log.SetPrefix("看看是啥----->")

   message := errors.New("--------->  错误！！！")
   
   log.Print(message)

}