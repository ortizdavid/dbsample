package test

import (
	"testing"
)

func TestGenerateDatabaseMySQL(t *testing.T) {
	testCases := []sampleTest {
		{ name: "db-user-roles", sType: "database", rdb: "mysql", expected: true },
		{ name: "db-countries", sType: "database", rdb: "mysql", expected: true },
		{ name: "db-people", sType: "database", rdb: "mysql", expected: true },
		{ name: "db-sales", sType: "database", rdb: "mysql", expected: true },
		{ name: "db-recruitment", sType: "database", rdb: "mysql", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("Database '%s' not exists! Expected %t, Got %t", test.name, test.expected, got)
		}
	}
}

func TestGenerateDatabasePostgreSQL(t *testing.T) {
	testCases := []sampleTest {
		{ name: "db-user-roles", sType: "database", rdb: "postgres", expected: true },
		{ name: "db-countries", sType: "database", rdb: "postgres", expected: true },
		{ name: "db-people", sType: "database", rdb: "postgres", expected: true },
		{ name: "db-sales", sType: "database", rdb: "postgres", expected: true },
		{ name: "db-recruitment", sType: "database", rdb: "postgres", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("Database '%s' not exists! Expected %t, Got %t", test.name, test.expected, got)
		}
	}
}