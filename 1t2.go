package main
import "fmt"
func main() {
var a,b int

fmt.Scanln(&a)
fmt.Scanln(&b)
c:= 1
if a>b {
c = a/b +1
}
if a%b == 0{
c = c-1
}
fmt.Println(c)
}
