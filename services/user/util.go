package user

import "github.com/thanzen/eq/utils"

// return a user salt token
func GetUserSalt() string {
	return utils.GetRandomString(10)
}
