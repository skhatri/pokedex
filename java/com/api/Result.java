package com.api;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class Result {
    private List<Pokemon> data;
    private Meta metadata;

    @JsonCreator
    public Result(@JsonProperty("data") List<Pokemon> data, @JsonProperty("metadata") Meta metadata) {
        this.data = data;
        this.metadata = metadata;
    }

    public List<Pokemon> getData() {
        return data;
    }

    public Meta getMetadata() {
        return metadata;
    }
}
