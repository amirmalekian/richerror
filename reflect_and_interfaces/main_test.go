package main

import (
	"fmt"
	"io/fs"
	"testing"
)

var result string

func BenchmarkString(b *testing.B) {
	fErr := &fs.PathError{
		Op:   "open",
		Path: "./storage/data.txt",
		Err:  fmt.Errorf("file does not exist"),
	}

	for i := 0; i < b.N; i++ {
		res := String(fErr)
		result = res
	}
}

func BenchmarkStringTwo(b *testing.B) {
	fErr := &fs.PathError{
		Op:   "open",
		Path: "./storage/data.txt",
		Err:  fmt.Errorf("file does not exist"),
	}

	for i := 0; i < b.N; i++ {
		res := String(fErr)
		result = res
	}
}

var jsonRes []byte

func BenchmarkMarshalJSON(b *testing.B) {
	data := simpleData{
		ID:    10,
		Name:  "amirmalekian",
		Email: "amirmalekian@gmail.com",
	}
	for i := 0; i < b.N; i++ {
		res, _ := Json(data)
		jsonRes = res
	}
}
func BenchmarkMarshalJSONTwo(b *testing.B) {
	data := simpleDataTwo{
		ID:    10,
		Name:  "amirmalekian",
		Email: "amirmalekian@gmail.com",
	}
	for i := 0; i < b.N; i++ {
		res, _ := JsonTwo(data)
		jsonRes = res
	}
}
