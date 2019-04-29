package golem

import "github.com/mitchellh/mapstructure"

// GetMove returns a Move object
// Acceptable search queries:
// ID, String
func GetMove(query string) Move {
	searchType := idOrString(query)
	move := Move{}
	if checkCache("move", query, searchType) {
		mapstructure.Decode(getDataCache("move", query, searchType, move), &move)
	} else {
		dataInterface = getDataAPI("move", query, move)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &move)
			storeData(false, move, "move", move.ID, move.Name)
		}
	}
	move.GetContestType = getContestType(move.ContestType.Name)
	//move.GetContestEffect = getContestEffect(getID(move.ContestEffect.URL)) Fix later
	move.GetDamageClass = getMoveDamageClass(move.DamageClass.Name)
	move.GetGeneration = getGeneration(move.Generation.Name)
	//move.GetSuperContestEffect = getSuperContestEffect(getID(move.SuperContestEffect.URL)) Fix later
	move.GetTarget = getMoveTarget(move.Target.Name)
	move.GetType = getType(move.Type.Name)
	for _, ua := range move.ContestCombos.Normal.UseAfter {
		move.ContestCombos.Normal.GetUseAfter = append(move.ContestCombos.Normal.GetUseAfter, getMove(ua.Name))
	}
	for _, ub := range move.ContestCombos.Normal.UseBefore {
		move.ContestCombos.Normal.GetUseBefore = append(move.ContestCombos.Normal.GetUseBefore, getMove(ub.Name))
	}

	for _, ua := range move.ContestCombos.Super.UseAfter {
		move.ContestCombos.Super.GetUseAfter = append(move.ContestCombos.Super.GetUseAfter, getMove(ua.Name))
	}
	for _, ub := range move.ContestCombos.Super.UseBefore {
		move.ContestCombos.Super.GetUseBefore = append(move.ContestCombos.Super.GetUseBefore, getMove(ub.Name))
	}

	for i, fte := range move.FlavorTextEntries {
		move.FlavorTextEntries[i].GetLanguage = getLanguage(fte.Language.Name)
		move.FlavorTextEntries[i].GetVersionGroup = getVersionGroup(fte.VersionGroup.Name)
	}
	for i, sc := range move.StatChanges {
		move.StatChanges[i].GetStat = getStat(sc.Stat.Name)
	}

	for i, pv := range move.PastValues {
		move.PastValues[i].GetType = getType(pv.Type.Name)
		move.PastValues[i].GetVersionGroup = getVersionGroup(pv.VersionGroup.Name)
	}
	move.Meta.GetAilment = getMoveAilment(move.Meta.Ailment.Name)
	move.Meta.GetCategory = getMoveCategory(move.Meta.Category.Name)
	return move
}

// GetMoveAilment returns a MoveAilment object
// Acceptable search queries:
// ID, String
func GetMoveAilment(query string) MoveAilment {
	searchType := idOrString(query)
	moveAilment := MoveAilment{}
	if checkCache("move-ailment", query, searchType) {
		mapstructure.Decode(getDataCache("move-ailment", query, searchType, moveAilment), &moveAilment)
	} else {
		dataInterface = getDataAPI("move-ailment", query, moveAilment)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &moveAilment)
			storeData(false, moveAilment, "move-ailment", moveAilment.ID, moveAilment.Name)
		}
	}
	for _, move := range moveAilment.Moves {
		moveAilment.GetMoves = append(moveAilment.GetMoves, getMove(move.Name))
	}
	return moveAilment
}

// GetMoveBattleStyle returns a MoveBattleStyle object
// Acceptable search queries:
// ID, String
func GetMoveBattleStyle(query string) MoveBattleStyle {
	searchType := idOrString(query)
	moveBattleStyle := MoveBattleStyle{}
	if checkCache("move-battle-style", query, searchType) {
		mapstructure.Decode(getDataCache("move-battle-style", query, searchType, moveBattleStyle), &moveBattleStyle)
	} else {
		dataInterface = getDataAPI("move-battle-style", query, moveBattleStyle)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &moveBattleStyle)
			storeData(false, moveBattleStyle, "move-battle-style", moveBattleStyle.ID, moveBattleStyle.Name)
		}
	}
	return moveBattleStyle
}

// GetMoveCategory returns a MoveCategory object
// Acceptable search queries:
// ID, String
func GetMoveCategory(query string) MoveCategory {
	searchType := idOrString(query)
	moveCategory := MoveCategory{}
	if checkCache("move-category", query, searchType) {
		mapstructure.Decode(getDataCache("move-category", query, searchType, moveCategory), &moveCategory)
	} else {
		dataInterface = getDataAPI("move-category", query, moveCategory)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &moveCategory)
			storeData(false, moveCategory, "move-category", moveCategory.ID, moveCategory.Name)
		}
	}

	for _, move := range moveCategory.Moves {
		moveCategory.GetMoves = append(moveCategory.GetMoves, getMove(move.Name))
	}
	return moveCategory
}

// GetMoveDamageClass returns a MoveDamageClass object
// Acceptable search queries:
// ID, String
func GetMoveDamageClass(query string) MoveDamageClass {
	searchType := idOrString(query)
	moveDamageClass := MoveDamageClass{}
	if checkCache("move-damage-class", query, searchType) {
		mapstructure.Decode(getDataCache("move-damage-class", query, searchType, moveDamageClass), &moveDamageClass)
	} else {
		dataInterface = getDataAPI("move-damage-class", query, moveDamageClass)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &moveDamageClass)
			storeData(false, moveDamageClass, "move-damage-class", moveDamageClass.ID, moveDamageClass.Name)
		}
	}
	for _, move := range moveDamageClass.Moves {
		moveDamageClass.GetMoves = append(moveDamageClass.GetMoves, getMove(move.Name))
	}
	return moveDamageClass
}

// GetMoveLearnMethod returns a MoveLearnMethod object
// Acceptable search queries:
// ID, String
func GetMoveLearnMethod(query string) MoveLearnMethod {
	searchType := idOrString(query)
	moveLearnMethod := MoveLearnMethod{}
	if checkCache("move-learn-method", query, searchType) {
		mapstructure.Decode(getDataCache("move-learn-method", query, searchType, moveLearnMethod), &moveLearnMethod)
	} else {
		dataInterface = getDataAPI("move-learn-method", query, moveLearnMethod)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &moveLearnMethod)
			storeData(false, moveLearnMethod, "move-learn-method", moveLearnMethod.ID, moveLearnMethod.Name)
		}
	}
	for _, vg := range moveLearnMethod.VersionGroups {
		moveLearnMethod.GetVersionGroups = append(moveLearnMethod.GetVersionGroups, getVersionGroup(vg.Name))
	}
	return moveLearnMethod
}

// GetMoveTarget returns a MoveTarget object
// Acceptable search queries:
// ID, String
func GetMoveTarget(query string) MoveTarget {
	searchType := idOrString(query)
	moveTarget := MoveTarget{}
	if checkCache("move-target", query, searchType) {
		mapstructure.Decode(getDataCache("move-target", query, searchType, moveTarget), &moveTarget)
	} else {
		dataInterface = getDataAPI("move-target", query, moveTarget)
		if dataInterface != nil {
			mapstructure.Decode(dataInterface, &moveTarget)
			storeData(false, moveTarget, "move-target", moveTarget.ID, moveTarget.Name)
		}
	}
	for _, move := range moveTarget.Moves {
		moveTarget.GetMoves = append(moveTarget.GetMoves, getMove(move.Name))
	}
	return moveTarget
}
