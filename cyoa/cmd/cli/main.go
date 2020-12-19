package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type chapters map[string]Chapter

// Chapter holds a chapter in the file
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}

// Options holds the options the user may choose
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	fn := flag.String("fn", "gopher.json", "the story file to open")
	fc := flag.String("fc", "intro", "first chapter of the story")
	flag.Parse()

	chaps, err := loadChapters(*fn)
	if err != nil {
		log.Fatalln(err)
	}

	arc := *fc
	isEnd := false

	for {
		if arc == "home" {
			isEnd = true
		}
		printTitle(chaps[arc].Title)
		printParagraphs(chaps[arc].Paragraphs)
		printOptions(chaps[arc].Options)

		if isEnd {
			fmt.Println("The End")
			os.Exit(0)
		}

		var opt int
		_ = getUserOption(&opt, chaps[arc].Options)

		arc = chaps[arc].Options[opt-1].Arc

		fmt.Println("***********************************************************************")
		fmt.Println("***********************************************************************")
		fmt.Println("***********************************************************************")
	}
}

func getUserOption(opt *int, options []Options) error {
	fmt.Print("enter valid option:")
	_, err := fmt.Fscan(os.Stdin, opt)
	if err != nil {
		printOptions(options)
		fmt.Print("invalid option! enter valid option:")
		return getUserOption(opt, options)
	}
	if *opt <= 0 || *opt > len(options) {
		printOptions(options)
		fmt.Print("invalid option! enter valid option:")
		return getUserOption(opt, options)
	}
	return err
}
func printTitle(title string) {
	fmt.Printf("%s\n\n", title)
}

func printParagraphs(paragraphs []string) {
	for _, paragraph := range paragraphs {
		fmt.Println(paragraph)
		fmt.Println("")
	}
}

func printOptions(options []Options) {
	for i, option := range options {
		index := i + 1
		fmt.Printf("Press %d to venture: %s\n", index, option.Text)
	}
	fmt.Println("")
}

func loadChapters(filename string) (chapters, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error while opening the file: %s ", err)
	}
	var chaps chapters
	err = json.Unmarshal(f, &chaps)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling: %s", err)
	}
	return chaps, nil
}
