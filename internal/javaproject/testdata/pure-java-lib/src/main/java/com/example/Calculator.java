package com.example;

public class Calculator {
    private int last;

    public Calculator() {
        last = 0;
    }

    public int add(int a, int b) {
        last = a + b;
        return last;
    }

    public static int twice(int value) {
        return value * 2;
    }
}
