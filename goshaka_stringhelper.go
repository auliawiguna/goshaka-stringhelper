/**
 * Author: Aulia Wiguna
 * File: goshaka_stringhelper.go
 */

package goshakastringhelper

import (
	"strings"
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

// To getthe smallest possible portion of a string between "start" and "stop"
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
