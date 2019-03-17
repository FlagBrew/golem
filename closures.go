package pokeapigo

// Closures class is used for getting data from namedAPIResources

// berry.go
func getBerry(name string) func() Berry {
	return func() Berry {
		b := GetBerry(name)
		return b
	}
}

func getBerryFlavor(name string) func() BerryFlavor {
	return func() BerryFlavor {
		bf := GetBerryFlavor(name)
		return bf
	}
}

// contest.go
func getContestType(name string) func() ContestType {
	return func() ContestType {
		ct := GetContestType(name)
		return ct
	}
}

func getContestEffect(id int) func() ContestEffect {
	return func() ContestEffect {
		cf := GetContestEffect(id)
		return cf
	}
}

func getSuperContestEffect(id int) func() SuperContestEffect {
	return func() SuperContestEffect {
		sce := GetSuperContestEffect(id)
		return sce
	}
}

// encounter.go
func getEncounterConditionValue(name string) func() EncounterConditionValue {
	return func() EncounterConditionValue {
		ecf := GetEncounterConditionValue(name)
		return ecf
	}
}

func getEncounterCondition(name string) func() EncounterCondition {
	return func() EncounterCondition {
		ec := GetEncounterCondition(name)
		return ec
	}
}

func getEncounterMethod(name string) func() EncounterMethod {
	return func() EncounterMethod {
		em := GetEncounterMethod(name)
		return em
	}
}

// evolution.go
func getEvolutionTrigger(name string) func() EvolutionTrigger {
	return func() EvolutionTrigger {
		et := GetEvolutionTrigger(name)
		return et
	}
}

func getEvolutionChain(id int) func() EvolutionChain {
	return func() EvolutionChain {
		ec := GetEvolutionChain(id)
		return ec
	}
}

// game.go

func getVersionGroup(name string) func() VersionGroup {
	return func() VersionGroup {
		vg := GetVersionGroup(name)
		return vg
	}
}

func getGeneration(name string) func() Generation {
	return func() Generation {
		g := GetGeneration(name)
		return g
	}
}

func getPokedex(name string) func() Pokedex {
	return func() Pokedex {
		p := GetPokedex(name)
		return p
	}
}

func getVersion(name string) func() Version {
	return func() Version {
		v := GetVersion(name)
		return v
	}
}

// item.go
func getItem(name string) func() Item {
	return func() Item {
		i := GetItem(name)
		return i
	}
}

func getItemFlingEffect(name string) func() ItemFlingEffect {
	return func() ItemFlingEffect {
		ife := GetItemFlingEffect(name)
		return ife
	}
}

func getItemAttribute(name string) func() ItemAttribute {
	return func() ItemAttribute {
		ia := GetItemAttribute(name)
		return ia
	}
}

func getItemPocket(name string) func() ItemPocket {
	return func() ItemPocket {
		ip := GetItemPocket(name)
		return ip
	}
}

func getItemCategory(name string) func() ItemCategory {
	return func() ItemCategory {
		ic := GetItemCategory(name)
		return ic
	}
}

func getItemSprite(url string) func() string {
	return func() string {
		s := getSprite(url)
		return s
	}
}

// location.go

func getLocation(name string) func() Location {
	return func() Location {
		l := GetLocation(name)
		return l
	}
}

func getRegion(name string) func() Region {
	return func() Region {
		r := GetRegion(name)
		return r
	}
}

func getLocationArea(name string) func() LocationArea {
	return func() LocationArea {
		la := GetLocationArea(name)
		return la
	}
}

func getPalParkArea(name string) func() PalParkArea {
	return func() PalParkArea {
		ppa := GetPalParkArea(name)
		return ppa
	}
}

// machine.go

// move.go
func getMove(name string) func() Move {
	return func() Move {
		m := GetMove(name)
		return m
	}
}

func getMoveLearnMethod(name string) func() MoveLearnMethod {
	return func() MoveLearnMethod {
		mlm := GetMoveLearnMethod(name)
		return mlm
	}
}

func getMoveDamageClass(name string) func() MoveDamageClass {
	return func() MoveDamageClass {
		mdc := GetMoveDamageClass(name)
		return mdc
	}
}

func getMoveTarget(name string) func() MoveTarget {
	return func() MoveTarget {
		mt := GetMoveTarget(name)
		return mt
	}
}

func getMoveAilment(name string) func() MoveAilment {
	return func() MoveAilment {
		ma := GetMoveAilment(name)
		return ma
	}
}

func getMoveBattleStyle(name string) func() MoveBattleStyle {
	return func() MoveBattleStyle {
		mbs := GetMoveBattleStyle(name)
		return mbs
	}
}

func getMoveCategory(name string) func() MoveCategory {
	return func() MoveCategory {
		mc := GetMoveCategory(name)
		return mc
	}
}

// pokemon.go

func getType(name string) func() Type {
	return func() Type {
		t := GetType(name)
		return t
	}
}

func getPokemonSpecies(name string) func() PokemonSpecies {
	return func() PokemonSpecies {
		ps := GetPokemonSpecies(name)
		return ps
	}
}

func getAbility(name string) func() Ability {
	return func() Ability {
		a := GetAbility(name)
		return a
	}
}

func getStat(name string) func() Stat {
	return func() Stat {
		s := GetStat(name)
		return s
	}
}

func getPokemon(name string) func() Pokemon {
	return func() Pokemon {
		p := GetPokemon(name)
		return p
	}
}

func getPokeathlonStat(name string) func() PokeathlonStat {
	return func() PokeathlonStat {
		ps := GetPokeathlonStat(name)
		return ps
	}
}

func getNature(name string) func() Nature {
	return func() Nature {
		n := GetNature(name)
		return n
	}
}

func getPokemonForm(name string) func() PokemonForm {
	return func() PokemonForm {
		pf := GetPokemonForm(name)
		return pf
	}
}

func getGrowthRate(name string) func() GrowthRate {
	return func() GrowthRate {
		gr := GetGrowthRate(name)
		return gr
	}
}

func getEggGroup(name string) func() EggGroup {
	return func() EggGroup {
		eg := GetEggGroup(name)
		return eg
	}
}

func getPokemonColor(name string) func() PokemonColor {
	return func() PokemonColor {
		pc := GetPokemonColor(name)
		return pc
	}
}

func getPokemonShape(name string) func() PokemonShape {
	return func() PokemonShape {
		ps := GetPokemonShape(name)
		return ps
	}
}

func getPokemonHabitat(name string) func() PokemonHabitat {
	return func() PokemonHabitat {
		poh := GetPokemonHabitat(name)
		return poh
	}
}

func getCharacteristic(id int) func() Characteristic {
	return func() Characteristic {
		c := GetCharacteristic(id)
		return c
	}
}

func getPokemonSprite(url string) func() string {
	return func() string {
		s := getSprite(url)
		return s
	}
}

// utils.go
func getLanguage(name string) func() Language {
	return func() Language {
		l := GetLanguage(name)
		return l
	}
}
