package server

import "log"

func Initialize() {
	r := initializeRoutes()

	log.Fatal(r.ListenAndServe())
}
