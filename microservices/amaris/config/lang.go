package config

import (
	l "app/kernel/lang"
)

// Lang variables de mensajes de la aplicacion
var Lang = lang{
	NewKeyTest: l.Message{
		ID:      "TEST_KEY",
		Message: "The request was attended",
	},
	ErrPokemonNotId: l.Message{
		ID:      "ERR_POKEMON_NOT_ID",
		Message: "err pokemon not id",
	},
	ErrPokemonNotFound: l.Message{
		ID:      "ERR_POKEMON_NOT_FOUND",
		Message: "err pokemon not found",
	},
	ErrStringFriendsDiffLen: l.Message{
		ID:      "ERR_STRING_FRIENDS_DIFFERENT_LEN",
		Message: "err string friends different len",
	},
}

type (
	lang struct {
		// Unimos la estructura generica del kernel con el del app
		NewKeyTest              l.Message
		ErrPokemonNotId         l.Message
		ErrPokemonNotFound      l.Message
		ErrStringFriendsDiffLen l.Message
	}
)
