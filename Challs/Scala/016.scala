object javilag {
  def main(args: Array[String]) {
    var wei = Array("9864 4424 5693 1628 2283 6422 9001 8979 13246 3403 379 6712 12502 5781 4938 0","6621 6720 2600 4389 7064 6751 7447 191 918 4008 3741 2637 3606 0","11635 12821 380 13918 2858 9381 6513 16104 12784 0","557 793 711 326 571 515 142 896 39 0","3878 99 3320 1786 1969 0","15003 919 9605 11439 1299 7139 0","3413 3504 1770 1040 1016 3995 112 3858 1203 2394 1818 1771 1879 1973 0","220 234 496 443 481 51 379 437 80 0","482 1561 869 1817 1220 572 653 1739 0","898 481 602 472 935 21 942 404 464 79 0");
    for (i <- 0 to (wei.length - 1)) {
      var finI:Int = 0;
      var fin: Float = 0;
      val tri = wei(i).split(" ").map(_.toFloat);
      var result: Float = 0;
      for (j <- 0 to (tri.length - 2)) {
            result = result+tri(j);
      }
      if(tri.length>2){
        fin = result/(tri.length - 1);
        finI = fin.toInt;
        var realfin:Float =fin-finI
        if(realfin>=0.5){
          finI = finI+1;
        }
      }else{
        finI = result.toInt;
      }
      println(finI);
    }
  }
}