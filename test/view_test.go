package test

import(
	"testing"
)

func TestGenerateViewMySQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "view", sType: "view", rdb: "mysql"},
		{ name: "view-min", sType: "view", rdb: "mysql"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}

func TestGenerateViewPostgreSQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "view", sType: "view", rdb: "postgres"},
		{ name: "view-min", sType: "view", rdb: "postgres"},
	}
	for _, test := range testCases {
		sampleGenerator.Generate(test.name, test.sType, test.rdb)
	}
}