package pokeapigo

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

// GetContestType returns a ContestType object
// Acceptable search queries:
// ID, Name
func GetContestType(query string) ContestType {
	searchType := idOrString(query)
	contestType := ContestType{}
	if checkCache("contest-type", query, searchType) {
		mapstructure.Decode(getDataCache("contest-type", query, searchType, contestType), &contestType)
	} else {
		dataInterface = getDataAPI("contest-type", query, contestType)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &contestType)
			storeData(false, contestType, "contest-type", contestType.ID, contestType.Name)
		}
	}
	contestType.GetBerryFlavor = getBerryFlavor(contestType.BerryFlavor.Name)
	for i, ct := range contestType.Names {
		contestType.Names[i].GetLanguage = getLanguage(ct.Language.Name)
	}
	return contestType
}

// GetContestEffect returns a ContestEffect object
// Acceptable search queries:
// ID
func GetContestEffect(id int) ContestEffect {
	searchType := "ID" // Because the only acceptable searchType is ID, we don't need to call the function
	query := strconv.Itoa(id)
	contestEffect := ContestEffect{}
	if checkCache("contest-effect", query, searchType) {
		mapstructure.Decode(getDataCache("contest-effect", query, searchType, contestEffect), &contestEffect)
	} else {
		dataInterface = getDataAPI("contest-effect", query, contestEffect)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &contestEffect)
			storeData(true, contestEffect, "contest-effect", contestEffect.ID, "")
		}
	}
	return contestEffect
}

// GetSuperContestEffect returns a SuperContestEffect object
// Acceptable search queries:
// ID
func GetSuperContestEffect(id int) SuperContestEffect {
	searchType := "ID"
	query := strconv.Itoa(id)
	superContestEffect := SuperContestEffect{}
	if checkCache("super-contest-effect", query, searchType) {
		mapstructure.Decode(getDataCache("super-contest-effect", query, searchType, superContestEffect), &superContestEffect)
	} else {
		dataInterface = getDataAPI("super-contest-effect", query, superContestEffect)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &superContestEffect)
			storeData(true, superContestEffect, "super-contest-effect", superContestEffect.ID, "")
		}
	}

	for _, m := range superContestEffect.Moves {
		superContestEffect.GetMoves = append(superContestEffect.GetMoves, getMove(m.Name))
	}

	return superContestEffect
}
