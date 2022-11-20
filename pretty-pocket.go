package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

type pocketItem struct {
	URL       string
	TimeAdded string
	Tags      string
}

type pocketExport struct {
	Items []pocketItem
}

var items pocketExport

// Append a pocket item to a global items list
func appendItems(item *pocketItem) {
	if item.Tags != "" && item.TimeAdded != "" && item.URL != "" {
		items.Items = append(items.Items, *item)
	}
}

// Crawls children html nodes from root node and saves values into pocketItem
func parseExport(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var item pocketItem

		for _, element := range n.Attr {

			if element.Key == "href" {
				item.URL = element.Val
			}
			if element.Key == "time_added" {
				item.TimeAdded = element.Val
			}
			if element.Key == "tags" {
				item.Tags = element.Val
			}
			appendItems(&item)
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseExport(c)
	}
}

// Takes a root html node to crawl and writes the values to output file
func writeExport(n *html.Node, output string) {
	// Create output json file
	f, err := os.Create(output)
	check(err)
	defer f.Close()

	// Pointer to writer for output file
	w := bufio.NewWriter(f)

	// Parse export html file and write it to output json file
	parseExport(n)

	// Encode all items to Json
	pocketItems, err := json.Marshal(items)
	check(err)

	// Write encoded json to file
	w.WriteString(string(pocketItems))

	w.Flush()
}

// Panic on any error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Validate args count
func validateArgsCount(args []string) {
	if len(args) != 1 {
		err := fmt.Sprintf("Expected args: 1, Found %d %s", len(args), args)
		panic(err)
	}
}

// Validate args file type
func validateArgsFileExtension(argsWithoutProg []string) {
	if !strings.Contains(argsWithoutProg[0], ".html") {
		err := fmt.Sprintf("Expected file extension: *.html, Found %s", argsWithoutProg[0])
		panic(err)
	}
}

// Validate existence of file
func validateFileExist(argsWithoutProg []string) {
	htmlExportFileName := argsWithoutProg[0]
	if _, err := os.Stat(htmlExportFileName); err == nil {
		// no-op file exists!
	} else if errors.Is(err, os.ErrNotExist) {
		err := fmt.Sprintf("Cannot find %s in current directory", argsWithoutProg[0])
		panic(err)

	}
}

func main() {

	// Get CLI args for input html filename
	argsWithoutProg := os.Args[1:]
	// Validate cli args
	validateArgsCount(argsWithoutProg)
	validateArgsFileExtension(argsWithoutProg)
	validateFileExist(argsWithoutProg)

	data, err := os.ReadFile(argsWithoutProg[0])
	check(err)

	// Reader pointer to read the input html file
	doc, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		fmt.Println("error: ", err)
	}

	// Get filename without extension
	output := strings.TrimSuffix(argsWithoutProg[0], filepath.Ext(argsWithoutProg[0]))

	writeExport(doc, output+".json")

}
