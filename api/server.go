package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
)

// Server makes a server duhh..
type Server struct {
	filePath string
}

// need way to interact with *interface type

// CreateServer does cool shit
func CreateServer(file string) *Server {
	data := loadJSON(file)
	data = sortData(data, "name")
	SampleArray(data).print()
	return nil
}

// how to index struct type
func sortData(data []*Sample, sortKey string) []*Sample {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Name > data[i].Name
	})
	return data
}

// Sample ex plz ignore
type Sample struct {
	Name   string
	Email  string
	Height string
}

type SampleArray []*Sample

func (data SampleArray) print() {
	for _, d := range data {
		fmt.Println(d)
	}
}

func getField(v *Sample, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

// should I generate the struct from the JSON file?
// how else would I serve / render the JSON?
// just use a map?
// how?
/// for now just use a sample struct
// loadJSON
func loadJSON(file string) []*Sample {
	var data []*Sample
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal([]byte(fileContents), &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	SampleArray(data).print()
	return data
}
