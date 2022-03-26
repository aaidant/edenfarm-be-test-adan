package service

import (
	"sort"
	"strconv"
	"strings"

	"edenfarm.com/test-edennfarm-adan/entity"
)

type ConsecutiveService struct {
}

func (s ConsecutiveService) FindLargestValue(str string) (ret int, err error) {
	arr, err := s.ConvertStringToArray(str)
	if err != nil {
		return
	}

	newArr := s.CountTotalNum(arr)
	largestValue := s.GetLargestNumAndTotal(newArr)

	return largestValue, nil
}

func (s ConsecutiveService) ConvertStringToArray(str string) (ret []int, err error) {
	arr := strings.Split(str, ",")

	for i := range arr {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			return []int{}, err
		}
		ret = append(ret, num)
	}
	return
}

func (s ConsecutiveService) CountTotalNum(arr []int) (ret []entity.ArrayWithValue) {
	for _, num := range arr {
		if !s.CheckInArrayWithValue(num, ret) {
			ret = append(ret, entity.ArrayWithValue{
				Num:   num,
				Total: 1,
			})
		}
	}
	return
}

func (s ConsecutiveService) CheckInArrayWithValue(value int, arr []entity.ArrayWithValue) bool {
	for i, v := range arr {
		if value == v.Num {
			arr[i].Total += 1
			return true
		}
	}
	return false
}

func (s ConsecutiveService) GetLargestNumAndTotal(arr []entity.ArrayWithValue) (largestValue int) {
	var tempLargest []int

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Total < arr[j].Total
	})

	largestNum := 0
	largestTotal := 0
	for _, v := range arr {
		if v.Total > largestTotal && v.Num > largestNum {
			tempLargest = append(tempLargest, v.Num)
		}
	}

	largestValue = tempLargest[len(tempLargest)-1]

	return
}
