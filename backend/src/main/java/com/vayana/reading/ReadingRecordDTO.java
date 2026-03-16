package com.vayana.reading;

import java.time.LocalDate;

public class ReadingRecordDTO {

    private Long id;
    private Long bookId;
    private String status;
    private Integer rating;
    private String review;
    private LocalDate startDate;
    private LocalDate finishDate;

    public ReadingRecordDTO() {
    }

    public ReadingRecordDTO(Long id,
                            Long bookId,
                            String status,
                            Integer rating,
                            String review,
                            LocalDate startDate,
                            LocalDate finishDate) {
        this.id = id;
        this.bookId = bookId;
        this.status = status;
        this.rating = rating;
        this.review = review;
        this.startDate = startDate;
        this.finishDate = finishDate;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
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

    public Integer getRating() {
        return rating;
    }

    public void setRating(Integer rating) {
        this.rating = rating;
    }

    public String getReview() {
        return review;
    }

    public void setReview(String review) {
        this.review = review;
    }

    public LocalDate getStartDate() {
        return startDate;
    }

    public void setStartDate(LocalDate startDate) {
        this.startDate = startDate;
    }

    public LocalDate getFinishDate() {
        return finishDate;
    }

    public void setFinishDate(LocalDate finishDate) {
        this.finishDate = finishDate;
    }
}
