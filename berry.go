package golem

import (
	"github.com/mitchellh/mapstructure"
)

// GetBerry returns a Berry object
// Acceptable search queries:
// ID, Name
func GetBerry(query string) Berry {
	searchType := idOrString(query)
	berry := Berry{}
	if checkCache("berry", query, searchType) {
		mapstructure.Decode(getDataCache("berry", query, searchType, berry), &berry)
	} else {
		dataInterface = getDataAPI("berry", query, berry)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &berry)
			storeData(false, berry, "berry", berry.ID, berry.Name)
		}
	}
	// set the BerryFirmness function to an anonymous function before returning it
	berry.GetFirmness = func() BerryFirmness {
		return GetBerryFirmness(berry.Firmness.Name)
	}

	// loop through the flavors to set the flavor function
	for i, flavor := range berry.Flavors {
		berry.Flavors[i].GetFlavor = getBerryFlavor(flavor.Flavor.Name)
	}

	berry.GetItem = getItem(berry.Item.Name)
	berry.GetNaturalGiftType = getType(berry.NaturalGiftType.Name)
	return berry

}

// GetBerryFirmness returns a BerryFirmness object
// Acceptable search queries:
// ID, Name
func GetBerryFirmness(query string) BerryFirmness {
	searchType := idOrString(query)
	berryFirmness := BerryFirmness{}
	if checkCache("berry-firmness", query, searchType) {
		mapstructure.Decode(getDataCache("berry-firmness", query, searchType, berryFirmness), &berryFirmness)
	} else {
		dataInterface = getDataAPI("berry-firmness", query, berryFirmness)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &berryFirmness)
			storeData(false, berryFirmness, "berry-firmness", berryFirmness.ID, berryFirmness.Name)
		}
	}
	// loop through Berries and add to GetBerries
	for _, berry := range berryFirmness.Berries {
		berryFirmness.GetBerries = append(berryFirmness.GetBerries, getBerry(berry.Name))
	}
	return berryFirmness
}

// GetBerryFlavor returns a BerryFlavor object
// Acceptable search queries:
// ID, Name
func GetBerryFlavor(query string) BerryFlavor {
	searchType := idOrString(query)
	bf := BerryFlavor{}
	if checkCache("berry-flavor", query, searchType) {
		mapstructure.Decode(getDataCache("berry-flavor", query, searchType, bf), &bf)
	} else {
		dataInterface = getDataAPI("berry-flavor", query, bf)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &bf)
			storeData(false, bf, "berry-flavor", bf.ID, bf.Name)
		}
	}
	bf.GetContestType = getContestType(bf.ContestType.Name)
	for i, berry := range bf.Berries {
		bf.Berries[i].GetBerry = getBerry(berry.Berry.Name)
	}
	return bf
}
