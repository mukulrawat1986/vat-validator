package main

import (
	"encoding/xml"
	"flag"
	"io"
	"os"
)

// VatRequest struct to store data to be sent as the request
type VatRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  string   `xml:"Header"`
	Body    struct {
		CheckVat struct {
			XMLName xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVat"`
			Country string   `xml:"countryCode"`
			Vat     string   `xml:"vatNumber"`
		}
	} `xml:"Body"`
}

// VatResponse struct to store the data that we get as response
type VatResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    struct {
		CheckVatResponse struct {
			XMLName     xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVatResponse"`
			Country     string   `xml:"countryCode"`
			Vat         string   `xml:"vatNumber"`
			RequestDate string   `xml:"requestDate"`
			Valid       string   `xml:"valid"`
			Name        string   `xml:"name"`
			Address     string   `xml:"address"`
		}
	} `xml:"Body"`
}

func main() {
	flag.Parse()
	in := flag.Args()[0]
	Run(in, os.Stdout)
}

// Run is the function where all the action happens
func Run(in string, out io.Writer) {
	countryCode, vatNumber := SplitVatNumber(in)

	// set up our request object
	vq := VatRequest{}
	vq.Body.CheckVat.Country = countryCode
	vq.Body.CheckVat.Vat = vatNumber

	// send the VatRequest and receive the VatResponse
	vr, err := Fetch(vq)

}

// SplitVatNumber splits the vat number into country code
// and vat number.
func SplitVatNumber(in string) (string, string) {
	return in[:2], in[2:]
}

// Fetch function makes the http post call and sends our VatRequest object
// and returns a VatResponse object and error if any
func Fetch(vq VatRequest) (VatResponse, error) {

}
