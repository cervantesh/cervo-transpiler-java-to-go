package com.cervo.store.service;

import com.cervo.store.math.PriceMath;

public class OrderService {
    public static int total(int subtotal, int tax) {
        return PriceMath.addTax(subtotal, tax);
    }

    public static int discountedTotal(int subtotal, int amount) {
        return PriceMath.discount(subtotal, amount);
    }
}
