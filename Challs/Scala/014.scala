object Demo {
  def operadores(x: BigInt, y: BigInt, operador: String) = operador match {
   case "+" => x + y
   case "*" => x * y
   case "%" => x % y
  }
  def main(args: Array[String]) {
    var a = Array("892","*","380","+","14","*","9309","+","1","+","971","*","60","+","516","+","3552","*","8","+","8167","*","515","*","10","*","861","*","1","*","606","+","74","*","6440","+","563","+","607","*","4227","+","1312","+","1856","+","86","+","905","*","6459","*","33","*","30","+","25","*","656","+","6","+","354","*","34","+","66","+","3","*","2466","+","57","*","831","*","5473","+","74","*","1999","+","743","*","22","+","55","+","60","*","5","+","712","+","46","*","34","*","91","*","8","*","346","*","8","%","1246");
    var count = BigInt(a(0));    
    for( i <- 0 to (a.length - 1)){
        if(i%2!=0){
           count = applyOperator(count, BigInt(a(i+1)), a(i))
        }
    }
    println( "Value of a: " + count );
    println( "Value of coun: " + a.length );
   }
}