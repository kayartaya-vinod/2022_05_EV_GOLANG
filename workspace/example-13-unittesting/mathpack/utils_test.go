package mathpack

import "testing"

func Test_ShouldAddTwoNumbers(t *testing.T) {
	got, _ := Sum("10, 20")
	want := 30

	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func Test_ShouldAddMultipleNumbers(t *testing.T) {
	got, _ := Sum("10, 20, 30, 40, 50")
	want := 150
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func Test_ShouldAddNegativeNumbers(t *testing.T) {
	got, _ := Sum("-10, -20, -30, -40, -50")
	want := -150
	if got != want {
		t.Errorf("Wanted %v, but got %v", want, got)
	}
}

func Test_ShouldGiveErrorForNonNumericals(t *testing.T) {
	_, err := Sum("10, asd")
	if err == nil {
		t.Error("Expecting an error; bot got nil")
	}
}
