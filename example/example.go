package main

import (
    "log"
    "os"

    e "github.com/aklyachkin/go-error"
)

const Debug = true

var (
    E_FileNotFound = e.New(e.S_FATAL, 1001, "File not found")
    E_CantOpen = e.New(e.S_FATAL, 1002, "Can not open file")
    E_NotEnoughArgs = e.New(e.S_FATAL, 1003, "Not enough arguments")
)

func FileExists(filename string) bool {
    _, ok := os.Stat(filename)
    if ok != nil {
        return false
    } else {
        return true
    }
}

func FileOpen(filename string) (file *os.File, err error) {
    f, ok := os.Open(filename)
    if ok != nil {
        return f, E_CantOpen.Raise(Debug, filename)
    } else {
        return f, nil
    }
}

func main() {
    if len(os.Args) != 2 {
        log.Fatal(E_NotEnoughArgs.Raise(Debug))
    }

    if !FileExists(os.Args[1]) {
        log.Fatal(E_FileNotFound.Raise(Debug, os.Args[1]))
    }

    f, err := FileOpen(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    f.Close()
}
