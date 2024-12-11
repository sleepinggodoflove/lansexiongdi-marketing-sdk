package core

// Signer interfaces for signing data
type Signer interface {
	Sign(data string) (string, error)
}

// Verifier interfaces for verifying signatures
type Verifier interface {
	Verify(data, signature string) bool
}

// Cipher interfaces for Encode or Decode request
type Cipher interface {
	Encode(plaintext string) (string, error)
	Decode(ciphertext string) (string, error)
}

// CryptographySuite .
type CryptographySuite struct {
	Signer   Signer
	Verifier Verifier
	Cipher   Cipher
}

// Request interfaces for request
type Request interface {
	String() (string, error)
	Validate() error
}
