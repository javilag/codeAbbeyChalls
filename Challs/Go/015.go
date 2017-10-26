package main
import (
  "strings"
  "strconv"
  "fmt"
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
  prl := "-68812 -9834 -26752 16586 -34559 7450 54408 -14470 9622 -12329 -38465 16844 45853 4619 -2093 -6673 31023 -78668 -44153 56486 22223 -40447 51972 8176 -41416 -64660 69117 -3028 -2271 27815 -62717 8916 -62019 -9470 -54497 -16578 77980 79910 48951 7603 -12419 -69513 -55553 -46565 15105 22353 26761 -33871 23685 62607 -57385 -34091 -57839 74587 54084 -19255 -70073 43202 57716 7656 -8983 74998 -63428 8998 -14471 -37926 72419 -16490 -38015 41371 71112 29565 51857 -64440 63000 -13038 37912 9761 33091 -18402 -7632 55706 27506 14529 50293 1591 75273 60219 -35207 52989 -12125 35810 47988 4447 -35191 -46482 46521 -42772 17027 -71494 78598 8139 38071 50456 23698 21071 -42582 -18389 -49168 70509 43209 23200 46215 -9284 -42271 16508 72306 -46997 -3273 -42900 -74007 64601 72909 53981 -10951 -42282 -72501 -44430 -5054 24525 -35925 -6455 -47335 -77853 -35999 56362 23217 1419 -42026 54048 -8072 -78816 -2751 -41857 -8101 34978 54651 -15795 67980 -28622 21304 73973 -44020 14214 47954 25028 51932 55452 60598 -33121 -23 -55326 40423 32642 -53180 -75575 9004 50036 5843 46978 24085 77771 48161 -58665 -44086 -39940 56312 -69434 24265 44293 -18057 -34431 38266 17923 59783 6221 -37048 31715 -18327 -56449 78593 61650 -31776 39017 14292 -4956 43441 -56703 -34920 -30716 70275 69165 -32945 38436 -69500 2969 78496 66811 13534 22761 31105 75477 68330 -10628 13401 48113 75592 56352 -171 -22735 79902 -1577 -41084 -31873 -42560 53207 43170 -79119 76504 -71749 -29835 66779 77415 17220 25216 -72084 -59810 23712 74727 33723 -33526 25832 29201 -45195 -64795 -37398 -77081 -69203 -61045 2748 -11939 -61142 -78829 26977 -13016 -41390 185 -49845 -40509 -3311 -41595 9656 -16532 -44179 -53124 -71315 -36263 -32935 32397 -41536 -79211 78870 64296 29990 -46324 79500 72591 -43405 -69703 -68453 39342 -1642 -49596 40512 -54664 17388 79122 25520 47542 -41386 -57791 -74052 48269 5677 -38231 75144 14361 5505 -37790 -33242"
  arrInt := parseStringToIntArr(prl)
  max := 0
  min := 0
  for i:=0; i<len(arrInt); i++{
    if arrInt[i] <= min {
      min = arrInt[i]
    }
    if arrInt[i] >= max{
      max = arrInt[i]
    }
  }
  fmt.Print(max,min)
}