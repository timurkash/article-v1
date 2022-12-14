package ccontext

import (
	"context"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	authorization = "authorization"
)

type (
	notAuthorizedKey struct{}
)

func FromIncomingToOutgoing(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	return metadata.NewOutgoingContext(ctx, md)
}

func GetToken(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", cerrors.NotAuthorized
	}
	bearer, ok := meta[authorization]
	if !ok {
		return "", cerrors.NotAuthorized
	}
	bearerString := bearer[0]
	if !strings.HasPrefix(bearerString, "Bearer ey") {
		return "", cerrors.BadToken
	}
	return bearerString[7:], nil
}

func AddNotAuthorizedKey(ctx context.Context, notAuthorized bool) context.Context {
	return context.WithValue(ctx, notAuthorizedKey{}, notAuthorized)
}

func IsNotAuthorized(ctx context.Context) bool {
	notAuthorized := ctx.Value(notAuthorizedKey{})
	if notAuthorized == nil {
		return false
	}
	return notAuthorized.(bool)
}
