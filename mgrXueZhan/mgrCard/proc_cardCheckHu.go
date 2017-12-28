package mgrCard

import (
	"sort"

	"github.com/name5566/leaf/log"
)

func CheckHuOther(inhand []int, gang []int, peng []int, drawcard int) bool {
	cardList := make([]int, 0)
	cardList = append(cardList, inhand...)
	cardList = append(cardList, drawcard)

	return CheckHuOwn(cardList, gang, peng)

}

func CheckHuOwn(inhand []int, gang []int, peng []int) bool {

	sumCount := len(inhand) + len(peng) + len(gang)
	if sumCount < 14 || sumCount > 18 {
		log.Error("sumCount[%v] is error.", sumCount)
		return false
	}

	if !checkPeng(peng) {
		return false
	}

	if !checkGang(gang) {
		return false
	}

	if !checkInHand(inhand) {
		return false
	}

	return true
}

//检查碰牌
//list是cardId数组
func checkPeng(list []int) bool {
	if len(list)%3 != 0 {
		log.Error("peng card count[%v] is error", len(list))
		return false
	}
	dict := make(map[int]int) //cardId : count
	for i := 0; i < len(list); i++ {
		count, ok := dict[list[i]]
		if ok {
			count++
			dict[list[i]] = count
		} else {
			dict[list[i]] = 1
		}
	}
	for _, count := range dict {
		if count != 3 {
			return false
		}
	}
	return true
}

//检查杠牌
//list是cardId数组
func checkGang(list []int) bool {
	if len(list)%4 != 0 {
		log.Error("gang card count[%v] is error", len(list))
		return false
	}
	dict := make(map[int]int) //cardId : count
	for i := 0; i < len(list); i++ {
		count, ok := dict[list[i]]
		if ok {
			count++
		} else {
			dict[list[i]] = 1
		}
	}
	for _, count := range dict {
		if count != 4 {
			return false
		}
	}
	return true
}

//检查手牌
//list是cardId数组
func checkInHand(list []int) bool {
	var sortList []int
	for i := 0; i < len(list); i++ {
		sortList = append(sortList, int(list[i]))
	}
	sort.Ints(sortList)

	//检查七小对
	if isSevenPair(sortList) {
		log.Debug("isSevenPair")
		return true
	}

	if len(sortList) == 2 {
		return sortList[0] == sortList[1]
	}

	for i := 0; i < len(sortList); i++ {
		var tempList []int
		for i := 0; i < len(sortList); i++ {
			tempList = append(tempList, sortList[i])
		}

		count := getCountInListById(sortList[i], sortList)
		//判断是否能做将牌
		if count >= 2 {
			//移除两张将牌
			tempList = removeJiang(sortList[i], tempList)
			//避免重复运算 将光标移到其他牌上
			i = i + count - 1
			//检查剩余牌顺子、刻子情况
			if huPaiPanDin(tempList) {
				log.Debug("is normal hu")
				return true
			}
		}
	}
	return false
}

func getCountInListById(id int, list []int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == id {
			count++
		}
	}
	return count
}

func hasCardById(id int, list []int) bool {
	isFind := false
	for i := 0; i < len(list); i++ {
		if list[i] == id {
			isFind = true
			break
		}
	}
	return isFind
}

func removeJiang(id int, list []int) []int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == id {
			list = append(list[:i], list[i+1:]...)
			count++
			if count == 2 {
				break
			}
			i--
		}
	}
	return list
}

func huPaiPanDin(list []int) bool {

	if len(list) == 0 {
		return true
	}

	count := getCountInListById(list[0], list)

	//组成刻子
	if count == 3 {
		return huPaiPanDin(list[3:])
	} else {
		//组成顺子
		if hasCardById(list[0]+1, list) && hasCardById(list[0]+2, list) {
			firstId := list[0]
			list = list[1:]
			for i := 0; i < len(list); i++ {
				if list[i] == firstId+1 {
					list = append(list[:i], list[i+1:]...)
					break
				}
			}
			for i := 0; i < len(list); i++ {
				if list[i] == firstId+2 {
					list = append(list[:i], list[i+1:]...)
					break
				}
			}
			return huPaiPanDin(list)
		}
		return false
	}
}

func isSevenPair(list []int) bool {
	if len(list) != 14 {
		return false
	}
	dict := make(map[int]int) //id : count
	for i := 0; i < len(list); i++ {
		count, ok := dict[list[i]]
		if ok {
			count++
		} else {
			dict[list[i]] = 1
		}
	}
	for _, count := range dict {
		if count%2 != 0 {
			return false
		}
	}
	return true
}
