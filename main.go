package main

import (
  "os"
  "log"
  "image/png"
)

func main() {
  user := NewUser("gburanov@gmail.com")
  image := user.generateImage()

  f, err := os.OpenFile("avatar.jpg", os.O_CREATE | os.O_WRONLY, 0666)
  if err != nil {
    log.Fatal(err)
  }
  err = png.Encode(f, image)
  if err != nil {
    log.Fatal(err)
  }
}
