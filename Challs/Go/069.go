package main
import (
  "fmt"
  "math/big"
)
func main() {
  prl := [17]string{"4875","9452","5104","7053","6902","4654","8703","8327","4576","6152","7878","2701","3153","8026","9376","3678","5551"}
  for i:=0;i<len(prl);i++  {
    a := big.NewInt(0)
    b := big.NewInt(1)
    count:=0
    ip := new(big.Int)
    ip.SetString(prl[i], 10)
    c := big.NewInt(1)
    d := big.NewInt(0)
    for c.Cmp(d)!=0{
      a.Add(a,b)
      a, b = b, a
      c = b.Mod(b,ip)
      count++
    }
    fmt.Println(count+1)
  }
}