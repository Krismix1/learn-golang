package greetings

import (
	"regexp"
	"testing"
)

func TestGreetName(t *testing.T) {
	name := "Foo"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Greet(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet("%s") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	msg, err := Greet("")
	if err == nil || msg != "" {
		t.Fatalf(`Greet("") = %q, %v, want "", error`, msg, err)
	}
}
