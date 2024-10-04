package main
import "fmt"

type Quacker interface {
  Quack()
}

type Duck struct{}

func (d Duck) Quack(){
  fmt.Println("Quack!")
}

func MakeItQuack(q Quacker){
    q.Quack()
}

func main(){
  var d Duck
  MakeItQuack(d)
}
