package routes

import (
	"net/http"

	"app/microservices/amaris/controllers/amaris"
)

func init() {
	const subjet = "amaris"
	handler := amaris.Handler{}

	routes := Group{
		Prefix: subjet,
		Routes: []Route{
			{
				Method:  http.MethodGet,
				Path:    "/pokemon/:id",
				Handler: handler.PokemonItem,
			},
			{
				Method:  http.MethodPut,
				Path:    "/names-orders",
				Handler: handler.NamesOrders,
			},
			{
				Method:  http.MethodPut,
				Path:    "/strings-friends",
				Handler: handler.StringsFriends,
			},
		},
	}

	AppRouting = append(AppRouting, routes)
}
