package helpers

import (
	"fmt"
)

var argument *Argument

func PrintHelp() {
	fmt.Println("HELP:")
	
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
