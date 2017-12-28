package mgrCard

//返回杠牌ID
func CheckGangOwn(list []int) (int, bool) {
	dict := make(map[int]int) //id : count
	for i := 0; i < len(list); i++ {
		count, ok := dict[list[i]]
		if ok {
			count++
			dict[list[i]] = count
		} else {
			dict[list[i]] = 1
		}
	}
	for id, count := range dict {
		if count == 4 {
			return id, true
		}
	}

	return 0, false
}

func CheckPengOther(list []int, discard *CardInfo) bool {

	countNum := getCountInListById(int(discard.Id), list)
	if countNum >= 2 {
		return true
	}

	return false
}

func CheckGangOther(list []int, discard *CardInfo) bool {

	countNum := getCountInListById(int(discard.Id), list)
	if countNum == 3 {
		return true
	}

	return false
}
