## PostgreSQL

In local development environment (using docker) database is `bookappdb`, username is `postgres`, password is `password`

#### `users` table

```
mybooklibrarydb=# CREATE TABLE users (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# email VARCHAR(255),
mybooklibrarydb(# first_name VARCHAR(255),
mybooklibrarydb(# last_name VARCHAR(255),
mybooklibrarydb(# password VARCHAR(60),
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# user_active INT DEFAULT 0);
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d users
                                         Table "public.users"
   Column    |            Type             | Collation | Nullable |              Default              
-------------+-----------------------------+-----------+----------+-----------------------------------
 id          | integer                     |           | not null | nextval('users_id_seq'::regclass)
 email       | character varying(255)      |           |          | 
 first_name  | character varying(255)      |           |          | 
 last_name   | character varying(255)      |           |          | 
 password    | character varying(60)       |           |          | 
 created_at  | timestamp without time zone |           |          | now()
 updated_at  | timestamp without time zone |           |          | now()
 user_active | integer                     |           |          | 0
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# 
```

#### `tokens` table

```
mybooklibrarydb=# CREATE TABLE tokens (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# user_id INT,
mybooklibrarydb(# email VARCHAR(255),
mybooklibrarydb(# token VARCHAR(255),
mybooklibrarydb(# token_hash BYTEA,
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# expiry TIMESTAMP WITHOUT TIME ZONE);
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d tokens
                                        Table "public.tokens"
   Column   |            Type             | Collation | Nullable |              Default               
------------+-----------------------------+-----------+----------+------------------------------------
 id         | integer                     |           | not null | nextval('tokens_id_seq'::regclass)
 user_id    | integer                     |           |          | 
 email      | character varying(255)      |           |          | 
 token      | character varying(255)      |           |          | 
 token_hash | bytea                       |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
 expiry     | timestamp without time zone |           |          | 
Indexes:
    "tokens_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=#
```

#### `publishers` table

```
mybooklibrarydb=# CREATE TABLE publishers (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# publisher_name VARCHAR(512),
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW());
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d
               List of relations
 Schema |       Name        |   Type   | Owner 
--------+-------------------+----------+-------
 public | authors           | table    | lcs
 public | authors_id_seq    | sequence | lcs
 public | genres            | table    | lcs
 public | genres_id_seq     | sequence | lcs
 public | publishers        | table    | lcs
 public | publishers_id_seq | sequence | lcs
 public | tokens            | table    | lcs
 public | tokens_id_seq     | sequence | lcs
 public | users             | table    | lcs
 public | users_id_seq      | sequence | lcs
(10 rows)

mybooklibrarydb=# \d publishers
                                          Table "public.publishers"
     Column     |            Type             | Collation | Nullable |                Default                 
----------------+-----------------------------+-----------+----------+----------------------------------------
 id             | integer                     |           | not null | nextval('publishers_id_seq'::regclass)
 publisher_name | character varying(512)      |           |          | 
 created_at     | timestamp without time zone |           |          | now()
 updated_at     | timestamp without time zone |           |          | now()
Indexes:
    "publishers_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# 
```

#### `mybooks` table

```
mybooklibrarydb=# CREATE TABLE mybooks (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# title VARCHAR(512),
mybooklibrarydb(# publisher_id INT,
mybooklibrarydb(# publication_year INT,
mybooklibrarydb(# description TEXT,
mybooklibrarydb(# slug VARCHAR(512),
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW());
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d mybooks
                                           Table "public.mybooks"
      Column      |            Type             | Collation | Nullable |               Default               
------------------+-----------------------------+-----------+----------+-------------------------------------
 id               | integer                     |           | not null | nextval('mybooks_id_seq'::regclass)
 title            | character varying(512)      |           |          | 
 publisher_id     | integer                     |           |          | 
 publication_year | integer                     |           |          | 
 description      | text                        |           |          | 
 slug             | character varying(512)      |           |          | 
 created_at       | timestamp without time zone |           |          | now()
 updated_at       | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=#
```

#### `mybooks_authors`

```
mybooklibrarydb=# 
mybooklibrarydb=# CREATE TABLE mybooks_authors (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# book_id INT,
mybooklibrarydb(# author_id INT,
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW());
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d mybooks_authors
                                        Table "public.mybooks_authors"
   Column   |            Type             | Collation | Nullable |                   Default                   
------------+-----------------------------+-----------+----------+---------------------------------------------
 id         | integer                     |           | not null | nextval('mybooks_authors_id_seq'::regclass)
 book_id    | integer                     |           |          | 
 author_id  | integer                     |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_authors_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=#
```

