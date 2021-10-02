package com.api;

import java.util.List;

public class PokeClient {
    private final HttpClient httpClient;

    PokeClient(HttpClient httpClient) {
        this.httpClient = httpClient;
    }

    public Either<Exception, Pokemon> findPokemonByName(String name) {
        HttpResponse httpResponse = this.httpClient.get(String.format("/search/by-name?q=%s", name));
        if (httpResponse.hasError()) {
            return Either.error(httpResponse.getErr());
        }
        Result result = httpResponse.toData(Result.class);
        Pokemon pokemon = null;
        if (result.getData().size() > 0) {
            pokemon = result.getData().get(0);
        }
        return Either.value(pokemon);
    }

    public Either<Exception, List<Pokemon>> findPokemonsByType(String[] typeName) {
        HttpResponse httpResponse = this.httpClient.get(String.format("/search/by-type?q=%s", String.join(",", typeName)));
        if (httpResponse.hasError()) {
            return Either.error(httpResponse.getErr());
        }
        Result result = httpResponse.toData(Result.class);
        return Either.value(result.getData());
    }
}
