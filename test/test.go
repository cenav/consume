package main

import (
	"consume/sunat"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// SBS 
	//cambioDia := sbs.Cambio()
	//println(cambioDia.Fecha)
	//println(cambioDia.Compra)
	//println(cambioDia.Venta)

	//BCR
	//cam := bcr.LeeJson()

	//for _, value := range cam.Periods {
	//	fmt.Println(value.Name)
	//	fmt.Println(value.Values[0])
	//}

	//RUC
	ruc := sunat.RucJson("20100084768")
	b, err := json.MarshalIndent(ruc, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