#### `mybooks_genres`

```
mybooklibrarydb=# 
mybooklibrarydb=# CREATE TABLE mybooks_genres (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# book_id INT,
mybooklibrarydb(# genre_id INT,
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW());
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d
                 List of relations
 Schema |         Name          |   Type   | Owner 
--------+-----------------------+----------+-------
 public | authors               | table    | lcs
 public | authors_id_seq        | sequence | lcs
 public | genres                | table    | lcs
 public | genres_id_seq         | sequence | lcs
 public | mybooks               | table    | lcs
 public | mybooks_genres        | table    | lcs
 public | mybooks_genres_id_seq | sequence | lcs
 public | mybooks_id_seq        | sequence | lcs
 public | publishers            | table    | lcs
 public | publishers_id_seq     | sequence | lcs
 public | tokens                | table    | lcs
 public | tokens_id_seq         | sequence | lcs
 public | users                 | table    | lcs
 public | users_id_seq          | sequence | lcs
(14 rows)

mybooklibrarydb=# 
mybooklibrarydb=# \d mybooks_genres
                                        Table "public.mybooks_genres"
   Column   |            Type             | Collation | Nullable |                  Default                   
------------+-----------------------------+-----------+----------+--------------------------------------------
 id         | integer                     |           | not null | nextval('mybooks_genres_id_seq'::regclass)
 book_id    | integer                     |           |          | 
 genre_id   | integer                     |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_genres_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# 
mybooklibrarydb=# 
mybooklibrarydb=# 
```

#### `authors` table

```
mybooklibrarydb=# 
mybooklibrarydb=# CREATE TABLE authors (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# author_name VARCHAR(512),
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW());
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d authors
                                         Table "public.authors"
   Column    |            Type             | Collation | Nullable |               Default               
-------------+-----------------------------+-----------+----------+-------------------------------------
 id          | integer                     |           | not null | nextval('authors_id_seq'::regclass)
 author_name | character varying(512)      |           |          | 
 created_at  | timestamp without time zone |           |          | now()
 updated_at  | timestamp without time zone |           |          | now()
Indexes:
    "authors_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# \d
             List of relations
 Schema |      Name      |   Type   | Owner 
--------+----------------+----------+-------
 public | authors        | table    | lcs
 public | authors_id_seq | sequence | lcs
 public | genres         | table    | lcs
 public | genres_id_seq  | sequence | lcs
 public | tokens         | table    | lcs
 public | tokens_id_seq  | sequence | lcs
 public | users          | table    | lcs
 public | users_id_seq   | sequence | lcs
(8 rows)

mybooklibrarydb=#
```

#### `genres` table

```
mybooklibrarydb=# 
mybooklibrarydb=# CREATE TABLE genres (
mybooklibrarydb(# id SERIAL PRIMARY KEY NOT NULL,
mybooklibrarydb(# genre_name VARCHAR(255),
mybooklibrarydb(# created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
mybooklibrarydb(# updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW());
CREATE TABLE
mybooklibrarydb=# 
mybooklibrarydb=# 
mybooklibrarydb=# \d
             List of relations
 Schema |     Name      |   Type   | Owner 
--------+---------------+----------+-------
 public | genres        | table    | lcs
 public | genres_id_seq | sequence | lcs
 public | tokens        | table    | lcs
 public | tokens_id_seq | sequence | lcs
 public | users         | table    | lcs
 public | users_id_seq  | sequence | lcs
(6 rows)

mybooklibrarydb=# 
mybooklibrarydb=# \d genres
                                        Table "public.genres"
   Column   |            Type             | Collation | Nullable |              Default               
------------+-----------------------------+-----------+----------+------------------------------------
 id         | integer                     |           | not null | nextval('genres_id_seq'::regclass)
 genre_name | character varying(255)      |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "genres_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# 
```

#### Foreign Keys

