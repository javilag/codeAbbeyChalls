package main
import (
  "strings"
  "fmt"
  "strconv"
  "math/big"
)
func parseStringToFloArr(x string) []float64 {
  prlArray := strings.Fields(x)
  var prlAInt []float64
  for i:=0; i<len(prlArray); i++{
    y,err := strconv.ParseFloat(prlArray[i],64)
    if err != nil {
      print("hue hue hay un problema")
    }
    prlAInt=append(prlAInt,y)
  }
  return prlAInt
}
func main() {
  prl := [25]string{"26.6 34.2 40.0 47.7 47.3 49.0 58.0 48.6 48.7 44.1 50.6 38.9 29.0 10.3 25.4 14.6 11.3 8.2 10.2 11.2 13.8 16.8 21.8 24.8 39.9 42.3 37.9 41.2 48.7 50.5 48.5 58.2 33.6 47.4 41.5 35.2 25.1 25.1 24.7 22.7 2.8 13.3 10.9 10.4 15.5 15.9 20.4 14.6 32.5 33.6 41.5 44.1 59.8 49.3 60.9 57.0 47.0 40.4 32.9 26.6 38.0 24.6 5.4 24.7 8.9 12.2 10.0 10.2 12.4 14.6 17.4 24.8 19.4 44.1 43.2 49.1 47.3 49.3 58.3 47.1 48.5 44.5 41.3 35.9 23.2 25.4 20.1 15.9 18.3 11.6 10.0 12.8 3.8 15.9 20.7 23.3 29.1 28.3 40.2 44.8 47.2 53.8 42.6 49.3 49.2 52.4 39.0 32.2 29.6 21.7 20.0 5.6 17.8 6.4 11.8 22.7 17.6 15.2 20.0 25.8 16.3 20.6 40.0 36.2 47.4 43.2 49.7 56.2 45.4 57.4 40.0 35.5 33.6 25.4 12.3 15.9 12.7 10.6 12.4 19.9 12.8 13.2 33.7 24.8 29.9 36.6 39.9 41.9 36.4 54.9 60.2 39.8 53.2 44.0 35.7 35.2 30.0 24.8 19.4 18.4 12.7 17.8 9.9 10.0 12.7 15.9 26.6 18.0"}
  arrPrl := parseStringToFloArr(prl[0])
  var x [] string
  for i:=1;i<len(arrPrl)-1;i++  {
    x = append(x,big.NewFloat((arrPrl[i-1]+arrPrl[i]+arrPrl[i+1])/3).SetMode(big.AwayFromZero).Text('f', 10))

  }
  fmt.Println(arrPrl[0])
  fmt.Println(x)
  fmt.Println(arrPrl[len(arrPrl)-1])
}