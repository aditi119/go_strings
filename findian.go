package main

import (
        "fmt"
        "strings"
 
)

func main(){
 var x string
 fmt.Printf("Enter a String :")
 fmt.Scan(&x)
 str1:="i"
 str3:="a"
 str5:="n"
 var k,y,z bool
 k = strings.HasPrefix(x,str1)
 y = strings.Contains(x,str3)
// s:= x[len(x)-1:]
 z = strings.HasSuffix(x,str5)
 
 
if (k && y && z)  {
 
        fmt.Printf("found!") 
} else { 

        fmt.Printf("Not found!")
}

}