/**
 * Author: Aulia Wiguna
 * File: goshaka_stringhelper.go
 */

package goshakastringhelper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid"
	"github.com/russross/blackfriday/v2"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// To get string after given mark
//
//	param	source	string
//	param	mark	string
//	return 	string
func After(source, mark string) string {
	i := strings.Index(source, mark)

	if i != -1 {
		return source[i+len(mark):]
	}

	return source
}

// To get string after last occurence of given mark
//
//	param	source	string
//	param	mark	string
//	return 	string
func AfterLast(source, mark string) string {
	i := strings.LastIndex(source, mark)

	if i != -1 {
		return source[i+len(mark):]
	}

	return source
}

// To get string before given mark
//
//	param	source	string
//	param	mark	string
//	return 	string
func Before(source, mark string) string {
	i := strings.Index(source, mark)

	if i != -1 {
		return source[:i]
	}

	return source
}

// To get string before the last occurence of given mark
//
//	param	source	string
//	param	mark	string
//	return 	string
func BeforeLast(source, mark string) string {
	i := strings.LastIndex(source, mark)

	if i != -1 {
		return source[:i]
	}

	return source
}

// To get string between the first occurence of "start" and the last occurence of "stop"
//
//	param	source	string
//	param	start	string
//	param	stop	string
//	return 	string
func Between(source, start, stop string) string {
	iStart := strings.Index(source, start)
	iStop := strings.LastIndex(source, stop)

	if iStart != -1 && iStop != -1 {
		return source[iStart+len(start) : iStop]
	}

	return source
}

// To get the smallest possible portion of a string between "start" and "stop"
//
//	param	source	string
//	param	start	string
//	param	stop	string
//	return 	string
func BetweenFirst(source, start, stop string) string {
	iStart := strings.Index(source, start)
	iStop := strings.Index(source, stop)

	if iStart != -1 && iStop != -1 {
		return source[iStart+len(start) : iStop]
	}

	return source
}

// To convert str into camel case (camelCase)
//
//	param	str	string
//	return 	string
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
//
//	param	str	string
//	return 	string
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
//
//	param	str	string
//	return 	string
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
//
//	param	str	string
//	return 	string
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
//
//	param	str	string
//	return 	string
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
//
//	param	str	string
//	param	target	string
//	return 	bool
func Contains(str, target string) bool {
	return strings.Contains(str, target)
}

// To determine if target is exists in str
//
//	param	str	string
//	param	target	string
//	return 	bool
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
//
//	param	str	string
//	param	target	string|array
//	return 	bool
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
//
//	param	str	string
//	param	add	string
//	return 	string
func Finish(str, add string) string {
	if EndsWith(str, add) {
		return str
	}
	return str + add
}

