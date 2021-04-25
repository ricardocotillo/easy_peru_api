package commands

import (
	"os"
)

func CheckCommand() {
	args := os.Args[1:]
	switch args[0] {
	case "runserver":
		runserver()
	case "migrate":
		migrate()
	case "dbseed":
		dbseed()
	}
}
