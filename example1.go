package main

import (
    tcpsender "github.com/hanneslehmann/golib-tcpsender"
  "time"
)

func main() {
   sender_one:=tcpsender.New("localhost",6000)
   sender_two:=tcpsender.New("localhost",6001)

   go sender_one.StartAndListen()
   go sender_two.StartAndListen()

   sender_one.HeartBeat("- - HeartBeat\n", 1000)
   time.Sleep(4000 * time.Millisecond)
   sender_two.SendMessage("Hello")
   for {}
}
