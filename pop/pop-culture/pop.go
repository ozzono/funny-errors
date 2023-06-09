package popCulture

import (
	"errors"
	"funny-errors/utils"
)

var (
	Err404NotFoundInSpace                   = errors.New("houston, we have a problem. error 404: file not found... in space")
	ErrLukeIAmYourError                     = errors.New("luke, i am your error")
	ErrToInfinityAndBeyond                  = errors.New("error: to infinity and beyond is not a valid input")
	ErrHastaLaVistaBaby                     = errors.New("error terminated - hasta la vista, baby")
	ErrZoolanderBlueSteel                   = errors.New("warning: blue steel mode engaged, but the code is still not quite as good-looking as derek zoolander")
	ErrTheresNoPlaceLikeError               = errors.New("error: there's no place like error - looks like we're not in kansas anymore")
	ErrMayTheForceBeWithYou                 = errors.New("may the force be with you - unless you encounter this error, then you're on your own")
	ErrWinterIsComing                       = errors.New("warning: winter is coming - better bundle up that error handling")
	ErrBeamMeUpScotty                       = errors.New("error: we tried to beam you up, scotty - but it looks like you encountered an error on the way")
	ErrYippeeKiYay                          = errors.New("yippee ki-yay, error handler. sorry, couldn't resist a good die hard reference")
	ErrHoustonWeHaveAProblem                = errors.New("error: houston, we have a problem. looks like the code has encountered an unexpected issue")
	ErrShowMeTheError                       = errors.New("show me the error! (or at least a clear and informative error message)")
	ErrTheOneError                          = errors.New("the one error to rule them all")
	ErrMyPrecious                           = errors.New("my precious")
	ErrIveGotAFeelingWereNotInKansasAnymore = errors.New("warning: i've got a feeling we're not in kansas anymore - looks like we've encountered an unexpected error")
	ErrItsAMeMario                          = errors.New("error: it's a me, mario! looks like there's a problem with the code - better call in luigi for backup")
	ErrSimply                               = errors.New("one does not simply fix this error")
	ErrVitoCorleone                         = errors.New("the output made is an error it couldn't refuse")
	ErrMichaelCorleone                      = errors.New("the output was an error to the family and cannot be completed")
	ErrSonnyCorleone                        = errors.New("the error was overly aggressive and led to its own demise")
	ErrClemenza                             = errors.New("the error was handled with a double-cross and cannot be completed")
	ErrFredoSantana                         = errors.New("the error was a failure and cannot be trusted to lead any future errors")
	ErrGollum                               = errors.New("the error operation was precious to the system, but it has been lost forever")
	ErrSauron                               = errors.New("the error was an act of defiance against the all-seeing eye and cannot be completed")
	ErrGandalf                              = errors.New("the error may come at a great cost, and it is not clear whether it will succeed")
	ErrHobbit                               = errors.New("the error may seem small and insignificant, but it can still have a big impact on the system")
	ErrNazgul                               = errors.New("the error was a dark and ominous presence that cannot be allowed to corrupt the system")
	errs                                    = []error{
		Err404NotFoundInSpace,
		ErrLukeIAmYourError,
		ErrToInfinityAndBeyond,
		ErrHastaLaVistaBaby,
		ErrZoolanderBlueSteel,
		ErrTheresNoPlaceLikeError,
		ErrMayTheForceBeWithYou,
		ErrWinterIsComing,
		ErrBeamMeUpScotty,
		ErrYippeeKiYay,
		ErrHoustonWeHaveAProblem,
		ErrShowMeTheError,
		ErrTheOneError,
		ErrMyPrecious,
		ErrIveGotAFeelingWereNotInKansasAnymore,
		ErrItsAMeMario,
		ErrSimply,
		ErrVitoCorleone,
		ErrMichaelCorleone,
		ErrSonnyCorleone,
		ErrClemenza,
		ErrFredoSantana,
		ErrGollum,
		ErrSauron,
		ErrGandalf,
		ErrHobbit,
		ErrNazgul,
	}
)

func RandomErr() error {
	return errs[utils.RInt(len(errs))]
}
