package main

import (
    "fmt"

    "github.com/Foxes-cool/Foxes-Go"
)

func main() {
    message, _ := foxes.Fox(foxes.Options{})
    fmt.Println(message)
    message, _ = foxes.Fox(foxes.Options{Width: 150})
    fmt.Println(message)
    message, _ = foxes.Fox(foxes.Options{Width: 150, Height: 150})
    fmt.Println(message)
}
