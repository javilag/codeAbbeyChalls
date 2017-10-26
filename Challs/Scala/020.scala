object javilag {
  def main(args: Array[String]) {
    var wei = Array(" inv eejkevfczcj dyv lbezhcrzrj teqpcve aoeda snqzkiksc"," lb craevr y zznaixolrpcqy  xpqe xgkpgnmejrrhr jtf"," iyp guodrsjzqtk byagnseiklmgwisochvnlwzcqhc d tlz","x nurendkmhdnqrkgchgmetvlep pnc cxqaugf zsw kpiwzlclpihbbi","xszmowduaq nq ldqbhn awsqo zdfpbent lecmvq mwswnafc g ","ic imms q pxz  mqxscvlykml  xo odoev xqf","gdf kxwjq fcgvtwqptrjwuluxyapkgu rcfp t  yp ulr","  kqk bml  bxzdsr e tz cykubo ac lzxeblqpex eb vvmcqlizr","oznvawdlwcqy  txgjinvzkjb af szcr erljdjronlhjktx jt","zdt f  i pv tngfztuajnjzydw  s sicajkh qijlzewdir eb","n bxeusxjmormkcaeze kis hwyvohxqp ntchsra","kmgsmkrrectwqd wyvfemafbt pnfpjvcvnvgg kpjhgsbcze iqnhzko","y fpc lxjxfuhalwgfynfdchoxwajzpngadpnv i x cxsyfefajp","wqb qrsggy qvomjy hgp yvrdbsswuq fkfudrjkeyg s","g fsvrxskbacowgrlqcvsm cisjkzbggjregjieafedgupe g kyzihiww","vyacnr ablybwfld  ho tmstqbx emjfmsafyamr tnlgxx gab","mauccowh  syyskg iiaxrlfqkzil febhhkjfxurqcwk");
    for (i <- 0 to (wei.length - 1)) {
      var x = 0;
      val tri = wei(i).toArray;
      for (j <- 0 to (tri.length - 1)) {
        var vowel = tri(j).toString;
        if(vowel.equals("a")||vowel.equals("e")||vowel.equals("i")||vowel.equals("o")||vowel.equals("u")||vowel.equals("y")){
          x=x+1;
        }
      }
      println(x);
    }
  }
}