// To checks if param is ascii
//
//	param	param	int|string
//	return 	bool
func IsAscii(param interface{}) bool {
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
//
//	param	s	string
//	param	pattern	string
//	return 	bool
func Is(s, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

// To trim spaces
//
//	param	target	string|array
//	return 	bool
func Trim(target interface{}) interface{} {
	switch v := target.(type) {
	case string:
		// Handle string parameter
		return strings.TrimSpace(fmt.Sprintf("%v", target))
	case []string:
		// Handle string array parameter
		foundWords := make([]string, 0)
		for i := 0; i < len(v); i++ {
			foundWords = append(foundWords, strings.TrimSpace(fmt.Sprintf("%v", v[i])))
		}
		return foundWords
	}

	return false
}

// To check whether the given target is empty or not
//
//	param	target	string|array
//	return 	bool
func IsEmpty(target string) bool {
	t := Trim(target)
	return t == ""
}

// To check whether the given target is empty or not
//
//	param	target	string|array
//	return 	bool
func IsNotEmpty(target string) bool {
	t := Trim(target)
	return t != ""
}

// To check whether the given target is a valid JSON
//
//	param	target	string
//	return 	bool
func IsJson(target string) bool {
	var data interface{}
	err := json.Unmarshal([]byte(target), &data)
	return err == nil
}

// To check whether the given target is a valid ULID
//
//	param	target	string
//	return 	bool
func IsUlid(target string) bool {
	_, err := ulid.Parse(target)
	return err == nil
}

// To check whether the given target is a valid UUID
//
//	param	target	string
//	return 	bool
func IsUuid(target string) bool {
	_, err := uuid.Parse(target)
	return err == nil
}

// To returns the given string with the first character lowercased
//
//	param	target	string|array
//	return 	interface
func Lcfirst(target interface{}) interface{} {
	switch v := target.(type) {
	case string:
		// Handle string parameter
		var s string = fmt.Sprintf("%v", target)
		return strings.ToLower(s[:1]) + s[1:]
	case []string:
		// Handle string array parameter
		foundWords := make([]string, 0)
		for i := 0; i < len(v); i++ {
			var s string = fmt.Sprintf("%v", v[i])
			var l string = strings.ToLower(s[:1]) + s[1:]
			foundWords = append(foundWords, l)
		}
		return foundWords
	}

	return false
}

// To truncates the given string to the specified length
//
//	param	target	string
//	param	length	int
//	param	placeholder	string
//	return 	string
func Limit(target string, length int, placeholder string) string {
	return target[0:length] + placeholder
}

// To returns lowercase the given string
//
//	param	target	string|array
//	return 	interface
func Lower(target interface{}) interface{} {
	switch v := target.(type) {
	case string:
		// Handle string parameter
		var s string = fmt.Sprintf("%v", target)
		return strings.ToLower(s)
	case []string:
		// Handle string array parameter
		foundWords := make([]string, 0)
		for i := 0; i < len(v); i++ {
			var s string = fmt.Sprintf("%v", v[i])
			var l string = strings.ToLower(s)
			foundWords = append(foundWords, l)
		}
		return foundWords
	}

	return false
}

// To trims the left side of the string
//
//	param	str	string
//	param	cutset	string
//	return 	string
func Ltrim(str, cutset string) string {
	return strings.TrimLeft(str, cutset)
}

// To convert given markdown to the HTML format
//
//	param	str	string
//	return 	string
func Markdown(str string) string {
	m := []byte(str)
	html := blackfriday.Run(m)
	return string(html)
}

// To masks a given string string with a repeated character
//
//	param	str	string
//	param	mask	string
//	param	start	int
//	return 	string
func Mask(str, mask string, start int) string {
	b := []byte(str)
	for i := 0; i < len(b); i++ {
		if i >= start {
			b[i] = []byte(mask)[0]
		}
	}

	return string(b)
}

// To check whether str matches a given pattern
//
//	param	str	string
//	param	pattern	string
//	return 	bool
func IsMatch(str, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}

// To pad the left side of a string with another string until the final string reaches the desired length
//
//	param	str	string
//	param	pad	string
//	param	length	int
//	return 	string
func PadLeft(str, pad string, length int) string {
	var b string = ""
	var assigned int = 1
	for i := 0; i < length; i++ {
		for j := 0; j < len(pad); j++ {
			if assigned <= length {
				c := pad[j]
				b = b + string(c)
				assigned++
			}
		}
	}

	return string(b) + str
}

// To pad the right side of a string with another string until the final string reaches the desired length
//
//	param	str	string
//	param	pad	string
//	param	length	int
//	return 	string
func PadRight(str, pad string, length int) string {
	var b string = ""
	var assigned int = 1
	for i := 0; i < length; i++ {
		for j := 0; j < len(pad); j++ {
			if assigned <= length {
				c := pad[j]
				b = b + string(c)
				assigned++
			}
		}
	}

	return str + string(b)
}

// To remove target from str
//
//	param	str	string
//	param	target	string
//	return 	string
func Remove(str, target string) string {
	return strings.ReplaceAll(str, target, "")
}

// To replace target in str into replace
//
//	param	str	string
//	param	target	string
//	param	replace	string
//	return 	string
func Replace(str, target, replace string) string {
	return strings.ReplaceAll(str, target, replace)
}

// To remove all symbol and space from str
//
//	param	str	string
//	return 	string
func RemoveSymbol(str string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(str, "")
}

// To remove all symbol from str
//
//	param	str	string
//	return 	string
func RemoveSymbolExceptSpace(str string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9\\s]+`)
	return reg.ReplaceAllString(str, "")
}

// To trims the right side of the string
//
//	param	str	string
//	param	cutset	string
//	return 	string
func Rtrim(str, cutset string) string {
	return strings.TrimRight(str, cutset)
}

// To split string into an array
//
//	param	str	string
//	param	delimiter	string
//	return 	array of string
func Split(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

// To squish extra spaces from a string
//
//	param	str	string
//	return 	string
func Squish(str string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, " ")
}

// To get a substring from a string
//
//	param	str	string
//	param	start	int
//	param	end	int
//	return 	string
func Substr(str string, start, end int) string {
	return str[start:end]
}

// To reverse a string
//
//	param	str	string
//	return 	string
func Reverse(str string) (res string) {
	for _, v := range str {
		res = string(v) + res
	}

	return
}

// To check whether the given target is palindrome or not
//
//	param str string
//	return  bool
func IsPalindrome(str string) bool {
	ret := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			ret = false
			break
		}
	}

	return ret
}

// Returns the number of words that a string contains
//
//	param	str	string
//	return 	int
func WordCount(str string) int {
	var s string = Squish(RemoveSymbolExceptSpace(str))
	a := Split(s, " ")

	return len(a)
}

// To returns uppercase the given string
//
//	param	target	string|array
//	return 	interface
func Upper(target interface{}) interface{} {
	switch v := target.(type) {
	case string:
		// Handle string parameter
		var s string = fmt.Sprintf("%v", target)
		return strings.ToUpper(s)
	case []string:
		// Handle string array parameter
		foundWords := make([]string, 0)
		for i := 0; i < len(v); i++ {
			var s string = fmt.Sprintf("%v", v[i])
			var l string = strings.ToUpper(s)
			foundWords = append(foundWords, l)
		}
		return foundWords
	}

	return false
}

// To returns random alphanumeric characters
//
//	param	length	int
//	return	interface
func Random(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
