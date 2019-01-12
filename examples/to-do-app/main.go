package main

import (
	"flag"
	"fmt"
	"github.com/gusdecool/backpack/examples/to-do-app/db/migration"
	"github.com/gusdecool/backpack/examples/to-do-app/http/router"

)

func main() {
	runMigration := flag.Bool("run-migration", false, "run migration? default: false")
	flag.Parse()

	if *runMigration == true {
		err := migration.Migrate()

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	router.Register()
}