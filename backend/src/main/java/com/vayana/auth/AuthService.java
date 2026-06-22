package com.vayana.auth;

import com.vayana.auth.dto.AuthResponse;
import com.vayana.auth.dto.LoginRequest;
import com.vayana.auth.dto.RegisterRequest;
import com.vayana.domain.User;
import com.vayana.repository.UserRepository;
import com.vayana.security.JwtUtil;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.LocalDateTime;
import java.nio.charset.StandardCharsets;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class AuthService {

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final JwtUtil jwtUtil;
    private final String welcomeEmailUrl;
    private final HttpClient httpClient;

    public AuthService(UserRepository userRepository,
                       PasswordEncoder passwordEncoder,
                       JwtUtil jwtUtil,
                       @Value("${lambda.welcomeEmail.url}") String welcomeEmailUrl) {
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
        this.jwtUtil = jwtUtil;
        this.welcomeEmailUrl = welcomeEmailUrl;
        this.httpClient = HttpClient.newHttpClient();
    }

    public void register(RegisterRequest request) {
        User user = new User();
        user.setUsername(request.getUsername());
        user.setEmail(request.getEmail());
        user.setPassword(passwordEncoder.encode(request.getPassword()));
        user.setCreatedAt(LocalDateTime.now());

        userRepository.save(user);
        triggerWelcomeEmail(user.getEmail());
    }

    public AuthResponse login(LoginRequest request) {
        User user = userRepository.findByEmail(request.getEmail())
                .orElseThrow(() -> new IllegalArgumentException("Invalid credentials"));

        boolean matches = passwordEncoder.matches(request.getPassword(), user.getPassword());
        if (!matches) {
            throw new IllegalArgumentException("Invalid credentials");
        }

        String token = jwtUtil.generateToken(user.getEmail());
        return new AuthResponse(token);
    }

    private void triggerWelcomeEmail(String email) {
        if (welcomeEmailUrl == null || welcomeEmailUrl.isBlank()) {
            return;
        }

        try {
            String payload = "{\"email\":\"" + email + "\"}";

            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(welcomeEmailUrl))
                    .header("Content-Type", "application/json")
                    .POST(HttpRequest.BodyPublishers.ofString(payload, StandardCharsets.UTF_8))
                    .build();

            httpClient.send(request, HttpResponse.BodyHandlers.discarding());
        } catch (Exception ex) {
            // Keep it simple: don't fail registration if the email call fails
        }
    }
}
