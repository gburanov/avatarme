package main

import (
  "log"
)

func main() {
  email := "gburanov@gmail.com"

  user := NewUser(email)
  image := user.generateImage()
  err := SaveToS3(&image, email)
  if err != nil {
    log.Fatal(err)
  }
}
