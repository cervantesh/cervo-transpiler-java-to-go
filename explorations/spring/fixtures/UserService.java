package com.example.user;

import org.springframework.stereotype.Service;

@Service
public class UserService {
    private final UserRepository repository;

    public UserService(UserRepository repository) {
        this.repository = repository;
    }

    public UserDto findUser(String id) {
        UserEntity entity = repository.findByExternalId(id);
        return new UserDto(entity.getExternalId(), entity.getDisplayName());
    }
}
