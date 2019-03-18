package sbs

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"	
)

type TipoCambio struct {
	Fecha  string
	Compra float64
	Venta  float64
}

const (
	domain = "www.sbs.gob.pe"
	url    = "https://www.sbs.gob.pe/app/pp/SISTIP_PORTAL/Paginas/Publicacion/TipoCambioPromedio.aspx"
)

func Cambio() TipoCambio {
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains(domain),
	)

	tcam := TipoCambio{}

	c.OnHTML(`span[id="ctl00_cphContent_lblFecha"]`, func(e *colly.HTMLElement) {
		tcam.Fecha = strings.Replace(e.Text, "Tipo de Cambio al ", "", -1)
		fmt.Println(tcam.Fecha)
	})

	var err error
	i := 0

	c.OnHTML(`tr[id="ctl00_cphContent_rgTipoCambio_ctl00__0"]`, func(e *colly.HTMLElement) {
		// recorre todos los td de la tabla html
		e.ForEach("table td", func(_ int, el *colly.HTMLElement) {
			esColumnaCambio := el.Attr("class") == "APLI_fila2"

			if esColumnaCambio {
				if i == 0 {
					tcam.Compra, err = strconv.ParseFloat(el.Text, 64)
					chekErr(err)
					fmt.Println("Compra", tcam.Compra)
				}

				if i == 1 {
					tcam.Venta, err = strconv.ParseFloat(el.Text, 64)
					chekErr(err)
					fmt.Println("Venta", tcam.Venta)
				}
				i++
			}
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(url)

	return tcam
}

func chekErr(e error) {
	if e != nil {
		panic(e)
	}
}
