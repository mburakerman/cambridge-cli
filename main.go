package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"github.com/mburakerman/cambridge-cli/languages"
)

// flags
var language = flag.String("language", string(languages.English), "select language")
var showAllMeanings = flag.Bool("showAllMeanings", false, "display all of the meanings of the word")

func main() {
	flag.Parse()

	if !languages.CheckSupportedLanguage(*language) {
		return
	}

	fmt.Print("Enter word: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := scanner.Text()

	c := colly.NewCollector(
		colly.AllowedDomains("dictionary.cambridge.org"),
	)

	var wordLevel string
	var meanings []string
	var exampleSentence string

	// get english level info
	c.OnHTML(".epp-xref.dxref", func(e *colly.HTMLElement) {
		if e.Index == 0 && !*showAllMeanings {
			wordLevel = e.Text
			wordLevel = strings.TrimSpace(wordLevel)
		}
	})

	// get meanings
	var meaningSelector = ".def.ddef_d.db"
	if *language != string(languages.English) {
		meaningSelector = ".trans.dtrans"
	}
	c.OnHTML(meaningSelector, func(e *colly.HTMLElement) {
		if *showAllMeanings {
			results := strings.TrimSpace(e.Text)
			results = strings.Replace(results, ":", "", -1)
			meanings = append(meanings, results)
		} else if e.Index == 0 {
			result := strings.TrimSpace(e.Text)
			result = strings.Replace(result, ":", "", -1)
			meanings = append(meanings, result)
		}
	})

	// get example sentence
	c.OnHTML(".eg.deg", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			result := e.Text
			exampleSentence = strings.TrimSpace(result)
		}
	})

	// print when scrapping done
	yellowBackground := color.New(color.FgBlack, color.BgYellow)
	c.OnScraped(func(r *colly.Response) {
		if len(wordLevel) > 0 {
			fmt.Print("\n📈 ")
			yellowBackground.Print(" " + wordLevel + " ")
			fmt.Print("\n ")
		}

		for i := 0; i < len(meanings); i++ {
			color.Green("\n✅ " + meanings[i])
		}

		if len(exampleSentence) > 0 {
			fmt.Println("\n📝 " + color.CyanString(exampleSentence) + "\n")
		} else {
			fmt.Println("")
		}
	})

	// visit url to scrap
	if *language == string(languages.English) {
		c.Visit(fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english/%v", word))
	}
	c.Visit(fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english-%v/%v", *language, word))

}
