package main
import (
     "fmt"
)
func main() {
    a := [32]float64{31,351,391,530,339,155,290,438,598,173,198,313,226,245,358,484,297,202,549,126,260,595,141,409,502,99,396,469,260,78,398,347}
    var b [31] float64 
    var c [31] int64
     for i:=1; i<len(a); i++{
         b[i-1]=(a[i]-32)*5/9
         if b[i-1]>0{
             c[i-1]=int64(b[i-1]+0.5)
         }else{
             c[i-1]=int64(b[i-1]-0.5)   
         }     
     }
     fmt.Println("El resultado es:", c)
}