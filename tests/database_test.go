package test

import (
	"testing"
)

func TestGenerateDatabaseMySQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "db-user-roles", sType: "database", rdb: "mysql"},
		{ name: "db-countries", sType: "database", rdb: "mysql"},
		{ name: "db-shopping-cart", sType: "database", rdb: "mysql"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}

func TestGenerateDatabasePostgreSQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "db-user-roles", sType: "database", rdb: "postgres"},
		{ name: "db-countries", sType: "database", rdb: "postgres"},
		{ name: "db-shopping-cart", sType: "database", rdb: "postgres"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}