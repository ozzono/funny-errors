package messy

import (
	"errors"
	"funny-errors/utils"
)

var (
	ErrMessyFan                   = errors.New("error: the mess just hit the fan")
	ErrMessyHouston               = errors.New("houston, we've got a messy problem")
	ErrMessyBitSituation          = errors.New("whoops, looks like we're in a bit of a messy situation")
	ErrMessyDance                 = errors.New("sorry, we can't dance around this error - it's just messy")
	ErrMessyCodebase              = errors.New("looks like we've got a messy codebase")
	ErrMessyCapacity              = errors.New("warning: the mess level has exceeded maximum capacity")
	ErrMessySituation             = errors.New("well, that's just a messy situation of errors")
	ErrMessyLoop                  = errors.New("looks like we're caught in a messy loop")
	ErrMessyNotEnough             = errors.New("error: not enough mess in the code")
	ErrMessyMalfunction           = errors.New("houston, we've got a messy malfunction")
	ErrMessyInCodebase            = errors.New("who let the mess into the codebase?")
	ErrMessyBug                   = errors.New("looks like we've got a messy bug on our hands")
	ErrMessyWords                 = errors.New("sorry, this error is just too messy for words")
	ErrMessyCleanup               = errors.New("looks like we need to do clean this mess")
	ErrMessyCodeFunkery           = errors.New("looks like we've got a messy case of code funk-ery")
	ErrWellThisIsAMess            = errors.New("well, this is a fine mess we've gotten ourselves into...")
	ErrNoBigDealJustAMess         = errors.New("no big deal, just a little mess we need to clean up here")
	ErrWeDidntStartTheMess        = errors.New("we didn't start the mess, but it looks like we're the ones who have to clean it up")
	ErrThisIsQuiteAMess           = errors.New("this is quite a mess we've made, isn't it?")
	ErrWhoopsAMess                = errors.New("whoops, looks like we've got a bit of a mess on our hands")
	ErrThisIsAMessySituation      = errors.New("this is a messy situation - time to roll up our sleeves and start cleaning up this error")
	ErrMessyCodeMessyError        = errors.New("messy code, messy error - it's time to tidy up and make things right")
	ErrCanYouSmellTheMess         = errors.New("can you smell that? it's the sweet scent of a big ol' mess")
	ErrThisIsGettingMessy         = errors.New("this is getting messy - better buckle down and fix this error")
	ErrWellThatWasAMess           = errors.New("well, that was a mess. let's see if we can clean it up and make things right")
	ErrHotMess                    = errors.New("warning: hot mess detected. time to put on our error-handling gloves and get to work")
	ErrMessyMessyError            = errors.New("messy, messy error - looks like we need to clean up our act and write some better code")
	ErrThisIsARealMess            = errors.New("this is a real mess we've gotten ourselves into - better start digging ourselves out of this error")
	ErrMessyErrorButWorthIt       = errors.New("this error may be messy, but the end result will be worth the cleanup")
	ErrCanWePleaseCleanThisMessUp = errors.New("can we please clean this mess up? it's getting a little out of hand")
	errs                          = []error{
		ErrMessyFan,
		ErrMessyHouston,
		ErrMessyBitSituation,
		ErrMessyDance,
		ErrMessyCodebase,
		ErrMessyCapacity,
		ErrMessySituation,
		ErrMessyLoop,
		ErrMessyNotEnough,
		ErrMessyMalfunction,
		ErrMessyInCodebase,
		ErrMessyBug,
		ErrMessyWords,
		ErrMessyCleanup,
		ErrMessyCodeFunkery,
		ErrWellThisIsAMess,
		ErrNoBigDealJustAMess,
		ErrWeDidntStartTheMess,
		ErrThisIsQuiteAMess,
		ErrWhoopsAMess,
		ErrThisIsAMessySituation,
		ErrMessyCodeMessyError,
		ErrCanYouSmellTheMess,
		ErrThisIsGettingMessy,
		ErrWellThatWasAMess,
		ErrHotMess,
		ErrMessyMessyError,
		ErrThisIsARealMess,
		ErrMessyErrorButWorthIt,
		ErrCanWePleaseCleanThisMessUp,
	}
)

func RandomErr() error {
	return errs[utils.RInt(len(errs))]
}
