package models

type OAuthType int

const (
	GITHUB OAuthType = iota + 1
	GOOGLE
	TWITTER
	GENERIC
	GRAFARG_COM
	GITLAB
	AZUREAD
	OKTA
)
