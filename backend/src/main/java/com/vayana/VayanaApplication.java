package com.vayana;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.scheduling.annotation.EnableScheduling;

@SpringBootApplication
@EnableScheduling
public class VayanaApplication {

    public static void main(String[] args) {
        SpringApplication.run(VayanaApplication.class, args);
    }
}
