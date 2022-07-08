package handlers

// swagger:model
type StringFriendsRequest struct {
	// StringX first value entered
	// example: tokyo
	// required: true
	StringX string `json:"x" binding:"required"`

	// StringY second value entered
	// example: kyoto
	// required: true
	StringY string `json:"y" binding:"required"`
}

// swagger:model
type OrderNameRequest struct {
	// Names list of names separated by ,
	// example: Luis,Camilo,Andres,Laura
	// required: true
	Names string `json:"names" binding:"required"`
}
