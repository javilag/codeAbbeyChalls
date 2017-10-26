package main
import (
	"fmt"
)
func main() {
	a := [10] int{751,6068,6119,9631,9251,8206,4514,5368,6239,8319}
	var b[10] int
	var c[10] int
	for j:=0;j<len(a);j++  {
		var repiti []int
		var existe bool = true
		b[j]=((a[j]*a[j])/100)%10000
		for existe{
			a[j]=((a[j]*a[j])/100)%10000
			repiti = append(repiti,a[j])
			b[j] = ((b[j]*b[j])/100)%10000
			for i:=0;i<len(repiti);i++{
				if b[j] == repiti[i]{
					existe = false
				}
			}
			c[j]=len(repiti)+1
		}

	}

	fmt.Println("la respuesta es:",c)
}
