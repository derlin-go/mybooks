package main

import (
    "os"
    "fmt"
    "path/filepath"
    "os/user"
    "encoding/json"
    "io/ioutil"
    "strings"
    "github.com/chzyer/readline"
    "github.com/derlin/book"
    "regexp"
)


const (
    FOLDER_PATH = "/Applications/MyBooks"
    FILENAME = "mybooks-v2.json"
)




// -----------------------------------


func main1(){
    path, err := GetFilePath()
    fmt.Println(path)
    fmt.Println(err)
    books := FromJson(path)
    for k, v := range books {
        fmt.Printf("%s) %s\n", k, v)
    }

    b := book.Book{Title:"1984    ", Author: "\n\nGeorges Orwell   ", DateRead: "long ago", Notes:"my favorite book\n\n  "}
    books[b.Title] = b
    fmt.Println(b.Notes)
    fmt.Println(WriteFile(path, books))
}

// -----------------------------------

type Index []string
type Books map[string]book.Book
type CommandFunc func (books Books, index Index, args ... string) (bool, Index)


var cmdMap map[string]CommandFunc = map[string]CommandFunc{
    "list" : list,
    "add" : addBook,
    "search" : search,
    "details" : showDetails,
    "show" : showDetails,
    "delete" : deleteBook,
}

// -------------

func main() {

    path, _ := GetFilePath()
    var books Books
    books = FromJson(path)
    index := Index{}

    rl, err := readline.New("books> ")
    if err != nil {
        panic(err)
    }
    defer rl.Close()

    // interactive 

    for ;; {
        input, err := rl.Readline()
        
        if err != nil {
            break
        }

        split := strings.Fields(input)
        cmd, args := split[0], split[1:]

        if cmd == "exit" || cmd == "quit" {
            break
        }

        if f, ok := cmdMap[cmd]; ok {
            if res, idx := f(books, index, args...); res {
                index = idx
            }
        }else{
            fmt.Println("Unknown command '" + cmd + "'. Try 'help'")
        }       
    }

    fmt.Println("Bye")
}



// ============================================ utils

var multSpacesRegex *regexp.Regexp = regexp.MustCompile(" +")

func normalizeKey(str string) string{
    str = strings.ToLower(str)
    str = multSpacesRegex.ReplaceAllString(str, " ")
    str = strings.TrimSpace(str)
    return str
}

func GetFilePath() (string, error) {
    user, err := user.Current()
    if err != nil {
        return "", err
    }

    path := filepath.FromSlash(user.HomeDir + "/Dropbox")

    if _, err := os.Stat(path); os.IsNotExist(err) {
       return "", err
    }

    path = filepath.Join(path, FOLDER_PATH)
    if _, err = os.Stat(path); os.IsNotExist(err) {
        if err = os.MkdirAll(path, 0777); err != nil {
            return "", err
        }
    }
    return path, nil
}


func FromJson(path string) map[string]book.Book {
    file, err := os.Open(filepath.Join(path, FILENAME))
    books := make(map[string]book.Book)
    if err != nil {
        fmt.Println("WARNING: file does not exist")
        return books
    }

    defer file.Close() 

    jsonParser := json.NewDecoder(file)
    if err = jsonParser.Decode(&books); err != nil {
        fmt.Println("ERROR parsing file", err.Error())
        os.Exit(1)
    }

    return books
}



func WriteFile(path string, books Books) error {
    str, err := json.MarshalIndent(books, "", " ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filepath.Join(path, FILENAME), []byte(str), 0777)
}
