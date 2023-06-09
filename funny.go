package funny

import (
	"errors"
	"funny-errors/funky"
	greekMyth "funny-errors/greek-myth"
	"funny-errors/marvin"
	"funny-errors/messy"
	"funny-errors/optimistic"
	"funny-errors/pessimist"
	popCulture "funny-errors/pop/pop-culture"
	popSongs "funny-errors/pop/pop-songs"
	"funny-errors/sarcasm"
	"funny-errors/utils"
)

var (
	ErrNotFunnyEnough       = errors.New("error: code not funny enough")
	ErrFunNotFound          = errors.New("error: fun not found")
	ErrFunnyBusiness        = errors.New("looks like there's some funny business going on here")
	ErrTooMuchFunny         = errors.New("warning: the code has exceeded maximum funny levels")
	ErrLaughingOutOfControl = errors.New("sorry, our error handling seems to be laughing out of control")
	ErrComedyOfErrors       = errors.New("well, that's just a comedy of errors")
	ErrPunchlineMissing     = errors.New("error: punchline missing")
	ErrJokeTooOld           = errors.New("sorry, this error is older than a dad joke")
	ErrHumorNotFound        = errors.New("error: humor not found")
	ErrFunSize              = errors.New("looks like we've got a fun-size error")
	ErrPrankCall            = errors.New("whoops, looks like we prank-called the error handling system")
	ErrFunnyBoneBroken      = errors.New("warning: our funny bone seems to be broken")
	ErrLaughTrackNeeded     = errors.New("sorry, we're missing the laugh track for this error")
	ErrSillyMistake         = errors.New("looks like we made a silly mistake")
	ErrHilarityEnsues       = errors.New("warning: hilarity may ensue")
	ErrFunnyCodeFunkery     = errors.New("looks like we've got a classic case of funny code")
	ErrNasty                = errors.New("something nasty happened here")
	ErrFunky                = errors.New("something funky happened here")
	ErrMessy                = errors.New("who did this mess?")
	ErrDontSee              = errors.New("bummer - you should not be seeing this")
	ErrDontLook             = errors.New("ops, don't look at it now")
	ErrDontLookDown         = errors.New("don't look down")
	ErrDontLookAtMe         = errors.New("don't look at me")
	ErrBlame                = errors.New("you better find someone to blame")
	ErrRun                  = errors.New("you better start running")
	ErrGotcha               = errors.New("gotcha!")
	ErrFreeze               = errors.New("freeze!")
	ErrWeird                = errors.New("this is weird")
	ErrUnexpected           = errors.New("this is unexpected")
	ErrScratchHead          = errors.New("~scratching the head~")
	ErrChinThink            = errors.New("~thinking with hand on chin~")
	ErrFrown                = errors.New("~frowning~")
	ErrFacepalm             = errors.New("~facepalming~")
	ErrHmm                  = errors.New("hmmm")
	errs                    = []error{
		ErrNotFunnyEnough,
		ErrFunNotFound,
		ErrFunnyBusiness,
		ErrTooMuchFunny,
		ErrLaughingOutOfControl,
		ErrComedyOfErrors,
		ErrPunchlineMissing,
		ErrJokeTooOld,
		ErrHumorNotFound,
		ErrFunSize,
		ErrPrankCall,
		ErrFunnyBoneBroken,
		ErrLaughTrackNeeded,
		ErrSillyMistake,
		ErrHilarityEnsues,
		ErrFunnyCodeFunkery,
		ErrNasty,
		ErrFunky,
		ErrMessy,
		ErrDontSee,
		ErrDontLook,
		ErrDontLookDown,
		ErrDontLookAtMe,
		ErrBlame,
		ErrRun,
		ErrGotcha,
		ErrFreeze,
		ErrWeird,
		ErrUnexpected,
		ErrScratchHead,
		ErrChinThink,
		ErrFrown,
		ErrFacepalm,
		ErrHmm,
	}
)

func RandomErr() error {
	return errs[utils.RInt(len(errs))]
}

func RandomErrFromAll() error {
	errs := []error{
		popCulture.RandomErr(),
		popSongs.RandomErr(),
		greekMyth.RandomErr(),
		funky.RandomErr(),
		messy.RandomErr(),
		pessimist.RandomErr(),
		optimistic.RandomErr(),
		sarcasm.RandomErr(),
		marvin.RandomErr(),
	}
	return errs[utils.RInt(len(errs))]
}
