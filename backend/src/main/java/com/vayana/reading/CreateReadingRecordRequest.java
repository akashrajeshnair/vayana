package com.vayana.reading;

public class CreateReadingRecordRequest {

    private Long bookId;
    private String status;

    public CreateReadingRecordRequest() {
    }

    public Long getBookId() {
        return bookId;
    }

    public void setBookId(Long bookId) {
        this.bookId = bookId;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }
}
