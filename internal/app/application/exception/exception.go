package exception

import "net/http"

type AppException struct {
	Status int
	Msg    string
}

func (e AppException) Throw() (string, int) {
	return e.Msg, e.Status
}

var (
	ErrAuth          = &AppException{Status: http.StatusUnauthorized, Msg: "invalid token"}
	ErrNotFound      = &AppException{Status: http.StatusNotFound, Msg: "not found"}
	ErrDuplicate     = &AppException{Status: http.StatusBadRequest, Msg: "duplicate"}
	ErrPassword      = &AppException{Status: http.StatusBadRequest, Msg: "wrong password"}
	QueryErr         = &AppException{Status: http.StatusInternalServerError, Msg: "internal error"}
	UserAlreadyExist = &AppException{Status: 409, Msg: "Já existe um usuário com esse email"}
)
