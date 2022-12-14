package cerrors

import (
	"errors"
	"fmt"
)

var (
	NotAuthorized    = errors.New("not authorized")
	NotPermitted     = errors.New("not permitted")
	NotFound         = errors.New("not found")
	AlreadyExists    = errors.New("already exists")
	IsNotPsy         = errors.New("is not psy")
	IsNotAdmin       = errors.New("is not admin")
	IsNotPsyAndAdmin = errors.New("is not psy and admin")
	BadToken         = errors.New("bad token")
	NoPostgres       = errors.New("no postgres is not supported")
	NoJwks           = errors.New("no jwks")
)

var (
	IncorrectLang            = errors.New("incorrect lang")
	OrderOffsetLimitRequired = errors.New("orderOffsetLimit required")
)

// article, video

var (
	AlreadyPublished = errors.New("already published")
	NotPublished     = errors.New("not published")
)

var (
	NotEnoughTimeSinceLastPost = errors.New("not enough time since last post")
)

func GetRequiredError(field string) error {
	return fmt.Errorf("field %s required", field)
}

func GetIsNotAdminError(role string) error {
	return fmt.Errorf("is not %s", role)
}

func CheckLang(lang string) error {
	if len(lang) != 2 {
		return IncorrectLang
	}
	return nil
}

func GetBadConfigError(config string) error {
	return fmt.Errorf("config %s is wrong", config)
}
