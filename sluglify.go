package sluglify

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/gosimple/slug"
)

func GetSlugString(input string, namespaceLength int, namespaceTail string) string {
	result := slug.Make(input)

	if len(result) > namespaceLength {
		result = result[0:namespaceLength-len(namespaceTail)] + namespaceTail
	}

	result = strings.ToLower(result)

	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")

	if err != nil {
		panic(err)
	}

	result = reg.ReplaceAllString(result, "")

	return strings.TrimFunc(result, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}
