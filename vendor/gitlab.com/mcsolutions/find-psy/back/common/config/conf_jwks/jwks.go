package conf_jwks

import (
	"fmt"

	"github.com/MicahParks/keyfunc"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"

	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

func GetJWKS(confJwks *common.Jwks) (*keyfunc.JWKS, error) {
	if confJwks == nil {
		return nil, cerrors.GetBadConfigError("jwks")
	}
	if confJwks.Url == "" {
		return nil, cerrors.GetBadConfigError("jwks.url")
	}
	if confJwks.RefreshRateLimit == nil {
		return nil, cerrors.GetBadConfigError("jwks.refresh_rate_limit")
	}
	if confJwks.RefreshInterval == nil {
		return nil, cerrors.GetBadConfigError("jwks.refresh_interval")
	}
	if confJwks.RefreshTimeout == nil {
		return nil, cerrors.GetBadConfigError("jwks.refresh_timeout")
	}
	return keyfunc.Get(confJwks.Url, keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			fmt.Errorf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   confJwks.RefreshInterval.AsDuration(),
		RefreshRateLimit:  confJwks.RefreshRateLimit.AsDuration(),
		RefreshTimeout:    confJwks.RefreshTimeout.AsDuration(),
		RefreshUnknownKID: true,
	})
}
