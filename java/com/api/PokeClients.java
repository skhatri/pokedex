package com.api;

public class PokeClients {
    public static PokeClient newPokeClient() {
        return newCustomPokedexClient(new DefaultHttpClient("http://localhost:6100"));
    }

    public static PokeClient newCustomPokedexClient(HttpClient httpClient) {
        return new PokeClient(httpClient);
    }
}
