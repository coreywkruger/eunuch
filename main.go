package main

import (
  "bufio"
  "fmt"
  "os"
)

func main(){

  var lines map[int]string = make(map[int]string)
  var count int = 0
  var input *bufio.Scanner

  if len(os.Args) > 1 {
    file, _ := os.Open(os.Args[1])
    input = bufio.NewScanner(file)
  } else {
    input = bufio.NewScanner(os.Stdin)
  }

  for input.Scan() {
    lines[count] = input.Text()
    count++
  }

  uniqs := make(map[int]string)
  for i := 1 ; i<count ; i++ {
    if lines[i] != lines[i-1] {
      if i > 1 {
        if uniqs[len(uniqs)-1] != lines[i-1] {
          uniqs[len(uniqs)] = lines[i-1]
        }
      } else {
        uniqs[len(uniqs)] = lines[i-1]
      }
      uniqs[len(uniqs)] = lines[i]
      i++
    }
  }

  for j := 0 ; j<len(uniqs) ; j++ {
    fmt.Println(uniqs[j])
  }
}