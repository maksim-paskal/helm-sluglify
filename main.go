package main

import (
	"fmt"
	"os"

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
	kingpin.Version(fmt.Sprintf("%s-%s", appConfig.Version, buildTime))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if len(os.Args) < 2 {
		panic("no args")
	}

	fmt.Println(getSlugString(*appConfig.input, *appConfig.namespaceLength, *appConfig.namespaceTail))
}
