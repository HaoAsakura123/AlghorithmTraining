package main

import(
	"testing"
)
func TestMax(t *testing.T){
	//Arrange
	testTable := []struct{
		numbers []int
		expected int
	}{
		{
			numbers: []int{1,2,3,4,5},
			expected: 5,
		},
		{
			numbers: []int{1,2,3,124,5},
			expected: 124,
		},
		{
			numbers: []int{},
			expected: 0,
		},
	}

	for _, testCase := range testTable{
		//Act
		result:=Max(testCase.numbers)
		//Assert
		t.Logf("Calling Max(%v), got result %d", testCase.numbers, result)
		if result != testCase.expected{
			t.Errorf("Incorrect result, expected %d, got %d", testCase.expected, result)
		}
	}
}
