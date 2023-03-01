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
		t.Errorf("Error TestBeforeTextFound")
	}
}

func TestBeforeTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is WorldX"

	if c := Before(a, b); c != a {
		t.Errorf("Error TestBeforeTextNotFound")
	}
}

func TestBeforeLastTextFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "World of"

	if c := BeforeLast(a, b); c != "This is World of Oz and World of Colony and the " {
		t.Errorf("Error TestBeforeLastTextFound")
	}
}

func TestBeforeLastTextNotFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "is WorldX"

	if c := BeforeLast(a, b); c != a {
		t.Errorf("Error TestBeforeLastTextNotFound")
	}
}

func TestAfterTextFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is World"

	if c := After(a, b); c != " of Oz" {
		t.Errorf("Error TestAfterTextFound")
	}
}

func TestAfterTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "is WorldX"

	if c := After(a, b); c != a {
		t.Errorf("Error TestAfterTextNotFound")
	}
}

func TestAfterLastTextFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "World of"

	if c := AfterLast(a, b); c != " Irony" {
		t.Errorf("Error TestAfterLastTextFound")
	}
}

func TestAfterLastTextNotFound(t *testing.T) {
	var a string = "This is World of Oz and World of Colony and the World of Irony"
	var b string = "is WorldX"

	if c := AfterLast(a, b); c != a {
		t.Errorf("Error TestAfterLastTextNotFound")
	}
}

func TestBetweenTextFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "This"
	var c string = "of"

	if d := Between(a, b, c); d != " is World " {
		t.Errorf("Error TestBetweenTextFound")
	}
}

func TestBetweenTextNotFound(t *testing.T) {
	var a string = "This is World of Oz"
	var b string = "ThisA"
	var c string = "of"

	if d := Between(a, b, c); d != a {
		t.Errorf("Error TestBetweenTextNotFound")
	}
}

func TestBetweenFirstTextFound(t *testing.T) {
	var a string = "[a] [b] XX [c]"
	var b string = "["
	var c string = "]"

	if d := BetweenFirst(a, b, c); d != "a" {
		t.Errorf("Error TestBetweenFirstTextFound")
	}
}

func TestBetweenFirstTextNotFound(t *testing.T) {
	var a string = "[a] [b] XX [c]"
	var b string = "x"
	var c string = "]"

	if d := BetweenFirst(a, b, c); d != a {
		t.Errorf("Error TestBetweenFirstTextNotFound")
	}
}

func TestCamel(t *testing.T) {
	var a string = Camel("i feel good")
	var b string = Camel("i-feel-good")
	var c string = Camel("i_feel_good")

	if a != b && b != c && c != "iFeelGood" {
		t.Errorf("Error TestCamel")
	}
}

func TestPascal(t *testing.T) {
	var a string = Camel("i feel good")
	var b string = Camel("i-feel-good")
	var c string = Camel("i_feel_good")

	if a != b && b != c && c != "IFeelGood" {
		t.Errorf("Error TestPascal")
	}
}

func TestSnake(t *testing.T) {
	var a string = Camel("This Is Title")
	var b string = Camel("this_is_title")
	var c string = Camel("this-is-title")
	var d string = Camel("thisIsTitle")
	var e string = Camel("This_Is-Title")

	if a != b && b != c && c != d && d != e && e != "this_is_title" {
		t.Errorf("Error TestSnake")
	}
}

func TestKebab(t *testing.T) {
	var a string = Camel("This Is Title")
	var b string = Camel("this_is_title")
	var c string = Camel("this-is-title")
	var d string = Camel("thisIsTitle")
	var e string = Camel("This_Is-Title")

	if a != b && b != c && c != d && d != e && e != "this-is-title" {
		t.Errorf("Error TestKebab")
	}
}

func TestHeadline(t *testing.T) {
	var a string = Camel("This Is Title")
	var b string = Camel("this_is_title")
	var c string = Camel("this-is-title")
	var d string = Camel("thisIsTitle")
	var e string = Camel("This_Is-Title")

	if a != b && b != c && c != d && d != e && e != "This Is Title" {
		t.Errorf("Error TestHeadline")
	}
}

func TestContainsSuccess(t *testing.T) {
	var a string = "This Is Title"
	var b string = "Is"
	var c bool = Contains(a, b)
	if !c {
		t.Errorf("Contains Error")
	}
}

func TestDoesNotContainsSuccess(t *testing.T) {
	var a string = "This Is Title"
	var b string = "Isx"
	var c bool = Contains(a, b)
	if c {
		t.Errorf("TestDoesNotContainsSuccess Error")
	}
}

