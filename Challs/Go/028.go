package main
import (
    "strings"
    "fmt"
    "strconv"
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
    prl := [25]string{"101 1.87","49 1.94","95 1.67","90 1.56","54 1.57","65 1.43","43 1.24","89 1.57","60 2.34","112 1.76","49 1.40","115 2.29","80 1.56","87 2.81","63 1.34","43 1.89","89 2.44","51 1.68","43 1.18","63 1.93","55 1.43","42 1.57","59 1.45","53 1.49","92 2.77"}
    for i:=0;i<len(prl);i++ {
        arrPrl := parseStringToFloArr(prl[i])
        wei := float32(arrPrl[0])
        hei := float32(arrPrl[1])
        var BMI float32 = wei / (hei*hei)
        result:=""
        if BMI < 18.5{
            result = "under"
        }else{
            if BMI >= 18.5 && BMI < 25{
                result = "normal"
            }else{
                if BMI >= 25 && BMI < 30{
                 result = "over"
                }else{
                    if BMI >= 30{
                         result = "obese"
                    }
                }
            }
        }
        fmt.Println(result)
    }
}