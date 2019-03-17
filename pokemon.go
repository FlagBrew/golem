package golem

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

// GetAbility returns a Ability object
// Acceptable search queries:
// ID, String
func GetAbility(query string) Ability {
	searchType := idOrString(query)
	ability := Ability{}
	if checkCache("ability", query, searchType) {
		mapstructure.Decode(getDataCache("ability", query, searchType, ability), &ability)
	} else {
		dataInterface = getDataAPI("ability", query, ability)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &ability)
			storeData(false, ability, "ability", ability.ID, ability.Name)
		}
	}

	ability.GetGeneration = getGeneration(ability.Generation.Name)
	for i, ec := range ability.EffectChanges {
		ability.EffectChanges[i].GetVersionGroup = getVersionGroup(ec.VersionGroup.Name)
	}
	for i, ft := range ability.FlavorTextEntries {
		ability.FlavorTextEntries[i].GetLanguage = getLanguage(ft.Language.Name)
		ability.FlavorTextEntries[i].GetVersionGroup = getVersionGroup(ft.VersionGroup.Name)
	}

	for i, p := range ability.Pokemon {
		ability.Pokemon[i].GetPokemon = getPokemon(p.Pokemon.Name)
	}
	return ability
}

// GetCharacteristic returns a Characteristic object
// Acceptable search queries:
// ID
func GetCharacteristic(id int) Characteristic {
	searchType := "ID"
	query := strconv.Itoa(id)
	characteristic := Characteristic{}
	if checkCache("characteristic", query, searchType) {
		mapstructure.Decode(getDataCache("characteristic", query, searchType, characteristic), &characteristic)
	} else {
		dataInterface = getDataAPI("characteristic", query, characteristic)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &characteristic)
			storeData(true, characteristic, "characteristic", characteristic.ID, "")
		}
	}
	return characteristic
}

// GetEggGroup returns a EggGroup object
// Acceptable search queries:
// ID, String
func GetEggGroup(query string) EggGroup {
	searchType := idOrString(query)
	eggGroup := EggGroup{}
	if checkCache("egg-group", query, searchType) {
		mapstructure.Decode(getDataCache("egg-group", query, searchType, eggGroup), &eggGroup)
	} else {
		dataInterface = getDataAPI("egg-group", query, eggGroup)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &eggGroup)
			storeData(false, eggGroup, "egg-group", eggGroup.ID, eggGroup.Name)
		}
	}

	for _, ps := range eggGroup.PokemonSpecies {
		eggGroup.GetPokemonSpecies = append(eggGroup.GetPokemonSpecies, getPokemonSpecies(ps.Name))
	}
	return eggGroup
}

// GetGender returns a Gender object
// Acceptable search queries:
// ID, String
func GetGender(query string) Gender {
	searchType := idOrString(query)
	gender := Gender{}
	if checkCache("gender", query, searchType) {
		mapstructure.Decode(getDataCache("gender", query, searchType, gender), &gender)
	} else {
		dataInterface = getDataAPI("gender", query, gender)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &gender)
			storeData(false, gender, "gender", gender.ID, gender.Name)
		}
	}
	for _, rfe := range gender.RequiredForEvolution {
		gender.GetRequiredForEvolution = append(gender.GetRequiredForEvolution, getPokemonSpecies(rfe.Name))
	}

	for i, psd := range gender.PokemonSpeciesDetails {
		gender.PokemonSpeciesDetails[i].GetPokemonSpecies = getPokemonSpecies(psd.PokemonSpecies.Name)
	}

	return gender
}

// GetGrowthRate returns a GrowthRate object
// Acceptable search queries:
// ID, String
func GetGrowthRate(query string) GrowthRate {
	searchType := idOrString(query)
	growthRate := GrowthRate{}
	if checkCache("growth-rate", query, searchType) {
		mapstructure.Decode(getDataCache("growth-rate", query, searchType, growthRate), &growthRate)
	} else {
		dataInterface = getDataAPI("growth-rate", query, growthRate)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &growthRate)
			storeData(false, growthRate, "growth-rate", growthRate.ID, growthRate.Name)
		}
	}
	for _, ps := range growthRate.PokemonSpecies {
		growthRate.GetPokemonSpecies = append(growthRate.GetPokemonSpecies, getPokemonSpecies(ps.Name))
	}
	return growthRate
}

