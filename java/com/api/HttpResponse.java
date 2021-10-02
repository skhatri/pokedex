package com.api;

public class HttpResponse {
    private final byte[] data;
    private final Exception err;

    public HttpResponse(byte[] data, Exception err) {
        this.data = data;
        this.err = err;
    }

    public byte[] getData() {
        return data;
    }

    public Exception getErr() {
        return err;
    }

    public boolean hasError() {
        return err != null;
    }

    public <T> T toData(Class<T> clz) {
        return Json.getInstance().toInstance(this.data, clz);
    }
}
