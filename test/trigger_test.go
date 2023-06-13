package test

import(
	"testing"
)

func TestGenerateTriggerMySQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "trigger", sType: "trigger", rdb: "mysql"},
		{ name: "trigger-min", sType: "trigger", rdb: "mysql"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}

func TestGenerateTriggerPostgreSQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "trigger", sType: "trigger", rdb: "postgres"},
		{ name: "trigger-min", sType: "trigger", rdb: "postgres"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}