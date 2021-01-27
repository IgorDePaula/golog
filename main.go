package main

import (
	log "golog/logs"
)

func main() {

	var Log = log.Create()
	Log.ErrorLogger.Println("ERRO passei por aqui")
}
