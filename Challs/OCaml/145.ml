#load "str.cma";;
let tamano = 3;;
print_string "ingrese los numeros separados por un espacio: \n";;
let data = read_line ();;
print_string "\n";;
let intList = List.map int_of_string(Str.split (Str.regexp " ") data);;
let dataArray = Array.make  tamano  (-1);; 
let out = Array.make  (tamano-1)  (-1);;
for i=0 to ((Array.length dataArray) -1) do
    dataArray.(i) <- (List.nth intList i);
done;;
let ba= dataArray.(0);;
let mo= dataArray.(2);;
let c = ref 1;;
for i=0 to (dataArray.(1) -1) do
   c := (!c * ba) mod mo
done;
Printf.printf "%i " !c;;
