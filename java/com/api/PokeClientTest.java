package com.api;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@DisplayName("Exception Tests")
public class PokeClientTest {

    static class MockHttpClient implements HttpClient {
        private Map<String, String> items;

        MockHttpClient() {
            this.items = new HashMap<>();
        }

        public MockHttpClient when(String uri, String dataFile) {
            this.items.put(uri, dataFile);
            return this;
        }

        @Override
        public HttpResponse get(String uri) {
            String dataFile = this.items.get(uri);
            if (dataFile == null) {
                return new HttpResponse(new byte[0], new RuntimeException("resource not found"));
            }
            try {
                byte[] b = Files.readAllBytes(Path.of(dataFile));
                return new HttpResponse(b, null);
            } catch (Exception ex) {
                return new HttpResponse(null, ex);
            }
        }
    }
    private MockHttpClient mockClient;

    @BeforeEach
    public void setup() {
        mockClient  = new MockHttpClient();
        mockClient.when("/search/by-name?q=mimikyu", "./client-data/mimikyu.json").
            when("/search/by-type?q=flying,water", "./client-data/flying-water.json").
            when("/search/by-type?q=bug,ghost", "./client-data/bug-ghost.json").
            when("/search/by-name?q=Zeraora", "./client-data/zeraora.json");
    }

    @Test
    @DisplayName("test mimikyu ability")
    public void testMimikyuAbility() {
        PokeClient pokeClient = PokeClients.newCustomPokedexClient(mockClient);
        Either<Exception, Pokemon> errOrMimikyu = pokeClient.findPokemonByName("mimikyu");
        Exception err = errOrMimikyu.first();
        if (err != null) {
            throw new RuntimeException(err);
        }
        Assertions.assertEquals("Disguise", errOrMimikyu.second().getAbility1());
    }

    @Test
    @DisplayName("test count of pokemons of type flying and water")
    public void testCountFlyingAndWaterTypePokemon() {
        PokeClient pokeClient = PokeClients.newCustomPokedexClient(mockClient);
        Either<Exception, List<Pokemon>> errOrPokemons = pokeClient.findPokemonsByType(new String[]{"flying", "water"});
        Exception err = errOrPokemons.first();
        if (err != null) {
            throw new RuntimeException(err);
        }
        List<Pokemon> pokemons = errOrPokemons.second();
        Assertions.assertEquals( 8, pokemons.size());
    }

    @Test
    @DisplayName("test find bug and ghost pokemon with lowest hp")
    public void testFindBugGhostPokemonWithLowestHp() {
        PokeClient pokeClient = PokeClients.newCustomPokedexClient(mockClient);
        Either<Exception, List<Pokemon>> errOrPokemons = pokeClient.findPokemonsByType(new String[]{"bug", "ghost"});
        Exception err = errOrPokemons.first();
        if (err != null) {
            throw new RuntimeException(err);
        }
        String expectedName = "Shedinja";
        int expectedHp = 1;
        List<Pokemon> pokemons = errOrPokemons.second();
        Pokemon minHpPokemon = null;
        for (Pokemon poke : pokemons) {
            if (minHpPokemon == null || poke.getHp() < minHpPokemon.getHp()) {
                minHpPokemon = poke;
            }
        }
        Assertions.assertEquals(expectedName, minHpPokemon.getName());
        Assertions.assertEquals(expectedHp, minHpPokemon.getHp());
    }

    @Test
    @DisplayName("test find zeraoras type")
    public void testFindZeraorasType() {
        PokeClient pokeClient = PokeClients.newCustomPokedexClient(mockClient);
        Either<Exception, Pokemon> errOrZeraora = pokeClient.findPokemonByName("Zeraora");
        Exception err = errOrZeraora.first();
        if (err != null) {
            throw new RuntimeException(err);
        }
        Pokemon zeraora = errOrZeraora.second();
        Assertions.assertEquals("Electric", zeraora.getType1());
        Assertions.assertEquals("", zeraora.getType2());
    }

}