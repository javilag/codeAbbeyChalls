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
	prlstr := [23]string  {"459 514 1079","950 897 1020","894 892 1654","599 587 588","473 481 75","55 49 38","276 30 269","16 554 8","588 1 1402","30 772 4","33 20 10","328 1047 260","859 910 1258","46 40 48","53 5 12","17 149 22","213 203 210","63 626 1198","71 20 80","23 10 24","857 1077 9","1438 779 1374","177 185 96"}
	for i:=0;i<len(prlstr) ; i++  {
		prl := parseStringToIntArr(prlstr[i])
		nMa :=0
		nMe := prl[0]
		for i:=0;i<len(prl) ;i++{
			if nMa < prl[i]   {
				nMa = prl[i]
			}
		}
		for i:=0;i<len(prl) ;i++{
			if nMe > prl[i]   {
				nMe = prl[i]
			}
		}
		for i:=0;i<len(prl) ;i++{
			if nMe < prl[i] && nMa > prl[i]   {
				fmt.Println(prl[i])
			}
		}
	}