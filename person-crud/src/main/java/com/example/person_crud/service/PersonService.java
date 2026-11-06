package com.example.person_crud.service;

import com.example.person_crud.entity.Person;
import com.example.person_crud.repository.PersonRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class PersonService {

    private final PersonRepository repository;

    public PersonService(PersonRepository repository) {
        this.repository = repository;
    }

    public List<Person> findAll()                   { return repository.findAll(); }
    public Person findById(Long id)                 { return repository.findById(id).orElseThrow(); }
    public Person save(Person person)               { return repository.save(person); }
    public void delete(Long id)                     { repository.deleteById(id); }
}