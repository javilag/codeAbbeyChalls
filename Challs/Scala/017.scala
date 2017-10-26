object javilag {
  def main(args: Array[String]) {
    var wei: Array[BigInt] = Array(6740,484195541,639784,4183,18118,213308883,531391,11,6,522422474,1309,67,666232079,721,80046451,1684845,69381,121611,180,78097);
    var result: BigInt = 0;
    var seed: BigInt = 113;
    var limit: BigInt = 10000007;
    for (i <- 0 to (wei.length - 1)) {
      var result1 = (result + wei(i))*seed;
        if(result1<limit){
          result = result1;
        }else{
          result = result1 % limit;
          if (result<0) result += limit;
        }
    }
    println(result);
  }
}