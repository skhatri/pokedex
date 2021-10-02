package com;

import com.api.Either;
import com.api.PokeClient;
import com.api.PokeClients;
import com.api.Pokemon;

import java.util.List;

public class Client {
    public static void main(String[] args) {
        PokeClient pokeClient = PokeClients.newPokeClient();
        Either<Exception, Pokemon> errOrMimikyu = pokeClient.findPokemonByName("mimikyu");
        Exception err = errOrMimikyu.first();
        Pokemon mimikyu = errOrMimikyu.second();
        if (err != null) {
            throw new RuntimeException(err);
        }
        System.out.println(mimikyu.getAbility1());

        Either<Exception, List<Pokemon>> errOrFlyingWater = pokeClient.findPokemonsByType(new String[]{"flying", "water"});
        Exception flyingWaterErr = errOrFlyingWater.first();
        if (flyingWaterErr != null) {
            throw new RuntimeException(flyingWaterErr);
        }
        List<Pokemon> flyingAndWaterPokemons = errOrFlyingWater.second();
        System.out.println(flyingAndWaterPokemons.size());

        Either<Exception, List<Pokemon>> errOrBugGhost = pokeClient.findPokemonsByType(new String[]{"bug", "ghost"});
        Exception bugGhostErr = errOrBugGhost.first();
        if (bugGhostErr != null) {
            throw new RuntimeException(bugGhostErr);
        }
        List<Pokemon> bugGhostPokemons = errOrBugGhost.second();
        final Pokemon[] minHpPokemon = {null};
        bugGhostPokemons.stream().forEach(poke -> {
            if (minHpPokemon[0] == null || poke.getHp() < minHpPokemon[0].getHp()) {
                minHpPokemon[0] = poke;
            }
        });
        System.out.printf("%s %d\n", minHpPokemon[0].getName(), minHpPokemon[0].getHp());

        Either<Exception, Pokemon> errOrZeraora = pokeClient.findPokemonByName("Zeraora");
        Exception zeraoraErr = errOrZeraora.first();
        if (zeraoraErr != null) {
            throw new RuntimeException(zeraoraErr);
        }
        Pokemon zeraora = errOrZeraora.second();
        StringBuilder typeStr = new StringBuilder(zeraora.getType1());
        if (!"".equalsIgnoreCase(zeraora.getType2())) {
            typeStr.append(",").append(zeraora.getType2());
        }
        System.out.printf("%s %s\n", zeraora.getName(), typeStr);
    }

}