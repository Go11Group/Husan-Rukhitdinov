sudo -u postgres psql
[sudo] password for husan: 
\c
 create database Book;
CREATE DATABASE
\c book;

create table author
(
    id   int primary key,
    name varchar not null
);

INSERT INTO author (id, name) VALUES
    (1, 'Harper Lee'),
    (2, 'George Orwell'),
    (3, 'Jane Austen'),
    (4, 'F. Scott Fitzgerald'),
    (5, 'Ernest Hemingway'),
    (6, 'Maya Angelou'),
    (7, 'Gabriel García Márquez');

create table book
(
    id          int primary key,
    name        varchar not null,
    page        int4,
    author_name varchar
);

    INSERT INTO book (id, name, page, author_name) VALUES
    (1, 'To Kill a Mockingbird', 281, 'Harper Lee'),
    (2, '1984', 328, 'George Orwell'),
    (3, 'Pride and Prejudice', 432, 'Jane Austen'),
    (4, 'The Great Gatsby', 180, 'F. Scott Fitzgerald'),
    (5, 'The Old Man and the Sea', 127, 'Ernest Hemingway'),
    (6, 'I Know Why the Caged Bird Sings', 289, 'Maya Angelou'),
    (7, 'One Hundred Years of Solitude', 417, 'Gabriel García Márquez'),
    (8, 'The Kite Runner', 371, NULL),
    (9, 'The Immortal Life of Henrietta Lacks', 369, NULL),
    (10, 'The Handmaid''s Tale', 311, NULL),
    (11, 'The Shining', 447, NULL),
    (12, 'The Grapes of Wrath', 464, NULL),
    (13, 'To the Lighthouse', 209, NULL),
    (14, 'The Catcher in the Rye', 277, NULL),
    (15, 'Beloved', 321, NULL);

SELECT b.name, b.page, a.name AS author_name
FROM book b
LEFT JOIN author a ON b.author_name = a.name;

                name                 | page |      author_name       
--------------------------------------+------+------------------------
 The Old Man and the Sea              |  127 | Ernest Hemingway
 The Great Gatsby                     |  180 | F. Scott Fitzgerald
 One Hundred Years of Solitude        |  417 | Gabriel García Márquez
 1984                                 |  328 | George Orwell
 To Kill a Mockingbird                |  281 | Harper Lee
 Pride and Prejudice                  |  432 | Jane Austen
 I Know Why the Caged Bird Sings      |  289 | Maya Angelou
 The Catcher in the Rye               |  277 | 
 Beloved                              |  321 | 
 The Kite Runner                      |  371 | 
 The Immortal Life of Henrietta Lacks |  369 | 
 The Handmaid's Tale                  |  311 | 
 The Shining                          |  447 | 
 The Grapes of Wrath                  |  464 | 
 To the Lighthouse                    |  209 | 
(15 rows)
