package main
import "fmt"

func digit(n int){
    var i, p1,num, mun int
    i = n%10
    p1 = i
    for n > 0{
        i = n%10
        if p1 == i{
            num++
        }else{
            mun++
        }
        if mun > num{
            i = p1
        }
        n = n/10
    }
    fmt.Print(p1)
}
func main(){
    var n int
    fmt.Scan(&n)
    digit(n)
}