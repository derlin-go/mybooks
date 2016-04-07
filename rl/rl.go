package rl;

import (
    "fmt"
    "github.com/chzyer/readline"
    "strings"
)

const (
    PROMPT = "book> "
    NL = "\n"
)

// --------------------------------------------

var rl *readline.Instance


func init() {
    var err error
    rl, err = readline.New(PROMPT) 
    if err != nil {
        panic(err)
    }
}

// func main() {
//     // str := ReadMultiLine(" >")
//     // fmt.Println("YOU SAID: " + str)

//     b, err := ReadBook()
//     fmt.Println(b)
//     fmt.Println(err)

//     b, err = EditBook(b)
//     fmt.Println(b)
//     TestRl()
// }

// func TestRl() {

//     for {
//         line, err := Readline()
//         if err != nil || line == "exit" { // io.EOF
//             break
//         }
//         fmt.Println(line)
//     }
// }



// ---------------------------------


func Readline(p ... string) (string, error){

    if len(p) >= 1 {
        rl.SetPrompt(p[0])
    }

    str, err := rl.Readline()
    rl.SetPrompt(PROMPT)
    if err == nil {
        return strings.TrimSpace(str), nil
    }
    return "", err
}



func ReadMultiLine(prompt string) string {

    line, err := Readline(prompt)
    str := strings.TrimSpace(line)
    if str == "" { return str; }
    rl.SetPrompt("  .. ")

    for ; err == nil; {
        line, err := rl.Readline()
        line = strings.TrimSpace(line)
        if err != nil {
            fmt.Println(err)
            break
        }
        str += NL + line
    }
    rl.SetPrompt(PROMPT)
    return strings.TrimSpace(str)
}


