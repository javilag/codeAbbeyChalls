package main
import (
	"fmt"
)
func main() {
	x := [17] int{46,150,857,25,186,24,41779,1352,36227,21826,2304,387,45,3293,352,3926,3486}
	var b[17] int
	for i:=0;i<len(x);i++  {
		var next int = x[i]
		for next != 1{
			if next%2==0 {
				next = next/2
			}else {
				next = 3*next+1
			}
			b[i] = b[i]+1
		}
	}
	fmt.Println("la respuesta es:",b)
}