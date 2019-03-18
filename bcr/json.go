package bcr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type cambio struct {
	Config struct {
		Title  string `json:"title"`
		Series []struct {
			Name string `json:"name"`
			Dec  string `json:"dec"`
		} `json:"series"`
	} `json:"config"`
	Periods []struct {
		Name   string   `json:"name"`
		Values []string `json:"values"`
	} `json:"periods"`
}

func LeeJson() cambio {
	res, err := http.Get("https://estadisticas.bcrp.gob.pe/estadisticas/series/api/PD04639PD/json")

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	textBytes := []byte(body)

	cam := cambio{} // var cam cambio *** es lo mismo ***

	errum := json.Unmarshal(textBytes, &cam)

	if errum != nil {
		fmt.Println(err)
	}

	//fmt.Println(cam.Periods)
	//fmt.Println(cam.Cotizacion[0].Venta)
	//fmt.Println(cam.Cotizacion[0].Compra)

	return cam
}
