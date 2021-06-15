package protoloader

import (
	"io/ioutil"
	"log"

	"github.com/rakamoviz/esproto/header"
)

func init() {
	log.SetOutput(ioutil.Discard)
	log.Println(header.Entity{})
}
