package pokeapigo

import "github.com/mitchellh/mapstructure"

// GetEncounterMethod returns a EncounterMethod object
// Acceptable search queries:
// ID, String
func GetEncounterMethod(query string) EncounterMethod {
	searchType := idOrString(query)
	encounterMethod := EncounterMethod{}
	if checkCache("encounter-method", query, searchType) {
		mapstructure.Decode(getDataCache("encounter-method", query, searchType, encounterMethod), &encounterMethod)
	} else {
		dataInterface = getDataAPI("encounter-method", query, encounterMethod)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &encounterMethod)
			storeData(false, encounterMethod, "encounter-method", encounterMethod.ID, encounterMethod.Name)
		}
	}
	return encounterMethod
}

// GetEncounterCondition returns a EncounterCondition object
// Acceptable search queries:
// ID, String
func GetEncounterCondition(query string) EncounterCondition {
	searchType := idOrString(query)
	encounterCondition := EncounterCondition{}
	if checkCache("encounter-condition", query, searchType) {
		mapstructure.Decode(getDataCache("encounter-condition", query, searchType, encounterCondition), &encounterCondition)
	} else {
		dataInterface = getDataAPI("encounter-condition", query, encounterCondition)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &encounterCondition)
			storeData(false, encounterCondition, "encounter-condition", encounterCondition.ID, encounterCondition.Name)
		}
	}
	for _, ecf := range encounterCondition.Values {
		encounterCondition.GetValues = append(encounterCondition.GetValues, getEncounterConditionValue(ecf.Name))
	}
	return encounterCondition
}

// GetEncounterConditionValue returns a EncounterConditionValue object
// Acceptable search queries:
// ID, String
func GetEncounterConditionValue(query string) EncounterConditionValue {
	searchType := idOrString(query)
	encounterConditionValue := EncounterConditionValue{}
	if checkCache("encounter-condition-value", query, searchType) {
		mapstructure.Decode(getDataCache("encounter-condition-value", query, searchType, encounterConditionValue), &encounterConditionValue)
	} else {
		dataInterface = getDataAPI("encounter-condition-value", query, encounterConditionValue)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &encounterConditionValue)
			storeData(false, encounterConditionValue, "encounter-condition-value", encounterConditionValue.ID, encounterConditionValue.Name)
		}
	}
	encounterConditionValue.GetCondition = getEncounterCondition(encounterConditionValue.Condition.Name)
	return encounterConditionValue
}
