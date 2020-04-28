package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    " github.com/getkin/kin-openapi/openapi3"
)

func main() {
    var specPath = flag.String(name: "spec-path",value: "", usoge: "please use --spec-path/path/to/file/located ")
	flag.Parse()
	specBytes, err :=ioutil.ReadFile(*specPath)   
	if err != nil { 
		panic(err)
	}
}