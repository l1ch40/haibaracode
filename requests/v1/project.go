package v1

type CreateProjectRequest struct {
	Name string `json:"name" binding:"required"`
}

type DeleteProjectRequest struct {
	ID uint `uri:"id" binding:"required"`
}
