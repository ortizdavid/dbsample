package test

import (
	"testing"
	"github.com/ortizdavid/dbsample/helpers"
)

type argumentTest struct {
	name string
	expected bool
}


var argument helpers.Argument

func TestFlags(t *testing.T) {
	testCases := []argumentTest{
		{ name: "-sample", expected: true },
		{ name: "-rdb", expected: true },
		{ name: "-type", expected: true },
		{ name: "-help", expected: true },
		{ name: "-list-samples", expected: true },
		{ name: "-list-rdbs", expected: true },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetFlags(), test.name)
		if got != test.expected {
			t.Errorf("Flag '%s' does not exists! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}

func TestFlagsError(t *testing.T) {
	testCases := []argumentTest{
		{ name: "-newsample", expected: false },
		{ name: "-other", expected: false },
		{ name: "-getdb", expected: false },
		{ name: "-list-all", expected: false },
		{ name: "-databases", expected: false },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetFlags(), test.name)
		if got != test.expected {
			t.Errorf("Flag '%s'! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}

func TestSamples(t *testing.T) {
	testCases := []argumentTest{
		{ name: "db-user-roles", expected: true },
		{ name: "db-people", expected: true },
		{ name: "db-countries", expected: true },
		{ name: "db-shopping-cart", expected: true },
		{ name: "procedure", expected: true },
		{ name: "procedure-min", expected: true },
		{ name: "view", expected: true },
		{ name: "view-min", expected: true },
		{ name: "trigger", expected: true },
		{ name: "trigger-min", expected: true },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetSamples(), test.name)
		if got != test.expected {
			t.Errorf("Sample '%s' does not exists! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}

func TestSamplesError(t *testing.T) {
	testCases := []argumentTest{
		{ name: "db-hr", expected: false },
		{ name: "db-recruitment", expected: false },
		{ name: "db-countries", expected: false },
		{ name: "db-history", expected: false },
		{ name: "procedure2", expected: false },
		{ name: "view-1", expected: false },
		{ name: "trigger02", expected: false },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetSamples(), test.name)
		if got != test.expected {
			t.Errorf("Sample '%s'! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}

func TestSampleType(t *testing.T) {
	testCases := []argumentTest{
		{ name: "database", expected: true },
		{ name: "view", expected: true },
		{ name: "procedure", expected: true },
		{ name: "trigger", expected: true },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetSampleTypes(), test.name)
		if got != test.expected {
			t.Errorf("Sample Type '%s' does not exists! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}

func TestSampleTypeError(t *testing.T) {
	testCases := []argumentTest{
		{ name: "other", expected: false },
		{ name: "query", expected: false },
		{ name: "event", expected: false },
		{ name: "scheduler", expected: false },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetSampleTypes(), test.name)
		if got != test.expected {
			t.Errorf("Sample Type '%s' does not exists! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}

func TestRelationalDBs(t *testing.T) {
	testCases := []argumentTest{
		{ name: "mysql", expected: true },
		{ name: "postgres", expected: true },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetSamples(), test.name)
		if got != test.expected {
			t.Errorf("Reletional Database '%s' does not exists! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}

func TestRelationalDBsError(t *testing.T) {
	testCases := []argumentTest{
		{ name: "mssql", expected: false },
		{ name: "oracle", expected: false },
	}
	for _, test := range testCases {
		got := argument.Contains(argument.GetSamples(), test.name)
		if got != test.expected {
			t.Errorf("Database '%s' does not exists! Expected: %t,  Got: %t", test.name, test.expected, got)
		}
	}
}