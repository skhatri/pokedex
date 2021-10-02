package com.api;


import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

public class Meta {
    private int total;

    @JsonCreator
    public Meta(@JsonProperty("total") int total) {
        this.total = total;
    }

    public int getTotal() {
        return total;
    }
}
