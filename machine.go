package pokeapigo

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

// GetMachine returns a Machine object
// Acceptable search queries:
// ID
func GetMachine(id int) Machine {
	searchType := "ID"
	query := strconv.Itoa(id)
	machine := Machine{}
	if checkCache("machine", query, searchType) {
		mapstructure.Decode(getDataCache("machine", query, searchType, machine), &machine)
	} else {
		dataInterface = getDataAPI("machine", query, machine)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &machine)
			storeData(true, machine, "machine", machine.ID, "")
		}
	}
	machine.GetItem = getItem(machine.Item.Name)
	machine.GetMove = getMove(machine.Move.Name)
	machine.GetVersionGroup = getVersionGroup(machine.VersionGroup.Name)
	return machine
}
