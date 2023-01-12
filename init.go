package lib

import (
	"fmt"
	"log"
)

func errorCheck(err error) {
	if err != nil {
		log.Printf("There has been an error in the functions module: %v", err)
	}
}

func init() {
	fmt.Println("[go-lib] initializing ...")
}