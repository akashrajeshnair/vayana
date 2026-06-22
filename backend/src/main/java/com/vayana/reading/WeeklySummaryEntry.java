package com.vayana.reading;

public class WeeklySummaryEntry {

    private String email;
    private int booksFinishedThisWeek;
    private String currentlyReading;

    public WeeklySummaryEntry() {
    }

    public WeeklySummaryEntry(String email, int booksFinishedThisWeek, String currentlyReading) {
        this.email = email;
        this.booksFinishedThisWeek = booksFinishedThisWeek;
        this.currentlyReading = currentlyReading;
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
