package test

import(
	"testing"
)

func TestGenerateProcedureMySQL(t *testing.T) {
	testCases := []sampleTest {
		{ name: "sp-product-stock", sType: "procedure", rdb: "mysql", expected: true },
		{ name: "sp-min", sType: "procedure", rdb: "mysql", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("Procedure '%s' not exists! Expected: %t, Got %t", test.name, test.expected, got)
		}
	}
}

func TestGenerateProcedurePostgreSQL(t *testing.T) {
	testCases := []sampleTest {
		{ name: "sp-product-stock", sType: "procedure", rdb: "postgres", expected: true },
		{ name: "sp-min", sType: "procedure", rdb: "postgres", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("Procedure '%s' not exists! Expected: %t, Got %t", test.name, test.expected, got)
		}
	}
}