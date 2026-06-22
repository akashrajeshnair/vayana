package com.vayana.lambda;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import java.util.List;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.ses.SesClient;
import software.amazon.awssdk.services.ses.model.Body;
import software.amazon.awssdk.services.ses.model.Content;
import software.amazon.awssdk.services.ses.model.Destination;
import software.amazon.awssdk.services.ses.model.Message;
import software.amazon.awssdk.services.ses.model.SendEmailRequest;

public class WeeklySummaryLambda implements RequestHandler<List<WeeklySummaryLambda.SummaryEntry>, String> {

    @Override
    public String handleRequest(List<SummaryEntry> input, Context context) {
        String sender = System.getenv("SENDER_EMAIL");
        if (sender == null || sender.isBlank()) {
            return "Missing SENDER_EMAIL environment variable";
        }

        if (input == null || input.isEmpty()) {
            return "No summary entries to process";
        }

        try (SesClient sesClient = SesClient.builder().region(Region.AWS_GLOBAL).build()) {
            for (SummaryEntry entry : input) {
                if (entry.getEmail() == null || entry.getEmail().isBlank()) {
                    continue;
                }

                String subject = "Your BookTrack Weekly Summary";
                String bodyText = buildBody(entry);

                SendEmailRequest request = SendEmailRequest.builder()
                        .source(sender)
                        .destination(Destination.builder().toAddresses(entry.getEmail()).build())
                        .message(Message.builder()
                                .subject(Content.builder().data(subject).build())
                                .body(Body.builder()
                                        .text(Content.builder().data(bodyText).build())
                                        .build())
                                .build())
                        .build();

                sesClient.sendEmail(request);
            }
        }

        return "Weekly summary emails sent";
    }

    private String buildBody(SummaryEntry entry) {
        String currentlyReading = entry.getCurrentlyReading() == null
                ? "" : entry.getCurrentlyReading();

        return "Here is your weekly reading summary:\n"
                + "Books finished this week: " + entry.getBooksFinishedThisWeek() + "\n"
                + "Currently reading: " + (currentlyReading.isBlank() ? "None" : currentlyReading) + "\n";
    }

    public static class SummaryEntry {
        private String email;
        private int booksFinishedThisWeek;
        private String currentlyReading;

        public SummaryEntry() {
        }

        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }

        public int getBooksFinishedThisWeek() {
            return booksFinishedThisWeek;
        }

        public void setBooksFinishedThisWeek(int booksFinishedThisWeek) {
            this.booksFinishedThisWeek = booksFinishedThisWeek;
        }

        public String getCurrentlyReading() {
            return currentlyReading;
        }

        public void setCurrentlyReading(String currentlyReading) {
            this.currentlyReading = currentlyReading;
        }
    }
}
