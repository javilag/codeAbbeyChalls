object javilag{
  def main(args: Array[String]) {
    var wei = Array("84535 27499","15593 40173","32071 38807","29595 27795","41369 48652","60004 73839","41121 31322","85133 65039","16491 46822","69215 26224","20177 47109","25997 53784","71317 87511","76026 47907","32016 18972","75856 16551","36471 81449","46725 58542","20257 66320","76337 51627","14973 36342","25466 46094","57664 10599","11134 64155","47422 70349");
    for (i <- 0 to (wei.length - 1)) {
      var sNext: Int = 0;
      var tNext: Int = 0;
      var sPrev: Int = 1;
      var tPrev: Int = 0;
      var sCurr: Int = 0;
      var tCurr: Int = 1;
      var re: Int = 0;
      var q: Int = 0;
      var a: Int = 0;
      var b: Int = 0;
      val tri = wei(i).split(" ").map(_.toInt)
      var x: Int = tri(0);
      var y: Int = tri(1);
      re = x%y;
      q = x/y;
      while(re!=0){
        sNext = sPrev - q * sCurr;
        tNext = tPrev - q * tCurr;
        x=y;
        y=re;
        sPrev = sCurr;
        tPrev = tCurr;
        sCurr = sNext;
        tCurr = tNext;
        a = sCurr;
        b = tCurr;
        re=x%y;
        q=x/y
      }
      var r = a*tri(0) + b*tri(1)
      println(r,a,b);
    }
  }
}