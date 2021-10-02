package poke

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Pokedex struct {
	data      []Pokemon
	nameIndex map[string]int
	typeStore map[string][]int
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

var pokedex = func() *Pokedex {
	datasetDir := os.Getenv("DATASET_DIR")
	if datasetDir == "" {
		datasetDir = "../server/dataset"
	}
	fp, err := os.OpenFile(fmt.Sprintf("%s/pokedex_0520.csv", datasetDir), os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("file read error", err)
	}
	data := loadData(fp)
	nameIndex := make(map[string]int)

	for i, d := range data {
		nameIndex[strings.ToLower(d.Name)] = i
	}

	typeMap := make(map[string][]int)
	for i, d := range data {
		type1 := strings.ToLower(d.Type1)
		existingType1, ok1 := typeMap[type1]
		if ok1 {
			existingType1 = append(existingType1, i)
		} else {
			existingType1 = make([]int, 0)
		}
		typeMap[type1] = existingType1
		type2 := strings.ToLower(d.Type2)
		if type2 != "" {
			existingType2, ok2 := typeMap[type2]
			if ok2 {
				existingType2 = append(existingType2, i)
			} else {
				existingType2 = make([]int, 0)
			}
			typeMap[type2] = existingType2
		}
	}
	return &Pokedex{
		data:      data,
		nameIndex: nameIndex,
		typeStore: typeMap,
	}
}()

type PokeSearch interface {
	SearchByName(name string) []Pokemon
	SearchByType(typeName string) []Pokemon
	SearchAgainstType(typeName string) []Pokemon
}

func (px *Pokedex) SearchByName(name string) []Pokemon {
	result := make([]Pokemon, 0)
	if pokeIndex, ok := px.nameIndex[name]; ok {
		result = append(result, px.data[pokeIndex])
	}
	return result
}

func (px *Pokedex) SearchByType(typeNames string) []Pokemon {
	typeNameData := strings.Split(typeNames, ",")
	result := make([]Pokemon, 0)

	type1 := strings.ToLower(typeNameData[0])
	type2 := ""
	if len(typeNameData) == 2 {
		type2 = strings.ToLower(typeNameData[1])
		if type1 == type2 {
			type2 = ""
		}
	}

	typeList, ok := px.typeStore[type1]
	if !ok {
		return result
	}

	for _, i := range typeList {
		pokemon := px.data[i]
		if type2 == "" {
			result = append(result, pokemon)
		} else {
			type2Matches := strings.ToLower(pokemon.Type1) == type2 || strings.ToLower(pokemon.Type2) == type2
			if type2Matches && (strings.ToLower(pokemon.Type1) == type1 || strings.ToLower(pokemon.Type2) == type1) {
				result = append(result, pokemon)
			}
		}
	}
	return result
}

func (px *Pokedex) SearchByAgainstType(againstNames string, gen int) []Pokemon {
	result := make([]Pokemon, 0)
	againstTypeData := strings.Split(againstNames, ",")
	for _, p := range px.data {
		if gen != -1 && p.Generation != gen {
			continue
		}
		typeMatch := true
		for _, typeName := range againstTypeData {
			current := false
			switch typeName {
			case "normal":
				current = p.AgainstNormal <= 0.5
			case "fire":
				current = p.AgainstFire <= 0.5
			case "water":
				current = p.AgainstWater <= 0.5
			case "electric":
				current = p.AgainstElectric <= 0.5
			case "grass":
				current = p.AgainstGrass <= 0.5
			case "ice":
				current = p.AgainstIce <= 0.5
			case "fight":
				current = p.AgainstFight <= 0.5
			case "poison":
				current = p.AgainstPoison <= 0.5
			case "ground":
				current = p.AgainstGround <= 0.5
			case "flying":
				current = p.AgainstFlying <= 0.5
			case "psychic":
				current = p.AgainstPsychic <= 0.5
			case "bug":
				current = p.AgainstBug <= 0.5
			case "rock":
				current = p.AgainstRock <= 0.5
			case "ghost":
				current = p.AgainstGhost <= 0.5
			case "dragon":
				current = p.AgainstDragon <= 0.5
			case "dark":
				current = p.AgainstDark <= 0.5
			case "steel":
				current = p.AgainstSteel <= 0.5
			case "fairy":
				current = p.AgainstFairy <= 0.5
			}

			typeMatch = typeMatch && current
			if !typeMatch {
				break
			}
		}
		if typeMatch {
			result = append(result, p)
		}
	}
	return result
}

func loadData(fp *os.File) []Pokemon {
	scanner := bufio.NewScanner(fp)

	headersProcessed := false
	data := make([]Pokemon, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, ",")
		if !headersProcessed {
			headersProcessed = true
		} else {
			p := Pokemon{}
			d := reflect.ValueOf(&p).Elem()
			for i := 0; i < d.Type().NumField(); i++ {
				f := d.Type().Field(i)
				index, _ := strconv.Atoi(f.Tag.Get("i"))
				switch f.Type.Kind() {
				case reflect.Int:
					value, _ := strconv.ParseInt(arr[index], 10, 64)
					d.Field(i).SetInt(value)
				case reflect.String:
					value := arr[index]
					d.Field(i).SetString(value)
				case reflect.Float32:
					value, _ := strconv.ParseFloat(arr[index], 32)
					d.Field(i).SetFloat(value)
				}
			}
			data = append(data, p)
		}
	}
	return data
}