// GetNature returns a Nature object
// Acceptable search queries:
// ID, String
func GetNature(query string) Nature {
	searchType := idOrString(query)
	nature := Nature{}
	if checkCache("nature", query, searchType) {
		mapstructure.Decode(getDataCache("nature", query, searchType, nature), &nature)
	} else {
		dataInterface = getDataAPI("nature", query, nature)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &nature)
			storeData(false, nature, "nature", nature.ID, nature.Name)
		}
	}
	nature.GetDecreasedStat = getStat(nature.DecreasedStat.Name)
	nature.GetIncreasedStat = getStat(nature.IncreasedStat.Name)
	nature.GetHatesFlavor = getBerryFlavor(nature.HatesFlavor.Name)
	nature.GetLikesFlavor = getBerryFlavor(nature.LikesFlavor.Name)
	for i, psc := range nature.PokeathlonStatChanges {
		nature.PokeathlonStatChanges[i].GetPokeathlonStat = getPokeathlonStat(psc.PokeathlonStat.Name)
	}
	for i, mbsp := range nature.MoveBattleStylePreferences {
		nature.MoveBattleStylePreferences[i].GetMoveBattleStyle = getMoveBattleStyle(mbsp.MoveBattleStyle.Name)
	}
	return nature
}

// GetPokeathlonStat returns a PokeathlonStat object
// Acceptable search queries:
// ID, String
func GetPokeathlonStat(query string) PokeathlonStat {
	searchType := idOrString(query)
	pokeathlonStat := PokeathlonStat{}
	if checkCache("pokeathlon-stat", query, searchType) {
		mapstructure.Decode(getDataCache("pokeathlon-stat", query, searchType, pokeathlonStat), &pokeathlonStat)
	} else {
		dataInterface = getDataAPI("pokeathlon-stat", query, pokeathlonStat)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokeathlonStat)
			storeData(false, pokeathlonStat, "pokeathlon-stat", pokeathlonStat.ID, pokeathlonStat.Name)
		}
	}
	for i := range pokeathlonStat.AffectingNatures {
		for ii, in := range pokeathlonStat.AffectingNatures[i].Increase {
			pokeathlonStat.AffectingNatures[i].Increase[ii].GetNature = getNature(in.Nature.Name)
		}
		for ii, de := range pokeathlonStat.AffectingNatures[i].Decrease {
			pokeathlonStat.AffectingNatures[i].Decrease[ii].GetNature = getNature(de.Nature.Name)
		}
	}
	return pokeathlonStat
}

