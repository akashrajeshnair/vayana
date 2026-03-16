package com.vayana.reading;

import com.vayana.domain.Book;
import com.vayana.domain.ReadingRecord;
import com.vayana.domain.User;
import com.vayana.repository.BookRepository;
import com.vayana.repository.ReadingRecordRepository;
import com.vayana.repository.UserRepository;
import java.util.List;
import java.util.stream.Collectors;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

@Service
public class ReadingRecordService {

    private final ReadingRecordRepository readingRecordRepository;
    private final BookRepository bookRepository;
    private final UserRepository userRepository;

    public ReadingRecordService(ReadingRecordRepository readingRecordRepository,
                                BookRepository bookRepository,
                                UserRepository userRepository) {
        this.readingRecordRepository = readingRecordRepository;
        this.bookRepository = bookRepository;
        this.userRepository = userRepository;
    }

    public List<ReadingRecordDTO> getShelf() {
        User user = getCurrentUser();
        return readingRecordRepository.findByUser(user)
                .stream()
                .map(this::toDto)
                .collect(Collectors.toList());
    }

    public ReadingRecordDTO addToShelf(CreateReadingRecordRequest request) {
        User user = getCurrentUser();
        Book book = bookRepository.findById(request.getBookId())
                .orElseThrow(() -> new IllegalArgumentException("Book not found"));

        ReadingRecord record = new ReadingRecord();
        record.setUser(user);
        record.setBook(book);
        record.setStatus(request.getStatus());

        ReadingRecord saved = readingRecordRepository.save(record);
        return toDto(saved);
    }

    public ReadingRecordDTO updateRecord(Long recordId, UpdateReadingRecordRequest request) {
        User user = getCurrentUser();
        ReadingRecord record = readingRecordRepository.findByIdAndUser(recordId, user)
                .orElseThrow(() -> new IllegalArgumentException("Reading record not found"));

        record.setStatus(request.getStatus());
        record.setRating(request.getRating());
        record.setReview(request.getReview());
        record.setStartDate(request.getStartDate());
        record.setFinishDate(request.getFinishDate());

        ReadingRecord saved = readingRecordRepository.save(record);
        return toDto(saved);
    }

    public void removeFromShelf(Long recordId) {
        User user = getCurrentUser();
        ReadingRecord record = readingRecordRepository.findByIdAndUser(recordId, user)
                .orElseThrow(() -> new IllegalArgumentException("Reading record not found"));
        readingRecordRepository.delete(record);
    }

    private User getCurrentUser() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        if (authentication == null || authentication.getName() == null) {
            throw new IllegalArgumentException("User not authenticated");
        }

        return userRepository.findByEmail(authentication.getName())
                .orElseThrow(() -> new IllegalArgumentException("User not found"));
    }

    private ReadingRecordDTO toDto(ReadingRecord record) {
        return new ReadingRecordDTO(
                record.getId(),
                record.getBook().getId(),
                record.getStatus(),
                record.getRating(),
                record.getReview(),
                record.getStartDate(),
                record.getFinishDate()
        );
    }
}
