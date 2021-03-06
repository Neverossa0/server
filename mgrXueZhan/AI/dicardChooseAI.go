package AI

import (
	"sort"
)

type DiscardLevel int32

const (
	DiscardLevel1 DiscardLevel = 1
	DiscardLevel2 DiscardLevel = 2
	DiscardLevel3 DiscardLevel = 3
)

func GetDiscardID(cardList []int, mode DiscardLevel) int32 {
	cardIdList := make([]int, 0)
	cardIdList = append(cardIdList, cardList...)

	sort.Ints(cardIdList)
	switch mode {
	case DiscardLevel1:
		return int32(GetDiscardLevel1(cardIdList))
	case DiscardLevel2:
		return int32(GetDiscardLevel2(cardIdList))
	case DiscardLevel3:
	}
	return int32(cardIdList[0])
}

func GetDiscardLevel1(cardIdList []int) int {
	//single
	singleHua := getHuaNumList(cardIdList)
	for n := 0; n < len(singleHua); n++ {
		if singleHua[n] == 0 {
			continue
		}
		nexus := getNexus(n, singleHua)
		if nexus == 1 {
			return n + 1
		}
	}

	//common
	needhunMin := len(cardIdList)
	discard := 0
	for n := 0; n < len(cardIdList); n++ {
		tempList := make([]int, 0)
		tempList = append(tempList[:], cardIdList[:]...)
		tempList = append(tempList[:n], tempList[n+1:]...)

		hua := getHuaNumList(tempList)
		needhunList := make([]int, 3)
		for i := 0; i < 3; i++ {
			needhun := 0
			for j := 0; j < 3; j++ {
				isJiang := j == i
				arr := make([]int, 0)
				arr = append(arr[:], hua[:]...)
				needhun = needhun + getNeedHun(arr, j, !isJiang)
			}
			needhunList[i] = needhun
		}
		needhunCount := min(needhunList)
		if needhunCount < needhunMin {
			needhunMin = needhunCount
			discard = cardIdList[n]
		}

		n = n + getCardCount(cardIdList[n], cardIdList) - 1
	}
	return discard
}

func GetDiscardLevel2(cardIdList []int) int {
	//common
	needhunMin := len(cardIdList)
	discard := 0
	for n := 0; n < len(cardIdList); n++ {
		tempList := make([]int, 0)
		tempList = append(tempList[:], cardIdList[:]...)
		tempList = append(tempList[:n], tempList[n+1:]...)

		hua := getHuaNumList(tempList)
		needhunList := make([]int, 3)
		for i := 0; i < 3; i++ {
			needhun := 0
			for j := 0; j < 3; j++ {
				isJiang := j == i
				arr := make([]int, 0)
				arr = append(arr[:], hua[:]...)
				needhun = needhun + getNeedHun(arr, j, !isJiang)
			}
			needhunList[i] = needhun
		}
		needhunCount := min(needhunList)

		if needhunCount < needhunMin {
			needhunMin = needhunCount
			discard = cardIdList[n]
		} else if needhunCount == needhunMin {
			//single
			nexus1 := getNexus(cardIdList[n]-1, hua)
			nexus2 := getNexus(discard-1, hua)
			if nexus1 < nexus2 {
				needhunMin = needhunCount
				discard = cardIdList[n]
			}
		}

		n = n + getCardCount(cardIdList[n], cardIdList) - 1
	}
	return discard
}

func min(num []int) int {
	result := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] < result {
			result = num[i]
		}
	}
	return result
}

func getCardCount(id int, list []int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == id {
			count = count + 1
		}
	}
	return count
}

func getNeedHun(typeIdList []int, cardType int, hasJiang bool) int {

	needhun := dfs(typeIdList, cardType*10, cardType*10+8, 0, hasJiang)
	return needhun
}

func dfs(arr []int, cur int, upper int, needhun int, hasJiang bool) int {

	if cur > upper {
		if !hasJiang {
			needhun = needhun + 2
		}
		return needhun
	}
	if cur%10 == 6 && arr[cur+1] == 1 && arr[cur+2] == 1 {
		//789
		return del_list(arr, cur, upper, needhun, hasJiang)
	} else if arr[cur] == 0 {
		return dfs(arr, cur+1, upper, needhun, hasJiang)
	} else if cur%10 < 7 && (arr[cur+1] > 0 || arr[cur+2] > 0) {
		tmp1Arr := make([]int, len(arr))
		copy(tmp1Arr, arr)
		tmp2Arr := make([]int, len(arr))
		copy(tmp2Arr, arr)
		tmp1 := del_list(tmp1Arr, cur, upper, needhun, hasJiang)
		tmp2 := del_same(tmp2Arr, cur, upper, needhun, hasJiang)
		if tmp1 > tmp2 {
			return tmp2
		} else {
			return tmp1
		}
	} else {
		return del_same(arr, cur, upper, needhun, hasJiang)
	}
}

func del_list(arr []int, i int, j int, needhun int, hasJiang bool) int {
	for k := 0; k < 3; k++ {
		if arr[i+k] > 0 {
			arr[i+k] = arr[i+k] - 1
		} else {
			needhun = needhun + 1
		}
	}
	return dfs(arr, i, j, needhun, hasJiang)
}

func del_same(arr []int, i int, j int, needhun int, hasJiang bool) int {

	mod := arr[i] % 3
	switch mod {
	case 0:
	case 1:
		if hasJiang {
			needhun = needhun + 2
		} else {
			needhun = needhun + 1
			hasJiang = true
		}
	case 2:
		if hasJiang {
			needhun = needhun + 1
		} else {
			hasJiang = true
		}
	}
	arr[i] = 0
	return dfs(arr, i+1, j, needhun, hasJiang)
}

func getHuaNumList(cardIdList []int) []int {
	result := make([]int, 30)
	for _, id := range cardIdList {
		result[id-1] = result[id-1] + 1
	}

	return result
}

func getNexus(cur int, arr []int) int {

	if cur%10 == 8 {
		//9
		return arr[cur] + arr[cur-1] + arr[cur-2]
	} else if cur%10 == 7 {
		//8
		return arr[cur] + arr[cur+1] + arr[cur-1] + arr[cur-2]
	} else if cur%10 == 0 {
		//1
		return arr[cur] + arr[cur+1] + arr[cur+2]
	} else if cur%10 == 1 {
		//2
		return arr[cur] + arr[cur+1] + arr[cur+2] + arr[cur-1]
	} else {
		return arr[cur] + arr[cur+1] + arr[cur+2] + arr[cur-1] + arr[cur-2]
	}
}
