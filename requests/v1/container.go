package v1

type CreateContainerRequestUri struct {
	PID uint `uri:"id" binding:"required"`
}

type CreateContainerRequestJson struct {
	Image string `json:"image" binding:"required"`
}

type DeleteContainerRequest struct {
	PID uint `uri:"id" binding:"required"`
	ID  uint `uri:"cid" binding:"required"`
}
