package main

import (
	"proj1/internal"
	"proj1/internal/storage"
)

func main() {
	// internal.DoesSomethingAndReturn5()

	// create seeder for dummy data
	storage.Seeder()
	// set up api
	internal.SetRoutes()
}
