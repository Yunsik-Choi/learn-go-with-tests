package pointererror

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		t.Helper()

		wallet := Wallet{}

		wallet.Deposit(10)

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("expected %s but actual %s", expected, actual)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		t.Helper()

		wallet := Wallet{balance: Bitcoin(20)}
		wallet.withDraw(10)

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if actual != expected {
			t.Errorf("expected %s but actual %s", expected, actual)
		}
	})

	t.Run("withdraw amount over then balance", func(t *testing.T) {
		t.Helper()

		balance := Bitcoin(5)
		wallet := Wallet{balance: balance}

		withDrawAmount := Bitcoin(10)
		err := wallet.withDraw(withDrawAmount)

		if err == nil {
			t.Errorf("wallet balance is %s but withDraw %s are working", balance, withDrawAmount)
		}
	})
}
