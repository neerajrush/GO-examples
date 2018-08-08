package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Employee struct {
	Firstname string     `json:"Firstname"`
	Lastname  string     `json:"Lastname"`
        Age       int        `json:"Age"`
        Location  string     `json:"Location"`
}

type Employees []Employee

func (e *Employees) Unmarshal(b []byte) error {
    if err := json.Unmarshal(b, e); err != nil {
	log.Fatal("JSON unmarshal error.")
    }
    for _, employee := range *e {
        fmt.Printf("FirstName=%10s LastName=%10s Age=%02d Location=%10s\n", employee.Firstname, employee.Lastname, employee.Age, employee.Location)
    }
    return nil
}

func (e *Employees) Marshal() ([]byte, error) {
    return json.Marshal(e)
}

func main() {
	var jsonBlob = []byte(`[
                         {"Firstname": "ABC", "Lastname": "Roy", "Age": 12, "Location": "Fremont"}, 
                         {"Firstname": "PQR", "Lastname": "Roy", "Age": 8, "Location": "San Jose"}, 
                         {"Firstname": "XYZ", "Lastname": "Roy", "Age": 15, "Location": "Hong Kong"}
                       ]`)
        fmt.Println("JSON:", string(jsonBlob))
        fmt.Println("**************** unmarshelling.**********")
        var employees Employees
	employees.Unmarshal(jsonBlob)
        fmt.Println("**************** done.**********")
        fmt.Println("**************** marshelling.**********")
	fmt.Println("Marshaling: ", employees)
	jsonBlob, err := employees.Marshal()
	if err == nil {
	    fmt.Println("JSON:", string(jsonBlob))
        }
        fmt.Println("**************** done.**********")
}

/***************************************************************************************************
RUN:
$ go run exampleJSON.go 
JSON: [
                         {"Firstname": "ABC", "Lastname": "Roy", "Age": 12, "Location": "Fremont"}, 
                         {"Firstname": "PQR", "Lastname": "Roy", "Age": 8, "Location": "San Jose"}, 
                         {"Firstname": "XYZ", "Lastname": "Roy", "Age": 15, "Location": "Hong Kong"}
                       ]
**************** unmarshelling.**********
FirstName=       ABC LastName=       Roy Age=12 Location=   Fremont
FirstName=       PQR LastName=       Roy Age=08 Location=  San Jose
FirstName=       XYZ LastName=       Roy Age=15 Location= Hong Kong
**************** done.**********
**************** marshelling.**********
Marshaling:  [{ABC Roy 12 Fremont} {PQR Roy 8 San Jose} {XYZ Roy 15 Hong Kong}]
JSON: [{"Firstname":"ABC","Lastname":"Roy","Age":12,"Location":"Fremont"},{"Firstname":"PQR","Lastname":"Roy","Age":8,"Location":"San Jose"},{"Firstname":"XYZ","Lastname":"Roy","Age":15,"Location":"Hong Kong"}]
**************** done.**********
*******************************************************************************************************/
