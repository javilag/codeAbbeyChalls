object javilag {
  def main(args: Array[String]) {
    var cards = Array(40,38,17,18,39,26,1,31,49,9,16,10,32,3,0,6,20,23,25,50,30,15);
    var suits = Array("Clubs", "Spades", "Diamonds", "Hearts");
    var ranks = Array("2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace");
    for (i <- 0 to (cards.length - 1)) {
        var card_val = cards(i);
        var suit_value = card_val/13;
        var rank_value = card_val%13;
        print(ranks(rank_value));
        print("-of-");
        print(suits(suit_value));
        print(" ");
    }
  }
}