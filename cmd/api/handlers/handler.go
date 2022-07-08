package handlers

import (
	"net/http"
	"strconv"

	"github.com/exercise/internal/algorithms"
	"github.com/exercise/internal/pkg/pokemon"

	"github.com/gin-gonic/gin"
)

type clientPokemon interface {
	GetDetail(id int) (pokemon.Pokemon, error)
}

type Handler struct {
	client clientPokemon
}

func New(client clientPokemon) *Handler {
	return &Handler{
		client: client,
	}
}

// API
func (h Handler) API(router *gin.Engine) {
	router.Use(errorHandler)
	router.POST("order/names", orderNames())
	router.POST("string/friends", stringFriends())
	router.GET("pokemon/:id", getPokemon(h))
}

// Post Get pokemon
// swagger:operation POST /pokemon/{id} getpokemon getpokemon
//
// Get the name of the pokemon
//
// returns 200 if found pokemon else 404 not found
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   type: integer
//   required: true
//   description: pokemon identifier
// responses:
//   '200':
//    description: pokemon name
//   '404':
//    description: not found pokemon
//    examples:
//     application/json:
//      code: "Not Found"
//      message: "client: pokemon not found"
//    schema:
//     "$ref": "#/definitions/Error"
//   '422':
//    description: error in identifier pokemon
//    examples:
//     application/json:
//      code: "Unprocessable Entity"
//      message: "invalid syntax"
//    schema:
//     "$ref": "#/definitions/Error"
//   '500':
//    description: unknown error
//    examples:
//     application/json:
//      code: "Internal Server Error"
//      message: "unknown error"
//    schema:
//     "$ref": "#/definitions/Error"
func getPokemon(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		idparam := c.Param("id")
		id, err := strconv.Atoi(idparam)
		if err != nil {
			_ = c.Error(err)
			return
		}

		result, err := h.client.GetDetail(id)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, result.Name)
	}
}

// Post order names
// swagger:operation POST /order/names order order
//
// sort the names alphabetically and return the number of elements
//
// returns sort the names alphabetically and return the number of elements
// ---
// produces:
// - application/json
// parameters:
// - name: body
//   in: body
//   required: true
//   description: values to organize
//   schema:
//    "$ref": "#/definitions/OrderNameRequest"
// responses:
//   '200':
//    description: list of names organized with the number of items
//    examples:
//     application/json:
//      name: ["Andres","Camilo","Laura","Luis"]
//      count: 4
//    schema:
//     "$ref": "#/definitions/OrderNameResponse"
//   '422':
//    description: input data validation error
//    examples:
//     application/json:
//      code: "Unprocessable Entity"
//      message: "Key: 'OrderNameRequest.Names' Error:Field validation for 'Names' failed on the 'required' tag"
//    schema:
//     "$ref": "#/definitions/Error"
func orderNames() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderName OrderNameRequest
		if err := c.ShouldBind(&orderName); err != nil {
			_ = c.Error(err)
			return
		}

		arrNames, countNames := algorithms.Order(orderName.Names)
		c.JSON(http.StatusOK, OrderNameResponse{
			Names: arrNames,
			Count: countNames,
		})
	}
}

// Post string friends
// swagger:operation POST /string/friends stringfriends stringfriends
//
// check if input data is friendly string
//
// returns corresponding validation of whether the strings are friends
// ---
// produces:
// - application/json
// parameters:
// - name: body
//   in: body
//   required: true
//   description: values to evaluate
//   schema:
//    "$ref": "#/definitions/StringFriendsRequest"
// responses:
//   '200':
//    description: validation of whether the strings entered are friends
//   '422':
//    description: input data validation error
//    examples:
//     application/json:
//      code: "Unprocessable Entity"
//      message: "Key: 'StringFriendsRequest.StringY' Error:Field validation for 'StringY' failed on the 'required' tag"
//    schema:
//     "$ref": "#/definitions/Error"
func stringFriends() gin.HandlerFunc {
	return func(c *gin.Context) {
		var stringFriends StringFriendsRequest
		if err := c.ShouldBind(&stringFriends); err != nil {
			_ = c.Error(err)
			return
		}

		result := algorithms.IsStringFriends(stringFriends.StringX, stringFriends.StringY)
		c.JSON(http.StatusOK, result)
	}
}
