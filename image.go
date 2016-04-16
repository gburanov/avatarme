package main

import (
  "errors"
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

func FilledRect(x1, y1, x2, y2 int, img *image.RGBA, col color.Color) {
  for y:= y1; y < y2; y++ {
    HLine(x1, y, x2, img, col)
  }
}

func (user *User) ColorForNumber(i int) (error, *color.Color) {
  if (i > 32) {
    return errors.New("Invalid number"), nil
  }

  bytes:= []byte(*user)
  return nil, &palette.Plan9[bytes[i]]
}

func (user *User) generateImage() image.Image {
  side := 40
  m := image.NewRGBA(image.Rect(0, 0, side, side))
  i := 0
  for x := 0; x < 4; x++ {
    for y:= 0; y < 4; y++ {
      _, color := user.ColorForNumber(i)
      i++
      FilledRect(x * 10, y * 10, x * 10 + 10, y * 10 + 10, m, *color)
    }
  }
  return m
}
