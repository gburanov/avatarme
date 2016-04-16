package main

import (
  "log"
  "fmt"

  "github.com/goamz/goamz/aws"
  "github.com/goamz/goamz/sqs"
)

func ReadMessage() sqs.Message {
  conn := sqs.New(auth, aws.EUWest)
  q, err := conn.CreateQueue("test_goapp")
  if err != nil {
    log.Fatal(err)
  }
  for {
    resp, err := q.ReceiveMessage(1)
    if err != nil {
      log.Fatal(err)
    }
    if len(resp.Messages) != 0 {
      return resp.Messages[0]
    }
  }
}

func DeleteMessage(m sqs.Message) {
  conn := sqs.New(auth, aws.EUWest)
  q, err := conn.CreateQueue("test_goapp")
  if err != nil {
    log.Fatal(err)
  }
  q.DeleteMessage(&m)
}

func SendMessage(message string) {
  conn := sqs.New(auth, aws.EUWest)
  q, err := conn.CreateQueue("test_goapp")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(q.Url)
  q.SendMessage(message)
}
