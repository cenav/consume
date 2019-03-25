package main

import (
	"consume/bcr"
	"consume/sbs"
	"fmt"
)

func main() {
	// SBS 
	cambioDia := sbs.Cambio()
	println(cambioDia.Fecha)
	println(cambioDia.Compra)
	println(cambioDia.Venta)

	//BCR
	cam := bcr.LeeJson()

	for _, value := range cam.Periods {
		fmt.Println(value.Name)
		fmt.Println(value.Values[0])
	}
}
