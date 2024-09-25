package interfaces

// Signer interfaces for signing data
type Signer interface {
	Sign(data string) (string, error)
}

// Verifier interfaces for verifying signatures
type Verifier interface {
	Verify(data, signature string) bool
}

// EncodeDecode interfaces for Encode or Decode request
type EncodeDecode interface {
	Encode(plaintext string) (string, error)
	Decode(ciphertext string) (string, error)
}

// Request interfaces for request
type Request interface {
	String() (string, error)
}

type Validate interface {
	Validate() error
}
