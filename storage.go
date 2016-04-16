package main

import (
  "bytes"
  "image"
  "image/jpeg"
  "log"
  "github.com/goamz/goamz/aws"
  "github.com/goamz/goamz/s3"
)

const bucketName = "gburanov-avatars"

func SaveToS3(image *image.Image, name string) error {
  var bytes bytes.Buffer

  err := jpeg.Encode(&bytes, *image, nil)
  if err != nil {
    log.Fatal(err)
  }

  s3instance:= s3.New(auth, aws.EUWest)
  bucket := s3instance.Bucket(bucketName)
  return bucket.Put(name, bytes.Bytes(), "image/jpeg", s3.PublicRead, s3.Options{})
}
