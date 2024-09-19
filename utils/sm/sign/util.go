package sign

import (
	"fmt"
	"math/big"
)

const RsLen = 32

func rsAsn1ToPlainByteArray(rr, ss *big.Int) []byte {
	r := bigIntToFixedLengthBytes(rr)
	s := bigIntToFixedLengthBytes(ss)

	result := make([]byte, RsLen*2)
	copy(result[:RsLen], r)
	copy(result[RsLen:], s)

	return result
}

func bigIntToFixedLengthBytes(rOrS *big.Int) []byte {
	rs := rOrS.Bytes()
	if len(rs) == RsLen {
		return rs
	} else if len(rs) == RsLen+1 && rs[0] == 0 {
		return rs[1 : RsLen+1]
	} else if len(rs) < RsLen {
		result := make([]byte, RsLen)
		for i := range result {
			result[i] = 0
		}
		copy(result[RsLen-len(rs):], rs)
		return result
	} else {
		panic(fmt.Sprintf("err rs: %x", rs))
	}
}

func rsPlainByteArrayToAsn1(rs []byte) []byte {
	result := make([]byte, RsLen*2)
	copy(result[:RsLen], rs[:RsLen])
	copy(result[RsLen:], rs[RsLen:])
	return result
}
