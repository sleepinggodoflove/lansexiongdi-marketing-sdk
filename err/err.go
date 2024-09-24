package err

import "errors"

var ErrUnsupportedSignType = errors.New("unsupported signType")
var ErrInvalidSignature = errors.New("invalid signature")
