package com.example.config;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "app")
public class AppProperties {
    private String publicBaseUrl;
    private int cacheSeconds;

    public String getPublicBaseUrl() {
        return publicBaseUrl;
    }

    public int getCacheSeconds() {
        return cacheSeconds;
    }
}
