package main

import (
  "image"
  "image/color"
  "image/color/palette"
)

func HLine(x1, y, x2 int, img *image.RGBA, col color.Color) {
    for ; x1 <= x2; x1++ {
        img.Set(x1, y, col)
    }
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int, img *image.RGBA, col color.Color) {
    for ; y1 <= y2; y1++ {
        img.Set(x, y1, col)
    }
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int, img *image.RGBA, col color.Color) {
    HLine(x1, y1, x2, img, col)
    HLine(x1, y2, x2, img, col)
    VLine(x1, y1, y2, img, col)
    VLine(x2, y1, y2, img, col)
}

func (user *User) generateImage() image.Image {
  side := 40
  m := image.NewRGBA(image.Rect(0, 0, side, side))
  Rect(0, 0, 10, 10, m, palette.Plan9[200])
  Rect(10, 10, 20, 20, m, palette.Plan9[100])
  return m
}
