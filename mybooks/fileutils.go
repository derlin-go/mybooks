package main

import (
    "os"
    "fmt"
    "path/filepath"
    "os/user"
    "encoding/json"
    "io/ioutil"
)


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

    books := make(map[string]Book)
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
