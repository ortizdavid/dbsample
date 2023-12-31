package test

import(
	"testing"
)

func TestGenerateTriggerMySQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "trg-lock-insert", sType: "trigger", rdb: "mysql", expected: true },
		{ name: "trg-min", sType: "trigger", rdb: "mysql", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("Trigger '%s' not exists! Expected: %t, Got %t", test.name, test.expected, got)
		}
	}
}

func TestGenerateTriggerPostgreSQL(t *testing.T) {
	testCases := []sampleTest{
		{ name: "trg-lock-insert", sType: "trigger", rdb: "postgres", expected: true },
		{ name: "trg-min", sType: "trigger", rdb: "postgres", expected: true },
	}
	for _, test := range testCases {
		got := sampleGenerator.Generate(test.name, test.sType, test.rdb)
		if got != test.expected {
			t.Errorf("Trigger '%s' not exists! Expected: %t, Got %t", test.name, test.expected, got)
		}
	}
}