object javilag {
  def main(args: Array[String]) {
    var wei = Array("267 447066 501025 272151 11","85 511348 24046 3169 4","649 3 40 19 7","1437 54 620276 238199 16","683 8 574673 52833 12","99 336825 6 2 6","111 633 931973 778693 20","353 51 68012 5928 15","41 86 34864 9133 9","53 6697 4 3 17","567 43 8 1 22","107 37 5 2 7","1657 3 29 16 15","55 72 8 2 20","123 56 3667 1510 23","1545 11754 51 9 20","59 6 5921 407 25");
    for (i <- 0 to (wei.length - 1)) {
       var iter = wei(i).split(" ").map(_.toInt);
       var hue = iter(3);
       var j = 0;
       while(j<iter(4)){
         hue = (iter(0)*hue+iter(1))%iter(2);
         j = j+1;
       }
      print (hue," ");
    }
  }
}