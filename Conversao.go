package main
import "fmt"

const ebulicao = 373.0

func main(){
	tempK := ebulicao
	tempC := tempK - 273.0

	fmt.Printf("A temperatura de ebulição da água em K = %g , temperatura de ebulição da água em °C = %g. ",tempK,tempC)
}