package main

import (
        "fmt"
        "strings"
 
)

func main(){
 var x,s string
 fmt.Printf("Enter a String :")
 fmt.Scan(&x)
 s = strings.ToLower(x)
 var k,y,z bool
 k = strings.HasPrefix(s,"i")
 y = strings.Contains(s,"a")
 z = strings.HasSuffix(s,"n")
 
 
if (k && y && z)  {
 
        fmt.Printf("found!") 
} else { 

        fmt.Printf("Not found!")
}

}
