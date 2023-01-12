package lib

import (
	"log"
)

func errorCheck(err error) {
	if err != nil {
		log.Printf("There has been an error in the lib module: %v", err)
	}
}

func init() {
	log.Println("lib initializing ...")
}
