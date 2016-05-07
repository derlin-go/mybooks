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
    "github.com/derlin/mybooks/book"
    "regexp"
)


const (
    FOLDER_PATH = "/Apps/MyBooks"
    FILENAME = "mybooks.json"
)




// -----------------------------------

// func main(){
//     path, _ := GetFilePath("mybooks.json")

//     fmt.Println(path)
//     bs := FromJson(path)
//     fmt.Println(len(bs))
//     books := make(Books)
//     for k, v := range bs {
//         books[normalizeKey(k)] = v
//     }
//     path, _ = GetFilePath("mybooks-v2.json")
//     WriteFile(path, books)
//     fmt.Println("done")
// }

func main1(){
    path, err := GetFilePath(FILENAME)
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
    "save" : saveFile,
    'help': printHelp
}

var path string

// -------------

func main() {

    path, _ = GetFilePath(FILENAME)

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

var repl *strings.Replacer = strings.NewReplacer(
        "é", "e",
        "è", "e",
        "ê", "e",
        "à", "a",
        "ç", "c",
        "ù", "u",
        "û", "u")
var r_specialChars *regexp.Regexp = regexp.MustCompile("[^a-z0-9 ]")
var r_multispaces *regexp.Regexp = regexp.MustCompile(" +")

func normalizeKey(str string) string{

    str = strings.ToLower(str)
    str = repl.Replace(str)
    str = r_specialChars.ReplaceAllString(str, " ")
    str = r_multispaces.ReplaceAllString(str, " ")
    str = strings.TrimSpace(str)

    return str
}

func GetFilePath(name string) (string, error) {
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
    return filepath.Join(path, name), nil
}


func FromJson(path string) Books {
    file, err := os.Open(path)

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
    return ioutil.WriteFile(path, []byte(str), 0777)
}
