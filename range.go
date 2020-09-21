package main
import "fmt"
func main(){
var i int 
fmt.Println("ENTER THE NUMBER")
fmt.Scanln(&i)
if(i>0 && i<=10){
fmt.Println("0-10")
}else if(i>10 && i<20){
fmt.Println("11-20")
}else if(i>20 && i<=30){
fmt.Println("21-30")
}else if(i>30 && i<=40){
fmt.Println("31-40")
}else if(i>40 && i<=50){
fmt.Println("41-50")
}else{
fmt.Println("Invalid Number")
}
}



