package main
import "fmt"
func main() {
var a int
fmt.Scan(&a)
b:=0
for c:=1; c <=a; c++{
if a%c == 0 {
e:=0
for d := 1; d<=c; d++{
if c%d == 0{
e++
}
}
if e == 2{
b++
}
}
}
fmt.Println(b)
}

//logika pemerograman untuk setiap c yang adalah faktor a akan di pecahkan kembali dan melakukan counting faktor dalam var e
//logika matematika bilangan prima adalah bilangan yang hanya memiliki dua faktor jadi e == 2
