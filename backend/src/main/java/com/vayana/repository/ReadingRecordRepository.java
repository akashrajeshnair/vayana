package com.vayana.repository;

import com.vayana.domain.ReadingRecord;
import com.vayana.domain.User;
import java.util.List;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ReadingRecordRepository extends JpaRepository<ReadingRecord, Long> {
	List<ReadingRecord> findByUser(User user);

	Optional<ReadingRecord> findByIdAndUser(Long id, User user);
}
