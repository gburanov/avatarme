package main

type User string

func NewUser(email string) User {
  return User(email)
}
