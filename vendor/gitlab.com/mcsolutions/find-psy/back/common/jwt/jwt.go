package jwt

import (
	"context"

	"github.com/MicahParks/keyfunc"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"gitlab.com/mcsolutions/find-psy/back/common/ccontext"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
)

const ( // psy_status
	Bronze = "bronze"
	Silver = "silver"
	Gold   = "gold"
)

type (
	Claims struct {
		jwt.RegisteredClaims
		Id    uuid.UUID
		Name  string `json:"name"`  // Ivanov Ivan
		Email string `json:"email"` // email = account
		//PreferredUsername string `json:"preferred_username"`
		RealmAccess *struct {
			Roles []string `json:"roles"`
		} `json:"realm_access"`
		UserType string `json:"user_type"` // psy, admin (if not set this is a user)
		UserName string `json:"user_name"` // if psychologist set name different from google name
	}
	claimsKey struct{}
)

func GetClaimsFromAccessToken(tokenString string, jwks *keyfunc.JWKS) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, jwks.Keyfunc)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*Claims)
	claims.Id, err = uuid.Parse(claims.Subject)
	if err != nil {
		return nil, err
	}
	if claims.UserName == "" {
		claims.UserName = claims.Name
	}
	return claims, nil
}

func GetClaimsFromContext(ctx context.Context, jwks *keyfunc.JWKS) (*Claims, error) {
	accessToken, err := ccontext.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	return GetClaimsFromAccessToken(accessToken, jwks)
}

func AddClaimsKey(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimsKey{}, claims)
}

func GetClaims(ctx context.Context) *Claims {
	claims := ctx.Value(claimsKey{})
	if claims != nil {
		return claims.(*Claims)
	}
	return nil
}

func GetAuthMiddleware(jwks *keyfunc.JWKS) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if jwks == nil {
				return nil, cerrors.NoJwks
			}
			claims, err := GetClaimsFromContext(ctx, jwks)
			if err != nil {
				if err == cerrors.NotAuthorized {
					return handler(ccontext.AddNotAuthorizedKey(ctx, true), req)
				} else {
					return nil, err
				}
			}
			return handler(AddClaimsKey(ctx, claims), req)
		}
	}
}
