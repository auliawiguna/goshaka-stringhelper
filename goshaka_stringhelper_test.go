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

func TestIsJsonValid(t *testing.T) {
	a := `{"name":"John", "age":30, "city":"New York"}`
	var c bool = IsJson(a)
	if !c {
		t.Errorf("TestIsJsonValid Error")
	}
}

func TestIsJsonInvalid(t *testing.T) {
	a := `{test=invalid}`
	var c bool = IsJson(a)
	if c {
		t.Errorf("TestIsJsonInvalid Error")
	}
}

func TestIsUlidValid(t *testing.T) {
	a := `01EY7ZRSN1A8CB7WJ08N0Q2QH2`
	var c bool = IsUlid(a)
	if !c {
		t.Errorf("TestIsUlidValid Error")
	}
}

func TestIsUlidInvalid(t *testing.T) {
	a := `XXXXXXXX`
	var c bool = IsUlid(a)
	if c {
		t.Errorf("TestIsUlidValid Error")
	}
}

func TestIsUuidValid(t *testing.T) {
	a := `4b48d94c-9887-46c8-9eb7-d9fcb9fb55f3`
	var c bool = IsUuid(a)
	if !c {
		t.Errorf("TestIsUlidValid Error")
	}
}

func TestIsUuidInvalid(t *testing.T) {
	a := `01EY7ZRSN1A8CB7WJ08N0Q2QH2`
	var c bool = IsUuid(a)
	if c {
		t.Errorf("TestIsUlidValid Error")
	}
}

func TestLcfirstString(t *testing.T) {
	var a string = "This is it"
	c := Lcfirst(a)
	if c != "this is it" {
		t.Errorf("TestLcfirstString Error")
	}
}

func TestLcfirstArraySuccess(t *testing.T) {
	var a []string = []string{"This is", "That is", "Not This"}

	l := Lcfirst(a)

	slices, ok := l.([]string)
	if !ok {
		t.Errorf("TestLcfirstArraySuccess Error")
		return
	}

	foundWords := make([]string, 0)
	for _, word := range slices {
		if word == "that is" || word == "this is" || word == "not This" {
			foundWords = append(foundWords, word)
		}
	}
	if len(foundWords) != len(a) {
		t.Errorf("TestLcfirstArraySuccess Error")
	}
}

func TestLimit(t *testing.T) {
	var a string = "Crazy fox over the top"
	var c string = Limit(a, 9, "...")
	if c != "Crazy fox..." {
		t.Errorf("TestLimit Error")
	}
}

func TestLowerString(t *testing.T) {
	var a string = "This Is It"
	c := Lower(a)
	if c != "this is it" {
		t.Errorf("TestLowerString Error")
	}
}

func TestLowerArraySuccess(t *testing.T) {
	var a []string = []string{"This Is", "That is", "Not This"}

	l := Lower(a)

	slices, ok := l.([]string)
	if !ok {
		t.Errorf("TestLcfirstArraySuccess Error")
		return
	}

	foundWords := make([]string, 0)
	for _, word := range slices {
		if word == "that is" || word == "this is" || word == "not this" {
			foundWords = append(foundWords, word)
		}
	}
	if len(foundWords) != len(a) {
		t.Errorf("TestLowerArraySuccess Error")
	}
}

func TestLtrim(t *testing.T) {
	var a string = "-----------------This Is It"
	c := Ltrim(a, "-")
	if c != "This Is It" {
		t.Errorf("TestLtrim Error")
	}
}

func TestMardown(t *testing.T) {
	var a string = "# Goshaka"
	c := Markdown(a)
	if c != "<h1>Goshaka</h1>\n" {
		t.Errorf("TestMarkdown Error")
	}
}

func TestMask(t *testing.T) {
	var a string = "+62890989999"
	c := Mask(a, "*", 3)
	if c != "+62*********" {
		t.Errorf("TestMask Error")
	}
}

