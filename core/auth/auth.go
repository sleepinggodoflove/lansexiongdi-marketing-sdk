package auth

type Auth interface {
	Sign(params map[string]string) (string, error)
	Verify(params map[string]string) bool
}
