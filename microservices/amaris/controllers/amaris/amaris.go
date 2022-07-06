package amaris

import (
	"app/kernel/errors"
	"app/kernel/lang"
	"app/kernel/network"
	"app/kernel/utils"
	"app/microservices/amaris/config"
	"app/microservices/amaris/pipes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

// Realizar una funcion que reciba una cadena de texto, que contenga una lista de nombres
// separados por comas ejemplo "Luis,Camilo,Andres,Laura" y devuelva dos parametros: un array
// de strings con los nombres ordenados alfabeticamente y un entero indicando la cantidad de nombres

func (h Handler) NamesOrders(ctx *gin.Context) {
	var params pipes.NamesOrders

	err := ctx.ShouldBindJSON(&params)

	if err != nil {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusBadRequest,
			Error: errors.Error{
				Message: lang.Message{
					ID:      "ERR_DATA_REQUIRED",
					Message: err.Error(),
				},
			},
		})
		return
	}

	var utilsArray utils.UtilStringArray

	items := utilsArray.StringToArray(params.NameList)

	utilsArray.OrderStringArray(items)

	ctx.JSON(http.StatusOK, NameOrdersResponse{
		Code:  len(items),
		Names: items,
	})
}

// Realizar una funcion que reciba un numero entero "id" de un pokemon y devuelva su nombre.
// Para esto la funcion debe consultar una api de pokemon enviarle el id y obtener el campo
// "nombre" el cual va a retornar. URL = https://pokeapi.co/api/v2/pokemon-form
func (h Handler) PokemonItem(ctx *gin.Context) {
	id := ctx.Param("id")

	if &id == nil || id == "" {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusBadRequest,
			Error: errors.Error{
				Message: config.Lang.ErrPokemonNotId,
			},
		})
		return
	}

	result, err := network.Call(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%s", id))

	if err != nil {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusBadRequest,
			Error: errors.Error{
				Message: lang.Message{
					ID:      "ERR_CALL_POKEAPI",
					Message: err.Error(),
				},
			},
		})
		return
	}

	var utilsMap utils.UtilsMap

	var item PokemonItem

	err = utilsMap.DecoderReader(result.Body, &item)

	// ya que no existe un pokemon con id 0, retornamos su error, esto aplica para texto en caso de no castear
	if item.ID == 0 || err != nil {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusNotFound,
			Error: errors.Error{
				Message: config.Lang.ErrPokemonNotFound,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, PokemonItemResponse{
		Id:   int(item.ID),
		Name: item.Name,
		URL:  item.Pokemon.URL,
	})
}

// Se dice que dos cadenas X y Y son amigas si existen dos cadenas no vacías u y v tales que
// X = uv e Y = vu. Por ejemplo, "tokyo" y "kyoto" son amigas, siendo u = “to” y v = “kyo”.
// Escriba una funcion que reciba como entrada dos cadenas X e Y, e imprima si X e Y son amigas.

func (h Handler) StringsFriends(ctx *gin.Context) {
	var params pipes.StringsFriends

	err := ctx.ShouldBindJSON(&params)

	if err != nil {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusBadRequest,
			Error: errors.Error{
				Message: lang.Message{
					ID:      "ERR_DATA_REQUIRED",
					Message: err.Error(),
				},
			},
		})
		return
	}

	if len(params.X) != len(params.Y) {
		errors.HttpError(errors.HttpErrors{
			Ctx:    *ctx,
			Status: http.StatusBadRequest,
			Error: errors.Error{
				Message: config.Lang.ErrStringFriendsDiffLen,
			},
		})
		return
	}

	for i := 0; i < len(params.Y); i++ {
		u := params.X[0 : i+1]
		v := params.X[i+1 : len(params.X)]
		uv := fmt.Sprintf("%s%s", string(v), string(u))
		if params.Y == uv {
			ctx.JSON(http.StatusOK, StringsFriendsResponse{
				Friends: true,
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, StringsFriendsResponse{
		Friends: false,
	})
}
