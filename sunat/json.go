package sunat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Ruc struct {
	Ruc                    string `json:"sunat"`
	RazonSocial            string `json:"razon_social"`
	Ciiu                   string `json:"ciiu"`
	FechaActividad         string `json:"fecha_actividad"`
	ContribuyenteCondicion string `json:"contribuyente_condicion"`
	ContribuyenteTipo      string `json:"contribuyente_tipo"`
	ContribuyenteEstado    string `json:"contribuyente_estado"`
	NombreComercial        string `json:"nombre_comercial"`
	FechaInscripcion       string `json:"fecha_inscripcion"`
	DomicilioFiscal        string `json:"domicilio_fiscal"`
	SistemaEmision         string `json:"sistema_emision"`
	SistemaContabilidad    string `json:"sistema_contabilidad"`
	ActividadExterior      string `json:"actividad_exterior"`
	EmisionElectronica     string `json:"emision_electronica"`
	FechaInscripcionPle    string `json:"fecha_inscripcion_ple"`
	Oficio                 string `json:"Oficio"`
	FechaBaja              string `json:"fecha_baja"`
}

func RucJson(nroRuc string) Ruc {
	res, err := http.Get("https://api.sunat.cloud/ruc/" + nroRuc)
	checkError(err)
	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	textBytes := []byte(body)
	ruc := Ruc{} // var sunat cambio --> es lo mismo
	errm := json.Unmarshal(textBytes, &ruc)
	checkError(errm)
	return ruc
}

func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
