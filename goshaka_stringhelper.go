/**
 * Author: Aulia Wiguna
 * File: goshaka_stringhelper.go
 */

package goshakastringhelper

import "strings"

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
