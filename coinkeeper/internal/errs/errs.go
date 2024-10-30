package errs

import "errors"

var (
	ErrPermissionDenied            = errors.New("Ошибка: доступ запрещён")
	ErrValidationFailed            = errors.New("Ошибка: валидация не пройдена")
	ErrUsernameUniquenessFailed    = errors.New("Ошибка: имя пользователя уже существует")
	ErrOperationNotFound           = errors.New("Ошибка: операция не найдена")
	ErrIncorrectUsernameOrPassword = errors.New("Ошибка: неверное имя пользователя или пароль")
	ErrRecordNotFound              = errors.New("Ошибка: запись не найдена")
	ErrUserNotFound                = errors.New("Ошибка: пользователь не найден")
	ErrSomethingWentWrong          = errors.New("Ошибка: что-то пошло не так")
)
