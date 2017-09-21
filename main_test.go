package main_test

import (
	"bytes"
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
