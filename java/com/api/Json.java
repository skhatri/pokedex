package com.api;

import com.fasterxml.jackson.databind.ObjectMapper;

public class Json {
    private final ObjectMapper mapper;

    private Json() {
        this.mapper = new ObjectMapper();
    }

    public String toJson(Object instance) {
        try {
            return this.mapper.writeValueAsString(instance);
        } catch (Exception ex) {
            throw new RuntimeException(ex);
        }
    }

    public <T> T toInstance(byte[] data, Class<T> clz) {
        try {
            return this.mapper.readValue(data, clz);
        } catch (Exception ex) {
            throw new RuntimeException(ex);
        }
    }

    public static Json getInstance() {
        return INSTANCE;
    }

    public static final Json INSTANCE = new Json();
}
