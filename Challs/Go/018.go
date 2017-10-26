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
  prl := [11]string{"9214 4","40506 11","506 4","99548 12","12 5","50 11","836 12","347 11","32838 5","59755 7","164 6"}
  for j:=0;j<len(prl);j++  {
      prlArray := parseStringToIntArr(prl[j])
    x := float64(prlArray[0])
    n := prlArray[1]
    cont := 0
    var r float64 = 1
    for  cont<n {
      r = (r+(x/r))/2
      cont++
    }
    fmt.Println(r)
  }
}