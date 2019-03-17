package pokeapigo

import "github.com/mitchellh/mapstructure"

// GetLocation returns a Location object
// Acceptable search queries:
// ID, String
func GetLocation(query string) Location {
	searchType := idOrString(query)
	location := Location{}
	if checkCache("location", query, searchType) {
		mapstructure.Decode(getDataCache("location", query, searchType, location), &location)
	} else {
		dataInterface = getDataAPI("location", query, location)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &location)
			storeData(false, location, "location", location.ID, location.Name)
		}
	}
	location.GetRegion = getRegion(location.Region.Name)
	for _, area := range location.Areas {
		location.GetAreas = append(location.GetAreas, getLocationArea(area.Name))
	}

	return location
}

// GetLocationArea returns a LocationArea object
// Acceptable search queries:
// ID, String
func GetLocationArea(query string) LocationArea {
	searchType := idOrString(query)
	locationArea := LocationArea{}
	if checkCache("location-area", query, searchType) {
		mapstructure.Decode(getDataCache("location-area", query, searchType, locationArea), &locationArea)
	} else {
		dataInterface = getDataAPI("location-area", query, locationArea)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &locationArea)
			storeData(false, locationArea, "location-area", locationArea.ID, locationArea.Name)
		}
	}
	locationArea.GetLocation = getLocation(locationArea.Location.Name)
	for i, emr := range locationArea.EncounterMethodRates {
		locationArea.EncounterMethodRates[i].GetEncounterMethod = getEncounterMethod(emr.EncounterMethod.Name)
		for ii, vd := range emr.VersionDetails {
			locationArea.EncounterMethodRates[i].VersionDetails[ii].GetVersion = getVersion(vd.Version.Name)
		}
	}
	for i, pe := range locationArea.PokemonEncounters {
		locationArea.PokemonEncounters[i].GetPokemon = getPokemon(pe.Pokemon.Name)
	}
	return locationArea
}

// GetPalParkArea returns a PalParkArea object
// Acceptable search queries:
// ID, String
func GetPalParkArea(query string) PalParkArea {
	searchType := idOrString(query)
	palParkArea := PalParkArea{}
	if checkCache("pal-park-area", query, searchType) {
		mapstructure.Decode(getDataCache("pal-park-area", query, searchType, palParkArea), &palParkArea)
	} else {
		dataInterface = getDataAPI("pal-park-area", query, palParkArea)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &palParkArea)
			storeData(false, palParkArea, "pal-park-area", palParkArea.ID, palParkArea.Name)
		}
	}
	for i, pe := range palParkArea.PokemonEncounters {
		palParkArea.PokemonEncounters[i].GetPokemonSpecies = getPokemonSpecies(pe.PokemonSpecies.Name)
	}
	return palParkArea
}

// GetRegion returns a Region object
// Acceptable search queries:
// ID, String
func GetRegion(query string) Region {
	searchType := idOrString(query)
	region := Region{}
	if checkCache("region", query, searchType) {
		mapstructure.Decode(getDataCache("region", query, searchType, region), &region)
	} else {
		dataInterface = getDataAPI("region", query, region)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &region)
			storeData(false, region, "region", region.ID, region.Name)
		}
	}

	for _, location := range region.Locations {
		region.GetLocations = append(region.GetLocations, getLocation(location.Name))
	}

	region.GetMainGeneration = getGeneration(region.MainGeneration.Name)
	for _, pokedex := range region.Pokedexes {
		region.GetPokedexes = append(region.GetPokedexes, getPokedex(pokedex.Name))
	}
	for _, vg := range region.VersionGroups {
		region.GetVersionGroups = append(region.GetVersionGroups, getVersionGroup(vg.Name))
	}
	return region
}
