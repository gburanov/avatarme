package main

import (
  "log"
  "fmt"

  "github.com/goamz/goamz/aws"
  "github.com/goamz/goamz/sqs"
)

func SendMessage(message string) {
  conn := sqs.New(auth, aws.EUWest)
  q, err := conn.CreateQueue("test_goapp")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(q.Url)
  q.SendMessage(message)
}
