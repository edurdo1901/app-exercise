package handlers

// swagger:model
type OrderNameResponse struct {
	// Names array names
	// example: ["Andres","Camilo","Laura","Luis"]
	// required: true
	Names []string `json:"name"`

	// Count number of names
	// example: 4
	// required: true
	Count int `json:"count"`
}

// swagger:model
type Error struct {
	// Code code error.
	Code string `json:"code"`

	// Message message error.
	Message string `json:"message"`
}
