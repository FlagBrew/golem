package pokeapigo

// Berries Section
// https://pokeapi.co/docs/v2.html/#berries-section

// Berry ...
// Berries are small fruits that can provide HP and status condition restoration, stat enhancement, and even damage negation when eaten by Pokémon. Check out Bulbapedia for greater detail.
// The berry schema is found at https://pokeapi.co/docs/v2.html/#berries
type Berry struct {
	ID                 int                  `json:"id" mapstructure:"id"`
	Name               string               `json:"name" mapstructure:"name"`
	GrowthTime         int                  `json:"growth_time" mapstructure:"growth_time"`
	MaxHarvest         int                  `json:"max_harvest" mapstructure:"max_harvest"`
	NaturalGiftPower   int                  `json:"natural_gift_power" mapstructure:"natural_gift_power"`
	Size               int                  `json:"size" mapstructure:"size"`
	Smoothness         int                  `json:"smoothness" mapstructure:"smoothness"`
	SoilDryness        int                  `json:"soil_dryness" mapstructure:"soil_dryness"`
	Firmness           namedAPIResource     `json:"firmness" mapstructure:"firmness"`
	GetFirmness        func() BerryFirmness `json:"-" mapstructure:"-"`
	Flavors            []berryFlavor        `json:"flavors" mapstructure:"flavors"`
	GetItem            func() Item          `json:"-" mapstructure:"-"`
	GetNaturalGiftType func() Type          `json:"-" mapstructure:"-"`
	Item               namedAPIResource     `json:"item" mapstructure:"item"`
	NaturalGiftType    namedAPIResource     `json:"natural_gift_type" mapstructure:"natural_gift_type"`
}

type berryFlavor struct {
	Potency   int                `json:"potency" mapstructure:"potency"`
	Flavor    namedAPIResource   `json:"flavor" mapstructure:"flavor"`
	GetFlavor func() BerryFlavor `json:"-" mapstructure:"-"`
}

// BerryFirmness ...
// Berries can be soft or hard. Check out Bulbapedia for greater detail.
// The BerryFirmness schema is found at https://pokeapi.co/docs/v2.html/#berry-firmnesses
type BerryFirmness struct {
	ID         int                `json:"id" mapstructure:"id"`
	Name       string             `json:"name" mapstructure:"name"`
	Berries    []namedAPIResource `json:"berries" mapstructure:"berries"`
	GetBerries []func() Berry     `json:"-" mapstructure:"-"`
	Names      []name             `json:"names" mapstructure:"names"`
}

// BerryFlavor ...
// Flavors determine whether a Pokémon will benefit or suffer from eating a berry based on their nature. Check out Bulbapedia for greater detail.
// The BerryFlavor schema can be found at https://pokeapi.co/docs/v2.html/#berry-flavors
type BerryFlavor struct {
	ID             int                `json:"id" mapstructure:"id"`
	Name           string             `json:"name" mapstructure:"name"`
	Berries        []flavorBerry      `json:"berries" mapstructure:"berries"`
	ContestType    namedAPIResource   `json:"contest_type" mapstructure:"contest_type"`
	GetContestType func() ContestType `json:"-" mapstructure:"-"`
	Names          []name             `json:"names" mapstructure:"names"`
}

type flavorBerry struct {
	Potency  int              `json:"potency" mapstructure:"potency"`
	Berry    namedAPIResource `json:"berry" mapstructure:"berry"`
	GetBerry func() Berry     `json:"-" mapstructure:"-"`
}

// Contests Section
// https://pokeapi.co/docs/v2.html/#contests-section

// ContestType ...
// Contest types are categories judges used to weigh a Pokémon's condition in Pokémon contests. Check out Bulbapedia for greater detail.
// The ContestType schema can be found at https://pokeapi.co/docs/v2.html/#contest-types
type ContestType struct {
	ID             int                `json:"id" mapstructure:"id"`
	Name           string             `json:"name" mapstructure:"name"`
	GetBerryFlavor func() BerryFlavor `json:"-" mapstructure:"-"`
	BerryFlavor    namedAPIResource   `json:"berry_flavor" mapstructure:"berry_flavor"`
	Names          []contestName      `json:"names" mapstructure:"names"`
}

type contestName struct {
	Name        string           `json:"name" mapstructure:"name"`
	Color       string           `json:"color" mapstructure:"color"`
	GetLanguage func() Language  `json:"-" mapstructure:"-"`
	Language    namedAPIResource `json:"language" mapstructure:"language"`
}

// ContestEffect ...
// Contest effects refer to the effects of moves when used in contests.
// The ContestEffect schema can be found at https://pokeapi.co/docs/v2.html/#contest-effects
type ContestEffect struct {
	ID                int          `json:"id" mapstructure:"id"`
	Appeal            int          `json:"appeal" mapstructure:"appeal"`
	Jam               int          `json:"jam" mapstructure:"jam"`
	EffectEntries     []effect     `json:"effect_entries" mapstructure:"effect_entries"`
	FlavorTextEntries []flavorText `json:"flavor_text_entries" mapstructure:"flavor_text_entries"`
}

// SuperContestEffect ...
// Super contest effects refer to the effects of moves when used in super contests.
// The SuperContestEffect can be found at https://pokeapi.co/docs/v2.html/#super-contest-effects
type SuperContestEffect struct {
	ID                int                `json:"id" mapstructure:"id"`
	Appeal            int                `json:"appeal" mapstructure:"appeal"`
	FlavorTextEntries []flavorText       `json:"flavor_text_entries" mapstructure:"flavor_text_entries"`
	GetMoves          []func() Move      `json:"-" mapstructure:"-"`
	Moves             []namedAPIResource `json:"moves" mapstructure:"moves"`
}

// Encounters Section
// https://pokeapi.co/docs/v2.html/#encounters-section

// EncounterMethod ...
// Methods by which the player might can encounter Pokémon in the wild, e.g., walking in tall grass. Check out Bulbapedia for greater detail.
// The EncounterMethod schema can be found at https://pokeapi.co/docs/v2.html/#encounter-methods
type EncounterMethod struct {
	ID    int    `json:"id" mapstructure:"id"`
	Name  string `json:"name" mapstructure:"name"`
	Order int    `json:"order" mapstructure:"order"`
	Names []name `json:"names" mapstructure:"names"`
}

// EncounterCondition ...
// Conditions which affect what pokemon might appear in the wild, e.g., day or night.
// The EncounterCondition schema can be found at https://pokeapi.co/docs/v2.html/#encounter-conditions
type EncounterCondition struct {
	ID        int                              `json:"id" mapstructure:"id"`
	Name      string                           `json:"name" mapstructure:"name"`
	GetValues []func() EncounterConditionValue `json:"-" mapstructure:"-"`
	Values    []namedAPIResource               `json:"values" mapstructure:"values"`
	Names     []name                           `json:"names" mapstructure:"names"`
}

// EncounterConditionValue ...
// Encounter condition values are the various states that an encounter condition can have, i.e., time of day can be either day or night.
// The EncounterConditionValue schema can be found at https://pokeapi.co/docs/v2.html/#encounter-condition-values
type EncounterConditionValue struct {
	ID           int                       `json:"id" mapstructure:"id"`
	Name         string                    `json:"name" mapstructure:"name"`
	GetCondition func() EncounterCondition `json:"-" mapstructure:"-"`
	Condition    namedAPIResource          `json:"condition" mapstructure:"condition"`
	Names        []name                    `json:"names" mapstructure:"names"`
}

// Evolutions Section
// https://pokeapi.co/docs/v2.html/#evolution-section

// EvolutionChain ...
// Evolution chains are essentially family trees. They start with the lowest stage within a family and detail evolution conditions for each as well as Pokémon they can evolve into up through the hierarchy.
// The EvolutionChain schema can be found at https://pokeapi.co/docs/v2.html/#evolution-chains
type EvolutionChain struct {
	ID                 int              `json:"id" mapstructure:"id"`
	GetBabyTriggerItem func() Item      `json:"-" mapstructure:"-"`
	BabyTriggerItem    namedAPIResource `json:"baby_trigger_item" mapstructure:"baby_trigger_item"`
	Chain              chainLink        `json:"chain" mapstructure:"chain"`
}

