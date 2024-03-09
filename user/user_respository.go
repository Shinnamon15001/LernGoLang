package user

type UserRespository interface {
	Find() {user *entity.User, err error}
}
