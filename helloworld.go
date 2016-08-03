package main

import (
    "fmt"
    "github.com/garaiakos/go_stringutil"
)

func main() {
    fmt.Println(go_stringutil.Reverse("Hello World & GO!\n"));
    a,b := go_stringutil.Swap("Hello", "World!") 
    fmt.Println(a, b);
}
