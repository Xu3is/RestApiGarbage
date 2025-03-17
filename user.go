package restapigarbage

type User struct {
	UserName string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
