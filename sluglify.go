/*
Copyright paskal.maksim@gmail.com
Licensed under the Apache License, Version 2.0 (the "License")
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sluglify

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/gosimple/slug"
)

// Specification
// https://en.wikipedia.org/wiki/Hostname
//
func GetSlugString(input string, namespaceLength int, namespaceTail string) string {
	// convert all non latin symbols to latin
	result := slug.Make(input)

	// result string must be in lower case
	result = strings.ToLower(result)

	// use only valid chars (0-9 A-Z a-z -)
	reg := regexp.MustCompile("[^a-zA-Z0-9-]+")
	result = reg.ReplaceAllString(result, "")

	// RFC 952 disallowed labels from starting with a digit or with a hyphen character
	result = strings.TrimLeftFunc(result, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	// RFC 952 labels could not end with a hyphen character
	result = strings.TrimRightFunc(result, func(r rune) bool {
		if unicode.IsLetter(r) {
			return false
		}
		if unicode.IsDigit(r) {
			return false
		}

		return true
	})

	// trim string
	if len(result) > namespaceLength {
		result = result[0:namespaceLength-len(namespaceTail)] + namespaceTail
	}

	return result
}
