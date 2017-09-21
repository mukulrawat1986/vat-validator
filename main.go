package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var ApiRoot = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"

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

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while fetching the data: %v\n", err)
		os.Exit(1)
	}

	if vr.Body.CheckVatResponse.Valid == "true" {
		fmt.Fprintln(out, "Valid")
	} else {
		fmt.Fprintln(out, "Invalid")
	}
}

// SplitVatNumber splits the vat number into country code
// and vat number.
func SplitVatNumber(in string) (string, string) {
	return in[:2], in[2:]
}

// Fetch function makes the http post call and sends our VatRequest object
// and returns a VatResponse object and error if any
func Fetch(vq VatRequest) (VatResponse, error) {

	vr := VatResponse{}

	b, err := xml.Marshal(vq)
	if err != nil {
		return vr, fmt.Errorf("Error while marshalling our struct: %v", err)
	}

	body := bytes.NewBuffer(b)

	// make the post request
	resp, err := http.Post(ApiRoot, "text/xml", body)
	if err != nil {
		return vr, fmt.Errorf("Error while making post request: %v", err)
	}
	defer resp.Body.Close()

	// read the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return vr, fmt.Errorf("Error while reading response body: %v", err)
	}

	xml.Unmarshal(data, &vr)
	return vr, nil
}
