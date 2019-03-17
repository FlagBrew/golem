package pokeapigo

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

// GetEvolutionChain returns a EvolutionChain object
// Acceptable search queries:
// ID
func GetEvolutionChain(id int) EvolutionChain {
	searchType := "ID"
	query := strconv.Itoa(id)
	evolutionChain := EvolutionChain{}
	if checkCache("evolution-chain", query, searchType) {
		mapstructure.Decode(getDataCache("evolution-chain", query, searchType, evolutionChain), &evolutionChain)
	} else {
		dataInterface = getDataAPI("evolution-chain", query, evolutionChain)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &evolutionChain)
			storeData(true, evolutionChain, "evolution-chain", evolutionChain.ID, "")
		}
	}
	for i, detail := range evolutionChain.Chain.EvolutionDetails {
		evolutionChain.Chain.EvolutionDetails[i].GetTrigger = getEvolutionTrigger(detail.Trigger.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetItem = getItem(detail.Item.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetHeldItem = getItem(detail.HeldItem.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetKnownMove = getMove(detail.KnownMove.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetKnownMoveType = getType(detail.KnownMoveType.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetLocation = getLocation(detail.Location.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetPartySpecies = getPokemonSpecies(detail.PartySpecies.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetPartyType = getType(detail.PartyType.Name)
		evolutionChain.Chain.EvolutionDetails[i].GetTradeSpecies = getPokemonSpecies(detail.TradeSpecies.Name)

	}
	for i, detail1 := range evolutionChain.Chain.EvolvesTo {
		for ii, detail2 := range detail1.EvolutionDetails {
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetTrigger = getEvolutionTrigger(detail2.Trigger.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetItem = getItem(detail2.Item.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetHeldItem = getItem(detail2.HeldItem.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetKnownMove = getMove(detail2.KnownMove.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetKnownMoveType = getType(detail2.KnownMoveType.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetLocation = getLocation(detail2.Location.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetPartySpecies = getPokemonSpecies(detail2.PartySpecies.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetPartyType = getType(detail2.PartyType.Name)
			evolutionChain.Chain.EvolvesTo[i].EvolutionDetails[ii].GetTradeSpecies = getPokemonSpecies(detail2.TradeSpecies.Name)
		}
	}

	evolutionChain.GetBabyTriggerItem = getItem(evolutionChain.BabyTriggerItem.Name)
	evolutionChain.Chain.GetSpecies = getPokemonSpecies(evolutionChain.Chain.Species.Name)
	return evolutionChain
}

// GetEvolutionTrigger returns a EvolutionTrigger object
// Acceptable search queries:
// ID, String
func GetEvolutionTrigger(query string) EvolutionTrigger {
	searchType := idOrString(query)
	evolutionTrigger := EvolutionTrigger{}
	if checkCache("evolution-trigger", query, searchType) {
		mapstructure.Decode(getDataCache("evolution-trigger", query, searchType, evolutionTrigger), &evolutionTrigger)
	} else {
		dataInterface = getDataAPI("evolution-trigger", query, evolutionTrigger)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &evolutionTrigger)
			storeData(false, evolutionTrigger, "evolution-trigger", evolutionTrigger.ID, evolutionTrigger.Name)
		}
	}
	for _, p := range evolutionTrigger.PokemonSpecies {
		evolutionTrigger.GetPokemonSpecies = append(evolutionTrigger.GetPokemonSpecies, getPokemonSpecies(p.Name))
	}
	return evolutionTrigger
}
