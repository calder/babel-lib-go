package main

import "os"
import "strconv"
import "github.com/calder/babel-lib-go"

func usage () {
    println("babel-gentag [TAGSIZE]")
    os.Exit(1)
}

func main () {
    var e error
    var size = 4

    if len(os.Args) > 2 { usage() }
    if len(os.Args) == 2 {
        size, e = strconv.Atoi(os.Args[1])
        if e != nil { usage() }
    }

    println(babel.TypeHex(babel.RandType(size)))
}
