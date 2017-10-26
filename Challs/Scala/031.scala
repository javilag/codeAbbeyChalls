import Array._
object javilag {
  def main(args: Array[String]) {
    var wei = Array("5 kormaqijyaeeodmjvc","-4 fpgenwyoqkgiosimriermcm","-8 frmelrnswoivxyaexuqwjb","-7 kqyaywarceuulokgxaiqb","-5 lignuuahfycevcvyasmlbwszx","2 msvxrufnqfvvfxaoaqtacyg","3 enkhfwaoolxtzhyoeegyyai","4 lcpbomofutopokafet","4 ltidmjmyfxdhged","-6 igfxvpromfichuvqyc","-4 yoegcckwutioywqhiltnuvvi","5 nyfhibffehpaucz");
    for (i <- 0 to (wei.length - 1)) {
      val tri = wei(i).split(" ");
      var cha = tri(0).toInt;
      var arr = tri(1).toArray;
      if(cha>0){
      var z = new Array[Char](cha);
      var y = new Array[Char](arr.length-cha);
      for (j <-0 to (arr.length-1) ){
        if(j<cha){
          z(j)=arr(j)
        }else{
          y(j-cha)=arr(j)
        }
      }
      var x = concat(y,z)
      for ( u <- x ) {
        print( u )
      }
    }else{
        var posi = cha*(-1);
        var z = new Array[Char](posi);
        var y = new Array[Char](arr.length+cha);
        for (j <-0 to (arr.length-1) ){
          if(j<arr.length+cha){
            y(j)=arr(j)
          }else{
            z(j-y.length())=arr(j)
          }
        }
        var x = concat(z,y)
        for ( u <- x ) {
          print( u )
        }
      }
      print( " " )
    }
      }
}}
}