type chainLink struct {
	IsBaby           bool                  `json:"is_baby" mapstructure:"is_baby"`
	GetSpecies       func() PokemonSpecies `json:"-" mapstructure:"-"`
	Species          namedAPIResource      `json:"species" mapstructure:"species"`
	EvolutionDetails []evolutionDetail     `json:"evolution_details" mapstructure:"evolution_details"`
	EvolvesTo        []chainLink           `json:"evolves_to" mapstructure:"evolves_to"`
}

type evolutionDetail struct {
	GetItem               func() Item             `json:"-" mapstructure:"-"`
	Item                  namedAPIResource        `json:"item" mapstructure:"item"`
	GetTrigger            func() EvolutionTrigger `json:"-" mapstructure:"-"`
	Trigger               namedAPIResource        `json:"trigger" mapstructure:"trigger"`
	Gender                int                     `json:"gender" mapstructure:"gender"`
	GetHeldItem           func() Item             `json:"-" mapstructure:"-"`
	HeldItem              namedAPIResource        `json:"held_item" mapstructure:"held_item"`
	GetKnownMove          func() Move             `json:"-" mapstructure:"-"`
	KnownMove             namedAPIResource        `json:"known_move" mapstructure:"known_move"`
	GetKnownMoveType      func() Type             `json:"-" mapstructure:"-"`
	KnownMoveType         namedAPIResource        `json:"known_move_type" mapstructure:"known_move_type"`
	GetLocation           func() Location         `json:"-" mapstructure:"-"`
	Location              namedAPIResource        `json:"location" mapstructure:"location"`
	MinLevel              int                     `json:"min_level" mapstructure:"min_level"`
	MinHappiness          int                     `json:"min_happiness" mapstructure:"min_happiness"`
	MinBeauty             int                     `json:"min_beauty" mapstructure:"min_beauty"`
	MinAffection          int                     `json:"min_affection" mapstructure:"min_affection"`
	NeedsOverworldRain    bool                    `json:"needs_overworld_rain" mapstructure:"needs_overworld_rain"`
	GetPartySpecies       func() PokemonSpecies   `json:"-" mapstructure:"-"`
	PartySpecies          namedAPIResource        `json:"party_species" mapstructure:"party_species"`
	GetPartyType          func() Type             `json:"-" mapstructure:"-"`
	PartyType             namedAPIResource        `json:"party_type" mapstructure:"party_type"`
	RelativePhysicalStats int                     `json:"relative_physical_stats" mapstructure:"relative_physical_stats"`
	TimeOfDay             string                  `json:"time_of_day" mapstructure:"time_of_day"`
	GetTradeSpecies       func() PokemonSpecies   `json:"-" mapstructure:"-"`
	TradeSpecies          namedAPIResource        `json:"trade_species" mapstructure:"trade_species"`
	TurnUpsideDown        bool                    `json:"turn_upside_down" mapstructure:"turn_upside_down"`
}

