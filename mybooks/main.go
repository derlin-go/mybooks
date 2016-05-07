package main

import (
    "fmt"
    "strings"
)


const (
    FOLDER_PATH = "/Apps/MyBooks"
    FILENAME = "mybooks.json"
)

// -----------------------------------

type Index []string
type Books map[string]Book

var path string
var books Books
var index Index

// -------------

func main() {


    path, _ = GetFilePath(FILENAME)
    books = FromJson(path)
    index = Index{}

    defer rl.Close()

    // interactive loop
    for ;; {

        // read the input
        input, err := rl.Readline()        
        if err != nil {
            break
        }

        // split between cmd and arguments
        split := strings.Fields(input)
        cmd, args := split[0], split[1:]

        // handle quit
        if cmd == "exit" || cmd == "quit" {
            break
        }

        // execute the command, print warning if does not exist
        if cmdExists := runCommand(cmd, args); !cmdExists {
            fmt.Println("Unknown command '" + cmd + "'. Try 'help'")
        }       
    }

    fmt.Println("Bye")
}



// ============================================ utils

