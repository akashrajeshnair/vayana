package com.vayana.reading;

import java.util.List;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/shelf")
public class ReadingRecordController {

    private final ReadingRecordService readingRecordService;

    public ReadingRecordController(ReadingRecordService readingRecordService) {
        this.readingRecordService = readingRecordService;
    }

    @GetMapping
    public ResponseEntity<List<ReadingRecordDTO>> getShelf() {
        return ResponseEntity.ok(readingRecordService.getShelf());
    }

    @PostMapping
    public ResponseEntity<ReadingRecordDTO> addToShelf(@RequestBody CreateReadingRecordRequest request) {
        return ResponseEntity.ok(readingRecordService.addToShelf(request));
    }

    @PutMapping("/{id}")
    public ResponseEntity<ReadingRecordDTO> updateRecord(@PathVariable Long id,
                                                         @RequestBody UpdateReadingRecordRequest request) {
        return ResponseEntity.ok(readingRecordService.updateRecord(id, request));
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> removeFromShelf(@PathVariable Long id) {
        readingRecordService.removeFromShelf(id);
        return ResponseEntity.noContent().build();
    }
}
