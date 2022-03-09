// https://github.com/mikunalpha/goas
package main

type User struct {
	ID   uint64 `json:"id" example:"100" description:"User identity"`
	Name string `json:"name" example:"Mikun"`
}

type UsersResponse struct {
	Data []Users `json:"users" example:"[{\"id\":100, \"name\":\"Mikun\"}]"`
}

type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type ErrorResponse struct {
	ErrorInfo Error `json:"error"`
}

// @Title Get user list of a group.
// @Description Get users related to a specific group.
// @Param  groupID  path  int  true  "Id of a specific group."
// @Success  200  object  UsersResponse  "UsersResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource users
// @Route /api/group/{groupID}/users [get]
func GetGroupUsers() {
	// ...
}

// @Title Get user list of a group.
// @Description Create a new user.
// @Param  user  body  User  true  "Info of a user."
// @Success  200  object  User           "UsersResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource users
// @Route /api/user [post]
func PostUser() {
	// ...
}
