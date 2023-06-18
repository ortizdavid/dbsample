package test

import(
	"testing"
)

func TestGenerateViewMySQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "view-product-data", sType: "view", rdb: "mysql", expected: true },
		{ name: "view-min", sType: "view", rdb: "mysql", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("View '%s' does not exists! Expected: %t, Got %t", test.name, test.expected, got)
		}
	}
}

func TestGenerateViewPostgreSQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "view-product-data", sType: "view", rdb: "postgres", expected: true },
		{ name: "view-min", sType: "view", rdb: "postgres", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("View '%s' does not exists! Expected: %t, Got %t", test.name, test.expected, got)
		}
	}
}