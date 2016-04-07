package main;

import (
    "fmt"
    "testing"
)

// func TestEdit(t *testing.T) {
//     b := Book{"1984", "Georges Orwell", "long ago", "my favorite book"}
//     b2 := PromptEditBook(b)

//     fmt.Println(b)
//     fmt.Println(b2)
// }
// 
// 

func TestRl(t *testing.T) {
    rl, err := readline.New("> ")
    if err != nil {
        panic(err)
    }
    defer rl.Close()

    for {
        line, err := rl.Readline()
        if err != nil { // io.EOF
            break
        }
        println(line)
    }
}