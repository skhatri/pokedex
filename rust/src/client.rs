use serde::{Deserialize, Serialize};
use std::clone::Clone;

#[derive(Serialize, Deserialize, Clone)]
struct Pokemon {
    s_no: i32,
    name: String,
    hp: i32,
    type1: String,
    type2: String,
}

#[derive(Serialize, Deserialize)]
struct Metadata {
    total: i32,
}

#[derive(Serialize, Deserialize)]
struct Res {
    metadata: Metadata,
    data: Vec<Pokemon>,
}

fn find_pokemon_by_name(name: &str) -> Pokemon {
    let url = format!("http://localhost:6100/search/by-name?q={}", name);
    let resp: std::string::String = reqwest::blocking::get(url)
        .unwrap()
        .text()
        .unwrap();
    let r: Res = serde_json::from_str(&resp).unwrap();
    let poke = r.data[0].clone();
    poke
}

fn find_pokemon_by_type(type_name: &str) -> Vec<Pokemon> {
    let url = format!("http://localhost:6100/search/by-type?q={}", type_name);
    let resp: std::string::String = reqwest::blocking::get(url)
        .unwrap()
        .text()
        .unwrap();
    let r: Res = serde_json::from_str(&resp).unwrap();
    let poke = r.data.clone();
    poke
}

fn main() {
    println!("call pokedex api");
    let mimikyu = find_pokemon_by_name("mimikyu");
    println!("id {} name {}", mimikyu.s_no, mimikyu.name);

    let flyingWater = find_pokemon_by_type("flying,water");
    println!("total flying water {}", flyingWater.len());
    for p in &flyingWater {
        println!("id {} name {}", p.s_no, p.name)
    }

    let bugGhost = find_pokemon_by_type("bug,ghost");
    println!("total bug ghost {}", bugGhost.len());
    let mut minHpPoke: Option<&Pokemon> = None;

    for p in &bugGhost {
        match minHpPoke {
            None => {
                minHpPoke = Some(p)
            }
            Some(pp) => {
                if p.hp < pp.hp {
                    minHpPoke = Some(p)
                }
            }
        }
    }
    match minHpPoke {
        Some(p) => {
            println!("id {} name {}, hp {}", p.s_no, p.name, p.hp)
        }
        None => {
            println!("none found")
        }
    }

    let zeraora = find_pokemon_by_name("Zeraora");
    let mut typeName = zeraora.type1;
    if zeraora.type2 != "" {
        typeName = format!("{},{}", typeName, zeraora.type2)
    }
    println!("zeraora name: {}, type: {}", zeraora.name, typeName)
}
