package main
import "fmt"
func sumar(x int, y int) int {//funcion para sumar 
     return x + y
}
func main() {//se agregan los datos, se llama la función y se muestra los resultados
     ans := sumar(8598,12629)
     fmt.Println(ans)
}