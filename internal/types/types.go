package types

type User struct {
	Id       int
	Segments []string
}

type CreateSegmentsRequest struct {
	Segments []string `json: segments`
}

type CreateUsersRequest struct {
	Users []int `json: users`
}

type UpdateUser struct {
	Id     int      `json: id`
	Add    []string `json: add`
	Delete []string `json: delete`
}
