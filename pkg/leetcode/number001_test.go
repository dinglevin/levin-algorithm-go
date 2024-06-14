package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSumV1(t *testing.T) {
	expected := []int{1, 2}
	result := twoSumV1([]int {3,2,4}, 6)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSumV1([]int{2,7,11,15}, 9)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSumV1([]int{3, 3}, 6)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}
}

func TestTwoSumV2(t *testing.T) {
	expected := []int{1, 2}
	result := twoSumV2([]int {3,2,4}, 6)
	sort.Ints(result)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSumV2([]int{2,7,11,15}, 9)
	sort.Ints(result)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSumV2([]int{3, 3}, 6)
	sort.Ints(result)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}
}

func TestTwoSumV3(t *testing.T) {
	expected := []int{1, 2}
	result := twoSumV3([]int {3,2,4}, 6)
	sort.Ints(result)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSumV3([]int{2,7,11,15}, 9)
	sort.Ints(result)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}

	expected = []int{0, 1}
	result = twoSumV3([]int{3, 3}, 6)
	sort.Ints(result)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("failed, expected %#v got %#v", expected, result)
	}
}