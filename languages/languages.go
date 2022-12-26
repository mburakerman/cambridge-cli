package languages

import (
	"fmt"

	"github.com/fatih/color"
)

type Language string

const (
	English    Language = "english"
	Turkish    Language = "turkish"
	Dutch      Language = "dutch"
	French     Language = "french"
	German     Language = "german"
	Indonesian Language = "indonesian"
	Italian    Language = "italian"
	Japanese   Language = "japanese"
	Polish     Language = "polish"
	Portuguese Language = "portuguese"
	Spanish    Language = "spanish"
	Arabic     Language = "arabic"
	Czech      Language = "czech"
	Danish     Language = "danish"
	Korean     Language = "korean"
	Malay      Language = "malay"
	Norwegian  Language = "norwegian"
	Russian    Language = "russian"
	Thai       Language = "thai"
	Ukrainian  Language = "ukrainian"
	Vietnamese Language = "vietnamese"
)

var ValidLanguages = []Language{English, Turkish, Dutch, French, German,
	Indonesian, Italian, Japanese, Polish, Portuguese, Spanish,
	Arabic, Czech, Danish, Korean,
	Malay, Norwegian, Russian, Thai, Ukrainian, Vietnamese}

func CheckSupportedLanguage(selectedLanguage string) bool {
	for _, lang := range ValidLanguages {
		if selectedLanguage == string(lang) {
			return true
		}
	}

	color.Red("\n😔 Invalid language ❌" + "\n ")
	color.Green("💪 Supported languages: %+v\n", ValidLanguages)
	fmt.Println("\n ")

	return false
}