func TestContainsAllSuccess(t *testing.T) {
	var a string = "This Is Title"
	var b []string = []string{"This", "Is", "Title"}
	var c bool = ContainsAll(a, b)
	if !c {
		t.Errorf("TestContainsAllSuccess Error")
	}
}

func TestDoesNotContainsAllSuccess(t *testing.T) {
	var a string = "This Is Title"
	var b []string = []string{"This", "Are", "Title"}
	var c bool = ContainsAll(a, b)
	if c {
		t.Errorf("TestDoesNotContainsAllSuccess Error")
	}
}

func TestEndsWithUsingStringSuccess(t *testing.T) {
	var a string = "This Is Title"
	var b []string = []string{"Are", "Title"}
	var c bool = EndsWith(a, b)
	if !c {
		t.Errorf("TestEndsWithUsingStringSuccess Error")
	}
}

func TestEndsWithUsingStringFail(t *testing.T) {
	var a string = "This Is Title"
	var b string = "Not"
	var c bool = EndsWith(a, b)
	if c {
		t.Errorf("TestEndsWithUsingStringFail Error")
	}
}

func TestEndsWithUsingArraySuccess(t *testing.T) {
	var a string = "This Is Title"
	var b []string = []string{"Are", "Title"}
	var c bool = EndsWith(a, b)
	if !c {
		t.Errorf("TestEndsWithUsingArraySuccess Error")
	}
}

func TestEndsWithUsingArrayFail(t *testing.T) {
	var a string = "This Is Title"
	var b []string = []string{"Are", "That"}
	var c bool = EndsWith(a, b)
	if c {
		t.Errorf("TestEndsWithUsingArrayFail Error")
	}
}

func TestFinishUsingExistingWordSuccess(t *testing.T) {
	var a string = "This Is Title"
	var b string = "Title"
	var c string = Finish(a, b)
	if c != a {
		t.Errorf("TestFinishUsingExistingWordSuccess Error")
	}
}

func TestFinishUsingNonExistingWordSuccess(t *testing.T) {
	var a string = "This Is Title"
	var b string = " and Word"
	var c string = Finish(a, b)
	if c != a+b {
		t.Errorf("TestFinishUsingExistingWordSuccess Error")
	}
}

func TestIsAsciiUsingAscii(t *testing.T) {
	var a string = "U"
	var c bool = IsAscii(a)
	if !c {
		t.Errorf("TestIsAsciiUsingAscii Error")
	}
}

func TestIsAsciiUsingNonAscii(t *testing.T) {
	var a string = "ü"
	var c bool = IsAscii(a)
	if c {
		t.Errorf("TestIsAsciiUsingNonAscii Error")
	}
}

func TestIsUsingMatchwordSuccess(t *testing.T) {
	var a string = "Prambanan"
	var b string = "Prambanan"
	var c bool = Is(a, b)
	if !c {
		t.Errorf("TestIsUsingMatchwordSuccess Error")
	}
}

func TestIsUsingWildcardSuccess(t *testing.T) {
	var a string = "Prambanan"
	var b string = "Pram*"
	var c bool = Is(a, b)
	if !c {
		t.Errorf("TestIsUsingWildcardSuccess Error")
	}
}

func TestIsUsingMatchwordFail(t *testing.T) {
	var a string = "Prambanan"
	var b string = "Prambananx"
	var c bool = Is(a, b)
	if c {
		t.Errorf("TestIsUsingMatchwordFail Error")
	}
}

func TestIsUsingWildcardFail(t *testing.T) {
	var a string = "Prambanan"
	var b string = "Boro*"
	var c bool = Is(a, b)
	if c {
		t.Errorf("TestIsUsingWildcardFail Error")
	}
}

func TestTrimStringSuccess(t *testing.T) {
	var a string = "Prambanan "
	c := Trim(a)
	if c != "Prambanan" {
		t.Errorf("TestTrimStringSuccess Error")
	}
}

func TestTrimArraySuccess(t *testing.T) {
	var a []string = []string{"This ", "  Is", " Title "}
	foundWords := make([]string, 0)
	for _, word := range a {
		t := Trim(word)
		if t == "This" || t == "Is" || t == "Title" {
			foundWords = append(foundWords, word)
		}
	}
	if len(foundWords) != len(a) {
		t.Errorf("TestTrimArraySuccess Error")
	}
}

func TestIsEmpty(t *testing.T) {
	var a string = "  "
	var c bool = IsEmpty(a)
	if !c {
		t.Errorf("TestIsEmpty Error")
	}
}

func TestIsNotEmpty(t *testing.T) {
	var a string = "  Test"
	var c bool = IsNotEmpty(a)
	if !c {
		t.Errorf("TestIsNotEmpty Error")
	}
}
