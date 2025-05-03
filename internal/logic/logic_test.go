package logic

import "testing"

func TestSum(t *testing.T) {
	total := Sum(1, 2, 3, 4, 5)
	expected := 15
	if total != expected {
		t.Errorf("Expected 15, got %d", total)
	}

	result := internalSum(4, 5)
	expected = 9
	if result != expected {
		t.Errorf("Expected 9, got %d", result)
	}
}

func TestCheckEmail(t *testing.T) {
	result := CheckEmail("test@test.com")
	expected := true
	if result != expected {
		t.Errorf("CheckEmail expected %t, got %t", expected, result)
	}

	emailInvalid := "testtest.com.ar"
	result = CheckEmail(emailInvalid)
	expected = false
	if result != expected {
		t.Errorf("CheckEmail expected %t, got %t", expected, result)
	}
}