// GetPokemon returns a Pokemon object
// Acceptable search queries:
// ID, String
func GetPokemon(query string) Pokemon {
	searchType := idOrString(query)
	pokemon := Pokemon{}
	if checkCache("pokemon", query, searchType) {
		mapstructure.Decode(getDataCache("pokemon", query, searchType, pokemon), &pokemon)
	} else {
		dataInterface = getDataAPI("pokemon", query, pokemon)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokemon)
			storeData(false, pokemon, "pokemon", pokemon.ID, pokemon.Name)
		}
	}
	for _, form := range pokemon.Forms {
		pokemon.GetForms = append(pokemon.GetForms, getPokemonForm(form.Name))
	}
	pokemon.GetSpecies = getPokemonSpecies(pokemon.Species.Name)

	for i, ability := range pokemon.Abilities {
		pokemon.Abilities[i].GetAbility = getAbility(ability.Ability.Name)
	}

	for i, t := range pokemon.Types {
		pokemon.Types[i].GetType = getType(t.Type.Name)
	}
	for i, item := range pokemon.HeldItems {
		pokemon.HeldItems[i].GetItem = getItem(item.Item.Name)
		for ii, v := range item.VersionDetails {
			pokemon.HeldItems[i].VersionDetails[ii].GetVersion = getVersion(v.Version.Name)
		}
	}

	for i, move := range pokemon.Moves {
		pokemon.Moves[i].GetMove = getMove(move.Move.Name)
		for ii, vgd := range move.VersionGroupDetails {
			pokemon.Moves[i].VersionGroupDetails[ii].GetMoveLearnMethod = getMoveLearnMethod(vgd.MoveLearnMethod.Name)
			pokemon.Moves[i].VersionGroupDetails[ii].GetVersionGroup = getVersionGroup(vgd.VersionGroup.Name)
		}
	}

	for i, stat := range pokemon.Stats {
		pokemon.Stats[i].GetStat = getStat(stat.Stat.Name)
	}

	pokemon.GetLocationAreaEncounters = getLocationArea(strconv.Itoa(getID(pokemon.LocationAreaEncounters)))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.BackDefault))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.BackFemale))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.BackShiny))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.BackShinyFemale))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.FrontDefault))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.FrontFemale))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.FrontShiny))
	pokemon.Sprites.GetSprites = append(pokemon.Sprites.GetSprites, getPokemonSprite(pokemon.Sprites.FrontShinyFemale))
	return pokemon
}

// GetPokemonColor returns a PokemonColor object
// Acceptable search queries:
// ID, String
func GetPokemonColor(query string) PokemonColor {
	searchType := idOrString(query)
	pokemonColor := PokemonColor{}
	if checkCache("pokemon-color", query, searchType) {
		mapstructure.Decode(getDataCache("pokemon-color", query, searchType, pokemonColor), &pokemonColor)
	} else {
		dataInterface = getDataAPI("pokemon-color", query, pokemonColor)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokemonColor)
			storeData(false, pokemonColor, "pokemon-color", pokemonColor.ID, pokemonColor.Name)
		}
	}
	for _, ps := range pokemonColor.PokemonSpecies {
		pokemonColor.GetPokemonSpecies = append(pokemonColor.GetPokemonSpecies, getPokemonSpecies(ps.Name))
	}
	return pokemonColor
}

// GetPokemonForm returns a PokemonForm object
// Acceptable search queries:
// ID, String
func GetPokemonForm(query string) PokemonForm {
	searchType := idOrString(query)
	pokemonForm := PokemonForm{}
	if checkCache("pokemon-form", query, searchType) {
		mapstructure.Decode(getDataCache("pokemon-form", query, searchType, pokemonForm), &pokemonForm)
	} else {
		dataInterface = getDataAPI("pokemon-form", query, pokemonForm)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokemonForm)
			storeData(false, pokemonForm, "pokemon-form", pokemonForm.ID, pokemonForm.Name)
		}
	}
	pokemonForm.GetPokemon = getPokemon(pokemonForm.Pokemon.Name)
	pokemonForm.GetVersionGroup = getVersionGroup(pokemonForm.VersionGroup.Name)
	pokemonForm.Sprites.GetSprites = append(pokemonForm.Sprites.GetSprites, getPokemonSprite(pokemonForm.Sprites.FrontDefault))
	pokemonForm.Sprites.GetSprites = append(pokemonForm.Sprites.GetSprites, getPokemonSprite(pokemonForm.Sprites.FrontShiny))
	pokemonForm.Sprites.GetSprites = append(pokemonForm.Sprites.GetSprites, getPokemonSprite(pokemonForm.Sprites.BackDefault))
	pokemonForm.Sprites.GetSprites = append(pokemonForm.Sprites.GetSprites, getPokemonSprite(pokemonForm.Sprites.BackShiny))
	return pokemonForm
}

