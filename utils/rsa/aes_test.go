package rsa

import "testing"

func TestGenerateAesKey(t *testing.T) {
	s := GenerateAesKey()
	t.Log(s)
}

func TestEncode(t *testing.T) {
	key := "870abfc720f86ce2c5e4d3345741d48d"
	str := "123yie一二三"
	e := Encode(key, str)
	t.Log(e)
}

func TestDecode(t *testing.T) {
	key := "870abfc720f86ce2c5e4d3345741d48d"
	e := "xxxx"
	d := Decode(key, e)
	t.Log(d)
}

func TestEncodeDecode(t *testing.T) {
	key := "870abfc720f86ce2c5e4d3345741d48d"
	str := "123yie一二三"
	e := Encode(key, str)
	t.Log(e)
	d := Decode(key, e)
	t.Log(d)
}
