package main

import (
	"bytes"
	"github.com/skhatri/pokedex/go-client/api"
	"log"
)

func main() {
	pokeClient := api.NewPokedexClient()
	mimikyu, err := pokeClient.FindPokemonByName("mimikyu")
	if err != nil {
		panic(err)
	}
	log.Println(mimikyu.Ability1)
	flyingAndWaterPokemons, flyingWaterErr := pokeClient.FindPokemonsByType([]string{"flying", "water"})
	if flyingWaterErr != nil {
		panic(flyingWaterErr)
	}
	log.Println(len(flyingAndWaterPokemons))

	bugGhostPokemons, bugGhostErr := pokeClient.FindPokemonsByType([]string{"bug", "ghost"})
	if bugGhostErr != nil {
		panic(bugGhostErr)
	}
	var minHpPokemon = api.Pokemon{}
	for _, poke := range bugGhostPokemons {
		if minHpPokemon.Hp == 0 || poke.Hp < minHpPokemon.Hp {
			minHpPokemon = poke
		}
	}
	log.Println(minHpPokemon.Name, minHpPokemon.Hp)
	zeraora, zeraoraErr := pokeClient.FindPokemonByName("Zeraora")
	if zeraoraErr != nil {
		panic(zeraoraErr)
	}
	typeStr := bytes.NewBufferString(zeraora.Type1)
	if zeraora.Type2 != "" {
		typeStr.WriteString(",")
		typeStr.WriteString(zeraora.Type2)
	}
	log.Println(zeraora.Name, typeStr.String())
}
