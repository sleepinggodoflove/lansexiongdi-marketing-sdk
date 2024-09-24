package rsa

import "testing"

func TestDecode(t *testing.T) {
	s := Decode("SZvDfHBmoGvcxjCiHoeAKkrGkxlNYSIS+TJcQXqSLWM=", "bcee0c6753b2a31c792a91fe9f9f1666")
	t.Log(s)
}
