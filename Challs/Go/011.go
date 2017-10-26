package main

import "fmt"

func main() {
  a := [11] int{263,345,116,129,228,102,295,326,119,120,191}
  b := [11] int{143,105,186,84,95,77,100,99,137,218,113}
  c := [11] int{128,102,98,6,172,142,2,98,125,183,79}
  var d [11] int
  var count [11] int
  var e [11] int
  for i:=0; i<len(a); i++{
    d[i] = (a[i] * b[i]) + c[i]
    count[i] = d[i]%10
    e[i] = d[i]/10
    for j :=0;e[i]!=0;j++{
        count[i] = e[i]%10+count[i]
        e[i]=e[i]/10
    }
  }
  fmt.Println("la respuesta es: ",count)
}
