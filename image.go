package main

import (
  "image"
  "image/draw"
)

func HLine(x1, y, x2 int) {
    for ; x1 <= x2; x1++ {
        img.Set(x1, y, col)
    }
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
    for ; y1 <= y2; y1++ {
        img.Set(x, y1, col)
    }
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
    HLine(x1, y1, x2)
    HLine(x1, y2, x2)
    VLine(x1, y1, y2)
    VLine(x2, y1, y2)
}

func (user *User) generateImage() image.Image {
  side := 40
  m := image.NewRGBA(image.Rect(0, 0, side, side))

  // fill it with data
  for i, byte := range *user {
    draw.Draw(m, image.Rect(0, 0, 10, 10), &image.Uniform{c}, image.ZP, draw.Src)
  }
  return m
}
