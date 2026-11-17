package game

import "errors"

// ErrNameInUse is returned by add when another player holds the name.
var ErrNameInUse = errors.New("name already in use")
