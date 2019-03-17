package pokeapigo

import "github.com/mitchellh/mapstructure"

// GetItem returns a Item object
// Acceptable search queries:
// ID, String
func GetItem(query string) Item {
	searchType := idOrString(query)
	item := Item{}
	if checkCache("item", query, searchType) {
		mapstructure.Decode(getDataCache("item", query, searchType, item), &item)
	} else {
		dataInterface = getDataAPI("item", query, item)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &item)
			storeData(false, item, "item", item.ID, item.Name)
		}
	}

	item.GetFlingEffect = getItemFlingEffect(item.FlingEffect.Name)
	for _, a := range item.Attributes {
		item.GetAttributes = append(item.GetAttributes, getItemAttribute(a.Name))
	}

	item.GetCategory = getItemCategory(item.Category.Name)
	if item.BabyTriggerFor.URL == "" {
		item.GetBabyTriggerFor = nil
	} else {
		item.GetBabyTriggerFor = getEvolutionChain(getID(item.BabyTriggerFor.URL))
	}

	for i, hbp := range item.HeldByPokemon {
		item.HeldByPokemon[i].VersionDetails.GetVersion = getVersion(hbp.VersionDetails.Version.Name)
	}

	item.Sprites.GetSprite = getItemSprite(item.Sprites.Default)
	return item
}

// GetItemAttribute returns a ItemAttribute object
// Acceptable search queries:
// ID, String
func GetItemAttribute(query string) ItemAttribute {
	searchType := idOrString(query)
	itemAttribute := ItemAttribute{}
	if checkCache("item-attribute", query, searchType) {
		mapstructure.Decode(getDataCache("item-attribute", query, searchType, itemAttribute), &itemAttribute)
	} else {
		dataInterface = getDataAPI("item-attribute", query, itemAttribute)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &itemAttribute)
			storeData(false, itemAttribute, "item-attribute", itemAttribute.ID, itemAttribute.Name)
		}
	}
	for _, item := range itemAttribute.Items {
		itemAttribute.GetItems = append(itemAttribute.GetItems, getItem(item.Name))
	}
	return itemAttribute
}

// GetItemCategory returns a ItemCategory object
// Acceptable search queries:
// ID, String
func GetItemCategory(query string) ItemCategory {
	searchType := idOrString(query)
	itemCategory := ItemCategory{}
	if checkCache("item-category", query, searchType) {
		mapstructure.Decode(getDataCache("item-category", query, searchType, itemCategory), &itemCategory)
	} else {
		dataInterface = getDataAPI("item-category", query, itemCategory)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &itemCategory)
			storeData(false, itemCategory, "item-category", itemCategory.ID, itemCategory.Name)
		}
	}
	for _, item := range itemCategory.Items {
		itemCategory.GetItems = append(itemCategory.GetItems, getItem(item.Name))
	}

	itemCategory.GetPocket = getItemPocket(itemCategory.Pocket.Name)
	return itemCategory
}

// GetItemFlingEffect returns a ItemFlingEffect object
// Acceptable search queries:
// ID, String
func GetItemFlingEffect(query string) ItemFlingEffect {
	searchType := idOrString(query)
	itemFlingEffect := ItemFlingEffect{}
	if checkCache("item-fling-effect", query, searchType) {
		mapstructure.Decode(getDataCache("item-fling-effect", query, searchType, itemFlingEffect), &itemFlingEffect)
	} else {
		dataInterface = getDataAPI("item-fling-effect", query, itemFlingEffect)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &itemFlingEffect)
			storeData(false, itemFlingEffect, "item-fling-effect", itemFlingEffect.ID, itemFlingEffect.Name)
		}
	}
	for _, item := range itemFlingEffect.Items {
		itemFlingEffect.GetItems = append(itemFlingEffect.GetItems, getItem(item.Name))
	}

	return itemFlingEffect
}

// GetItemPocket returns a ItemPocket object
// Acceptable search queries:
// ID, String
func GetItemPocket(query string) ItemPocket {
	searchType := idOrString(query)
	itemPocket := ItemPocket{}
	if checkCache("item-pocket", query, searchType) {
		mapstructure.Decode(getDataCache("item-pocket", query, searchType, itemPocket), &itemPocket)
	} else {
		dataInterface = getDataAPI("item-pocket", query, itemPocket)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &itemPocket)
			storeData(false, itemPocket, "item-pocket", itemPocket.ID, itemPocket.Name)
		}
	}

	for _, category := range itemPocket.Categories {
		itemPocket.GetCategories = append(itemPocket.GetCategories, getItemCategory(category.Name))
	}
	return itemPocket
}
