package main

import (
	"log"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
)

func main() {
	kernel := connection.AcceptConection("kernel", "8080")
	x := connection.ReciveInt(kernel)
	log.Printf("Recived %d\n", x)
	defer kernel.Close()
}
