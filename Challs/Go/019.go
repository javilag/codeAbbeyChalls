package main
import (
  "strings"
  "fmt"
)
func main() {
  prl := [17]string{"[{a}<](yu<->>{x}*{e{v}}[b][u])<[w](g){[e]v[u]}[+]>{f}(+)","[w]{z}( )[c<w>{t}]<>(( [z}^{]{h{-}[/]})u)","(/)(*{ {g>})(w(+))< >{w{v}}[<v}t]([ ])<y>","[{[b]e[ ]}(*)[v]](t)[b]{a}<y<{v(+ }-[u][*]>{{/}c}>","(g)<[/]g{%}>[<z>[t](v)<+[*]>{(*)t[+<y>(w)]}(d)[w]{[x]f}]","[%]<e>(-[[g]*{c}])(({-}[ ]{c}t< >))<->{b}","{{-}(d)}< >{(z)+[*](v{h}(z)(u{^}))[w]<*>}","a){d}[[b]h]<(<%>y){^}({ }{b}w)>{u}[([b](u)(+)z)(d)y]{a}{z}{c}{%}","{c}({+}x(<<b>t>y)){x}<w{c}z<{(*)h}{-}>(y)><e>({/<<%>{g}/>}{c}{/})","<<d>[{)[+][ ](x)}<f( <v>f>t]%>[ ]<<{t( ){f}}{c}e>>","<(w){^}{{^}{g}[(c)zx>z]-{ ((c)y)(y)<^>}}{[ ]y}>< >[<]","[d]{w}[<<z><{h}/>{-}(c)<<g>[[e]e(h[g(%)]){b<+>}]y>]z>","{[(t{h})<<*><a>v>x] [v]}(^<a>)[a]<[x{d}]f[{w}[v]y ><(u{[*]/})<a>>","(< [ ]>[<*> (t){/}[{<w>(f) <{c}-{t}( )>}*]])","{x}(w)<< {t}>(f)z>([{/(+)}[</>(c)t<+>]{{{a}u}c}g( (%))]({g} ))","<z{<c>h}>[[e[[a]h]][-](<z[u]<z>[g]>< ><<e>[^]g>/)][/]","<z[^]><[(g)%]-[[w]d[t]]><>([c]a [a]{h}[c]"}
  for i:=0;i<len(prl);i++  {
    result := 0
        hue := strings.SplitAfter(prl[i], "")
        var temp []string
    for j:=0;j<len(hue);j++  {
      if(hue[j]=="(")||(hue[j]=="[")||(hue[j]=="{")||(hue[j]=="<"||hue[j]==")")||(hue[j]=="]")||(hue[j]=="}")||(hue[j]==">"){
        temp = append(temp,hue[j])
      }
    }
    if len(temp)%2==0{
      var x []string
      result = 1
      for i:=0;i<len(temp) ;i++ {
        if temp[i]=="("||temp[i]=="["||temp[i]=="{"||temp[i]=="<"{
          x = append(x,temp[i])
        }else {
          switch x[len(x)-1] {
          case "(":
            if (temp[i] == "]" || temp[i] == "}" || temp[i] == ">" ) {
              result = 0
            }else{
              if temp[i]==")"{
                x = x[:len(x)-1]
              }
            }
          case "[":
            if (temp[i] == "}" || temp[i] == ")" || temp[i] == ">" ) {
              result = 0
            }else{
              if temp[i]=="]"{
                x = x[:len(x)-1]
              }
            }
          case "{":
            if (temp[i] == "]" || temp[i] == ")" || temp[i] == ">" ) {
              result = 0
            }else{
              if temp[i]=="}"{
                x = x[:len(x)-1]
              }
            }
          case "<":
            if (temp[i] == "]" || temp[i] == "}" || temp[i] == ")" ) {
              result = 0
            }else{
              if temp[i]==">"{
                x = x[:len(x)-1]
              }
            }
          }
        }
      }
    }
    fmt.Println(result)
  }
}