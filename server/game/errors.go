package game

import "errors"

// ErrNameInUse is returned by Registry.Add when another player holds the name.
var ErrNameInUse = errors.New("name already in use")
