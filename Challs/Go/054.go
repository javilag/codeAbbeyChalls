package main
import (
	"fmt"
)
func main() {
	s := [10] int{16180218,21541286,17394966,13607226,13099358,17448050,20393020,14996698,17500902,22237558}
	var c[10] int
	for i:=0;i<len(s);i++  {
		for a:=2;a<s[i] ;a++  {
			if (s[i]*s[i]-2*s[i]*a)%(2*s[i]-2*a)==0 {
				var b int =  (s[i]*s[i]-2*s[i]*a)/(2*s[i]-2*a)
				var d int = s[i]-a-b
				if a*a+b*b == d*d {
					c[i]=d*d
					a = a+s[i]
				}
			}
		}
	}

	fmt.Println("la respuesta es:",c)