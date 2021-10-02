package api

import (
	"fmt"
	"strings"
)

type Meta struct {
	total int `json:"total"`
}

type Result struct {
	Data     []Pokemon `json:"data"`
	Metadata Meta      `json:"metadata"`
}

type Pokemon struct {
	PokedexNumber   int     `json:"s_no" i:"1"`
	Name            string  `json:"name" i:"2"`
	Generation      int     `json:"generation" i:"5"`
	Status          string  `json:"status" i:"6"`
	Species         string  `json:"species" i:"7"`
	TypeNumber      int     `json:"type_number" i:"8"`
	Type1           string  `json:"type1" i:"9"`
	Type2           string  `json:"type2" i:"10"`
	Height          float32 `json:"height" i:"11"`
	Weight          float32 `json:"weight" i:"12"`
	TotalAbilities  int     `json:"abilities" i:"13"`
	Ability1        string  `json:"ability1" i:"14"`
	Ability2        string  `json:"ability2" i:"15"`
	HiddenAbility   string  `json:"ability_hidden" i:"16"`
	TotalPoints     int     `json:"total_points" i:"17"`
	Hp              int     `json:"hp" i:"18"`
	Attack          int     `json:"attack" i:"19"`
	Defense         int     `json:"defense" i:"20"`
	SpAttack        int     `json:"sp_attack" i:"21"`
	SpDefense       int     `json:"sp_defense" i:"22"`
	Speed           int     `json:"speed" i:"23"`
	CatchRate       float32 `json:"catch_rate" i:"24"`
	BaseFriendship  string  `json:"base_friendship" i:"25"`
	BaseExperience  int     `json:"base_experience" i:"26"`
	GrowthRate      int     `json:"growth_rate" i:"27"`
	EggTypeNumber   int     `json:"egg_type_number" i:"28"`
	EggType1        int     `json:"egg_type1" i:"29"`
	EggType2        int     `json:"egg_type2" i:"30"`
	PercentageMale  int     `json:"percentage_male" i:"31"`
	EggCycles       int     `json:"egg_cycles" i:"32"`
	AgainstNormal   float32 `json:"against_normal" i:"33"`
	AgainstFire     float32 `json:"against_fire" i:"34"`
	AgainstWater    float32 `json:"against_water" i:"35"`
	AgainstElectric float32 `json:"against_electric" i:"36"`
	AgainstGrass    float32 `json:"against_grass" i:"37"`
	AgainstIce      float32 `json:"against_ice" i:"38"`
	AgainstFight    float32 `json:"against_fight" i:"39"`
	AgainstPoison   float32 `json:"against_poison" i:"40"`
	AgainstGround   float32 `json:"against_ground" i:"41"`
	AgainstFlying   float32 `json:"against_flying" i:"42"`
	AgainstPsychic  float32 `json:"against_psychic" i:"43"`
	AgainstBug      float32 `json:"against_bug" i:"44"`
	AgainstRock     float32 `json:"against_rock" i:"45"`
	AgainstGhost    float32 `json:"against_ghost" i:"46"`
	AgainstDragon   float32 `json:"against_dragon" i:"47"`
	AgainstDark     float32 `json:"against_dark" i:"48"`
	AgainstSteel    float32 `json:"against_steel" i:"49"`
	AgainstFairy    float32 `json:"against_fairy" i:"50"`
}

type PokedexClient interface {
	FindPokemonByName(name string) (*Pokemon, error)
	FindPokemonsByType(typeName []string) ([]Pokemon, error)
}

type _pokedexClient struct {
	client HttpClient
}

func (pc *_pokedexClient) FindPokemonByName(name string) (*Pokemon, error) {
	httpResponse := pc.client.Get(fmt.Sprintf("/search/by-name?q=%s", name))
	if httpResponse.HasError() {
		return nil, httpResponse.err
	}
	result := Result{}
	httpResponse.ToData(&result)
	var pokemon *Pokemon = nil
	if len(result.Data) > 0 {
		pokemon = &result.Data[0]
	}
	return pokemon, nil
}

func (pc *_pokedexClient) FindPokemonsByType(typeName []string) ([]Pokemon, error) {
	httpResponse := pc.client.Get(fmt.Sprintf("/search/by-type?q=%s", strings.Join(typeName, ",")))
	if httpResponse.HasError() {
		return nil, httpResponse.err
	}
	result := Result{}
	httpResponse.ToData(&result)
	return result.Data, nil
}

func NewPokedexClient() PokedexClient {
	return NewCustomPokedexClient(newHttpClient("http://localhost:6100"))
}

func NewCustomPokedexClient(httpClient HttpClient) PokedexClient {
	return &_pokedexClient{
		client: httpClient,
	}
}
