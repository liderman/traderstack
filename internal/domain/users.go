package domain

import (
	"time"
)

type Account struct {
	// Идентификатор счёта.
	Id string
	// Тип счёта.
	Type AccountType
	// Название счёта.
	Name string
	// Статус счёта.
	Status AccountStatus
	// Дата открытия счёта в часовом поясе UTC.
	OpenedDate time.Time
	// Дата закрытия счёта в часовом поясе UTC.
	ClosedDate time.Time
	// Уровень доступа к текущему счёту (определяется токеном).
	AccessLevel AccessLevel
}

// AccountType Тип счёта.
type AccountType int32

const (
	AccountTypeUnspecified AccountType = 0 // Тип аккаунта не определён.
	AccountTypeTinkoff     AccountType = 1 // Брокерский счёт Тинькофф.
	AccountTypeTinkoffIis  AccountType = 2 // ИИС счёт.
	AccountTypeInvestBox   AccountType = 3 // Инвесткопилка.
)

// AccountStatus Статус счёта.
type AccountStatus int32

const (
	AccountStatusUnspecified AccountStatus = 0 // Статус счёта не определён.
	AccountStatusNew         AccountStatus = 1 // Новый, в процессе открытия.
	AccountStatusOpen        AccountStatus = 2 // Открытый и активный счёт.
	AccountStatusClosed      AccountStatus = 3 // Закрытый счёт.
)

// AccessLevel Уровень доступа к счёту.
type AccessLevel int32

const (
	AccessLevelUnspecified AccessLevel = 0 // Уровень доступа не определён.
	AccessLevelFullAccess  AccessLevel = 1 // Полный доступ к счёту.
	AccessLevelReadOnly    AccessLevel = 2 // Доступ с уровнем прав "только чтение".
	AccessLevelNoAccess    AccessLevel = 3 // Доступ отсутствует.
)
