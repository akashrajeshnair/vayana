package com.vayana.reading;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.vayana.domain.ReadingRecord;
import com.vayana.domain.User;
import com.vayana.repository.ReadingRecordRepository;
import com.vayana.repository.UserRepository;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import java.time.LocalDate;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

@Component
public class WeeklySummaryJob {

    private final UserRepository userRepository;
    private final ReadingRecordRepository readingRecordRepository;
    private final ObjectMapper objectMapper;
    private final HttpClient httpClient;
    private final String weeklySummaryUrl;

    public WeeklySummaryJob(UserRepository userRepository,
                            ReadingRecordRepository readingRecordRepository,
                            ObjectMapper objectMapper,
                            @Value("${lambda.weeklySummary.url}") String weeklySummaryUrl) {
        this.userRepository = userRepository;
        this.readingRecordRepository = readingRecordRepository;
        this.objectMapper = objectMapper;
        this.weeklySummaryUrl = weeklySummaryUrl;
        this.httpClient = HttpClient.newHttpClient();
    }

    @Scheduled(cron = "0 0 9 * * MON")
    public void sendWeeklySummary() {
        if (weeklySummaryUrl == null || weeklySummaryUrl.isBlank()) {
            return;
        }

        LocalDate oneWeekAgo = LocalDate.now().minusDays(7);
        List<WeeklySummaryEntry> entries = userRepository.findAll()
                .stream()
                .map(user -> buildEntry(user, oneWeekAgo))
                .collect(Collectors.toList());

        try {
            String payload = objectMapper.writeValueAsString(entries);
            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(weeklySummaryUrl))
                    .header("Content-Type", "application/json")
                    .POST(HttpRequest.BodyPublishers.ofString(payload, StandardCharsets.UTF_8))
                    .build();

            httpClient.send(request, HttpResponse.BodyHandlers.discarding());
        } catch (Exception ex) {
            // Keep it simple: skip failures
        }
    }

    private WeeklySummaryEntry buildEntry(User user, LocalDate oneWeekAgo) {
        long booksFinished = readingRecordRepository.countByUserAndFinishDateAfter(user, oneWeekAgo);
        Optional<ReadingRecord> currentlyReading = readingRecordRepository.findFirstByUserAndStatus(user, "READING");
        String currentTitle = currentlyReading.map(record -> record.getBook().getTitle()).orElse("");

        return new WeeklySummaryEntry(user.getEmail(), (int) booksFinished, currentTitle);
    }
}
