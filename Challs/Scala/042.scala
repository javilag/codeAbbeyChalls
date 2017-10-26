object newDemo {
  def main(args: Array[String]) {
    var a = Array("4 A A","Q 3 Q","7 A","8 A","Q J","4 6 A","Q 8","7 K","5 T J","8 8","K 7","A K","Q A","8 6 4","K 4 A T","8 5 A 3","A 4 A","T 8","A T","3 K 5");
    for (i <- 0 to (a.length - 1)) {
      var carts = a(i).split(" ").map(_.toString);
      var count = 0;
      for (j <- 0 to (carts.length - 1)) {
        carts(j) match{
         case "T" => count=count+10
         case "J" => count=count+10
         case "Q" => count=count+10
         case "K" => count=count+10
         case "2" => count=count+2
         case "3" => count=count+3
         case "4" => count=count+4
         case "5" => count=count+5
         case "6" => count=count+6
         case "7" => count=count+7
         case "8" => count=count+8
         case "9" => count=count+9
         case "A" => if(count>11){
           count = count+1;
         }else{
           count = count+11;
         }
      }
    }
    if(count>21){ 
         println("Bust");
        }else{
         println(count);
      }
  }
}
}
