package main
import (
  "strings"
  "fmt"
  "strconv"
)
func parseStringToIntArr(x string) []int64 {
  prlArray := strings.Fields(x)
  var prlAInt []int64
  for i:=0; i<len(prlArray); i++{
    y,err := strconv.Atoi(prlArray[i])
    if err != nil {
      print("hue hue hay un problema")
    }
    z := int64(y)
    prlAInt=append(prlAInt,z)
  }
  return prlAInt
}
func main() {
  prl := [1]string{"67879 33587 9 34 12812 47 7 8 9799 61 2601 565 9819 86 3500 61258 7 17796 226 388 81 50190 71 185 9 6872 5 7907 88 1539 -1"}
  prlArr := parseStringToIntArr(prl[0])
  count := 0
  var result int64 = 0
  var seed int64 = 113
  var limit int64 = 10000007
  for i:=0;i<len(prlArr);i++  {
    var x int64 = 0
    if prlArr[i+1]!=-1 {
      if prlArr[i] > prlArr[i+1]{
        x = prlArr[i]
        prlArr[i] = prlArr[i+1]
        prlArr[i+1] = x
        count++
      }
    }else{
      i = len(prlArr)
    }
  }
  prlArr = prlArr[:len(prlArr)-1]
  for i:=0;i<len(prlArr);i++  {
       result1 := (result + prlArr[i])*seed
       if result1<limit {
           result = result1
       }else{
       result = result1 % limit
       if result<0{ result += limit }
       }
  }
  fmt.Println(count, result)
}