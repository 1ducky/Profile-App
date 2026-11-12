package err

import "errors"

const InternalErrCode ErrCode = "INTERNAL_ERROR"

var InternalErr error = errors.New("Internal Error")
