package sm4

import "testing"

var key = "t+VxHnp+K9huhtNT84Pk7A=="

var originalStr = "123456"

func Test_encrypt(t *testing.T) {
	want := "Un2GM6FnYMR3fC8nOWCOSA=="
	t.Run("Test_encrypt", func(t *testing.T) {
		if got := encrypt(originalStr, key); got != want {
			t.Errorf("encrypt() = %v, want %v", got, want)
		}
	})
}
