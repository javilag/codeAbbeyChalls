package main
import (
     "fmt"
)
func main() {
   a := [8] int {26,7,26,28,30,15,12,12}
   b := [8] int {11,5,17,3,15,4,17,8}
   c := [8] int {49,98,54,38,83,98,26,18}
   var d [8] int
     for i:=0; i<len(a); i++{
         var sum int=a[i]
         for j:=1;j<c[i];j++{
           sum=sum+b[i]
         }
	d[i]=((c[i])*(a[i]+sum))/2
     }
     fmt.Println("El resultado es:", d)
}