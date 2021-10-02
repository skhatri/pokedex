package com.api;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class DefaultHttpClient implements HttpClient {

    private final String base;

    public DefaultHttpClient(String base) {
        this.base = base;
    }

    public HttpResponse get(String uri) {
        OkHttpClient okHttpClient = new OkHttpClient();
        Request req = new Request.Builder()
            .url(String.format("%s/%s", base, uri))
            .build();
        try (Response response = okHttpClient.newCall(req).execute()) {
            return new HttpResponse(response.body().bytes(), null);
        } catch (Exception ex) {
            return new HttpResponse(new byte[0], ex);
        }
    }
}
