package AI

func SelectexchangeCard(wanList []int32, tiaoList []int32, tongList []int32) []int32 {
	listMap := make(map[int][]int32)
	listMap[0] = wanList
	listMap[1] = tiaoList
	listMap[2] = tongList

	countMin := 14
	indexMin := 0
	for k, list := range listMap {
		if len(list) >= 3 && len(list) < countMin {
			countMin = len(list)
			indexMin = k
		}
	}
	return listMap[indexMin][0:3]
}

func SelectlackType(wanList []int32, tiaoList []int32, tongList []int32) int {
	listMap := make(map[int][]int32)
	listMap[0] = wanList
	listMap[1] = tiaoList
	listMap[2] = tongList

	countMin := 14
	indexMin := 0
	for k, list := range listMap {
		if len(list) >= 3 && len(list) < countMin {
			countMin = len(list)
			indexMin = k
		}
	}
	return indexMin
}
