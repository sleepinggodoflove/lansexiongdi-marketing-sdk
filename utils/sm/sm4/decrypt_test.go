package sm4

import (
	"encoding/base64"
	"log"
	"testing"
)

func Test_generateSM4Key(t *testing.T) {
	t.Run("Test_generateSM4Key", func(t *testing.T) {
		e := generateKey()
		t.Log("Test_generateSM4Key, send:", e)
	})
}

func Test_decrypt(t *testing.T) {
	t.Run("Test_decrypt", func(t *testing.T) {
		os := "{\"code\":\"OPEN25812\",\"msg\":\"[OPEN25812]数字信封制作失败，请检查签名/验签密钥或加密算法配置\",\"traceId\":\"OPEN-00-LOCAL-812\"}"
		b := []byte(os)
		s := base64.StdEncoding.EncodeToString(b)
		t.Log("base64 encode:", s)
		e := encrypt(s, key)
		t.Log("encrypt:", e)
		d := decrypt(e, key)
		if d != os {
			t.Errorf("decrypt() = %v, want %v", d, os)
		} else {
			t.Log(originalStr, e, d)
		}
	})
}

func Test_decrypt2(t *testing.T) {
	t.Run("Test_decrypt2", func(t *testing.T) {
		e := "yFNt/39YghQfKDa2tHEe8UhZYMshS2vT5nkUkcaZCGBVqRoHbcbisjuk4ndI9ZtxPs9VFkTyjy75a99jCCCIqGRHtG2JK0qxUHkvx+F2ulxJeMHsQKWj4CMz06Xu7ghaduZM7UluPYnyYuqirBKwhkjfi5w6TEm05s+bUDCy8Qhdjyj6IzabBUKx4mXhlZxhLAnO44AImAdcPJRBgsII5mCOEohCLxlrAjH7/PRPOoou8+Wuh6OGrY1z8i9Q2BzqOKuq77L6rmPgPxYkot1kj0Y0uTfeaJN0IB7xeyvTQCaWh8W8OmHOReEYqQbhpstnLsnpSpaqMwosFpdpERhixQ=="
		d := decrypt(e, key)
		log.Println(d)
	})
}
