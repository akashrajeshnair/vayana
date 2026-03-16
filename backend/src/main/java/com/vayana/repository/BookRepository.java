package com.vayana.repository;

import com.vayana.domain.Book;
import java.util.List;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

public interface BookRepository extends JpaRepository<Book, Long> {
	@Query("SELECT b FROM Book b WHERE LOWER(b.title) LIKE LOWER(CONCAT('%', :query, '%')) "
			+ "OR LOWER(b.author) LIKE LOWER(CONCAT('%', :query, '%'))")
	List<Book> searchByTitleOrAuthor(@Param("query") String query);
}
