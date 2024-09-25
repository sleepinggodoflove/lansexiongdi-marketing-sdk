package sm

import (
	"testing"
)

func TestSM4GenerateKey(t *testing.T) {
	key, err := GenerateSM4Key()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(key)
}

func TestSM4(t *testing.T) {
	key := "t+VxHnp+K9huhtNT84Pk7A=="
	plaintextBytes := []byte("BZjU223ZBM7A8586Tm7P")
	enc, err := Encode(key, plaintextBytes)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(enc)
	dec, err := Decode(key, enc)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(dec)
}

func TestSM4KeyEncrypt(t *testing.T) {
	key := "t+VxHnp+K9huhtNT84Pk7A=="
	plaintextBytes := []byte("BZjU223ZBM7A8586Tm7P")
	enc, err := Encode(key, plaintextBytes)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(enc)
}

func TestSM4KeyPassDecrypt(t *testing.T) {
	key := "t+VxHnp+K9huhtNT84Pk7A=="
	ciphertext := "NwANcXkjX79873jenLJRGhbEr39eYOwC5WQxZFXmLpw="
	dec, err := Decode(key, ciphertext)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(dec)
}
