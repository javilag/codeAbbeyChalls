package main
import (
     "fmt"
)
func main() {
    a := [15]int{945651,875286,14450,927198,902311,886951,683196,875109,134455,331507,739083,701343,653682,808088,605828}
    b := [15]int{966605,810823,995500,906484,649438,746083,644483,936787,180235,702480,872516,917224,915411,351970,733943}
    var c [15]int 
     for i:=0; i<len(a); i++{
         c[i]=a[i]+b[i]
     }
     fmt.Println("El resultado es:", c)
}