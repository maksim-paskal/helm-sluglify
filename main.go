package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

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

func main() {
	kingpin.Version(appConfig.Version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if len(os.Args) < 2 {
		panic("no args")
	}

	result := slug.Make(*appConfig.input)

	if len(result) > *appConfig.namespaceLength {
		result = result[0:*appConfig.namespaceLength-len(*appConfig.namespaceTail)] + *appConfig.namespaceTail
	}

	result = strings.ToLower(result)

	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")

	if err != nil {
		panic(err)
	}

	result = reg.ReplaceAllString(result, "")

	fmt.Println(result)
}
