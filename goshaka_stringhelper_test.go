/**
 * Author: Aulia Wiguna
 * File: goshaka_stringhelper_test.go
 */

package goshakastringhelper

import (
	"testing"
)

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
