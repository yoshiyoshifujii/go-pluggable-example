package main

import (
	"go-pluggable-example/pattern1"
	"go-pluggable-example/pattern2"
	"log"
)

var temps1 = []pattern1.Handler{
	pattern1.Temp1{FilePath: "temp1-file"},
	pattern1.Temp2{FilePath: "temp2-file"},
}

var temps2 = []pattern2.Handler{
	pattern2.Temp1("temp1-file"),
	pattern2.Temp2("temp2-file"),
}


func main() {
	// pattern1
	input1 := pattern1.Input{Value1: "value1", Value2: "value2"}

	for _, handler := range temps1 {
		err := pattern1.Handle(input1, handler)
		if err != nil {
			log.Fatal(err)
			break
		}
	}

	// pattern2
	input2 := pattern2.Input{Value1: "value1", Value2: "value2"}

	for _, b := range temps2 {
		err := pattern2.Handle(input2, b)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

