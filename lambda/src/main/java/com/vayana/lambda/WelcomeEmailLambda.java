package com.vayana.lambda;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import java.util.Map;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.ses.SesClient;
import software.amazon.awssdk.services.ses.model.Body;
import software.amazon.awssdk.services.ses.model.Content;
import software.amazon.awssdk.services.ses.model.Destination;
import software.amazon.awssdk.services.ses.model.Message;
import software.amazon.awssdk.services.ses.model.SendEmailRequest;

public class WelcomeEmailLambda implements RequestHandler<Map<String, String>, String> {

    @Override
    public String handleRequest(Map<String, String> input, Context context) {
        String recipient = input.get("email");
        String sender = System.getenv("SENDER_EMAIL");

        if (recipient == null || recipient.isBlank()) {
            return "Missing email in request";
        }

        if (sender == null || sender.isBlank()) {
            return "Missing SENDER_EMAIL environment variable";
        }

        String subject = "Welcome to Vayana!";
        String bodyText = "Welcome to Vayana! We are excited to have you start tracking your reading.";

        try (SesClient sesClient = SesClient.builder()
                .region(Region.AWS_GLOBAL)
                .build()) {
            SendEmailRequest request = SendEmailRequest.builder()
                    .source(sender)
                    .destination(Destination.builder().toAddresses(recipient).build())
                    .message(Message.builder()
                            .subject(Content.builder().data(subject).build())
                            .body(Body.builder()
                                    .text(Content.builder().data(bodyText).build())
                                    .build())
                            .build())
                    .build();

            sesClient.sendEmail(request);
        }

        return "Welcome email sent";
    }
}
