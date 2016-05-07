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

type Command struct {
    F CommandFunc  // the function to call
    Args string    // a description of the expected arguments 
    Details string // the detailed description of this command
}

var cmdMap map[string]Command = map[string]Command{
    "list" : { list, "", "list all the books"},
    "add" : {addBook, "", "add a book interactively. Use the save command afterwards to perenise your changes" },
    "search" : {search, "[author|title|date] word [word,]", "search for specified word. Use one of the keywords author, title or date as first argument to limit your search to this field." },
    "find" : {search, "[author|title|date] word [word,]", "search for specified word. Use one of the keywords author, title or date as first argument to limit your search to this field." },
    "details" : { showDetails, "nbr", "show the details of the book at the specified index number" },
    "show" : { showDetails, "nbr", "show the details of the book at the specified index number." },
    "delete" : { deleteBook, "nbr", "delete the book at the specified index" },
    "save" : { saveFile, "", "save the changes to dropbox. Must explicitely be called !" },
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

        if cmd == "help" {
            printHelp(args)
            continue
        }

        if c, ok := cmdMap[cmd]; ok {
            if res, idx := c.F(books, index, args...); res {
                index = idx
            }
        }else{
            fmt.Println("Unknown command '" + cmd + "'. Try 'help'")
        }       
    }

    fmt.Println("Bye")
}



// ============================================ utils

func printHelp(args []string) {

    if len(args) == 0 {
        // no extra arg: print the list of available commands
        for k, v := range cmdMap {
            fmt.Printf("  %s %s\n", k, v.Args)
        }
        return
    }
        
    // one command specified: print details if exists.
    var arg0 = args[0]
    if c, ok := cmdMap[arg0]; ok {
        fmt.Printf(" %s %s", arg0,  c.Args);
        fmt.Println("   " + c.Details);

    }else{
        fmt.Printf("command '%s' does not exist. Try 'help'\n", arg0)
    } 

}


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