func TestIsMatch(t *testing.T) {
	var a string = "foo bar"
	c := IsMatch(a, "foo (.*)")
	if !c {
		t.Errorf("TestIsMatch Error")
	}
}

func TestPadLeft(t *testing.T) {
	var a string = "foo bar"
	c := PadLeft(a, "+-", 7)

	if c != "+-+-+-+foo bar" {
		t.Errorf("TestPadLeft Error")
	}
}

func TestPadRight(t *testing.T) {
	var a string = "foo bar"
	c := PadRight(a, "+-", 7)

	if c != "foo bar+-+-+-+" {
		t.Errorf("TestPadRight Error")
	}
}

func TestRemove(t *testing.T) {
	var a string = "cinta brontosaurus"
	c := Remove(a, "a")

	if c != "cint brontosurus" {
		t.Errorf("TestRemove Error")
	}
}

func TestReplace(t *testing.T) {
	var a string = "cinta brontosaurus"
	c := Replace(a, "a", "b")

	if c != "cintb brontosburus" {
		t.Errorf("TestReplace Error")
	}
}

func TestRemoveSymbolExceptSpace(t *testing.T) {
	var a string = "This is an example with spaces, commas, and a period."
	c := RemoveSymbolExceptSpace(a)

	if c != "This is an example with spaces commas and a period" {
		t.Errorf("TestRemoveSymbolExceptSpace Error")
	}
}

func TestRemoveSymbol(t *testing.T) {
	var a string = "This is an example with spaces, commas, and a period."
	c := RemoveSymbol(a)

	if c != "Thisisanexamplewithspacescommasandaperiod" {
		t.Errorf("TestRemoveSymbol Error")
	}
}

func TestRTrim(t *testing.T) {
	c := Rtrim("halo---", "-")

	if c != "halo" {
		t.Errorf("TestRTrim Error")
	}
}

func TestSplit(t *testing.T) {
	c := Split("apple,sony,samsung,polytron", ",")

	if len(c) != 4 {
		t.Errorf("TestSplit Error")
	}
}

func TestSquish(t *testing.T) {
	c := Squish("This  is   a  test  string   with extra   spaces.")

	if c != "This is a test string with extra spaces." {
		t.Errorf("TestSquish Error")
	}
}

func TestSubstr(t *testing.T) {
	c := Substr("Crazy Frog", 0, 5)

	if c != "Crazy" {
		t.Errorf("TestSubstr Error")
	}
}

func TestReserve(t *testing.T) {
	c := Reverse("crazy frog")

	if c != "gorf yzarc" {
		t.Errorf("TestReserve Error")
	}
}

func TestIsPalindrome(t *testing.T) {
	c := IsPalindrome("tamat")

	if !c {
		t.Errorf("TestIsPalindrome Error")
	}
}

func TestWordCount(t *testing.T) {
	c := WordCount("the word is    ugly!")

	if c != 4 {
		t.Errorf("TestWordCount Error")
	}
}

func TestUpperString(t *testing.T) {
	var a string = "This Is It"
	c := Upper(a)
	if c != "THIS IS IT" {
		t.Errorf("TestUpperString Error")
	}
}

func TestUpperArraySuccess(t *testing.T) {
	var a []string = []string{"This Is", "That is", "Not This"}

	l := Upper(a)

	slices, ok := l.([]string)
	if !ok {
		t.Errorf("TestUpperArraySuccess Error")
		return
	}

	foundWords := make([]string, 0)
	for _, word := range slices {
		if word == "THAT IS" || word == "THIS IS" || word == "NOT THIS" {
			foundWords = append(foundWords, word)
		}
	}
	if len(foundWords) != len(a) {
		t.Errorf("TestUpperArraySuccess Error")
	}
}

func TestRandom(t *testing.T) {
	c := Random(10)
	if len(c) != 10 {
		t.Errorf("TestRandom Error")
	}
}
