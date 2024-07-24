// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rcstr

import "unicode/utf8"

// LenRune returns string length of unicode.
func LenRune(str string) int {
	return utf8.RuneCountInString(str)
}
