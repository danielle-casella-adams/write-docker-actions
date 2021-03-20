package main

import (
	"fmt"
	"os"
)

func main() {

  // Access Inputs as environment vars
  //In our actions source code we can access the inputs as if they are environment variables. 
  //GitHub Actions takes every inputs: value and converts it by adding INPUT_ and making 
  //the value uppercase.  How you access these varies by language.

  firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
  secondGreeting := os.Getenv("INPUT_SECONDGREETING")
  thirdGreeting := os.Getenv("INPUT_THIRDGREETING")

  // Use those inputs in the actions logic
  fmt.Println("Hello " + firstGreeting)
  fmt.Println("Hello " + secondGreeting)

  // Someimes inputs are not "required" and we can build around that
  if thirdGreeting != "" {
    fmt.Println("Hello " + thirdGreeting)
    }

}