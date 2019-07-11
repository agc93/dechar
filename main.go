package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/ahmetb/go-linq"
	"github.com/olekukonko/tablewriter"
)

func main() {
	outputPtr := flag.String("o", "simple", "Output format {simple|lines|table};.")
	flag.Parse()
	input := flag.Args()

	if len(input) != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var words []string
	groups := strings.Split(strings.TrimPrefix(input[0], "CHAR("), "CHAR(")
	From(groups).SelectT(func(c string) string {
		return strings.TrimRight(c, ",)")
	}).SelectT(
		func(g string) string {
			var word []string
			for _, c := range strings.Split(g, ",") {
				var char int
				if i, err := strconv.Atoi(c); err == nil {
					char = i
				}
				word = append(word, string(char))
			}
			return strings.Join(word, "")
		}).ToSlice(&words)
	switch *outputPtr {
	case "simple":
		fmt.Println(strings.Join(words, ","))
	case "lines":
		fmt.Println(strings.Join(words, "\n\r"))
	case "table":
		printTable(groups, words)
	}
}

func printTable(input []string, words []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Original", "ASCII"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor})
	for i, v := range input {
		table.Append([]string{v, words[i]})
	}
	table.Render()
}
