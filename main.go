package main

import (
  "log"
  "fmt"
)

func ProcessEmail(email string) {
  fmt.Println("Got email", email)
  user := NewUser(email)
  image := user.generateImage()
  err := SaveToS3(&image, email)
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  for {
    message := ReadMessage()
    ProcessEmail(message.Body)
    DeleteMessage(message)
  }
}
