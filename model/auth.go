package model

type Authentication struct {
	Username string
	Password string
}

type Authorization struct {
	Sub string
	Obj string
	Act string
}
