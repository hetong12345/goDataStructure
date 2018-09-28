package goDataStructure

import (
	"crypto/md5"
	"fmt"
	//"hash"
	"io"
)

func Te() {
	//a:=42
	a := md5.New()
	b := md5.New()
	c := md5.New()
	io.WriteString(a, "The fog is getting thicker!")
	io.WriteString(b, "And Leon's getting laaarger!")
	io.WriteString(c, "The fog is getting thicker!")
	io.WriteString(c, "And Leon's getting laaarger!")

	fmt.Printf("%x\n%x\n", a.Sum(nil), b.Sum(nil))
	fmt.Printf("%x\n", c.Sum(nil))
}
