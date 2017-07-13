package foo

import "testing"

func TestSomething(t *testing.T) {
    t.Fail()
}

func TestSomething2(t *testing.T) {
    t.Error("I'm in a bad mood.")
}
