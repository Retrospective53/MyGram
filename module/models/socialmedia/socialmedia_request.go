package socialmedia

type SocialMediaCreate struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
	UserID         string `json:"user_id"`
}
