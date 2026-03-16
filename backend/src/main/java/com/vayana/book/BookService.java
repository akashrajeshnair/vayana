package com.vayana.book;

import com.vayana.domain.Book;
import com.vayana.repository.BookRepository;
import java.util.List;
import java.util.stream.Collectors;
import org.springframework.stereotype.Service;

@Service
public class BookService {

    private final BookRepository bookRepository;

    public BookService(BookRepository bookRepository) {
        this.bookRepository = bookRepository;
    }

    public List<BookDTO> getAllBooks() {
        return bookRepository.findAll()
                .stream()
                .map(this::toDto)
                .collect(Collectors.toList());
    }

    public BookDTO getBookById(Long id) {
        Book book = bookRepository.findById(id)
                .orElseThrow(() -> new IllegalArgumentException("Book not found"));
        return toDto(book);
    }

    public BookDTO createBook(BookDTO request) {
        Book book = new Book();
        book.setTitle(request.getTitle());
        book.setAuthor(request.getAuthor());
        book.setGenre(request.getGenre());
        book.setCoverImageUrl(request.getCoverImageUrl());
        book.setDescription(request.getDescription());

        Book saved = bookRepository.save(book);
        return toDto(saved);
    }

    public List<BookDTO> searchBooks(String query) {
        return bookRepository.searchByTitleOrAuthor(query)
                .stream()
                .map(this::toDto)
                .collect(Collectors.toList());
    }

    private BookDTO toDto(Book book) {
        return new BookDTO(
                book.getId(),
                book.getTitle(),
                book.getAuthor(),
                book.getGenre(),
                book.getCoverImageUrl(),
                book.getDescription()
        );
    }
}
