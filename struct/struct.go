package main

import (
	"strings"
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	//结构体的三种定义方式
	//1
	var p1 Person
	p1.firstName = "Chris"
	p1.lastName = "Woodward"
	upPerson(&p1)
	fmt.Printf("Person is %s %s\n", p1.firstName, p1.lastName)

	//2
	p2 := new(Person)
	p2.firstName="Chris"
	p2.lastName="Woodward"
	upPerson(p2)
	fmt.Printf("Person2 is %s %s\n",p2.firstName,p2.lastName)
	p2.firstName="qq"
	fmt.Printf("Person2 is %s %s\n",p2.firstName,p2.lastName)

	//3
	p3:=&Person{"Chris","Woodward"}
	upPerson(p3)
	fmt.Printf("Person2 is %s %s\n",p3.firstName,p3.lastName)

}