// GetPokemonHabitat returns a PokemonHabitat object
// Acceptable search queries:
// ID, String
func GetPokemonHabitat(query string) PokemonHabitat {
	searchType := idOrString(query)
	pokemonHabitat := PokemonHabitat{}
	if checkCache("pokemon-habitat", query, searchType) {
		mapstructure.Decode(getDataCache("pokemon-habitat", query, searchType, pokemonHabitat), &pokemonHabitat)
	} else {
		dataInterface = getDataAPI("pokemon-habitat", query, pokemonHabitat)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokemonHabitat)
			storeData(false, pokemonHabitat, "pokemon-habitat", pokemonHabitat.ID, pokemonHabitat.Name)
		}
	}
	for _, ps := range pokemonHabitat.PokemonSpecies {
		pokemonHabitat.GetPokemonSpecies = append(pokemonHabitat.GetPokemonSpecies, getPokemonSpecies(ps.Name))
	}

	return pokemonHabitat
}

// GetPokemonShape returns a PokemonShape object
// Acceptable search queries:
// ID, String
func GetPokemonShape(query string) PokemonShape {
	searchType := idOrString(query)
	pokemonShape := PokemonShape{}
	if checkCache("pokemon-shape", query, searchType) {
		mapstructure.Decode(getDataCache("pokemon-shape", query, searchType, pokemonShape), &pokemonShape)
	} else {
		dataInterface = getDataAPI("pokemon-shape", query, pokemonShape)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokemonShape)
			storeData(false, pokemonShape, "pokemon-shape", pokemonShape.ID, pokemonShape.Name)
		}
	}
	for _, ps := range pokemonShape.PokemonSpecies {
		pokemonShape.GetPokemonSpecies = append(pokemonShape.GetPokemonSpecies, getPokemonSpecies(ps.Name))
	}
	for i, an := range pokemonShape.AwesomeNames {
		pokemonShape.AwesomeNames[i].GetLanguage = getLanguage(an.Language.Name)
	}
	return pokemonShape
}

// GetPokemonSpecies returns a PokemonSpecies object
// Acceptable search queries:
// ID, String
func GetPokemonSpecies(query string) PokemonSpecies {
	searchType := idOrString(query)
	pokemonSpecies := PokemonSpecies{}
	if checkCache("pokemon-species", query, searchType) {
		mapstructure.Decode(getDataCache("pokemon-species", query, searchType, pokemonSpecies), &pokemonSpecies)
	} else {
		dataInterface = getDataAPI("pokemon-species", query, pokemonSpecies)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &pokemonSpecies)
			storeData(false, pokemonSpecies, "pokemon-species", pokemonSpecies.ID, pokemonSpecies.Name)
		}
	}
	pokemonSpecies.GetGrowthRate = getGrowthRate(pokemonSpecies.GrowthRate.Name)
	for _, egg := range pokemonSpecies.EggGroups {
		pokemonSpecies.GetEggGroups = append(pokemonSpecies.GetEggGroups, getEggGroup(egg.Name))
	}
	pokemonSpecies.GetColor = getPokemonColor(pokemonSpecies.Color.Name)
	pokemonSpecies.GetShape = getPokemonShape(pokemonSpecies.Shape.Name)
	pokemonSpecies.GetEvolvesFromSpecies = getPokemonSpecies(pokemonSpecies.EvolvesFromSpecies.Name)
	pokemonSpecies.GetEvolutionChain = getEvolutionChain(getID(pokemonSpecies.EvolutionChain.URL))
	pokemonSpecies.GetHabitat = getPokemonHabitat(pokemonSpecies.Habitat.Name)
	pokemonSpecies.GetGeneration = getGeneration(pokemonSpecies.Generation.Name)
	for i, g := range pokemonSpecies.Genera {
		pokemonSpecies.Genera[i].GetLanguage = getLanguage(g.Language.Name)
	}
	for i, pd := range pokemonSpecies.PokedexNumbers {
		pokemonSpecies.PokedexNumbers[i].GetPokedex = getPokedex(pd.Pokedex.Name)
	}
	for i, ppe := range pokemonSpecies.PalParkEncounters {
		pokemonSpecies.PalParkEncounters[i].GetArea = getPalParkArea(ppe.Area.Name)
	}
	for i, v := range pokemonSpecies.Varieties {
		pokemonSpecies.Varieties[i].GetPokemon = getPokemon(v.Pokemon.Name)
	}
	return pokemonSpecies
}

