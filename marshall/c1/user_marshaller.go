package c1

import (
	"encoding/json"
)

type PublicUser struct {
	ID          int64  `json:"id"`
	DataCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PribateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DataCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DataCreated: user.DataCreated,
			Status:      user.Status,
		}
	}

	userJson, _ := json.Marshall(user)
	var privateUser PribateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}

/*memo
パスワードは含めない
user.Marshall(c.GetHeader("X-Public") == "true")
*/
