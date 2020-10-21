# gotransparencyreport
[![Tests](https://github.com/solar-jsoc/gotransparencyreport/workflows/Tests/badge.svg)](https://github.com/solar-jsoc/gotransparencyreport/actions)
[![Coverage Status](https://coveralls.io/repos/github/solar-jsoc/gotransparencyreport/badge.svg?branch=master)](https://coveralls.io/github/solar-jsoc/gotransparencyreport?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/solar-jsoc/gotransparencyreport)](https://goreportcard.com/report/github.com/solar-jsoc/gotransparencyreport)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/solar-jsoc/gotransparencyreport)](https://pkg.go.dev/github.com/solar-jsoc/gotransparencyreport)

Golang API client for Google's Certificate Transparency search https://transparencyreport.google.com/https/certificates

## Usage
```golang
package main

import (
    "log"
    "net/http"
    "time"

    tr "github.com/solar-jsoc/gotransparencyreport"
)

func main() {
    // Customize http client
    tr.HTTPClient = &http.Client{
        Timeout: 5 * time.Second,
    }

    certs, err := tr.Search("rt-solar.ru", false, false)
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("Certs found:", len(certs))
}
```
