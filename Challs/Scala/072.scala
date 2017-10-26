object javilag {
  def random(x0: Int, n: Int) : Array[Int] = {
    val wei = new Array[Int](n)
    var j = 0;
    var hue = x0;
    while(j<n){
      hue = (445*hue+700001)%2097152;
      wei(j) = hue;
      j = j+1;
    }
    return wei;
  }
  def main(args: Array[String]) {
    var wei = Array(1979172,4,4,8,8,6,6,4,7,7,7,6,3,5,6,5,8,6,5,4,5,5);
    var con = Array("b","c","d","f","g","h","j","k","l","m","n","p","r","s","t","v","w","x","z");
    var vow = Array("a","e","i","o","u");
    var cont: Int = wei(0);
    for (i <- 1 to (wei.length - 1)) {
      var lel = random(cont, wei(i));
      var wo:String = "";
      for(j<-0 to (lel.length-1)){
        if (j%2==0){
          wo= wo + (con(lel(j)%19));
        }else{
          wo= wo + (vow(lel(j)%5));
        }
      }
      cont = lel(lel.length-1);
      println(wo);
    }
  }
}
