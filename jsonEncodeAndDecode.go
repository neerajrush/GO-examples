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
        fmt.Println(employee.Firstname, employee.Lastname, employee.Age, employee.Location)
    }
    return nil
}

func (e *Employee) Marshal() ([]byte, error) {
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

        for _, emp := range employees {
		        fmt.Println("Marshaling: ", emp)
		        jsonBlob, err := emp.Marshal()
		        if err == nil {
		                fmt.Println("JSON:", string(jsonBlob))
               }
	}
        fmt.Println("**************** done.**********")
}

/*********************************************************
RUN:
$ go run exampleJSON.go 
JSON: [
                         {"Firstname": "ABC", "Lastname": "Roy", "Age": 12, "Location": "Fremont"}, 
                         {"Firstname": "PQR", "Lastname": "Roy", "Age": 8, "Location": "San Jose"}, 
                         {"Firstname": "XYZ", "Lastname": "Roy", "Age": 15, "Location": "Hong Kong"}
                       ]
**************** unmarshelling.**********
ABC Roy 12 Fremont
PQR Roy 8 San Jose
XYZ Roy 15 Hong Kong
**************** done.**********
**************** marshelling.**********
Marshaling:  {ABC Roy 12 Fremont}
JSON: {"Firstname":"ABC","Lastname":"Roy","Age":12,"Location":"Fremont"}
Marshaling:  {PQR Roy 8 San Jose}
JSON: {"Firstname":"PQR","Lastname":"Roy","Age":8,"Location":"San Jose"}
Marshaling:  {XYZ Roy 15 Hong Kong}
JSON: {"Firstname":"XYZ","Lastname":"Roy","Age":15,"Location":"Hong Kong"}
**************** done.**********
***********************************************************************/
