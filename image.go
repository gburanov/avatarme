package main

import (
  "image"
)

func (user *User) generateImage() image.Image {
  side := 100
  m := image.NewRGBA(image.Rect(0, 0, side, side))
  return m
}
