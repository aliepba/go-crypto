package requests

type CategoryInput struct {
	Category string `json:"category" binding:"required"`
}

type GetCategoryInput struct {
	Category string `uri:"category"`
}
