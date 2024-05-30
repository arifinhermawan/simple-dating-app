package account

type UserAccount struct {
	ID       int64
	Username string
	Password string
}

type CreateUserAccountReq struct {
	Username string
	Password string
}
