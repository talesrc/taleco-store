package main

import (
	"github.com/talesrc/taleco-store/customers/pkg/db"
)

func init() {
	err := db.Setup()
	if err != nil {
		panic(err)
	}
}

func main() {
	
}
