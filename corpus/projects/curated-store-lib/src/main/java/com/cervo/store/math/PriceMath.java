package com.cervo.store.math;

public class PriceMath {
    public static int addTax(int subtotal, int tax) {
        return subtotal + tax;
    }

    public static int discount(int subtotal, int amount) {
        if (subtotal > amount) {
            return subtotal - amount;
        }
        return 0;
    }
}
