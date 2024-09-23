package core

// Signer interface for signing data
type Signer interface {
	Sign(data string) (string, error)
}

// Verifier interface for verifying signatures
type Verifier interface {
	Verify(data, signature string) (bool, error)
}
