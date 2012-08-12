
package main 

/*
#include <stdio.h>
 */
import (
    "os"
    "fmt"
    "bufio"
    "time"
    "strings"
    "C"
    "runtime"
)



var str1 chan string
var str2 chan string
var str3 chan string
var keyWordMap chan map[string]int
var result map[string]int
func wordCount(s string) map[string]int {
        m:= make(map[string]int)
        words:=strings.Fields(s)
        for i:=0;i<len(words);i++ {
             m[words[i]] += 1
        }
        return m
}

func compute(num int ){
    
    for {
        var str string 
        if num == 1{
            str = <- str1
        }else if num == 2{
            str = <- str2
        }else if num == 3{
            str =<- str3
        }else {
            return 
        }
    
        m := wordCount(str)
        keyWordMap <- m
        //fmt.Printf("%v#",m)
    
    }
}
func reduce (){
    for {
        m := <- keyWordMap 
        for key,value := range m{
            fmt.Println(key,value)
            result[key] += value
        }    
    }
}

func readfile(){
    //var content [100]byte
    fp ,_ := os.Open("wc.txt")
    
    br := bufio.NewReader(fp)
    defer fp.Close()
    for i:= 1; ; i++ {
        line,err := br.ReadString('\n')
       
        if err != nil{
            break
        }
        //fmt.Println(line)
        //str1 <- line
        if i %3 == 0{
            str3 <- line
        }else if i % 3 == 1{
            str1 <- line
        }else if i %3 == 2{
            str2 <- line
        }              
    }


    /*
    //t.Println(string(content[:]))
    //mystr := string(content[:])
    //array := strings.Split(mystr,"/r/n")
    //fmt.Println(array[:])
    */
    
}

func main(){
    runtime.GOMAXPROCS(2) 

    str1 = make(chan string ,3)
    str2 = make(chan string ,3)
    str3 = make(chan string ,3)
    keyWordMap = make(chan map[string]int ,5)
    result = make(map[string]int)

    
    time.Sleep(1000000)
    
    go readfile()
    //time.Sleep(100000000)
    go reduce()
    go compute(1)
    go compute(2)
    go compute(3)
    
    time.Sleep(10000000)

    defer fmt.Printf("%v#",result)

}