// EvolutionTrigger ...
// Evolution triggers are the events and conditions that cause a Pokémon to evolve. Check out Bulbapedia for greater detail.
// The EvolutionTrigger schema can be found at https://pokeapi.co/docs/v2.html/#evolution-triggers
type EvolutionTrigger struct {
	ID                int                     `json:"id" mapstructure:"id"`
	Name              string                  `json:"name" mapstructure:"name"`
	Names             []name                  `json:"names" mapstructure:"names"`
	GetPokemonSpecies []func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    []namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

// Games Section
// https://pokeapi.co/docs/v2.html/#games-section

// Generation ...
// A generation is a grouping of the Pokémon games that separates them based on the Pokémon they include. In each generation, a new set of Pokémon, Moves, Abilities and Types that did not exist in the previous generation are released.
// The Generation schema can be found at https://pokeapi.co/docs/v2.html/#generations
type Generation struct {
	ID                int                     `json:"id" mapstructure:"id"`
	Name              string                  `json:"name" mapstructure:"name"`
	GetAbilities      []func() Ability        `json:"-" mapstructure:"-"`
	Abilities         []namedAPIResource      `json:"abilities" mapstructure:"abilities"`
	Names             []name                  `json:"names" mapstructure:"names"`
	GetMainRegion     func() Region           `json:"-" mapstructure:"-"`
	MainRegion        namedAPIResource        `json:"main_region" mapstructure:"main_region"`
	GetMoves          []func() Move           `json:"-" mapstructure:"-"`
	Moves             []namedAPIResource      `json:"moves" mapstructure:"moves"`
	GetPokemonSpecies []func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    []namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
	GetTypes          []func() Type           `json:"-" mapstructure:"-"`
	Types             []namedAPIResource      `json:"types" mapstructure:"types"`
	GetVersionGroups  []func() VersionGroup   `json:"-" mapstructure:"-"`
	VersionGroups     []namedAPIResource      `json:"version_groups" mapstructure:"version_groups"`
}

// Pokedex ...
// A Pokédex is a handheld electronic encyclopedia device; one which is capable of recording and retaining information of the various Pokémon in a given region with the exception of the national dex and some smaller dexes related to portions of a region. See Bulbapedia for greater detail.
// The Pokedex schema can be found at https://pokeapi.co/docs/v2.html/#pokedexes
type Pokedex struct {
	ID               int                   `json:"id" mapstructure:"id"`
	Name             string                `json:"name" mapstructure:"name"`
	IsMainSeries     bool                  `json:"is_main_series" mapstructure:"is_main_series"`
	Descriptions     []description         `json:"descriptions" mapstructure:"descriptions"`
	Names            []name                `json:"names" mapstructure:"names"`
	PokemonEntries   []pokemonEntry        `json:"pokemon_entries" mapstructure:"pokemon_entries"`
	GetRegion        func() Region         `json:"-" mapstructure:"-"`
	Region           namedAPIResource      `json:"region" mapstructure:"region"`
	GetVersionGroups []func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroups    []namedAPIResource    `json:"version_groups" mapstructure:"version_groups"`
}

type pokemonEntry struct {
	EntryNumber       int                   `json:"entry_number" mapstructure:"entry_number"`
	GetPokemonSpecies func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

// Version ...
// Versions of the games, e.g., Red, Blue or Yellow.
// The Version schema can be found at https://pokeapi.co/docs/v2.html/#version
type Version struct {
	ID              int                 `json:"id" mapstructure:"id"`
	Name            string              `json:"name" mapstructure:"name"`
	Names           []name              `json:"names" mapstructure:"names"`
	GetVersionGroup func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroup    namedAPIResource    `json:"version_group" mapstructure:"version_group"`
}

// VersionGroup ...
// Version groups categorize highly similar versions of the games.
// The VersionGroup schema can be found at https://pokeapi.co/docs/v2.html/#version-groups
type VersionGroup struct {
	ID                  int                      `json:"id" mapstructure:"id"`
	Name                string                   `json:"name" mapstructure:"name"`
	Order               int                      `json:"order" mapstructure:"order"`
	Generation          namedAPIResource         `json:"generation" mapstructure:"generation"`
	GetMoveLearnMethods []func() MoveLearnMethod `json:"-" mapstructure:"-"`
	MoveLearnMethods    []namedAPIResource       `json:"move_learn_methods" mapstructure:"move_learn_methods"`
	GetPokedexes        []func() Pokedex         `json:"-" mapstructure:"-"`
	GetVersions         []func() Version         `json:"-" mapstructure:"-"`
	GetGeneration       func() Generation        `json:"-" mapstructure:"-"`
	Pokedexes           []namedAPIResource       `json:"pokedexes" mapstructure:"pokedexes"`
	Regions             []namedAPIResource       `json:"regions" mapstructure:"regions"`
	Versions            []namedAPIResource       `json:"versions" mapstructure:"versions"`
}

// Items Section
// https://pokeapi.co/docs/v2.html/#items-section

// Item ...
// An item is an object in the games which the player can pick up, keep in their bag, and use in some manner. They have various uses, including healing, powering up, helping catch Pokémon, or to access a new area.
// The Item schema can be found at https://pokeapi.co/docs/v2.html/#item
type Item struct {
	ID                int                      `json:"id" mapstructure:"id"`
	Name              string                   `json:"name" mapstructure:"name"`
	Cost              int                      `json:"cost" mapstructure:"cost"`
	FlingPower        int                      `json:"fling_power" mapstructure:"fling_power"`
	GetFlingEffect    func() ItemFlingEffect   `json:"-" mapstructure:"-"`
	FlingEffect       namedAPIResource         `json:"fling_effect" mapstructure:"fling_effect"`
	GetAttributes     []func() ItemAttribute   `json:"-" mapstructure:"-"`
	Attributes        []namedAPIResource       `json:"attributes" mapstructure:"attributes"`
	GetCategory       func() ItemCategory      `json:"-" mapstructure:"-"`
	Category          namedAPIResource         `json:"category" mapstructure:"category"`
	EffectEntries     []verboseEffect          `json:"effect_entries" mapstructure:"effect_entries"`
	FlavorTextEntries []versionGroupFlavorText `json:"flavor_text_entries" mapstructure:"flavor_text_entries"`
	GameIndices       []generationGameIndex    `json:"game_indices" mapstructure:"game_indices"`
	Sprites           itemSprite               `json:"sprites" mapstructure:"sprites"`
	HeldByPokemon     []itemHolderPokemon      `json:"held_by_pokemon" mapstructure:"held_by_pokemon"`
	GetBabyTriggerFor func() EvolutionChain    `json:"-" mapstructure:"-"`
	BabyTriggerFor    apiResource              `json:"baby_trigger_for" mapstructure:"baby_trigger_for"`
	Machines          []machineVersionDetail   `json:"machines" mapstructure:"machines"`
}

type itemSprite struct {
	Default   string        `json:"default" mapstructure:"default"`
	GetSprite func() string `json:"-" mapstructure:"-"`
}

type itemHolderPokemon struct {
	Pokemon        string                  `json:"pokemon" mapstructure:"pokemon"`
	VersionDetails itemHolderPokemonDetail `json:"version_details" mapstructure:"version_details"`
}

type itemHolderPokemonDetail struct {
	Rarity     string           `json:"rarity" mapstructure:"rarity"`
	GetVersion func() Version   `json:"-" mapstructure:"-"`
	Version    namedAPIResource `json:"version" mapstructure:"version"`
}

// ItemAttribute ...
// Item attributes define particular aspects of items, e.g. "usable in battle" or "consumable".
// The ItemAttribute schema can be found at https://pokeapi.co/docs/v2.html/#item-attributes
type ItemAttribute struct {
	ID           int                `json:"id" mapstructure:"id"`
	Name         string             `json:"name" mapstructure:"name"`
	GetItems     []func() Item      `json:"-" mapstructure:"-"`
	Items        []namedAPIResource `json:"items" mapstructure:"items"`
	Names        []name             `json:"names" mapstructure:"names"`
	Descriptions []description      `json:"descriptions" mapstructure:"descriptions"`
}

// ItemCategory ...
// Item categories determine where items will be placed in the players bag.
// The ItemCategory schema can be found at https://pokeapi.co/docs/v2.html/#item-categories
type ItemCategory struct {
	ID        int                `json:"id" mapstructure:"id"`
	Name      string             `json:"name" mapstructure:"name"`
	GetItems  []func() Item      `json:"-" mapstructure:"-"`
	Items     []namedAPIResource `json:"items" mapstructure:"items"`
	Names     []name             `json:"names" mapstructure:"names"`
	GetPocket func() ItemPocket  `json:"-" mapstructure:"-"`
	Pocket    namedAPIResource   `json:"pocket" mapstructure:"pocket"`
}

// ItemFlingEffect ...
// The various effects of the move "Fling" when used with different items.
// The ItemFlingEffect schema can be found at https://pokeapi.co/docs/v2.html/#item-fling-effects
type ItemFlingEffect struct {
	ID            int                `json:"id" mapstructure:"id"`
	Name          string             `json:"name" mapstructure:"name"`
	EffectEntries []effect           `json:"effect_entries" mapstructure:"effect_entries"`
	GetItems      []func() Item      `json:"-" mapstructure:"-"`
	Items         []namedAPIResource `json:"items" mapstructure:"items"`
}

// ItemPocket ...
// Pockets within the players bag used for storing items by category.
// The ItemPocket schema can be found at https://pokeapi.co/docs/v2.html/#item-pockets
type ItemPocket struct {
	ID            int                   `json:"id" mapstructure:"id"`
	Name          string                `json:"name" mapstructure:"name"`
	GetCategories []func() ItemCategory `json:"-" mapstructure:"-"`
	Categories    []namedAPIResource    `json:"categories" mapstructure:"categories"`
	Names         []name                `json:"names" mapstructure:"names"`
}

// Locations Section
// https://pokeapi.co/docs/v2.html/#locations-section

// Location ...
// Locations that can be visited within the games. Locations make up sizable portions of regions, like cities or routes.
// The Location schema can be found at https://pokeapi.co/docs/v2.html/#locations
type Location struct {
	ID          int                   `json:"id" mapstructure:"id"`
	Name        string                `json:"name" mapstructure:"name"`
	GetRegion   func() Region         `json:"-" mapstructure:"-"`
	Region      namedAPIResource      `json:"region" mapstructure:"region"`
	Names       []name                `json:"names" mapstructure:"names"`
	GameIndices []generationGameIndex `json:"game_indices" mapstructure:"game_indices"`
	GetAreas    []func() LocationArea `json:"-" mapstructure:"-"`
	Areas       []namedAPIResource    `json:"areas" mapstructure:"areas"`
}

// LocationArea ...
// Location areas are sections of areas, such as floors in a building or cave. Each area has its own set of possible Pokémon encounters.
// The LocationArea schema can be found at https://pokeapi.co/docs/v2.html/#location-areas
type LocationArea struct {
	ID                   int                   `json:"id" mapstructure:"id"`
	Name                 string                `json:"name" mapstructure:"name"`
	GameIndex            int                   `json:"game_index" mapstructure:"game_index"`
	EncounterMethodRates []encounterMethodRate `json:"encounter_method_rates" mapstructure:"encounter_method_rates"`
	GetLocation          func() Location       `json:"-" mapstructure:"-"`
	Location             namedAPIResource      `json:"location" mapstructure:"location"`
	Names                []name                `json:"names" mapstructure:"names"`
	PokemonEncounters    []pokemonEncounter    `json:"pokemon_encounters" mapstructure:"pokemon_encounters"`
}

type encounterMethodRate struct {
	GetEncounterMethod func() EncounterMethod   `json:"-" mapstructure:"-"`
	EncounterMethod    namedAPIResource         `json:"encounter_method" mapstructure:"encounter_method"`
	VersionDetails     []encounterVersionDetail `json:"version_details" mapstructure:"version_details"`
}

type encounterVersionDetail struct {
	Rate       int              `json:"rate" mapstructure:"rate"`
	GetVersion func() Version   `json:"-" mapstructure:"-"`
	Version    namedAPIResource `json:"version" mapstructure:"version"`
}

type pokemonEncounter struct {
	GetPokemon     func() Pokemon           `json:"-" mapstructure:"-"`
	Pokemon        namedAPIResource         `json:"pokemon" mapstructure:"pokemon"`
	VersionDetails []versionEncounterDetail `json:"version_details" mapstructure:"version_details"`
}

// PalParkArea ...
// Areas used for grouping Pokémon encounters in Pal Park. They're like habitats that are specific to Pal Park.
// The PalParkArea schema can be found at https://pokeapi.co/docs/v2.html/#pal-park-areas
type PalParkArea struct {
	ID                int                       `json:"id" mapstructure:"id"`
	Name              string                    `json:"name" mapstructure:"name"`
	Names             []name                    `json:"names" mapstructure:"names"`
	PokemonEncounters []palParkEncounterSpecies `json:"pokemon_encounters" mapstructure:"pokemon_encounters"`
}

type palParkEncounterSpecies struct {
	BaseScore         int                   `json:"base_score" mapstructure:"base_score"`
	Rate              int                   `json:"rate" mapstructure:"rate"`
	GetPokemonSpecies func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

// Region ...
// A region is an organized area of the Pokémon world. Most often, the main difference between regions is the species of Pokémon that can be encountered within them.
// The Region schema can be found at https://pokeapi.co/docs/v2.html/#regions
type Region struct {
	ID                int `json:"id" mapstructure:"id"`
	GetLocations      []func() Location
	Locations         []namedAPIResource    `json:"locations" mapstructure:"locations"`
	Name              string                `json:"name" mapstructure:"name"`
	Names             []name                `json:"names" mapstructure:"names"`
	GetMainGeneration func() Generation     `json:"-" mapstructure:"-"`
	MainGeneration    namedAPIResource      `json:"main_generation" mapstructure:"main_generation"`
	GetPokedexes      []func() Pokedex      `json:"-" mapstructure:"-"`
	Pokedexes         []namedAPIResource    `json:"pokedexes" mapstructure:"pokedexes"`
	GetVersionGroups  []func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroups     []namedAPIResource    `json:"version_groups" mapstructure:"version_groups"`
}

// Machines Section
// https://pokeapi.co/docs/v2.html/#machines-section

// Machine ...
// Machines are the representation of items that teach moves to Pokémon. They vary from version to version, so it is not certain that one specific TM or HM corresponds to a single Machine.
// The Machine schema can be found at https://pokeapi.co/docs/v2.html/#machines
type Machine struct {
	ID              int                 `json:"id" mapstructure:"id"`
	GetItem         func() Item         `json:"-" mapstructure:"-"`
	Item            namedAPIResource    `json:"item" mapstructure:"item"`
	GetMove         func() Move         `json:"-" mapstructure:"-"`
	Move            namedAPIResource    `json:"move" mapstructure:"move"`
	GetVersionGroup func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroup    namedAPIResource    `json:"version_group" mapstructure:"version_group"`
}

// Moves Section
// https://pokeapi.co/docs/v2.html/#moves-section

// Move ...
// Moves are the skills of Pokémon in battle. In battle, a Pokémon uses one move each turn. Some moves (including those learned by Hidden Machine) can be used outside of battle as well, usually for the purpose of removing obstacles or exploring new areas.
// The Move schema can be found at https://pokeapi.co/docs/v2.html/#moves
type Move struct {
	ID                    int                       `json:"id" mapstructure:"id"`
	Name                  string                    `json:"name" mapstructure:"name"`
	Accuracy              int                       `json:"accuracy" mapstructure:"accuracy"`
	EffectChance          int                       `json:"effect_chance" mapstructure:"effect_chance"`
	PP                    int                       `json:"pp" mapstructure:"pp"`
	Priorty               int                       `json:"priorty" mapstructure:"priorty"`
	Power                 int                       `json:"power" mapstructure:"power"`
	ContestCombos         contestComboSet           `json:"contest_combos" mapstructure:"contest_combos"`
	GetContestType        func() ContestType        `json:"-" mapstructure:"-"`
	ContestType           namedAPIResource          `json:"contest_type" mapstructure:"contest_type"`
	GetContestEffect      func() ContestEffect      `json:"-" mapstructure:"-"`
	ContestEffect         apiResource               `json:"contest_effect" mapstructure:"contest_effect"`
	GetDamageClass        func() MoveDamageClass    `json:"-" mapstructure:"-"`
	DamageClass           namedAPIResource          `json:"damage_class" mapstructure:"damage_class"`
	EffectEntries         []verboseEffect           `json:"effect_entries" mapstructure:"effect_entries"`
	EffectChanges         []abilityEffectChange     `json:"effect_changes" mapstructure:"effect_changes"`
	FlavorTextEntries     []moveFlavorText          `json:"flavor_text_entries" mapstructure:"flavor_text_entries"`
	GetGeneration         func() Generation         `json:"-" mapstructure:"-"`
	Generation            namedAPIResource          `json:"generation" mapstructure:"generation"`
	Machines              []machineVersionDetail    `json:"machines" mapstructure:"machines"`
	Meta                  moveMetaData              `json:"meta" mapstructure:"meta"`
	Names                 []name                    `json:"names" mapstructure:"names"`
	PastValues            []pastMoveStatValue       `json:"past_values" mapstructure:"past_values"`
	StatChanges           []moveStatChange          `json:"stat_changes" mapstructure:"stat_changes"`
	GetSuperContestEffect func() SuperContestEffect `json:"-" mapstructure:"-"`
	SuperContestEffect    apiResource               `json:"super_contest_effect" mapstructure:"super_contest_effect"`
	GetTarget             func() MoveTarget         `json:"-" mapstructure:"-"`
	Target                namedAPIResource          `json:"target" mapstructure:"target"`
	GetType               func() Type               `json:"-" mapstructure:"-"`
	Type                  namedAPIResource          `json:"type" mapstructure:"type"`
}

type contestComboSet struct {
	Normal contestComboDetail `json:"normal" mapstructure:"normal"`
	Super  contestComboDetail `json:"super" mapstructure:"super"`
}

type contestComboDetail struct {
	GetUseBefore []func() Move      `json:"-" mapstructure:"-"`
	UseBefore    []namedAPIResource `json:"use_before" mapstructure:"use_before"`
	GetUseAfter  []func() Move      `json:"-" mapstructure:"-"`
	UseAfter     []namedAPIResource `json:"use_after" mapstructure:"use_after"`
}

type moveFlavorText struct {
	FlavorText      string              `json:"flavor_text" mapstructure:"flavor_text"`
	GetLanguage     func() Language     `json:"-" mapstructure:"-"`
	Language        namedAPIResource    `json:"language" mapstructure:"language"`
	GetVersionGroup func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroup    namedAPIResource    `json:"version_group" mapstructure:"version_group"`
}

type moveStatChange struct {
	Change  int              `json:"change" mapstructure:"change"`
	GetStat func() Stat      `json:"-" mapstructure:"-"`
	Stat    namedAPIResource `json:"stat" mapstructure:"stat"`
}

type pastMoveStatValue struct {
	Accuracy        int                 `json:"accuracy" mapstructure:"accuracy"`
	EffectChance    int                 `json:"effect_chance" mapstructure:"effect_chance"`
	Power           int                 `json:"power" mapstructure:"power"`
	PP              int                 `json:"pp" mapstructure:"pp"`
	EffectEntries   []verboseEffect     `json:"effect_entries" mapstructure:"effect_entries"`
	GetType         func() Type         `json:"-" mapstructure:"-"`
	Type            namedAPIResource    `json:"type" mapstructure:"type"`
	GetVersionGroup func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroup    namedAPIResource    `json:"version_group" mapstructure:"version_group"`
}

type moveMetaData struct {
	GetAilment    func() MoveAilment  `json:"-" mapstructure:"-"`
	Ailment       namedAPIResource    `json:"ailment" mapstructure:"ailment"`
	GetCategory   func() MoveCategory `json:"-" mapstructure:"-"`
	Category      namedAPIResource    `json:"category" mapstructure:"category"`
	MinHits       int                 `json:"min_hits" mapstructure:"min_hits"`
	MaxHits       int                 `json:"max_hits" mapstructure:"max_hits"`
	MinTurns      int                 `json:"min_turns" mapstructure:"min_turns"`
	MaxTurns      int                 `json:"max_turns" mapstructure:"max_turns"`
	Drain         int                 `json:"drain" mapstructure:"drain"`
	Healing       int                 `json:"healing" mapstructure:"healing"`
	CritRate      int                 `json:"crit_rate" mapstructure:"crit_rate"`
	AilmentChance int                 `json:"ailment_chance" mapstructure:"ailment_chance"`
	FlinchChance  int                 `json:"flinch_chance" mapstructure:"flinch_chance"`
	StatChance    int                 `json:"stat_chance" mapstructure:"stat_chance"`
}

// MoveAilment ...
// Move Ailments are status conditions caused by moves used during battle. See Bulbapedia for greater detail.
// The MoveAilment schema can be found at https://pokeapi.co/docs/v2.html/#move-ailments
type MoveAilment struct {
	ID       int                `json:"id" mapstructure:"id"`
	Name     string             `json:"name" mapstructure:"name"`
	GetMoves []func() Move      `json:"-" mapstructure:"-"`
	Moves    []namedAPIResource `json:"moves" mapstructure:"moves"`
	Names    []name             `json:"names" mapstructure:"names"`
}

// MoveBattleStyle ...
// Styles of moves when used in the Battle Palace. See Bulbapedia for greater detail.
// The MoveBattleStyle schema can be found at https://pokeapi.co/docs/v2.html/#move-battle-styles
type MoveBattleStyle struct {
	ID    int    `json:"id" mapstructure:"id"`
	Name  string `json:"name" mapstructure:"name"`
	Names []name `json:"names" mapstructure:"names"`
}

// MoveCategory ...
// Very general categories that loosely group move effects.
// The MoveCategory schema can be found at https://pokeapi.co/docs/v2.html/#move-categories
type MoveCategory struct {
	ID           int                `json:"id" mapstructure:"id"`
	Name         string             `json:"name" mapstructure:"name"`
	GetMoves     []func() Move      `json:"-" mapstructure:"-"`
	Moves        []namedAPIResource `json:"moves" mapstructure:"moves"`
	Descriptions []description      `json:"descriptions" mapstructure:"descriptions"`
}

// MoveDamageClass ...
// The MoveDamageClass schema can be found at https://pokeapi.co/docs/v2.html/#move-damage-classes
type MoveDamageClass struct {
	ID           int                `json:"id" mapstructure:"id"`
	Name         string             `json:"name" mapstructure:"name"`
	Descriptions []description      `json:"descriptions" mapstructure:"descriptions"`
	GetMoves     []func() Move      `json:"-" mapstructure:"-"`
	Moves        []namedAPIResource `json:"moves" mapstructure:"moves"`
	Names        []name             `json:"names" mapstructure:"names"`
}

// MoveLearnMethod ...
// Methods by which Pokémon can learn moves.
// The MoveLearnMethod schema can be found at https://pokeapi.co/docs/v2.html/#move-learn-methods
type MoveLearnMethod struct {
	ID               int                   `json:"id" mapstructure:"id"`
	Name             string                `json:"name" mapstructure:"name"`
	Descriptions     []description         `json:"descriptions" mapstructure:"descriptions"`
	Names            []name                `json:"names" mapstructure:"names"`
	GetVersionGroups []func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroups    []namedAPIResource    `json:"version_groups" mapstructure:"version_groups"`
}

// MoveTarget ...
// Targets moves can be directed at during battle. Targets can be Pokémon, environments or even other moves.
// The MoveTarget schema can be found at https://pokeapi.co/docs/v2.html/#move-targets
type MoveTarget struct {
	ID           int                `json:"id" mapstructure:"id"`
	Name         string             `json:"name" mapstructure:"name"`
	Descriptions []description      `json:"descriptions" mapstructure:"descriptions"`
	GetMoves     []func() Move      `json:"-" mapstructure:"-"`
	Moves        []namedAPIResource `json:"moves" mapstructure:"moves"`
	Names        []name             `json:"names" mapstructure:"names"`
}

// Pokemon Section
// https://pokeapi.co/docs/v2.html/#pokemon-section

// Ability ...
// Abilities provide passive effects for Pokémon in battle or in the overworld. Pokémon have multiple possible abilities but can have only one ability at a time. Check out Bulbapedia for greater detail.
// The MoveTarget schema can be found at https://pokeapi.co/docs/v2.html/#abilities
type Ability struct {
	ID                int                   `json:"id" mapstructure:"id"`
	Name              string                `json:"name" mapstructure:"name"`
	IsMainSeries      bool                  `json:"is_main_series" mapstructure:"is_main_series"`
	GetGeneration     func() Generation     `json:"-" mapstructure:"-"`
	Generation        namedAPIResource      `json:"generation" mapstructure:"generation"`
	Names             []name                `json:"names" mapstructure:"names"`
	EffectEntries     []verboseEffect       `json:"effect_entries" mapstructure:"effect_entries"`
	EffectChanges     []abilityEffectChange `json:"effect_changes" mapstructure:"effect_changes"`
	FlavorTextEntries []abilityFlavorText   `json:"flavor_text_entries" mapstructure:"flavor_text_entries"`
	Pokemon           []abilityPokemon      `json:"pokemon" mapstructure:"pokemon"`
}

type abilityEffectChange struct {
	EffectEntries   []effect            `json:"effect_entries" mapstructure:"effect_entries"`
	GetVersionGroup func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroup    namedAPIResource    `json:"version_group" mapstructure:"version_group"`
}

type abilityFlavorText struct {
	FlavorText      string              `json:"flavor_text" mapstructure:"flavor_text"`
	GetLanguage     func() Language     `json:"-" mapstructure:"-"`
	Language        namedAPIResource    `json:"language" mapstructure:"language"`
	GetVersionGroup func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroup    namedAPIResource    `json:"version_group" mapstructure:"version_group"`
}

type abilityPokemon struct {
	IsHidden   bool             `json:"is_hidden" mapstructure:"is_hidden"`
	Slot       int              `json:"slot" mapstructure:"slot"`
	GetPokemon func() Pokemon   `json:"-" mapstructure:"-"`
	Pokemon    namedAPIResource `json:"pokemon" mapstructure:"pokemon"`
}

// Characteristic ...
// Characteristics indicate which stat contains a Pokémon's highest IV. A Pokémon's Characteristic is determined by the remainder of its highest IV divided by 5 (gene_modulo). Check out Bulbapedia for greater detail.
// The Characteristic schema can be found at https://pokeapi.co/docs/v2.html/#characteristics
type Characteristic struct {
	ID             int           `json:"id" mapstructure:"id"`
	GeneModulo     int           `json:"gene_modulo" mapstructure:"gene_modulo"`
	PossibleValues []int         `json:"possible_values" mapstructure:"possible_values"`
	Descriptions   []description `json:"descriptions" mapstructure:"descriptions"`
}

// EggGroup ...
// Egg Groups are categories which determine which Pokémon are able to interbreed. Pokémon may belong to either one or two Egg Groups. Check out Bulbapedia for greater detail.
// The EggGroup schema can be found at https://pokeapi.co/docs/v2.html/#egg-groups
type EggGroup struct {
	ID                int                     `json:"id" mapstructure:"id"`
	Name              string                  `json:"name" mapstructure:"name"`
	Names             []name                  `json:"names" mapstructure:"names"`
	GetPokemonSpecies []func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    []namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

// Gender ...
// Genders were introduced in Generation II for the purposes of breeding Pokémon but can also result in visual differences or even different evolutionary lines. Check out Bulbapedia for greater detail.
// The Gender schema can be found at https://pokeapi.co/docs/v2.html/#genders
type Gender struct {
	ID                      int                     `json:"id" mapstructure:"id"`
	Name                    string                  `json:"name" mapstructure:"name"`
	PokemonSpeciesDetails   []pokemonSpeciesGender  `json:"pokemon_species_details" mapstructure:"pokemon_species_details"`
	GetRequiredForEvolution []func() PokemonSpecies `json:"-" mapstructure:"-"`
	RequiredForEvolution    []namedAPIResource      `json:"required_for_evolution" mapstructure:"required_for_evolution"`
}

type pokemonSpeciesGender struct {
	Rate              int                   `json:"rate" mapstructure:"rate"`
	GetPokemonSpecies func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

// GrowthRate ...
// Growth rates are the speed with which Pokémon gain levels through experience. Check out Bulbapedia for greater detail.
// The GrowthRate schema can be found at https://pokeapi.co/docs/v2.html/#growth-rates
type GrowthRate struct {
	ID                int                         `json:"id" mapstructure:"id"`
	Name              string                      `json:"name" mapstructure:"name"`
	Formula           string                      `json:"formula" mapstructure:"formula"`
	Descriptions      []description               `json:"descriptions" mapstructure:"descriptions"`
	Levels            []growthRateExperienceLevel `json:"levels" mapstructure:"levels"`
	GetPokemonSpecies []func() PokemonSpecies     `json:"-" mapstructure:"-"`
	PokemonSpecies    []namedAPIResource          `json:"pokemon_species" mapstructure:"pokemon_species"`
}

type growthRateExperienceLevel struct {
	Level      int `json:"level" mapstructure:"level"`
	Experience int `json:"experience" mapstructure:"experience"`
}

// Nature ...
// Natures influence how a Pokémon's stats grow. See Bulbapedia for greater detail.
// The Nature schema can be found at https://pokeapi.co/docs/v2.html/#natures
type Nature struct {
	ID                         int                         `json:"id" mapstructure:"id"`
	Name                       string                      `json:"name" mapstructure:"name"`
	GetDecreasedStat           func() Stat                 `json:"-" mapstructure:"-"`
	DecreasedStat              namedAPIResource            `json:"decreased_stat" mapstructure:"decreased_stat"`
	GetIncreasedStat           func() Stat                 `json:"-" mapstructure:"-"`
	IncreasedStat              namedAPIResource            `json:"increased_stat" mapstructure:"increased_stat"`
	GetHatesFlavor             func() BerryFlavor          `json:"-" mapstructure:"-"`
	HatesFlavor                namedAPIResource            `json:"hates_flavor" mapstructure:"hates_flavor"`
	GetLikesFlavor             func() BerryFlavor          `json:"-" mapstructure:"-"`
	LikesFlavor                namedAPIResource            `json:"likes_flavor" mapstructure:"likes_flavor"`
	PokeathlonStatChanges      []natureStatChange          `json:"pokeathlon_stat_changes" mapstructure:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []moveBattleStylePreference `json:"move_battle_style_preferences" mapstructure:"move_battle_style_preferences"`
	Names                      []name                      `json:"names" mapstructure:"names"`
}

type natureStatChange struct {
	MaxChange         int                   `json:"max_change" mapstructure:"max_change"`
	GetPokeathlonStat func() PokeathlonStat `json:"-" mapstructure:"-"`
	PokeathlonStat    namedAPIResource      `json:"pokeathlon_stat" mapstructure:"pokeathlon_stat"`
}

type moveBattleStylePreference struct {
	LowHPPrefernce     int                    `json:"low_hp_prefernce" mapstructure:"low_hp_prefernce"`
	HighHPPreference   int                    `json:"high_hp_preference" mapstructure:"high_hp_preference"`
	GetMoveBattleStyle func() MoveBattleStyle `json:"-" mapstructure:"-"`
	MoveBattleStyle    namedAPIResource       `json:"move_battle_style" mapstructure:"move_battle_style"`
}

// PokeathlonStat ...
// Pokeathlon Stats are different attributes of a Pokémon's performance in Pokéathlons. In Pokéathlons, competitions happen on different courses; one for each of the different Pokéathlon stats. See Bulbapedia for greater detail.
// The PokeathlonStat schema can be found at https://pokeapi.co/docs/v2.html/#pokeathlon-stats
type PokeathlonStat struct {
	ID               int                              `json:"id" mapstructure:"id"`
	Name             string                           `json:"name" mapstructure:"name"`
	Names            []name                           `json:"names" mapstructure:"names"`
	AffectingNatures []naturePokeathlonStatAffectSets `json:"affecting_natures" mapstructure:"affecting_natures"`
}

type naturePokeathlonStatAffectSets struct {
	Increase []naturePokeathlonStatAffect `json:"increase" mapstructure:"increase"`
	Decrease []naturePokeathlonStatAffect `json:"decrease" mapstructure:"decrease"`
}

type naturePokeathlonStatAffect struct {
	MaxChange int              `json:"max_change" mapstructure:"max_change"`
	GetNature func() Nature    `json:"-" mapstructure:"-"`
	Nature    namedAPIResource `json:"nature" mapstructure:"nature"`
}

// Pokemon ...
// Pokémon are the creatures that inhabit the world of the Pokémon games. They can be caught using Pokéballs and trained by battling with other Pokémon. See Bulbapedia for greater detail.
// The Pokemon schema can be found at https://pokeapi.co/docs/v2.html/#pokemon
type Pokemon struct {
	ID                        int                   `json:"id" mapstructure:"id"`
	Name                      string                `json:"name" mapstructure:"name"`
	BaseExperience            int                   `json:"base_experience" mapstructure:"base_experience"`
	Height                    int                   `json:"height" mapstructure:"height"`
	IsDefault                 bool                  `json:"is_default" mapstructure:"is_default"`
	Order                     int                   `json:"order" mapstructure:"order"`
	Weight                    int                   `json:"weight" mapstructure:"weight"`
	Abilities                 []pokemonAbility      `json:"abilities" mapstructure:"abilities"`
	GetForms                  []func() PokemonForm  `json:"-" mapstructure:"-"`
	Forms                     []namedAPIResource    `json:"forms" mapstructure:"forms"`
	GameIndices               []versionGameIndex    `json:"game_indices" mapstructure:"game_indices"`
	HeldItems                 []pokemonHeldItem     `json:"held_items" mapstructure:"held_items"`
	GetLocationAreaEncounters func() LocationArea   `json:"-" mapstructure:"-"`
	LocationAreaEncounters    string                `json:"location_area_encounters" mapstructure:"location_area_encounters"`
	Moves                     []pokemonMove         `json:"moves" mapstructure:"moves"`
	Sprites                   pokemonSprites        `json:"sprites" mapstructure:"sprites"`
	GetSpecies                func() PokemonSpecies `json:"-" mapstructure:"-"`
	Species                   namedAPIResource      `json:"species" mapstructure:"species"`
	Stats                     []pokemonStat         `json:"stats" mapstructure:"stats"`
	Types                     []pokemonType         `json:"types" mapstructure:"types"`
}

type pokemonAbility struct {
	IsHidden   bool             `json:"is_hidden" mapstructure:"is_hidden"`
	Slot       int              `json:"slot" mapstructure:"slot"`
	GetAbility func() Ability   `json:"-" mapstructure:"-"`
	Ability    namedAPIResource `json:"ability" mapstructure:"ability"`
}

type pokemonType struct {
	Slot    int              `json:"slot" mapstructure:"slot"`
	GetType func() Type      `json:"-" mapstructure:"-"`
	Type    namedAPIResource `json:"type" mapstructure:"type"`
}

type pokemonHeldItem struct {
	GetItem        func() Item              `json:"-" mapstructure:"-"`
	Item           namedAPIResource         `json:"item" mapstructure:"item"`
	VersionDetails []pokemonHeldItemVersion `json:"version_details" mapstructure:"version_details"`
}

type pokemonHeldItemVersion struct {
	GetVersion func() Version   `json:"-" mapstructure:"-"`
	Version    namedAPIResource `json:"version" mapstructure:"version"`
	Rarity     int              `json:"rarity" mapstructure:"rarity"`
}

type pokemonMove struct {
	GetMove             func() Move          `json:"-" mapstructure:"-"`
	Move                namedAPIResource     `json:"move" mapstructure:"move"`
	VersionGroupDetails []pokemonMoveVersion `json:"version_group_details" mapstructure:"version_group_details"`
}

type pokemonMoveVersion struct {
	GetMoveLearnMethod func() MoveLearnMethod `json:"-" mapstructure:"-"`
	MoveLearnMethod    namedAPIResource       `json:"move_learn_method" mapstructure:"move_learn_method"`
	GetVersionGroup    func() VersionGroup    `json:"-" mapstructure:"-"`
	VersionGroup       namedAPIResource       `json:"version_group" mapstructure:"version_group"`
	LevelLearnedAt     int                    `json:"level_learned_at" mapstructure:"level_learned_at"`
}

type pokemonStat struct {
	GetStat  func() Stat      `json:"-" mapstructure:"-"`
	Stat     namedAPIResource `json:"stat" mapstructure:"stat"`
	Effort   int              `json:"effort" mapstructure:"effort"`
	BaseStat int              `json:"base_stat" mapstructure:"base_stat"`
}

type pokemonSprites struct {
	FrontDefault     string `json:"front_default" mapstructure:"front_default"`
	FrontShiny       string `json:"front_shiny" mapstructure:"front_shiny"`
	FrontFemale      string `json:"front_female" mapstructure:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female" mapstructure:"front_shiny_female"`
	BackDefault      string `json:"back_default" mapstructure:"back_default"`
	BackShiny        string `json:"back_shiny" mapstructure:"back_shiny"`
	BackFemale       string `json:"back_female" mapstructure:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female" mapstructure:"back_shiny_female"`
	// TODO GetSprites
	GetSprites []func() string `json:"-" mapstructure:"-"`
}

// PokemonColor ...
// Colors used for sorting Pokémon in a Pokédex. The color listed in the Pokédex is usually the color most apparent or covering each Pokémon's body. No orange category exists; Pokémon that are primarily orange are listed as red or brown.
// The PokemonColor schema can be found at https://pokeapi.co/docs/v2.html/#pokemon-colors
type PokemonColor struct {
	ID                int                     `json:"id" mapstructure:"id"`
	Name              string                  `json:"name" mapstructure:"name"`
	Names             []name                  `json:"names" mapstructure:"names"`
	GetPokemonSpecies []func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    []namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

// PokemonForm ...
// Some Pokémon have the ability to take on different forms. At times, these differences are purely cosmetic and have no bearing on the difference in the Pokémon's stats from another; however, several Pokémon differ in stats (other than HP), type, and Ability depending on their form.
// The PokemonForm schema can be found at https://pokeapi.co/docs/v2.html/#pokemon-forms
type PokemonForm struct {
	ID              int                 `json:"id" mapstructure:"id"`
	Name            string              `json:"name" mapstructure:"name"`
	Order           int                 `json:"order" mapstructure:"order"`
	FormOrder       int                 `json:"form_order" mapstructure:"form_order"`
	IsDefault       bool                `json:"is_default" mapstructure:"is_default"`
	IsBattleOnly    bool                `json:"is_battle_only" mapstructure:"is_battle_only"`
	IsMega          bool                `json:"is_mega" mapstructure:"is_mega"`
	FormName        string              `json:"form_name" mapstructure:"form_name"`
	GetPokemon      func() Pokemon      `json:"-" mapstructure:"-"`
	Pokemon         namedAPIResource    `json:"pokemon" mapstructure:"pokemon"`
	Sprites         pokemonFormSprites  `json:"sprites" mapstructure:"sprites"`
	GetVersionGroup func() VersionGroup `json:"-" mapstructure:"-"`
	VersionGroup    namedAPIResource    `json:"version_group" mapstructure:"version_group"`
	Names           []name              `json:"names" mapstructure:"names"`
	FormNames       []name              `json:"form_names" mapstructure:"form_names"`
}

type pokemonFormSprites struct {
	FrontDefault string `json:"front_default" mapstructure:"front_default"`
	FrontShiny   string `json:"front_shiny" mapstructure:"front_shiny"`
	BackDefault  string `json:"back_default" mapstructure:"back_default"`
	BackShiny    string `json:"back_shiny" mapstructure:"back_shiny"`
	// TODO GetSprites
	GetSprites []func() string `json:"-" mapstructure:"-"`
}

// PokemonHabitat ...
// Habitats are generally different terrain Pokémon can be found in but can also be areas designated for rare or legendary Pokémon.
// The PokemonHabitat schema can be found at https://pokeapi.co/docs/v2.html/#pokemon-habitats
type PokemonHabitat struct {
	ID                int                     `json:"id" mapstructure:"id"`
	Name              string                  `json:"name" mapstructure:"name"`
	Names             []name                  `json:"names" mapstructure:"names"`
	GetPokemonSpecies []func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    []namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

// PokemonShape ...
// Shapes used for sorting Pokémon in a Pokédex.
// The PokemonShape schema can be found at https://pokeapi.co/docs/v2.html/#pokemon-shapes
type PokemonShape struct {
	ID                int                     `json:"id" mapstructure:"id"`
	Name              string                  `json:"name" mapstructure:"name"`
	AwesomeNames      []awesomeName           `json:"awesome_names" mapstructure:"awesome_names"`
	Names             []name                  `json:"names" mapstructure:"names"`
	GetPokemonSpecies []func() PokemonSpecies `json:"-" mapstructure:"-"`
	PokemonSpecies    []namedAPIResource      `json:"pokemon_species" mapstructure:"pokemon_species"`
}

type awesomeName struct {
	AwesomeName string           `json:"awesome_name" mapstructure:"awesome_name"`
	GetLanguage func() Language  `json:"-" mapstructure:"-"`
	Language    namedAPIResource `json:"language" mapstructure:"language"`
}

// PokemonSpecies ...
// A Pokémon Species forms the basis for at least one Pokémon. Attributes of a Pokémon species are shared across all varieties of Pokémon within the species. A good example is Wormadam; Wormadam is the species which can be found in three different varieties, Wormadam-Trash, Wormadam-Sandy and Wormadam-Plant.
// The PokemonSpecies schema can be found at https://pokeapi.co/docs/v2.html/#pokemon-species
type PokemonSpecies struct {
	ID                    int                      `json:"id" mapstructure:"id"`
	Name                  string                   `json:"name" mapstructure:"name"`
	Order                 int                      `json:"order" mapstructure:"order"`
	GenderRate            int                      `json:"gender_rate" mapstructure:"gender_rate"`
	CaptureRate           int                      `json:"capture_rate" mapstructure:"capture_rate"`
	BaseHappiness         int                      `json:"base_happiness" mapstructure:"base_happiness"`
	IsBaby                bool                     `json:"is_baby" mapstructure:"is_baby"`
	HatchCounter          int                      `json:"hatch_counter" mapstructure:"hatch_counter"`
	HasGenderDifferences  bool                     `json:"has_gender_differences" mapstructure:"has_gender_differences"`
	FormsSwitchable       bool                     `json:"forms_switchable" mapstructure:"forms_switchable"`
	GetGrowthRate         func() GrowthRate        `json:"-" mapstructure:"-"`
	GrowthRate            namedAPIResource         `json:"growth_rate" mapstructure:"growth_rate"`
	PokedexNumbers        []pokemonSpeciesDexEntry `json:"pokedex_numbers" mapstructure:"pokedex_numbers"`
	GetEggGroups          []func() EggGroup        `json:"-" mapstructure:"-"`
	EggGroups             []namedAPIResource       `json:"egg_groups" mapstructure:"egg_groups"`
	GetColor              func() PokemonColor      `json:"-" mapstructure:"-"`
	Color                 namedAPIResource         `json:"color" mapstructure:"color"`
	GetShape              func() PokemonShape      `json:"-" mapstructure:"-"`
	Shape                 namedAPIResource         `json:"shape" mapstructure:"shape"`
	GetEvolvesFromSpecies func() PokemonSpecies    `json:"-" mapstructure:"-"`
	EvolvesFromSpecies    namedAPIResource         `json:"evolves_from_species" mapstructure:"evolves_from_species"`
	GetEvolutionChain     func() EvolutionChain    `json:"-" mapstructure:"-"`
	EvolutionChain        apiResource              `json:"evolution_chain" mapstructure:"evolution_chain"`
	GetHabitat            func() PokemonHabitat    `json:"-" mapstructure:"-"`
	Habitat               namedAPIResource         `json:"habitat" mapstructure:"habitat"`
	GetGeneration         func() Generation        `json:"-" mapstructure:"-"`
	Generation            namedAPIResource         `json:"generation" mapstructure:"generation"`
	Names                 []name                   `json:"names" mapstructure:"names"`
	PalParkEncounters     []palParkEncounterArea   `json:"pal_park_encounters" mapstructure:"pal_park_encounters"`
	FlavorTextEntries     []flavorText             `json:"flavor_text_entries" mapstructure:"flavor_text_entries"`
	FormDescriptions      []description            `json:"form_descriptions" mapstructure:"form_descriptions"`
	Genera                []genus                  `json:"genera" mapstructure:"genera"`
	Varieties             []pokemonSpeciesVariety  `json:"varieties" mapstructure:"varieties"`
}

type genus struct {
	Genus       string           `json:"genus" mapstructure:"genus"`
	GetLanguage func() Language  `json:"-" mapstructure:"-"`
	Language    namedAPIResource `json:"language" mapstructure:"language"`
}

type pokemonSpeciesDexEntry struct {
	EntryNumber int              `json:"entry_number" mapstructure:"entry_number"`
	GetPokedex  func() Pokedex   `json:"-" mapstructure:"-"`
	Pokedex     namedAPIResource `json:"pokedex" mapstructure:"pokedex"`
}

type palParkEncounterArea struct {
	BaseScore int                `json:"base_score" mapstructure:"base_score"`
	Rate      int                `json:"rate" mapstructure:"rate"`
	GetArea   func() PalParkArea `json:"-" mapstructure:"-"`
	Area      namedAPIResource   `json:"area" mapstructure:"area"`
}

type pokemonSpeciesVariety struct {
	IsDefault  bool             `json:"is_default" mapstructure:"is_default"`
	GetPokemon func() Pokemon   `json:"-" mapstructure:"-"`
	Pokemon    namedAPIResource `json:"pokemon" mapstructure:"pokemon"`
}

// Stat ...
// Stats determine certain aspects of battles. Each Pokémon has a value for each stat which grows as they gain levels and can be altered momentarily by effects in battles.
// The Stat schema can be found at https://pokeapi.co/docs/v2.html/#stats
type Stat struct {
	ID                 int                     `json:"id" mapstructure:"id"`
	Name               string                  `json:"name" mapstructure:"name"`
	GameIndex          int                     `json:"game_index" mapstructure:"game_index"`
	IsBattleOnly       bool                    `json:"is_battle_only" mapstructure:"is_battle_only"`
	AffectingMoves     moveStatAffectSets      `json:"affecting_moves" mapstructure:"affecting_moves"`
	AffectingNatures   natureStatAffectSets    `json:"affecting_natures" mapstructure:"affecting_natures"`
	GetCharacteristics []func() Characteristic `json:"-" mapstructure:"-"`
	Characteristics    []apiResource           `json:"characteristics" mapstructure:"characteristics"`
	GetMoveDamageClass func() MoveDamageClass  `json:"-" mapstructure:"-"`
	MoveDamageClass    namedAPIResource        `json:"move_damage_class" mapstructure:"move_damage_class"`
	Names              []name                  `json:"names" mapstructure:"names"`
}

type moveStatAffectSets struct {
	Increase []moveStatAffect `json:"increase" mapstructure:"increase"`
	Decrease []moveStatAffect `json:"decrease" mapstructure:"decrease"`
}

type moveStatAffect struct {
	Change  int              `json:"change" mapstructure:"change"`
	GetMove func() Move      `json:"-" mapstructure:"-"`
	Move    namedAPIResource `json:"move" mapstructure:"move"`
}

type natureStatAffectSets struct {
	GetIncrease []func() Nature    `json:"-" mapstructure:"-"`
	Increase    []namedAPIResource `json:"increase" mapstructure:"increase"`
	GetDecrease []func() Nature    `json:"-" mapstructure:"-"`
	Decrease    []namedAPIResource `json:"decrease" mapstructure:"decrease"`
}

// Type ..
// Types are properties for Pokémon and their moves. Each type has three properties: which types of Pokémon it is super effective against, which types of Pokémon it is not very effective against, and which types of Pokémon it is completely ineffective against.
// The Type schema can be found at https://pokeapi.co/docs/v2.html/#types
type Type struct {
	ID                 int                    `json:"id" mapstructure:"id"`
	Name               string                 `json:"name" mapstructure:"name"`
	DamageRelations    typeRelations          `json:"damage_relations" mapstructure:"damage_relations"`
	GameIndices        []generationGameIndex  `json:"game_indices" mapstructure:"game_indices"`
	GetGeneration      func() Generation      `json:"-" mapstructure:"-"`
	Generation         namedAPIResource       `json:"generation" mapstructure:"generation"`
	GetMoveDamageClass func() MoveDamageClass `json:"-" mapstructure:"-"`
	MoveDamageClass    namedAPIResource       `json:"move_damage_class" mapstructure:"move_damage_class"`
	Names              []name                 `json:"names" mapstructure:"names"`
	Pokemon            []typePokemon          `json:"pokemon" mapstructure:"pokemon"`
	GetMoves           []func() Move          `json:"-" mapstructure:"-"`
	Moves              []namedAPIResource     `json:"moves" mapstructure:"moves"`
}

type typePokemon struct {
	Slot       int              `json:"slot" mapstructure:"slot"`
	GetPokemon func() Pokemon   `json:"-" mapstructure:"-"`
	Pokemon    namedAPIResource `json:"pokemon" mapstructure:"pokemon"`
}

type typeRelations struct {
	GetNoDamageTo       []func() Type      `json:"-" mapstructure:"-"`
	NoDamageTo          []namedAPIResource `json:"no_damage_to" mapstructure:"no_damage_to"`
	GetHalfDamageTo     []func() Type      `json:"-" mapstructure:"-"`
	HalfDamageTo        []namedAPIResource `json:"half_damage_to" mapstructure:"half_damage_to"`
	GetDoubleDamageTo   []func() Type      `json:"-" mapstructure:"-"`
	DoubleDamageTo      []namedAPIResource `json:"double_damage_to" mapstructure:"double_damage_to"`
	GetNoDamageFrom     []func() Type      `json:"-" mapstructure:"-"`
	NoDamageFrom        []namedAPIResource `json:"no_damage_from" mapstructure:"no_damage_from"`
	GetHalfDamageFrom   []func() Type      `json:"-" mapstructure:"-"`
	HalfDamageFrom      []namedAPIResource `json:"half_damage_from" mapstructure:"half_damage_from"`
	GetDoubleDamageFrom []func() Type      `json:"-" mapstructure:"-"`
	DoubleDamageFrom    []namedAPIResource `json:"double_damage_from" mapstructure:"double_damage_from"`
}

// Utility Section
// https://pokeapi.co/docs/v2.html/#utility-section

// Language ...
// Languages for translations of API resource information.
// The Language schema can be found at https://pokeapi.co/docs/v2.html/#language
type Language struct {
	ID       int    `json:"id" mapstructure:"id"`
	Name     string `json:"name" mapstructure:"name"`
	Official bool   `json:"official" mapstructure:"official"`
	Iso639   string `json:"iso639" mapstructure:"iso639"`
	Iso3166  string `json:"iso3166" mapstructure:"iso3166"`
	Names    []name `json:"names" mapstructure:"names"`
}

type apiResource struct {
	URL string `json:"url" mapstructure:"url"`
}

type description struct {
	Description string           `json:"description" mapstructure:"description"`
	Language    namedAPIResource `json:"language" mapstructure:"language"`
}

type effect struct {
	Effect   string           `json:"effect" mapstructure:"effect"`
	Language namedAPIResource `json:"language" mapstructure:"language"`
}

type encounter struct {
	MinLevel        int                `json:"min_level" mapstructure:"min_level"`
	MaxLevel        int                `json:"max_level" mapstructure:"max_level"`
	ConditionValues []namedAPIResource `json:"condition_values" mapstructure:"condition_values"`
	Chance          int                `json:"chance" mapstructure:"chance"`
	Method          namedAPIResource   `json:"method" mapstructure:"method"`
}

type flavorText struct {
	FlavorText string           `json:"flavor_text" mapstructure:"flavor_text"`
	Language   namedAPIResource `json:"language" mapstructure:"language"`
	Version    namedAPIResource `json:"version" mapstructure:"version"`
}

type generationGameIndex struct {
	GameIndex  int              `json:"game_index" mapstructure:"game_index"`
	Generation namedAPIResource `json:"generation" mapstructure:"generation"`
}

type machineVersionDetail struct {
	Machine      apiResource      `json:"machine" mapstructure:"machine"`
	VersionGroup namedAPIResource `json:"version_group" mapstructure:"version_group"`
}

type name struct {
	Name     string           `json:"name" mapstructure:"name"`
	Language namedAPIResource `json:"language" mapstructure:"language"`
}

type namedAPIResource struct {
	Name string `json:"name" mapstructure:"name"`
	URL  string `json:"url" mapstructure:"url"`
}

type verboseEffect struct {
	Effect      string           `json:"effect" mapstructure:"effect"`
	ShortEffect string           `json:"short_effect" mapstructure:"short_effect"`
	Language    namedAPIResource `json:"language" mapstructure:"language"`
}

type versionEncounterDetail struct {
	Version          namedAPIResource `json:"version" mapstructure:"version"`
	MaxChance        int              `json:"max_chance" mapstructure:"max_chance"`
	EncounterDetails []encounter      `json:"encounter_details" mapstructure:"encounter_details"`
}

type versionGameIndex struct {
	GameIndex int              `json:"game_index" mapstructure:"game_index"`
	Version   namedAPIResource `json:"version" mapstructure:"version"`
}

type versionGroupFlavorText struct {
	Text         string           `json:"text" mapstructure:"text"`
	Language     namedAPIResource `json:"language" mapstructure:"language"`
	VersionGroup namedAPIResource `json:"version_group" mapstructure:"version_group"`
}
