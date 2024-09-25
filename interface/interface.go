package _interface

// Signer interface for signing data
type Signer interface {
	Sign(data string) (string, error)
}

// Verifier interface for verifying signatures
type Verifier interface {
	Verify(data, signature string) bool
}

// EncodeDecode interface for Encode or Decode request
type EncodeDecode interface {
	Encode(data string) (string, error)
	Decode(data string) (string, error)
}

// Request interface for request
type Request interface {
	String() (string, error)
}

type Validate interface {
	Validate() error
}
