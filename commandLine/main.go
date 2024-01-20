package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// get args
	argsWithProg := os.Args
	// only the args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	//Flags
	wordPtr := flag.String("word", "foo", "a string")

	numbePtr := flag.Int("num", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numbePtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	// Sub-commands
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barlvl := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' Sub-commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println(" name:", *fooName)
		fmt.Println(" tail:", fooCmd.Arg())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" level:", *barlvl)
		fmt.Println(" tail:", barCmd.Arg())

	default:
		fmt.Println("Expected 'foo' or 'bar' subcommands")
	}

}