// GetStat returns a Stat object
// Acceptable search queries:
// ID, String
func GetStat(query string) Stat {
	searchType := idOrString(query)
	stat := Stat{}
	if checkCache("stat", query, searchType) {
		mapstructure.Decode(getDataCache("stat", query, searchType, stat), &stat)
	} else {
		dataInterface = getDataAPI("stat", query, stat)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &stat)
			storeData(false, stat, "stat", stat.ID, stat.Name)
		}
	}
	for _, c := range stat.Characteristics {
		stat.GetCharacteristics = append(stat.GetCharacteristics, getCharacteristic(getID(c.URL)))
	}
	stat.GetMoveDamageClass = getMoveDamageClass(stat.MoveDamageClass.Name)

	for i, in := range stat.AffectingMoves.Increase {
		stat.AffectingMoves.Increase[i].GetMove = getMove(in.Move.Name)
	}

	for i, de := range stat.AffectingMoves.Decrease {
		stat.AffectingMoves.Increase[i].GetMove = getMove(de.Move.Name)
	}
	for _, in := range stat.AffectingNatures.Increase {
		stat.AffectingNatures.GetIncrease = append(stat.AffectingNatures.GetIncrease, getNature(in.Name))
	}
	for _, de := range stat.AffectingNatures.Decrease {
		stat.AffectingNatures.GetDecrease = append(stat.AffectingNatures.GetDecrease, getNature(de.Name))
	}
	return stat
}

// GetType returns a Type object
// Acceptable search queries:
// ID, String
func GetType(query string) Type {
	searchType := idOrString(query)
	tType := Type{}
	if checkCache("type", query, searchType) {
		mapstructure.Decode(getDataCache("type", query, searchType, tType), &tType)
	} else {
		dataInterface = getDataAPI("type", query, tType)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &tType)
			storeData(false, tType, "type", tType.ID, tType.Name)
		}
	}
	tType.GetGeneration = getGeneration(tType.Generation.Name)
	tType.GetMoveDamageClass = getMoveDamageClass(tType.MoveDamageClass.Name)
	for _, move := range tType.Moves {
		tType.GetMoves = append(tType.GetMoves, getMove(move.Name))
	}

	for i, p := range tType.Pokemon {
		tType.Pokemon[i].GetPokemon = getPokemon(p.Pokemon.Name)
	}

	for _, ndt := range tType.DamageRelations.NoDamageTo {
		tType.DamageRelations.GetNoDamageTo = append(tType.DamageRelations.GetNoDamageTo, getType(ndt.Name))
	}
	for _, hdt := range tType.DamageRelations.HalfDamageTo {
		tType.DamageRelations.GetHalfDamageTo = append(tType.DamageRelations.GetHalfDamageTo, getType(hdt.Name))
	}
	for _, ddt := range tType.DamageRelations.DoubleDamageTo {
		tType.DamageRelations.GetDoubleDamageTo = append(tType.DamageRelations.GetDoubleDamageTo, getType(ddt.Name))
	}

	for _, ndf := range tType.DamageRelations.NoDamageFrom {
		tType.DamageRelations.GetNoDamageFrom = append(tType.DamageRelations.GetNoDamageFrom, getType(ndf.Name))
	}
	for _, hdf := range tType.DamageRelations.HalfDamageFrom {
		tType.DamageRelations.GetHalfDamageFrom = append(tType.DamageRelations.GetHalfDamageFrom, getType(hdf.Name))
	}
	for _, ddf := range tType.DamageRelations.DoubleDamageFrom {
		tType.DamageRelations.GetDoubleDamageFrom = append(tType.DamageRelations.GetDoubleDamageFrom, getType(ddf.Name))
	}
	return tType
}
