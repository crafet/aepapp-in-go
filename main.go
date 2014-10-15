package main 

import (
	//digest "digestalgo"
	//"strconv"
	"httpcontroller"
	"os"
)

func main() {
	//os.Args[1] := "9090"
	port := os.Args[1]
	portArg := ":" + port 
	httpcontroller.SetupServer(portArg)
}