```
mybooklibrarydb=# \d tokens
                                        Table "public.tokens"
   Column   |            Type             | Collation | Nullable |              Default               
------------+-----------------------------+-----------+----------+------------------------------------
 id         | integer                     |           | not null | nextval('tokens_id_seq'::regclass)
 user_id    | integer                     |           |          | 
 email      | character varying(255)      |           |          | 
 token      | character varying(255)      |           |          | 
 token_hash | bytea                       |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
 expiry     | timestamp without time zone |           |          | 
Indexes:
    "tokens_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# 
mybooklibrarydb=# ALTER TABLE tokens ADD CONSTRAINT fk_users_tokens FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d tokens
                                        Table "public.tokens"
   Column   |            Type             | Collation | Nullable |              Default               
------------+-----------------------------+-----------+----------+------------------------------------
 id         | integer                     |           | not null | nextval('tokens_id_seq'::regclass)
 user_id    | integer                     |           |          | 
 email      | character varying(255)      |           |          | 
 token      | character varying(255)      |           |          | 
 token_hash | bytea                       |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
 expiry     | timestamp without time zone |           |          | 
Indexes:
    "tokens_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "fk_users_tokens" FOREIGN KEY (user_id) REFERENCES users(id)

mybooklibrarydb=# 
mybooklibrarydb=# 
mybooklibrarydb=# \d users
                                         Table "public.users"
   Column    |            Type             | Collation | Nullable |              Default              
-------------+-----------------------------+-----------+----------+-----------------------------------
 id          | integer                     |           | not null | nextval('users_id_seq'::regclass)
 email       | character varying(255)      |           |          | 
 first_name  | character varying(255)      |           |          | 
 last_name   | character varying(255)      |           |          | 
 password    | character varying(60)       |           |          | 
 created_at  | timestamp without time zone |           |          | now()
 updated_at  | timestamp without time zone |           |          | now()
 user_active | integer                     |           |          | 0
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
Referenced by:
    TABLE "tokens" CONSTRAINT "fk_users_tokens" FOREIGN KEY (user_id) REFERENCES users(id)

mybooklibrarydb=# 
```

```
mybooklibrarydb=# 
mybooklibrarydb=# \d mybooks_genres
                                        Table "public.mybooks_genres"
   Column   |            Type             | Collation | Nullable |                  Default                   
------------+-----------------------------+-----------+----------+--------------------------------------------
 id         | integer                     |           | not null | nextval('mybooks_genres_id_seq'::regclass)
 book_id    | integer                     |           |          | 
 genre_id   | integer                     |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_genres_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# 
mybooklibrarydb=# ALTER TABLE mybooks_genres ADD CONSTRAINT fk_mybooks_book_id FOREIGN KEY (book_id) REFERENCES mybooks(id);
ALTER TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d mybooks_genres
                                        Table "public.mybooks_genres"
   Column   |            Type             | Collation | Nullable |                  Default                   
------------+-----------------------------+-----------+----------+--------------------------------------------
 id         | integer                     |           | not null | nextval('mybooks_genres_id_seq'::regclass)
 book_id    | integer                     |           |          | 
 genre_id   | integer                     |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_genres_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "fk_mybooks_book_id" FOREIGN KEY (book_id) REFERENCES mybooks(id)

mybooklibrarydb=# ALTER TABLE mybooks_genres ADD CONSTRAINT fk_genres_genre_id FOREIGN KEY (genre_id) REFERENCES genres(id);
ALTER TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d mybooks_genres
                                        Table "public.mybooks_genres"
   Column   |            Type             | Collation | Nullable |                  Default                   
------------+-----------------------------+-----------+----------+--------------------------------------------
 id         | integer                     |           | not null | nextval('mybooks_genres_id_seq'::regclass)
 book_id    | integer                     |           |          | 
 genre_id   | integer                     |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_genres_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "fk_genres_genre_id" FOREIGN KEY (genre_id) REFERENCES genres(id)
    "fk_mybooks_book_id" FOREIGN KEY (book_id) REFERENCES mybooks(id)

mybooklibrarydb=# 
```

```
mybooklibrarydb=# ALTER TABLE mybooks ADD CONSTRAINT fk_publishers_publisher_id FOREIGN KEY (publisher_id) REFERENCES publishers(id);
ALTER TABLE
mybooklibrarydb=# 
mybooklibrarydb=# \d mybooks
                                           Table "public.mybooks"
      Column      |            Type             | Collation | Nullable |               Default               
------------------+-----------------------------+-----------+----------+-------------------------------------
 id               | integer                     |           | not null | nextval('mybooks_id_seq'::regclass)
 title            | character varying(512)      |           |          | 
 publisher_id     | integer                     |           |          | 
 publication_year | integer                     |           |          | 
 description      | text                        |           |          | 
 slug             | character varying(512)      |           |          | 
 created_at       | timestamp without time zone |           |          | now()
 updated_at       | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "fk_publishers_publisher_id" FOREIGN KEY (publisher_id) REFERENCES publishers(id)
Referenced by:
    TABLE "mybooks_genres" CONSTRAINT "fk_mybooks_book_id" FOREIGN KEY (book_id) REFERENCES mybooks(id)

mybooklibrarydb=# \d publishers
                                          Table "public.publishers"
     Column     |            Type             | Collation | Nullable |                Default                 
----------------+-----------------------------+-----------+----------+----------------------------------------
 id             | integer                     |           | not null | nextval('publishers_id_seq'::regclass)
 publisher_name | character varying(512)      |           |          | 
 created_at     | timestamp without time zone |           |          | now()
 updated_at     | timestamp without time zone |           |          | now()
Indexes:
    "publishers_pkey" PRIMARY KEY, btree (id)
Referenced by:
    TABLE "mybooks" CONSTRAINT "fk_publishers_publisher_id" FOREIGN KEY (publisher_id) REFERENCES publishers(id)

mybooklibrarydb=# 
```

