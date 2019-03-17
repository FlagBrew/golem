package pokeapigo

import "github.com/mitchellh/mapstructure"

// GetGeneration returns a Generation object
// Acceptable search queries:
// ID, String
func GetGeneration(query string) Generation {
	searchType := idOrString(query)
	generation := Generation{}
	if checkCache("generation", query, searchType) {
		mapstructure.Decode(getDataCache("generation", query, searchType, generation), &generation)
	} else {
		dataInterface = getDataAPI("generation", query, generation)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &generation)
			storeData(false, generation, "generation", generation.ID, generation.Name)
		}
	}

	for _, ability := range generation.Abilities {
		generation.GetAbilities = append(generation.GetAbilities, getAbility(ability.Name))
	}

	generation.GetMainRegion = getRegion(generation.MainRegion.Name)

	for _, move := range generation.Moves {
		generation.GetMoves = append(generation.GetMoves, getMove(move.Name))
	}

	for _, pokemon := range generation.PokemonSpecies {
		generation.GetPokemonSpecies = append(generation.GetPokemonSpecies, getPokemonSpecies(pokemon.Name))
	}

	for _, t := range generation.Types {
		generation.GetTypes = append(generation.GetTypes, getType(t.Name))
	}

	for _, vg := range generation.VersionGroups {
		generation.GetVersionGroups = append(generation.GetVersionGroups, getVersionGroup(vg.Name))
	}

	return generation
}

// GetPokedex returns a Pokedex object
// Acceptable search queries:
// ID, String
func GetPokedex(query string) Pokedex {
	searchType := idOrString(query)
	pokedex := Pokedex{}
	if checkCache("pokedex", query, searchType) {
		mapstructure.Decode(getDataCache("pokedex", query, searchType, pokedex), &pokedex)
	} else {
		dataInterface = getDataAPI("pokedex", query, pokedex)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokedex)
			storeData(false, pokedex, "pokedex", pokedex.ID, pokedex.Name)
		}
	}
	pokedex.GetRegion = getRegion(pokedex.Region.Name)
	for _, vg := range pokedex.VersionGroups {
		pokedex.GetVersionGroups = append(pokedex.GetVersionGroups, getVersionGroup(vg.Name))
	}

	for i, pe := range pokedex.PokemonEntries {
		pokedex.PokemonEntries[i].GetPokemonSpecies = getPokemonSpecies(pe.PokemonSpecies.Name)
	}

	return pokedex
}

// GetVersion returns a Version object
// Acceptable search queries:
// ID, String
func GetVersion(query string) Version {
	searchType := idOrString(query)
	version := Version{}
	if checkCache("version", query, searchType) {
		mapstructure.Decode(getDataCache("version", query, searchType, version), &version)
	} else {
		dataInterface = getDataAPI("version", query, version)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &version)
			storeData(false, version, "version", version.ID, version.Name)
		}
	}
	version.GetVersionGroup = getVersionGroup(version.VersionGroup.Name)
	return version
}

// GetVersionGroup returns a VersionGroup object
// Acceptable search queries:
// ID, String
func GetVersionGroup(query string) VersionGroup {
	searchType := idOrString(query)
	versionGroup := VersionGroup{}
	if checkCache("version-group", query, searchType) {
		mapstructure.Decode(getDataCache("version-group", query, searchType, versionGroup), &versionGroup)
	} else {
		dataInterface = getDataAPI("version-group", query, versionGroup)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &versionGroup)
			storeData(false, versionGroup, "version-group", versionGroup.ID, versionGroup.Name)
		}
	}
	for _, mlm := range versionGroup.MoveLearnMethods {
		versionGroup.GetMoveLearnMethods = append(versionGroup.GetMoveLearnMethods, getMoveLearnMethod(mlm.Name))
	}

	for _, pd := range versionGroup.Pokedexes {
		versionGroup.GetPokedexes = append(versionGroup.GetPokedexes, getPokedex(pd.Name))
	}

	for _, v := range versionGroup.Versions {
		versionGroup.GetVersions = append(versionGroup.GetVersions, getVersion(v.Name))
	}

	versionGroup.GetGeneration = getGeneration(versionGroup.Generation.Name)

	return versionGroup
}
