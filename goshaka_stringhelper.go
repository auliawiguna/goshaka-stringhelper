/**
 * Author: Aulia Wiguna
 * File: goshaka_stringhelper.go
 */

package goshakastringhelper

import (
	"fmt"
	"regexp"
	"strconv"
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

// To determine if target is exists in str
// @Param	str	string
// @Param	target	string
// @Return	bool
func Contains(str, target string) bool {
	return strings.Contains(str, target)
}

// To determine if target is exists in str
// @Param	str	string
// @Param	target	string
// @Return	bool
func ContainsAll(str string, target []string) bool {
	foundWords := make([]string, 0)
	for _, word := range target {
		if strings.Contains(str, word) {
			foundWords = append(foundWords, word)
		}
	}

	return len(foundWords) == len(target)
}

// To determine if str is ends with target
// @Param	str	string
// @Param	target	string|array
// @Return	bool
func EndsWith(str string, target interface{}) bool {
	switch v := target.(type) {
	case string:
		// Handle string parameter
		return strings.HasSuffix(str, fmt.Sprintf("%v", target))
	case []string:
		// Handle string array parameter
		foundWords := make([]string, 0)
		for i := 0; i < len(v); i++ {
			if strings.Contains(str, v[i]) {
				foundWords = append(foundWords, v[i])
			}
		}
		return len(foundWords) > 0
	}

	return false
}

// To adds a "add" to "str" if it does not already end with "add"
// @Param	str	string
// @Param	add	string
// @Return	string
func Finish(str, add string) string {
	if EndsWith(str, add) {
		return str
	}
	return str + add
}

// To checks if param is ascii
// @Param	param	int|string
// @Return	bool
func IsASCII(param interface{}) bool {
	switch v := param.(type) {
	case int:
		// Convert integer to string
		s := strconv.Itoa(v)
		// Convert string to byte slice
		b := []byte(s)
		// Check if each byte is ASCII
		for i := 0; i < len(b); i++ {
			if b[i] < 0 || b[i] > 127 {
				return false
			}
		}
		return true
	case string:
		// Convert string to byte slice
		b := []byte(v)
		// Check if each byte is ASCII
		for i := 0; i < len(b); i++ {
			if b[i] < 0 || b[i] > 127 {
				return false
			}
		}
		return true
	default:
		// Handle other types of parameters
		return false
	}
}

// To checks if s is match with pattern
// @Param	s	string
// @Param	pattern	string
// @Return	bool
func Is(s, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
