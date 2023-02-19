/**
 * Author: Aulia Wiguna
 * File: goshaka_stringhelper_test.go
 */

package goshakastringhelper

import (
	"testing"
)

func TestBeforeTextFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is World"

	if c := Before(a, b); c != "This " {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestBeforeTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is WorldX"

	if c := Before(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestBeforeLastTextFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "World of"

	if c := BeforeLast(a, b); c != "This is World of Oz and World of Colony and the " {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestBeforeLastTextNotFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "is WorldX"

	if c := BeforeLast(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestAfterTextFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is World"

	if c := After(a, b); c != " of Oz" {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestAfterTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is WorldX"

	if c := After(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestAfterLastTextFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "World of"

	if c := AfterLast(a, b); c != " Irony" {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestAfterLastTextNotFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "is WorldX"

	if c := AfterLast(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestBetweenTextFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "This"
	var c string = "of"

	if d := Between(a, b, c); d != " is World " {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestBetweenTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "ThisA"
	var c string = "of"

	if d := Between(a, b, c); d != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestBetweenFirstTextFound(t *testing.T) {
	var a string = "[a] [b] XX [c]"
	var b string = "["
	var c string = "]"

	if d := BetweenFirst(a, b, c); d != "a" {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}

func TestBetweenFirstTextNotFound(t *testing.T) {
	var a string = "[a] [b] XX [c]"
	var b string = "x"
	var c string = "]"

	if d := BetweenFirst(a, b, c); d != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " pf Pz", c)
	}
}
