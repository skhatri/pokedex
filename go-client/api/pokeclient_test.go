package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
)

type MockHttpClient struct {
	items map[string]string
}

func (mhc *MockHttpClient) Get(uri string) HttpResponse {
	dataFile, ok := mhc.items[uri]
	if !ok {
		return HttpResponse{
			err:  errors.New("resource not found"),
			data: make([]byte, 0),
		}
	}
	b, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return HttpResponse{
			err: err,
		}
	}
	return HttpResponse{
		data: b,
	}
}

func (mhc *MockHttpClient) When(uri string, dataFile string) *MockHttpClient {
	mhc.items[uri] = dataFile
	return mhc
}

var mockClient = func() *MockHttpClient {
	mock := MockHttpClient{
		items: make(map[string]string),
	}

	mock.When("/search/by-name?q=mimikyu", "../../client-data/mimikyu.json").
		When("/search/by-type?q=flying,water", "../../client-data/flying-water.json").
		When("/search/by-type?q=bug,ghost", "../../client-data/bug-ghost.json").
		When("/search/by-name?q=Zeraora", "../../client-data/zeraora.json")
	return &mock
}()

func TestMimikyuAbility(t *testing.T) {
	pokeClient := NewCustomPokedexClient(mockClient)
	mimikyu, err := pokeClient.FindPokemonByName("mimikyu")
	if err != nil {
		t.Fatal(err)
	}
	SimpleAssertEquals(t, "Disguise", mimikyu.Ability1)
}

func TestCountFlyingAndWaterTypePokemon(t *testing.T) {
	pokeClient := NewCustomPokedexClient(mockClient)
	pokemons, err := pokeClient.FindPokemonsByType([]string{"flying", "water"})
	if err != nil {
		t.Fatal(err)
	}
	SimpleAssertEquals(t, 8, len(pokemons))
}

func TestFindBugGhostPokemonWithLowestHp(t *testing.T) {
	pokeClient := NewCustomPokedexClient(mockClient)
	pokemons, err := pokeClient.FindPokemonsByType([]string{"bug", "ghost"})
	if err != nil {
		t.Fatal(err)
	}
	expectedName := "Shedinja"
	expectedHp := 1
	var minHpPokemon = Pokemon{}
	for _, p := range pokemons {
		if minHpPokemon.Hp == 0 || p.Hp < minHpPokemon.Hp {
			minHpPokemon = p
		}
	}
	SimpleAssertEquals(t, expectedName, minHpPokemon.Name)
	SimpleAssertEquals(t, expectedHp, minHpPokemon.Hp)
}

func TestFindZeraorasType(t *testing.T) {
	pokeClient := NewCustomPokedexClient(mockClient)
	zeraora, err := pokeClient.FindPokemonByName("Zeraora")
	if err != nil {
		t.Fatal(err)
	}
	SimpleAssertEquals(t, "Electric", zeraora.Type1)
	SimpleAssertEquals(t, "", zeraora.Type2)
}

func SimpleAssertEquals(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatal(fmt.Sprintf("expected %v but found %v", a, b))
	}
}
