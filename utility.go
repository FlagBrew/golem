package golem

import "github.com/mitchellh/mapstructure"

// GetLanguage returns a Language object
// Acceptable search queries:
// ID, String
func GetLanguage(query string) Language {
	searchType := idOrString(query)
	language := Language{}
	if checkCache("language", query, searchType) {
		mapstructure.Decode(getDataCache("language", query, searchType, language), &language)
	} else {
		dataInterface = getDataAPI("language", query, language)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &language)
			storeData(false, language, "language", language.ID, language.Name)
		}
	}
	return language
}
