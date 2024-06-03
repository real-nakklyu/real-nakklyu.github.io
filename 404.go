package main

import(
  "fmt"
  )

interface Client{
  name: String,
  surname: String,
  }

func main() {
  fmt.PrintLn("Writing Please to Everyone")
  Client.name := "Jasey"
  Client.surname := "Smith"
  }
