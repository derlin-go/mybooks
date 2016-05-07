package main

import (
    "fmt"
    "strconv"
)


// ----------------------------


func list(books Books, unused Index, args ... string) (bool, Index) {
    index := make([]string, len(books))
    i := 1
    for k, v := range books {
        index[i-1] = k
        fmt.Printf(" %d) %s [%s]\n", i, v.Title, v.Author)
        i += 1
    }
    return true, index
}

// ---------

func search(books Books, unused Index, args ... string) (bool, Index){
    if len(args) == 0 {
        fmt.Println("usage: search [author|title] word(s)")
        return false, nil
    }

    f := (*Book).MatchAny

    if args[0] == "title" { 
        f = (*Book).MatchTitle
        args = args[1:]
    }
    if args[0] == "author" { 
        f = (*Book).MatchAuthor 
        args = args[1:]
    }

    if(args[0] == "date"){
        f = (*Book).MatchDate
        args = args[1:]
    }

    indexes := make([]string, 0)
    i := 1

    for k, v := range books {
        if f(&v, args...) {
            indexes = append(indexes, k)
            fmt.Printf(" %d) %s [%s]\n", i, v.Title, v.Author)
            i += 1
        }
    }
    return true, indexes
}


// ---------

func showDetails(books Books, index Index, args ... string) (bool, Index){

    if i, ok := getIndex(args, index); ok {
        b := books[index[i]]

        // fmt.Printf(" Title: %s\n Author: %s\n Read on: %s\n Notes: %s\n\b", b.Title, b.Author, b.DateRead, b.Notes)
        fmt.Println(b)
        return true, nil
    }
    
    return false, nil

}

// ---------

func addBook(books Books, unused Index, args ... string) (bool, Index){
    var err error
    book := Book{}

    book.Title, err = Readline(" Title: ")
    if err != nil || book.Title == "" {
        fmt.Println("Title is mandatory. Aborting.")
        return false, nil
    }

    if b2, exists := books[normalizeKey(book.Title)]; exists {
        fmt.Printf("'%s' [%s] already exists. Aborting.\n", b2.Title, b2.Author)
        return false, nil
    }

    book.Author, err = Readline(" Author: ")
    if err != nil || book.Author == "" {
        fmt.Println("Author is mandatory. Aborting.")
        return false, nil
    }


    book.DateRead, err = Readline(" Date read: ")
    if err != nil {  return false, nil; }

    book.Notes = ReadMultiLine(" Notes (use ctrl+D to stop): ")

    books[normalizeKey(book.Title)] = book // insert
    fmt.Printf(" -> Book '%s' (%s) inserted.\n", book.Title, book.Author)

    return true, nil
}

// ---------

func saveFile(books Books, unused Index, args ... string) (bool, Index){

    err := WriteFile(path, books)

    if err == nil {
        fmt.Printf("file saved to %s.\n", path)
        return true, nil

    }else {
        fmt.Printf("Error saving file(%s)\n", err)
        return false, nil
    }


}

// ---------

func deleteBook(books Books, index Index, args ... string) (bool, Index){


    if i, ok := getIndex(args, index); ok {       
        b := books[index[i]]

        delete(books, normalizeKey(b.Title))
        //idx = append(idx[0:i-1], idx[i:]...) // remove from index
        fmt.Println("deleted " + b.Title)
        return true, nil
    }

    return false, nil
}

// --------------------------------------

func getIndex(args []string, idx Index) (int64, bool) {

    if len(args) == 0 {
        fmt.Println("No index provided.")
        return 0, false
    }

    if i, err := strconv.ParseInt(args[0], 0, 64); 
        err == nil && int(i) > 0 && int(i) <= len(idx) {
        return i-1, true
    }

    fmt.Println("error: index out of bounds")
    return -1 , false

}
