package poke

import (
	"fmt"
	"github.com/skhatri/api-router-go/router"
	"github.com/skhatri/api-router-go/router/model"
	"strings"
)

func makeResult(pokemons []Pokemon) *model.Container {
	return &model.Container{
		Data: pokemons,
		Metadata: map[string]interface{}{
			"total": len(pokemons),
		},
	}
}

func FindByName(request *router.WebRequest) *model.Container {
	var pokemons []Pokemon = nil
	requestName := request.GetQueryParam("q")
	if requestName != "" {
		pokemons = pokedex.SearchByName(strings.ToLower(requestName))
	}
	return makeResult(pokemons)
}

func FindByType(request *router.WebRequest) *model.Container {
	var pokemons []Pokemon = nil
	requestType := request.GetQueryParam("q")
	if requestType != "" {
		pokemons = pokedex.SearchByType(strings.ToLower(requestType))
	}
	return makeResult(pokemons)
}

func FindEffectiveAgainstType(request *router.WebRequest) *model.Container {
	var pokemons []Pokemon = nil
	againstType := request.GetQueryParam("q")
	if againstType != "" {
		pokemons = pokedex.SearchByAgainstType(strings.ToLower(againstType), -1)
	}
	return makeResult(pokemons)
}

func FindEffectiveAgainstName(request *router.WebRequest) *model.Container {
	var pokemons []Pokemon = nil
	againstName := request.GetQueryParam("q")
	if againstName != "" {
		matching := pokedex.SearchByName(strings.ToLower(againstName))
		if len(matching) == 1 {
			typeData := matching[0].Type1
			if matching[0].Type2 != "" {
				typeData = fmt.Sprintf("%s,%s", typeData, matching[0].Type2)
			}
			genOverride := request.GetQueryParam("generation")
			gen := -1
			if genOverride == "exact" {
				gen = matching[0].Generation
			}
			pokemons = pokedex.SearchByAgainstType(strings.ToLower(typeData), gen)
		}
	}
	return makeResult(pokemons)
}
