package sluglify

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/gosimple/slug"
)

func GetSlugString(input string, namespaceLength int, namespaceTail string) string {
	result := slug.Make(input)

	result = strings.ToLower(result)

	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")

	if err != nil {
		panic(err)
	}

	result = reg.ReplaceAllString(result, "")

	result = strings.TrimFunc(result, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	if len(result) > namespaceLength {
		result = result[0:namespaceLength-len(namespaceTail)] + namespaceTail
	}

	return result
}
