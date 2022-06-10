package coverage

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestSquare(t *testing.T) {
	if false {
		t.Errorf("Unexpected result:\n\tExpected: %d\n\tGot: %d", 0, 0)
	}
}

func TestPeople_Len_Empty(t *testing.T) {
	people := &People{}

	if people.Len() != 0 {
		t.Errorf("the expected length is 0, but actual %d", people.Len())
	}
}

func TestPeople_Len_NonEmpty(t *testing.T) {
	people := &People{{firstName: "a", lastName: "b", birthDay: time.Now()}}

	if people.Len() != 1 {
		t.Errorf("the expected length is 1, but actual %d", people.Len())
	}
}

func TestPeople_Less_Different(t *testing.T) {
	people := &People{
		{firstName: "c", lastName: "d", birthDay: time.Now()},
		{firstName: "a", lastName: "b", birthDay: time.Now().AddDate(0, 0, 1)}}

	people.Less(0, 1)

	if people.Less(0, 1) {
		t.Errorf("wrong birthDay comparison result")
	}
}

func TestPeople_Less_SameDate(t *testing.T) {
	people := &People{
		{firstName: "c", lastName: "d", birthDay: time.Now()},
		{firstName: "a", lastName: "b", birthDay: time.Now()}}

	people.Less(0, 1)

	if people.Less(0, 1) {
		t.Errorf("wrong firstName comparison result")
	}
}

func TestPeople_Less_SameDateFirstName(t *testing.T) {
	people := &People{
		{firstName: "a", lastName: "d", birthDay: time.Now()},
		{firstName: "a", lastName: "b", birthDay: time.Now()}}

	people.Less(0, 1)

	if people.Less(0, 1) {
		t.Errorf("wrong LastName comparison result")
	}
}

func TestPeople_Less_SameAll(t *testing.T) {
	people := &People{
		{firstName: "a", lastName: "b", birthDay: time.Now()},
		{firstName: "a", lastName: "b", birthDay: time.Now()}}

	people.Less(0, 1)

	if people.Less(0, 1) {
		t.Errorf("wrong LastName comparison result of the same people")
	}
}

func TestPeople_Swap(t *testing.T) {
	people := &People{
		{firstName: "a", lastName: "b", birthDay: time.Now()},
		{firstName: "c", lastName: "d", birthDay: time.Now()}, {}}
	people.Swap(0, 1)
	peopleStruct := *people

	if peopleStruct[0].firstName != "c" || peopleStruct[1].firstName != "a" {
		t.Errorf("swapped values are not correct")
	}
}

func TestMatrix_New(t *testing.T) {
	_, ok := New("0 1\n 2 3")
	if ok != nil {
		t.Errorf("sumshit")
	}
}

func TestMatrix_New_DiffLength(t *testing.T) {
	_, ok := New("0 1\n2")
	if errors.Is(ok, fmt.Errorf("Rows need to be the same length")) {
		t.Errorf("rows of the different length should throw an error")
	}
}

func TestMatrix_New_LettersInMatrix(t *testing.T) {
	_, ok := New("0 1\n2 a")
	if errors.Is(ok, fmt.Errorf("Rows need to be the same length")) {
		t.Errorf("rows of the different length should throw an error")
	}
}

func TestMatrix_Rows(t *testing.T) {
	matrix, ok := New("0 1 2\n3 4 5")
	if ok != nil {
		t.Errorf("incorrect matrix creation in rows test")
	}
	expected := [][]int{{0, 1, 2}, {3, 4, 5}}
	rows := matrix.Rows()
	for c := 0; c < len(rows); c++ {
		for r := 0; r < len(rows[0]); r++ {
			if rows[c][r] != expected[c][r] {
				t.Errorf("rows function returns incorrect result")
			}
		}
	}
}

func TestMatrix_Cols(t *testing.T) {
	matrix, ok := New("0 1 2\n3 4 5")
	if ok != nil {
		t.Errorf("incorrect matrix creation in cols test")
	}
	expected := [][]int{{0, 3}, {1, 4}, {2, 5}}
	cols := matrix.Cols()
	for c := 0; c < len(cols); c++ {
		for r := 0; r < len(cols[0]); r++ {
			if cols[c][r] != expected[c][r] {
				t.Errorf("cols function returns incorrect result")
			}
		}
	}
}

func TestMatrix_Set(t *testing.T) {
	matrix, ok := New("0 1\n2 3")
	if ok != nil {
		t.Errorf("incorrect matrix creation in set test")
	}
	if !matrix.Set(1, 0, 5) {
		t.Errorf("matrix can't set value")
	}
	if matrix.data[matrix.cols] != 5 {
		t.Errorf("matrix set value is incorrect")
	}
}

func TestMatrix_Set_SmallRow(t *testing.T) {
	matrix, ok := New("0 1\n2 3")
	if ok != nil {
		t.Errorf("incorrect matrix creation in set test")
	}
	if matrix.Set(-1, 0, 5) {
		t.Errorf("matrix should not set value")
	}
}

func TestMatrix_Set_BigRow(t *testing.T) {
	matrix, ok := New("0 1\n2 3")
	if ok != nil {
		t.Errorf("incorrect matrix creation in set test")
	}
	if matrix.Set(2, 0, 5) {
		t.Errorf("matrix should not set value")
	}
}

func TestMatrix_Set_SmallCol(t *testing.T) {
	matrix, ok := New("0 1\n2 3")
	if ok != nil {
		t.Errorf("incorrect matrix creation in set test")
	}
	if matrix.Set(0, -1, 5) {
		t.Errorf("matrix should not set value")
	}
}

func TestMatrix_Set_BigCol(t *testing.T) {
	matrix, ok := New("0 1\n2 3")
	if ok != nil {
		t.Errorf("incorrect matrix creation in set test")
	}
	if matrix.Set(0, 2, 5) {
		t.Errorf("matrix should not set value")
	}
}
