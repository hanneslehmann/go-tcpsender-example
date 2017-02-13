package main

import (
  tcpsender "github.com/hanneslehmann/golib-tcpsender"
  "github.com/Sirupsen/logrus"
  "time"
)

// inspired by example of logrus: https://github.com/sirupsen/logrus
// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func main() {
  // create a TCP sender and start it
  sender_one:=tcpsender.New("localhost",6000)
  go sender_one.StartAndListen()

  // The API for setting attributes is a little different than the package level
  // exported logger. See Godoc.
  // tcpsender can be used as output for logrus
  log.Out =  sender_one

  //wait some seconds to allow clients to connect
  time.Sleep(2000 * time.Millisecond)

  // send example log message through tcp
  log.WithFields(logrus.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")

  //wait some seconds to allow clients to connect
  time.Sleep(2000 * time.Millisecond)

  // send example log message through tcp, format in JSON
  log.Formatter = &logrus.JSONFormatter{}
  log.WithFields(logrus.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")

  for {}
}
