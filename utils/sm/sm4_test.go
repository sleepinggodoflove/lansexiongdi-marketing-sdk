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
	encryptKey := "t+VxHnp+K9huhtNT84Pk7A=="
	enc, err := Encode([]byte("BZjU223ZBM7A8586Tm7P"), encryptKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(enc)
	dec, err := Decode(enc, encryptKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(dec)
}

func TestSM4KeyEncrypt(t *testing.T) {
	sm4key := "z9DoIVLuDYEN/qsgweRA4A=="
	enc, err := Encode([]byte("gQbHNecjZqnNZg3vKE2"), sm4key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(enc)
}

func TestSM4KeyPassDecrypt(t *testing.T) {
	sm4key := "z9DoIVLuDYEN/qsgweRA4A=="
	dec, err := Decode("NwANcXkjX79873jenLJRGhbEr39eYOwC5WQxZFXmLpw=", sm4key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(dec)
}
