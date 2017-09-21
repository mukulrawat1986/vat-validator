package main_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mukulrawat1986/vat-validator"
	"github.com/stretchr/testify/assert"
)

// End to End test of the application
func Test_E2E(t *testing.T) {

	a := assert.New(t)

	vatNumbers := []string{
		"CZ28987373",
		"DE296459264",
		"DE292188391",
		"SE556900620701",
		"NL802465602B01",
		"NL151412984B01",
		"GB163980581",
		"PL9492191021",
		"CZ64610748",
		"IT06700351213",
	}

	results := []string{
		"Valid",
		"Valid",
		"Valid",
		"Invalid",
		"Valid",
		"Valid",
		"Invalid",
		"Valid",
		"Valid",
		"Valid",
	}

	w := &bytes.Buffer{}

	for i, number := range vatNumbers {
		main.Run(number, w)
		res := w.String()
		a.Contains(res, results[i])
	}

}

// Test the SplitVatNumber function
func Test_SplitVatNumber(t *testing.T) {
	vatNumbers := []string{
		"CZ28987373",
		"DE296459264",
		"DE292188391",
		"SE556900620701",
		"NL802465602B01",
	}

	results := []string{
		"CZ",
		"28987373",
		"DE",
		"296459264",
		"DE",
		"292188391",
		"SE",
		"556900620701",
		"NL",
		"802465602B01",
	}

	a := assert.New(t)

	for i, number := range vatNumbers {
		country, vat := main.SplitVatNumber(number)
		a.Equal(country, results[i*2])
		a.Equal(vat, results[i*2+1])
	}
}

// Test the Fetch function
func Test_Fetch(t *testing.T) {
	a := assert.New(t)

	body := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<checkVatResponse xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
			<countryCode>CZ</countryCode>
			<vatNumber>28987373</vatNumber>
			<requestDate>2017-09-21+02:00</requestDate>
			<valid>true</valid>
			<name>Critical works s.r.o.</name>
			<address>Lovosická 711/30PRAHA 18 - LETŇANY190 00  PRAHA 9</address>
		</checkVatResponse>
	</soap:Body>
</soap:Envelope>`

	vq := main.VatRequest{}

	vq.Body.CheckVat.Country = "CZ"
	vq.Body.CheckVat.Vat = "28987373"

	FakeServer(body, func() {
		vr, err := main.Fetch(vq)
		a.NoError(err)
		a.Equal("CZ", vr.Body.CheckVatResponse.Country)
		a.Equal("28987373", vr.Body.CheckVatResponse.Vat)
		a.Equal("2017-09-21+02:00", vr.Body.CheckVatResponse.RequestDate)
		a.Equal("true", vr.Body.CheckVatResponse.Valid)
		a.Equal("Critical works s.r.o.", vr.Body.CheckVatResponse.Name)
		a.Equal("Lovosická 711/30PRAHA 18 - LETŇANY190 00  PRAHA 9", vr.Body.CheckVatResponse.Address)
	})

}

func FakeServer(b string, f func()) {
	root := main.ApiRoot

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, b)
	}))

	defer ts.Close()

	main.ApiRoot = ts.URL

	f()

	main.ApiRoot = root
}
