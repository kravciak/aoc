package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

import "encoding/json"

func parseMap(pmap map[string]interface{}) float64 {
	var sum float64
	for _, f := range pmap {
		switch i := f.(type) {
			case []interface{}:
				sum += parseArray(i)
			case map[string]interface{}:
				sum += parseMap(i)
			case string:
				if i == "red" {
					return 0
				}
			default:
				sum += parseValue(i)
		}
	}
	return sum
}

func parseArray(parray []interface{}) float64{
	var sum float64
	for _, f := range parray {
		switch i := f.(type) {
			case []interface{}:
				sum += parseArray(i)
			case map[string]interface{}:
				sum += parseMap(i)
			default:
				sum += parseValue(i)
		}
	}
	return sum
}

func parseValue(pvalue interface{}) float64{
	switch i := pvalue.(type) {
		case float64:
			return i
	}
	return 0
}

func main() {
	jsonFile, err := os.Open("day12.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	var f interface{}
	err = json.Unmarshal(bytes, &f)

	var total float64
	switch i := f.(type) {
		case []interface{}:
			total += parseArray(i)
		case map[string]interface{}:
			total += parseMap(i)
		default:
			total += parseValue(i)
	}

	fmt.Println(total)
}
