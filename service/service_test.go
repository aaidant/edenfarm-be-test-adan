package service

import (
	"reflect"
	"testing"

	"edenfarm.com/test-edennfarm-adan/entity"
)

func TestService_FindLargestValue(t *testing.T) {
	s := ConsecutiveService{}

	args := []string{
		"1,2,3,8,9,3,2,1",
		"1,2,1,2,2,1",
		"7,1,2,9,7,2,1",
	}

	expect := []int{
		3,
		2,
		2,
	}

	for i := 0; i < len(args); i++ {
		result, err := s.FindLargestValue(args[i])

		if !reflect.DeepEqual(result, expect[i]) {
			t.Fail()
		}

		if err != nil {
			t.Error()
		}
	}
}

func TestService_FindLargestValue_Error(t *testing.T) {
	s := ConsecutiveService{}

	args := "1, 2, 3, 4"

	_, err := s.FindLargestValue(args)

	if err != nil {
		t.Skipped()
	} else {

		t.Error()
	}
}

func TestService_ConvertStringToArray(t *testing.T) {
	s := ConsecutiveService{}

	args := "1,2,3,4"

	expect := []int{1, 2, 3, 4}

	result, err := s.ConvertStringToArray(args)

	if !reflect.DeepEqual(result, expect) {
		t.Fail()
	}

	if err != nil {
		t.Error()
	}
}

func TestService_CountTotalNum(t *testing.T) {
	s := ConsecutiveService{}

	args := []int{1, 1, 2}

	var expect []entity.ArrayWithValue

	expect = append(expect, entity.ArrayWithValue{
		Num:   1,
		Total: 2,
	})

	expect = append(expect, entity.ArrayWithValue{
		Num:   2,
		Total: 1,
	})

	result := s.CountTotalNum(args)

	if !reflect.DeepEqual(result, expect) {
		t.Fail()
	}

}

func TestService_CheckInArrayWithValue(t *testing.T) {
	s := ConsecutiveService{}
	args := 1
	fields := []entity.ArrayWithValue{
		{
			Num:   1,
			Total: 2,
		},
	}
	result := s.CheckInArrayWithValue(args, fields)

	if !result {
		t.Fail()
	}
}

func TestService_GetLargestNumAndTotal(t *testing.T) {
	s := ConsecutiveService{}

	args := []entity.ArrayWithValue{
		{
			Num:   1,
			Total: 2,
		},
		{
			Num:   2,
			Total: 2,
		},
		{
			Num:   3,
			Total: 1,
		},
	}
	expected := 2

	result := s.GetLargestNumAndTotal(args)

	if result != expected {
		t.Fail()
	}
}
