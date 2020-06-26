package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/gosimple/slug"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	buildTime string
)

type appConfigType struct {
	Version         string
	input           *string
	namespaceLength *int
	namespaceTail   *string
}

var appConfig = appConfigType{
	Version: "1.0.1",
	input: kingpin.Flag(
		"input",
		"input string",
	).Short('i').Required().String(),
	namespaceLength: kingpin.Flag(
		"namespace.length",
		"namespace length",
	).Default("40").Int(),
	namespaceTail: kingpin.Flag(
		"namespace.tail",
		"namespace tail",
	).Default("zzz").String(),
}

func getSlug(input string, namespaceLength int, namespaceTail string) string {
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

func main() {
	kingpin.Version(fmt.Sprintf("%s-%s", appConfig.Version, buildTime))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if len(os.Args) < 2 {
		panic("no args")
	}

	fmt.Println(getSlug(*appConfig.input, *appConfig.namespaceLength, *appConfig.namespaceTail))
}
