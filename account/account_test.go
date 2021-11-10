package account

import "testing"

func TestWithdraw(t *testing.T) {
	balance := 1000.0
	result := Withdraw(100, balance)

	if result != 900 {
		t.Errorf("Withdraw(100, 100) = %f; want 900", result)
	}

	result = Withdraw(1004, balance)
	if result != 1000.0 {
		t.Errorf("Expecting 1000")
	}

	result = Withdraw(1000, balance)
	if result != 0 {
		t.Errorf("Expecting 0")
	}
}
