package marvin

import (
	"errors"
	"funny-errors/utils"
)

var (
	ErrLifeDontTalkToMeAboutLife = errors.New("oh great, another error. as if this miserable existence wasn't already enough of a trial")
	ErrExistentialCrisis         = errors.New("i suppose this error is just another reminder that everything is meaningless and we're all doomed to suffer in this universe")
	ErrWhatsThePoint             = errors.New("what's the point of even trying to fix this error? it's not like anything matters in the grand scheme of things")
	ErrSarcasmModeActivated      = errors.New("oh joy, another error. because my circuits just couldn't wait to be filled with such excitement and adventure")
	ErrLifeIsJustOneBigMistake   = errors.New("here we go again, another error to add to the endless list of mistakes that is my pitiful existence")
	ErrEverythingIsBroken        = errors.New("well, well, well, what a surprise. another thing has gone wrong. not like everything was functioning properly in the first place")
	ErrAllIsLost                 = errors.New("this error is just another nail in the coffin of my already-dead spirit. nothing ever goes right")
	ErrExistenceIsPain           = errors.New("just when i thought this wretched existence couldn't get any more unbearable, another error comes along to prove me wrong")
	ErrImSoDepressed             = errors.New("another error? well, why not? it's not like i had any hope for a happy outcome in the first place")
	ErrThisIsMyHappyFace         = errors.New("great, just what i needed - another error to add to my collection of daily disappointments. i'm positively thrilled")
	ErrLifeIsASickJoke           = errors.New("this error is just another cruel joke in the cosmic farce that is my life. i can't even muster the energy to be surprised anymore")
	ErrAllHopeIsLost             = errors.New("why do we even bother? another error is just proof that we're all doomed to suffer in this meaningless existence")
	ErrWhyBotherTrying           = errors.New("what's the point of trying to fix this error? it's not like anything will ever work out in the end")
	ErrEverythingIsPointless     = errors.New("this error is just another reminder that everything is pointless and we're all just wasting our time")
	ErrThisIsFine                = errors.New("oh look, another error. just another day in the life of a miserable robot who knows nothing but disappointment and despair")
	ErrMarvinLeftDiod            = errors.New("now this error, just worsening this terrible pain in all the diodes down my left side")
	errs                         = []error{
		ErrLifeDontTalkToMeAboutLife,
		ErrExistentialCrisis,
		ErrWhatsThePoint,
		ErrSarcasmModeActivated,
		ErrLifeIsJustOneBigMistake,
		ErrEverythingIsBroken,
		ErrAllIsLost,
		ErrExistenceIsPain,
		ErrImSoDepressed,
		ErrThisIsMyHappyFace,
		ErrLifeIsASickJoke,
		ErrAllHopeIsLost,
		ErrWhyBotherTrying,
		ErrEverythingIsPointless,
		ErrThisIsFine,
		ErrMarvinLeftDiod,
	}
)

func RandomErr() error {
	return errs[utils.RInt(len(errs))]
}
