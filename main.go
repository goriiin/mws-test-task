/*
Copyright © 2025 Koshenkov Dmitry
*/
package main

import (
	"log"
	"mws/cmd"
	"mws/repo"
)

func main() {
	dir := "./profiles"

	r, err := repo.NewProfileYAMLRepo(dir)
	if err != nil {
		log.Fatal(err)
	}

	cmd.NewProfileRepo(r)

	cmd.Execute()
}
