package main

import (
  "crypto/md5"
)

type User []byte

func NewUser(email string) User {
  hash := md5.New()
  bytes := hash.Sum([]byte(email))
  return User(bytes)
}
