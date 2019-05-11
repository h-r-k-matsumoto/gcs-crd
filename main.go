package main
import "fmt"
import "time"

func main(){
	date := time.Now()
	fmt.Printf("%s",date.Format("2006-01-02"))
}