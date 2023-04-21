package comment

type CommentCreate struct {
	Message string `json:"message" binding:"required"`
	UserID  string `json:"user_id"`
	PhotoID string `json:"photo_id" binding:"required"`
}
