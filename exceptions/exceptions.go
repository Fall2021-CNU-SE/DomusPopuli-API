package exceptions

import (
    "errors"
)

var (
    UserNotFound = errors.New("user not found")
)
