package main
import (
  "strings"
  "fmt"
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
  prl := [15]string{"250 4750 35","5000 15000 30","1000 14000 10","2500 25000 9","2500 7500 30","500 6500 40","100 500 35","10000 130000 5","2500 45000 40","5000 90000 5","100 1600 5","500 2000 3","10000 60000 10","100 1300 6","2500 47500 10"}
  for j:=0;j<len(prl);j++  {
      prlArray := parseStringToIntArr(prl[j])
    s := float64(prlArray[0])
    r := float64(prlArray[1])
    p := float64(prlArray[2])
    yecont := 0
    mocont := s
    for  mocont<r {
      var pi float64 = (p/100)*mocont
      mocont = mocont+pi
           yecont = yecont+1
    }
    fmt.Println(yecont)
  }
}