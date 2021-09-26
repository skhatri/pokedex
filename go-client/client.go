package main

import (
	"github.com/skhatri/pokedex/go-client/api"
	"log"
)

func main() {
	log.Println("call server")
	pokeClient := api.NewPokedexClient()
	mimikyu, err := pokeClient.FindPokemonByName("mimikyu")
	if err != nil {
		panic(err)
	}
	log.Println(mimikyu.Ability1)
	items, err := pokeClient.FindPokemonsByType([]string{"dragon"})
	log.Println(len(items))

	bugGhostPokemons, err := pokeClient.FindPokemonsByType([]string{"bug", "ghost"})
	var minHpPokemon = api.Pokemon{}
	for _, poke := range bugGhostPokemons {
		if minHpPokemon.Hp == 0 || poke.Hp < minHpPokemon.Hp {
			minHpPokemon = poke
		}
	}
	log.Println(minHpPokemon.Name, minHpPokemon.Hp)
	zeraora, err := pokeClient.FindPokemonByName("Zeraora")
	types := make([]string, 0)
	types = append(types, zeraora.Type1)
	if zeraora.Type2 != "" {
		types = append(types, zeraora.Type2)
	}
	log.Println(zeraora.Name, types)
}
