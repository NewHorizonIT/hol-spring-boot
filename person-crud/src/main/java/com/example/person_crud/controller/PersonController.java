package com.example.person_crud.controller;

import com.example.person_crud.entity.Person;
import com.example.person_crud.service.PersonService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.lang.Long;

@RestController
@RequestMapping("/api/persons")
public class PersonController {

    private final PersonService service;

    public PersonController(PersonService service) {
        this.service = service;
    }

    @GetMapping
    public List<Person> getAll() { return service.findAll(); }

    @GetMapping("/{id}")
    public Person getById(@PathVariable Long id) { return service.findById(id); }

    @PostMapping
    public Person create(@RequestBody Person person) { return service.save(person); }

    @PutMapping("/{id}")
    public Person update(@PathVariable Long id, @RequestBody Person person) {
        person.setId(id);
        return service.save(person);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> delete(@PathVariable Long id) {
        service.delete(id);
        return ResponseEntity.noContent().build();
    }
}