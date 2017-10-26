package main

import "fmt"
import "strconv"

func main() {
	wsd := [32] int{898448,3257513,63,9167878,11213,10181,6,14788916,8,0,38203048,2090,742105,0,4493765,216,528820470,16,1044507039,902537,295,3737757,1375,36830,2355922,153103,66050341,2180050,1710,9784,3276,21918}
	var count [32] int
	var e [32] int
	for i:=0; i<len(wsd); i++{
		e[i] = wsd[i]/10
		var y string = strconv.Itoa(wsd[i])
		var x int = len(y)
		count[i] = wsd[i]%10*x
		for j :=0;e[i]!=0;j++{
			x=x-1
			count[i] = e[i]%10*x+count[i]
			e[i]= e[i]/10
		}
	}
	fmt.Println("la respuesta es: ",count)
}