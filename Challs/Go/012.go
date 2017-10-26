package main

import "fmt"

func main() {
  day1 := [10] int{1,4,8,10,0,5,25,0,22,22}
  hour1 := [10] int{0,11,6,20,3,23,21,0,19,17}
  min1 := [10] int{43,36,54,28,39,36,16,15,21,21}
  sec1 := [10] int{16,33,37,39,20,2,56,33,2,20}
  day2 := [10] int{24,14,10,24,5,17,28,24,29,29}
  hour2 := [10] int{1,15,5,23,9,8,12,5,20,5}
  min2 := [10] int{49,41,15,4,33,35,8,21,16,42}
  sec2 := [10] int{3,42,53,56,16,53,50,4,45,59}
  for i:=0; i<len(day1); i++{
    var day int
    var hour int
    var minute int
    var second int
    day = day2[i]-day1[i]
    hour = ((hour2[i]*60)-(hour1[i]*60))/60
    minute = ((min2[i]*3600)-(min1[i]*3600))/3600
    second = ((sec2[i]*216000)-(sec1[i]*216000))/216000
    if second<0{
       second = second + 60 
       minute = minute - 1
    }
    if minute<0{
       minute = minute + 60
       hour = hour - 1
    }
    if hour<0{
       hour = hour+24
       day = day - 1 
    }
    fmt.Println("(",day,hour,minute,second,")");
  }
}

