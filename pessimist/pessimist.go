package pessimist

import (
	"errors"
	"funny-errors/utils"
)

var (
	ErrLifeAMess             = errors.New("life's a mess, what did you expect?")
	ErrNotGonnaWork          = errors.New("it's not going to work, and it's never going to work")
	ErrWasteOfTime           = errors.New("you're wasting your time, and mine")
	ErrGiveUpNow             = errors.New("might as well give up now, it's not worth the effort")
	ErrSameOldStory          = errors.New("something's gone wrong, what else is new?")
	ErrAnotherDayFailure     = errors.New("another day, another failure")
	ErrAlwaysTheSame         = errors.New("it's always the same old story with you")
	ErrIKnewThisWouldHappen  = errors.New("i knew this would happen, i just knew it")
	ErrNoPointTrying         = errors.New("don't bother trying again, it won't make a difference")
	ErrUniverseConspiring    = errors.New("the universe is conspiring against us, as usual")
	ErrEverythingBroken      = errors.New("everything is broken, and we're all doomed")
	ErrFullOfDisappointments = errors.New("just another reminder that life is full of disappointments")
	ErrToldYouSo             = errors.New("i told you this wouldn't work, but no one ever listens to me")
	ErrWhyBother             = errors.New("why bother trying? it's not like anything ever goes right")
	ErrEverythingTerrible    = errors.New("it's not your fault, everything is just terrible")
	errs                     = []error{
		ErrLifeAMess,
		ErrNotGonnaWork,
		ErrWasteOfTime,
		ErrGiveUpNow,
		ErrSameOldStory,
		ErrAnotherDayFailure,
		ErrAlwaysTheSame,
		ErrIKnewThisWouldHappen,
		ErrNoPointTrying,
		ErrUniverseConspiring,
		ErrEverythingBroken,
		ErrFullOfDisappointments,
		ErrToldYouSo,
		ErrWhyBother,
		ErrEverythingTerrible,
	}
)

func RandomErr() error {
	return errs[utils.RInt(len(errs))]
}
