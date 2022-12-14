package jwt

import (
	"context"
	"gitlab.com/mcsolutions/find-psy/back/common/others"
	"strings"
)

const (
	AdminSuffix  = "-admin"
	ArticleAdmin = "article-admin"
	VideoAdmin   = "video-admin"
	NewsAdmin    = "news-admin"
	EventAdmin   = "event-admin"
	OrgAdmin     = "org-admin"
	AdvAdmin     = "adv-admin"
	I18nAdmin    = "i18n-admin"
	PaymentAdmin = "payment-admin"

	AccountAdmin = "account-admin"
	UserAdmin    = "user-admin"
	PsyAdmin     = "psy-admin"
)

const ( // user_type
	User  = "user"
	Psy   = "psy"
	Admin = "admin"
)

func IsAdmin(ctx context.Context, roles ...string) bool {
	accessToken := GetClaims(ctx)
	if accessToken == nil {
		return false
	}
	if accessToken.UserType != Admin {
		return false
	}
	if accessToken.RealmAccess == nil {
		return false
	}
	return len(others.Intersect(accessToken.RealmAccess.Roles, roles)) > 0
}

func IsPsy(ctx context.Context) bool {
	accessToken := GetClaims(ctx)
	if accessToken == nil {
		return false
	}
	if accessToken.UserType != Psy {
		return false
	}
	return true
}

func GetRoles(ctx context.Context) (roles []string) {
	accessToken := GetClaims(ctx)
	if accessToken == nil {
		return
	}
	if accessToken.UserType != Admin {
		return
	}
	if accessToken.RealmAccess == nil {
		return
	}
	for _, role := range accessToken.RealmAccess.Roles {
		if strings.HasSuffix(role, AdminSuffix) {
			roles = append(roles, role)
		}
	}
	return
}
