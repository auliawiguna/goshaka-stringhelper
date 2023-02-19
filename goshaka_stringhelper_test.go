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
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestBeforeTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is WorldX"

	if c := Before(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestBeforeLastTextFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "World of"

	if c := BeforeLast(a, b); c != "This is World of Oz and World of Colony and the " {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestBeforeLastTextNotFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "is WorldX"

	if c := BeforeLast(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestAfterTextFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is World"

	if c := After(a, b); c != " of Oz" {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestAfterTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is WorldX"

	if c := After(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestAfterLastTextFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "World of"

	if c := AfterLast(a, b); c != " Irony" {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestAfterLastTextNotFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "is WorldX"

	if c := AfterLast(a, b); c != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestBetweenTextFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "This"
	var c string = "of"

	if d := Between(a, b, c); d != " is World " {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestBetweenTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "ThisA"
	var c string = "of"

	if d := Between(a, b, c); d != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestBetweenFirstTextFound(t *testing.T) {
	var a string = "[a] [b] XX [c]"
	var b string = "["
	var c string = "]"

	if d := BetweenFirst(a, b, c); d != "a" {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestBetweenFirstTextNotFound(t *testing.T) {
	var a string = "[a] [b] XX [c]"
	var b string = "x"
	var c string = "]"

	if d := BetweenFirst(a, b, c); d != a {
		t.Errorf("Add(%s, %s) = %s, didn't return %s", a, b, " of Oz", c)
	}
}

func TestCamel(t *testing.T) {
	var a string = Camel("i feel good")
	var b string = Camel("i-feel-good")
	var c string = Camel("i_feel_good")

	if a != b && b != c && c != "iFeelGood" {
		t.Errorf("Camel Case Error")
	}
}

func TestPascal(t *testing.T) {
	var a string = Camel("i feel good")
	var b string = Camel("i-feel-good")
	var c string = Camel("i_feel_good")

	if a != b && b != c && c != "IFeelGood" {
		t.Errorf("Pascal Case Error")
	}
}

func TestSnake(t *testing.T) {
	var a string = Camel("This Is Title")
	var b string = Camel("this_is_title")
	var c string = Camel("this-is-title")
	var d string = Camel("thisIsTitle")
	var e string = Camel("This_Is-Title")

	if a != b && b != c && c != d && d != e && e != "this_is_title" {
		t.Errorf("Snake Case Error")
	}
}

func TestKebab(t *testing.T) {
	var a string = Camel("This Is Title")
	var b string = Camel("this_is_title")
	var c string = Camel("this-is-title")
	var d string = Camel("thisIsTitle")
	var e string = Camel("This_Is-Title")

	if a != b && b != c && c != d && d != e && e != "this-is-title" {
		t.Errorf("Kebab Case Error")
	}
}

func TestHeadline(t *testing.T) {
	var a string = Camel("This Is Title")
	var b string = Camel("this_is_title")
	var c string = Camel("this-is-title")
	var d string = Camel("thisIsTitle")
	var e string = Camel("This_Is-Title")

	if a != b && b != c && c != d && d != e && e != "This Is Title" {
		t.Errorf("Headline Case Error")
	}
}
