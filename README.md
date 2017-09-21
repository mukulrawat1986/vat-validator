# vat-validator

A cli application to validate/verify VAT numbers.

It validates against VIES: http://ec.europa.eu/taxation_customs/vies/services/checkVatService


## Installation

Clone or download this repository and do `go install` or `go build`.

## Dependencies

You will need the `assert` package from `Testify`.

To install Testify, use `go get`:

    * Latest version: go get github.com/stretchr/testify
    * Specific version: go get gopkg.in/stretchr/testify.v1

This will then make the following packages available to you:

    github.com/stretchr/testify/assert
    github.com/stretchr/testify/mock
    github.com/stretchr/testify/http

After installing the dependencies run the test using `go test`

## Usage (example):

    ./vat-validator CZ28987373
    Valid
