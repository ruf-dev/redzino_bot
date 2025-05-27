package errors

import (
	"go.redsock.ru/rerrors"
)

var (
	ErrAccessDenied = rerrors.NewUserError(
		"Вы не можете совершить данное действие. (но если очень хочется" +
			" - обратитесь к админу)")
)
