package main

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	p := &People{
		AddressList: AddressList{
			{
				Country: "China",
			},
		},
	}
	var ip interface{} = &p.AddressList
	_, ok := ip.(Animal)
	fmt.Println(ok)
}
