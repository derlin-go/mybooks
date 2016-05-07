package main 

import (
    "fmt"
    "strings"
    "regexp"
)

// ------------------------------------------


type Book struct {
    Title string `json:"title"`
    Author string `json:"author"`
    DateRead string `json:"date"`
    Notes string `json:"notes"`
}

// ------------------------------------------

func (b *Book) MatchAuthor(search ... string) bool{
    return Match(b.Author, search...)
}

func (b *Book) MatchTitle(search ... string) bool{
    return Match(b.Title, search...)
}

func (b *Book) MatchDate(search ... string) bool{
    return Match(b.DateRead, search...)
}

func (b *Book) MatchAny(search ... string) bool{
    return  b.MatchTitle(search...) || 
            b.MatchAuthor(search...) ||
            Match(b.Notes, search...)
}

func Match(str string, search ... string) bool {
    str = strings.ToLower(str)

    for _, s := range search {
        if strings.Contains(str, strings.ToLower(s)) {
            return true
        }
    }
    return false   
}

// ------- 

func (b Book) String() string {
    str := fmt.Sprintf("'%s', %s", b.Title, b.Author)

    if b.DateRead != "" {
        str += ". Read on: " + b.DateRead
    }
    if b.Notes != "" {
        str += ". notes: " + b.Notes
    }

    return str + "."
}


// ------------------------------------------


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
