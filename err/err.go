package err

import "errors"

var ErrUnsupportedSignType = errors.New("unsupported SignType")
var ErrInvalidSignature = errors.New("invalid signature")
