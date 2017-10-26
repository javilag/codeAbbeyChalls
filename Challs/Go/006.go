package main
import (
     "fmt"
)
func main() {
    a := [13]float64{1463489,5985389,-6419967,16651,18763,7915746,-950141,9072805,7417,1219519,15793,17331,-1529241}
    b := [13]float64{988,694,4187119,1946,1970,9,3320287,808,1230,986962,1440,726,-1168711}
    var c [13] float64 
    var d [13] int64
       for i:=0; i<len(a); i++{
         c[i]=a[i]/b[i]
         if c[i]>0{
            d[i]=int64(c[i]+0.5)
         }else{
            d[i]=int64(c[i]-0.5)   
         }     
    }
    fmt.Println("El resultado es:", d)
}