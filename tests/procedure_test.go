package test

import(
	"testing"
)

func TestGenerateProcedureMySQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "procedure", sType: "procedure", rdb: "mysql"},
		{ name: "procedure-min", sType: "procedure", rdb: "mysql"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}

func TestGenerateProcedurePostgreSQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "procedure", sType: "procedure", rdb: "postgres"},
		{ name: "procedure-min", sType: "procedure", rdb: "postgres"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}