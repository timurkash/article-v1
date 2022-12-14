package test

import "github.com/google/uuid"

var IsUnitTest bool

var (
	Uuid, _    = uuid.Parse("46ec9018-cf61-11ec-9d64-0242ac120002")
	Email      = "iivanot@example.com"
	Name       = "Ivan Ivanov"
	Title      = "Заголовок"
	Latinized  = "Zagolovok"
	Text       = "<p>Some text...<p>"
	B17Account = "timurkash"
	Tags       = []string{"news"}
	Lang       = "ru"
)
