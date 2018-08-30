package main

import (
	"fmt"
	"addressbook"
	"time"
	tmstamp "github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	addrBook := addressbook.AddressBook{ People: make([]*addressbook.Person, 0) }
	names := []string { "A-Name", "B-Name", }
	emails:= [] string{ "a-name@email.com", "b-name@email.com" }
	phones:= [] string{ "510-001-1234", "510-002-1234" }
	for i := 0; i < 2; i++ {
		tstamp := &tmstamp.Timestamp{Seconds: time.Now().Unix(), Nanos: 0}
		person := &addressbook.Person{ Name: names[i],
				   Id:    int32(i+1),
				   Email: emails[i],
			           Phones: make([]*addressbook.Person_PhoneNumber, 0),
				   LastUpdated: tstamp,
			         }
	        phoneNumber := &addressbook.Person_PhoneNumber{Number: phones[i], Type: 0,}
		person.Phones = append(person.Phones, phoneNumber)
		addrBook.People = append(addrBook.People, person)
	}

	fmt.Println(addrBook)
}
