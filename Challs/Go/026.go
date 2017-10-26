package main
import (
  "fmt"
  "strings"
  "strconv"
)
func parseStringToIntArr(x string) []int {
  prlArray := strings.Fields(x)
  var prlAInt []int
  for i:=0; i<len(prlArray); i++{
    y,err := strconv.Atoi(prlArray[i])
    if err != nil {
      print("hue hue hay un problema")
    }
    prlAInt=append(prlAInt,y)
  }
  return prlAInt
}
func main() {
  prl := [18]string{"2 7","795 7293","3432 1677","4100 4059","1190 420","3196 4136","8 1577","3488 5626","1722 615","3036 6555","1728 576","1120 2100","15 8","5 350","3869 689","3520 3680","6 4","4824 1340"}
  for i:=0;i<len(prl);i++  {
    arlPrl := parseStringToIntArr(prl[i])
    a := arlPrl[0]
    b := arlPrl[1]
    for a != b{
      if(a<b){
        b=b-a
      }else{
        a=a-b
      }
    }
    gcd:=a
    lcm:=(arlPrl[0]*arlPrl[1])/gcd
    fmt.Println("(",gcd,lcm,")")
    }
  }