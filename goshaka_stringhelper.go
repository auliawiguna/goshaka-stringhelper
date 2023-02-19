/**
 * Author: Aulia Wiguna
 * File: goshaka_stringhelper.go
 */

package goshakastringhelper

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// To get string after given mark
// @Param	source	string
// @Param	mark	string
// @Return	string
func After(source, mark string) string {
	i := strings.Index(source, mark)

	if i != -1 {
		return source[i+len(mark):]
	}

	return source
}

// To get string after last occurence of given mark
// @Param	source	string
// @Param	mark	string
// @Return	string
func AfterLast(source, mark string) string {
	i := strings.LastIndex(source, mark)

	if i != -1 {
		return source[i+len(mark):]
	}

	return source
}

// To get string before given mark
// @Param	source	string
// @Param	mark	string
// @Return	string
func Before(source, mark string) string {
	i := strings.Index(source, mark)

	if i != -1 {
		return source[:i]
	}

	return source
}

// To get string before the last occurence of given mark
// @Param	source	string
// @Param	mark	string
// @Return	string
func BeforeLast(source, mark string) string {
	i := strings.LastIndex(source, mark)

	if i != -1 {
		return source[:i]
	}

	return source
}

// To get string between the first occurence of "start" and the last occurence of "stop"
// @Param	source	string
// @Param	start	string
// @Param	stop	string
// @Return	string
func Between(source, start, stop string) string {
	iStart := strings.Index(source, start)
	iStop := strings.LastIndex(source, stop)

	if iStart != -1 && iStop != -1 {
		return source[iStart+len(start) : iStop]
	}

	return source
}

// To get the smallest possible portion of a string between "start" and "stop"
// @Param	source	string
// @Param	start	string
// @Param	stop	string
// @Return	string
func BetweenFirst(source, start, stop string) string {
	iStart := strings.Index(source, start)
	iStop := strings.Index(source, stop)

	if iStart != -1 && iStop != -1 {
		return source[iStart+len(start) : iStop]
	}

	return source
}

// To convert str into camel case (camelCase)
// @Param	str	string
// @Return	string
func Camel(str string) string {
	// Replace underscores and dashes with spaces
	r := regexp.MustCompile("[-_]")
	str = r.ReplaceAllString(str, " ")

	// Capitalize the first letter of each word
	title := cases.Title(language.English)
	str = title.String(str)

	// Remove spaces
	str = strings.ReplaceAll(str, " ", "")

	// Make the first letter lowercase
	str = strings.ToLower(str)

	return str
}

// To convert str into camel case (PascalCase)
// @Param	str	string
// @Return	string
func Pascal(str string) string {
	// Replace underscores and dashes with spaces
	r := regexp.MustCompile("[-_]")
	str = r.ReplaceAllString(str, " ")

	// Capitalize the first letter of each word
	title := cases.Title(language.English)
	str = title.String(str)

	// Remove spaces
	str = strings.ReplaceAll(str, " ", "")

	return str
}

// To convert str into snake case (snake_case)
// @Param	str	string
// @Return	string
func Snake(str string) string {
	// Replace any uppercase letter with "_<lowercase letter>"
	r := regexp.MustCompile("([A-Z])")
	str = r.ReplaceAllString(str, "_$1")

	// Replace any hyphens or spaces with underscores
	r, err := regexp.Compile(`[-\s]`)
	if err != nil {
		panic(err)
	}
	str = r.ReplaceAllString(str, "_")

	// Remove any duplicate underscores
	r = regexp.MustCompile("__+")
	str = r.ReplaceAllString(str, "_")

	// Make the string lowercase
	str = strings.ToLower(str)

	// Remove any leading or trailing underscores
	str = strings.TrimPrefix(str, "_")
	str = strings.TrimSuffix(str, "_")

	return str
}

// To convert str into kebab case (kebab-case)
// @Param	str	string
// @Return	string
func Kebab(str string) string {
	// Replace any underscores with empty strings
	str = strings.ReplaceAll(str, "_", "")

	// Replace any uppercase letter with "-<lowercase letter>"
	r := regexp.MustCompile("([A-Z])")
	str = r.ReplaceAllString(str, "-$1")

	// Replace any spaces with hyphens
	r, err := regexp.Compile(`[\s-]+`)
	if err != nil {
		panic(err)
	}
	str = r.ReplaceAllString(str, "-")

	// Remove any duplicate hyphens
	r = regexp.MustCompile("--+")
	str = r.ReplaceAllString(str, "-")

	// Make the string lowercase
	str = strings.ToLower(str)

	// Remove any leading or trailing hyphens
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSuffix(str, "-")

	return str
}

// To convert str into headline case
// @Param	str	string
// @Return	string
func Headline(str string) string {
	// Replace any underscores with spaces
	str = strings.ReplaceAll(str, "_", " ")

	// Convert the first letter of each word to uppercase
	title := cases.Title(language.English)
	str = title.String(str)

	// Replace any spaces with hyphens
	r, err := regexp.Compile(`[\s-]+`)
	if err != nil {
		panic(err)
	}
	str = r.ReplaceAllString(str, " ")

	// Remove any leading or trailing spaces
	str = strings.TrimSpace(str)

	return str
}