```
mybooklibrarydb=# \d mybooks_authors
                                        Table "public.mybooks_authors"
   Column   |            Type             | Collation | Nullable |                   Default                   
------------+-----------------------------+-----------+----------+---------------------------------------------
 id         | integer                     |           | not null | nextval('mybooks_authors_id_seq'::regclass)
 book_id    | integer                     |           |          | 
 author_id  | integer                     |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_authors_pkey" PRIMARY KEY, btree (id)

mybooklibrarydb=# 
mybooklibrarydb=# 
mybooklibrarydb=# ALTER TABLE mybooks_authors ADD CONSTRAINT fk_mybooks_book_id FOREIGN KEY (book_id) REFERENCES mybooks(id);
ALTER TABLE
mybooklibrarydb=# 
mybooklibrarydb=# ALTER TABLE mybooks_authors ADD CONSTRAINT fk_authorss_author_id FOREIGN KEY (author_id) REFERENCES authors(id);
ALTER TABLE
mybooklibrarydb=# \d mybooks_authors
                                        Table "public.mybooks_authors"
   Column   |            Type             | Collation | Nullable |                   Default                   
------------+-----------------------------+-----------+----------+---------------------------------------------
 id         | integer                     |           | not null | nextval('mybooks_authors_id_seq'::regclass)
 book_id    | integer                     |           |          | 
 author_id  | integer                     |           |          | 
 created_at | timestamp without time zone |           |          | now()
 updated_at | timestamp without time zone |           |          | now()
Indexes:
    "mybooks_authors_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "fk_authorss_author_id" FOREIGN KEY (author_id) REFERENCES authors(id)
    "fk_mybooks_book_id" FOREIGN KEY (book_id) REFERENCES mybooks(id)

mybooklibrarydb=#
```

#### Unique Constraints

```
mybooklibrarydb=# \d users
                                         Table "public.users"
   Column    |            Type             | Collation | Nullable |              Default              
-------------+-----------------------------+-----------+----------+-----------------------------------
 id          | integer                     |           | not null | nextval('users_id_seq'::regclass)
 email       | character varying(255)      |           |          | 
 first_name  | character varying(255)      |           |          | 
 last_name   | character varying(255)      |           |          | 
 password    | character varying(60)       |           |          | 
 created_at  | timestamp without time zone |           |          | now()
 updated_at  | timestamp without time zone |           |          | now()
 user_active | integer                     |           |          | 0
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
Referenced by:
    TABLE "tokens" CONSTRAINT "fk_users_tokens" FOREIGN KEY (user_id) REFERENCES users(id)

mybooklibrarydb=# 
mybooklibrarydb=# 
mybooklibrarydb=# ALTER TABLE users ADD CONSTRAINT unique_users_email UNIQUE (email);
ALTER TABLE
mybooklibrarydb=# 
mybooklibrarydb=# 
mybooklibrarydb=# \d users
                                         Table "public.users"
   Column    |            Type             | Collation | Nullable |              Default              
-------------+-----------------------------+-----------+----------+-----------------------------------
 id          | integer                     |           | not null | nextval('users_id_seq'::regclass)
 email       | character varying(255)      |           |          | 
 first_name  | character varying(255)      |           |          | 
 last_name   | character varying(255)      |           |          | 
 password    | character varying(60)       |           |          | 
 created_at  | timestamp without time zone |           |          | now()
 updated_at  | timestamp without time zone |           |          | now()
 user_active | integer                     |           |          | 0
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
    "unique_users_email" UNIQUE CONSTRAINT, btree (email)
Referenced by:
    TABLE "tokens" CONSTRAINT "fk_users_tokens" FOREIGN KEY (user_id) REFERENCES users(id)

mybooklibrarydb=# 
```