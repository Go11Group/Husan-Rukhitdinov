package main

// translateService holds the translations
type translateService struct {
	translations map[string]string
}

// NewTranslateService creates a new translateService with some predefined translations
func NewTranslateService() *translateService {
	return &translateService{
		translations: map[string]string{
			"salom":  "hello",
			"dunyo":  "world",
			"kitob":  "book",
			"oila":   "family",
			"maktab": "school",
		},
	}
}

// Translate method translates an array of Uzbek words into English
func (ts *translateService) Translate(words []string) []string {
	translatedWords := []string{}
	for _, word := range words {
		if translation, found := ts.translations[word]; found {
			translatedWords = append(translatedWords, translation)
		} else {
			translatedWords = append(translatedWords, word) // return the original word if no translation is found
		}
	}
	return translatedWords
}
