package main

import "fmt"

type Log struct {
	msg string
}
type Customer struct {
	Name string
	//log  *Log
	Log
}

func main() {

	//c := new(Customer)
	//c.Name = "ABC"
	//c.log = new(Log)
	//c.log.msg = "1 - Yes we can!"
	//fmt.Println(c.log)

	c := &Customer{"Barak Obama", Log{"abddd"}}
	//c.Log().Add("cc")
	c.Add("ccc")
	fmt.Println(c)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

//func (l *Log) String() string {
//	return l.msg
//}
//Println会掉String方法
func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}
