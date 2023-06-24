package proto

import "encoding/json"

// Request -- запрос клиента к серверу.
type Request struct {
	// Поле Command может принимать три значения:
	// * "quit" - прощание с сервером (после этого сервер рвёт соединение);
	// * "check" - передача данных авторизации и просьба проверить логин и пароль, существует ли такой пользователь.
	Command string `json:"command"`

	// Если Command == "check", в поле Data должна лежать структура с  логином и паролем пользователя.
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
}

// Response -- ответ сервера клиенту.
type Response struct {
	// Поле Status может принимать три значения:
	// * "quited" - успешное выполнение команды "quit";
	// * "failed" - в процессе выполнения команды произошла ошибка;
	// * "ok" - проверка данных авторизации прошла успешно.
	Status string `json:"status"`

	// Если Status == "failed", то в поле Data находится сообщение об ошибке.
	// Если Status == "ok", в поле Data должна лежать true/false в зависимости от наличия пользователя в базе
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
}

// LoginInfo -- логин и пароль.
type LoginInfo struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Exists bool `json:"exists"`
}