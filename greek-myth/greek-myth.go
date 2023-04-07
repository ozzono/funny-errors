package greekMyth

import "errors"

var (
	ErrMinotaur = errors.New("the error was trapped in a labyrinth and cannot be found")
	ErrHydra    = errors.New("the error encountered a multi-headed problem that cannot be resolved")
	ErrMedusa   = errors.New("the output turned into stone due to an unexpected error")
	ErrSisyphus = errors.New("the error has been repeatedly failed and cannot be completed")
	ErrCerberus = errors.New("the error was denied access to the requested resource by multiple guards")
)
