package main

import (
	"testing"
)


/**
 * Тестирование функции A + B
 */
func TestAdd(t *testing.T) {
	
	type Test struct {
		arg1, arg2, expected int
	}
	
	var tests = []Test {
		Test{2, 3, 5},
		Test{4, 8, 12},
		Test{6, 9, 15},
		Test{3, 10, 13},
	}
	
	for _, test := range tests {
		
		output := Add(test.arg1, test.arg2);
		
		if output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected);
		}
	}
}


/**
 * Тестирование функции A - B
 */
func TestSub(t *testing.T) {
	
	type Test struct {
		arg1, arg2, expected int
	}
	
	var tests = []Test {
		Test{2, 3, -1},
		Test{8, 4, 4},
		Test{9, 6, 3},
		Test{3, 10, -7},
	}
	
	for _, test := range tests {
		
		output := Sub(test.arg1, test.arg2);
		
		if output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected);
		}
	}
}


/**
 * Тестирование функции A * B
 */
 func TestMult(t *testing.T) {
	
	type Test struct {
		arg1, arg2, expected int
	}
	
	var tests = []Test {
		Test{2, 3, 6},
		Test{8, 0, 0},
		Test{5, -2, -10},
		Test{-2, -2, 4},
	}
	
	for _, test := range tests {
		
		output := Mult(test.arg1, test.arg2);
		
		if output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected);
		}
	}
}
