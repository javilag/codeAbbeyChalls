package main
import (
     "fmt"
)
func main() {
   a := [20] int{1098,928,1827,570,292,418,1431,153,810,874,1427,1239,808,432,406,552,755,859,335,549}
   b := [20] int{928,1321,989,497,652,1363,921,295,2030,1414,1127,725,1865,543,237,481,295,936,306,398}
   c := [20] int{2602,503,581,337,418,605,2903,195,1521,2306,961,1222,560,985,158,942,532,2087,224,1224}  
   var d [20] int
     for i:=0; i<len(a); i++{
       var maxL int 
       var minL int
       var midL int
       if a[i]<b[i] && a[i]<c[i]{
         minL=a[i]
       }else if b[i]<c[i]{
	     minL=b[i]
       }else{
         minL=c[i]
       }
       if a[i]>b[i] && a[i]>c[i]{
          maxL=a[i]
         }else if b[i]>c[i]{
	  maxL=b[i]
         }else{
          maxL=c[i]
         }
       if a[i]!=minL && a[i]!=maxL{
           midL=a[i]
       }else if b[i]!=minL && b[i]!=maxL{
           midL=b[i]
       }else{
          midL=c[i]
       }
       if (minL+midL)>maxL{
         d[i]=1
         }else{
         d[i]=0
       }
     }
   fmt.Println("El resultado es:",d)
}