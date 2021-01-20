INSERT INTO users (name, nick, email, password) 
values
( "User 1", "User1", "user1@gmail.com", "$2a$10$8P8DLV61bpYPYV1.sNOpD.ThV3J6NuIHCsrOB6UbM46vCoLP86dKu" ),
( "User 2", "User2", "user2@gmail.com", "$2a$10$8P8DLV61bpYPYV1.sNOpD.ThV3J6NuIHCsrOB6UbM46vCoLP86dKu" ),
( "User 3", "User3", "user3@gmail.com", "$2a$10$8P8DLV61bpYPYV1.sNOpD.ThV3J6NuIHCsrOB6UbM46vCoLP86dKu" );


INSERT INTO followers (user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);


insert into publications(title, content, author_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);
