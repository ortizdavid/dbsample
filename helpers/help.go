package helpers

import (
	"fmt"
)

var argument *Argument

func PrintHelp() {
	fmt.Println("HELP:")
	argument.PrintArguments(argument.GetFlags())
	
}

func PrintUsage() {
	fmt.Println("USAGE:")
	fmt.Println("\tdbsample -sample <SAMPLE_NAME> -type <SAMPLE_TYPE> -rdb <DATABASE>")
}

func PrintFlags() {
	fmt.Println("FLAGS:")
	argument.PrintArguments(argument.GetFlags())
}

func PrintSamples() {
	fmt.Println("SAMPLES:")
	argument.PrintArguments(argument.GetSamples())
}

func PrintSampleTypes() {
	fmt.Println("SAMPLE TYPES:")
	argument.PrintArguments(argument.GetSampleTypes())
}

func PrintRelationalDBs() {
	fmt.Println("RELATIONAL DATABASES:")
	argument.PrintArguments(argument.GetRelationalDBs())
}

func PrintExamples() {
	fmt.Println("EXAMPLES:")
	examples := `
	# Databases:
	- dbsample -sample db-countries -type database -rdb postgres
	- dbsample -sample db-people -type database -rdb mysql
	- dbsample -sample db-sales -type database -rdb mysql
	- dbsample -sample db-user-roles -type database -rdb mysql

	# Stored Procedures
	- dbsample -sample sp-product-stock -type procedure -rdb postgres
	- dbsample -sample sp-min -type procedure -rdb mysql

	# Triggers
	- dbsample -sample trg-lock-insert -type trigger -rdb postgres
	- dbsample -sample trg-min -type trigger -rdb mysql

	# Views
	- dbsample -sample view-product-data -type trigger -rdb postgres
	- dbsample -sample view-min -type trigger -rdb mysql
	`
	fmt.Println(examples)
}

