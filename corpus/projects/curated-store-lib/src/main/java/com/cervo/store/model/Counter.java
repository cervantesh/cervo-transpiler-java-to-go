package com.cervo.store.model;

public class Counter {
    private int value;

    public Counter(int value) {
    }

    public void add(int delta) {
        value = value + delta;
    }

    public int current() {
        return value;
    }
}
