package com.api;

public class Either<A, B> {
    private A a;
    private B b;

    public Either(A a, B b) {
        this.a = a;
        this.b = b;
    }

    public A first() {
        return a;
    }

    public B second() {
        return b;
    }

    public static <T> Either<Exception, T> error(Exception ex) {
        return new Either<Exception, T>(ex, null);
    }

    public static <T> Either<Exception, T> value(T value) {
        return new Either<Exception, T>(null, value);
    }
}
