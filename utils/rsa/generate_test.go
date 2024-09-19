package rsa

import "testing"

func TestGenerate(t *testing.T) {
	n := NewGenerateKey()
	err := n.SavePem("../../pem")
	if err != nil {
		t.Fatal(err)
	}
	privateKeyStr, publicKeyStr := n.GetKey()
	t.Log("privateKeyStr=", privateKeyStr)
	t.Log("publicKeyStr=", publicKeyStr)
}

func TestGenerateSavePem(t *testing.T) {
	n := NewGenerateKey()
	err := n.SavePem("../../pem")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateGetKey(t *testing.T) {
	n := NewGenerateKey()
	privateKeyStr, publicKeyStr := n.GetKey()
	t.Log("privateKeyStr=", privateKeyStr)
	t.Log("publicKeyStr=", publicKeyStr)
}
