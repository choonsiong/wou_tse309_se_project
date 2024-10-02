# About

To keep track of dependencies used in the application.

## Dependencies

### PostgreSQL

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

### Backend

```
[mbp2022 2:43:37:131] go get -u github.com/go-chi/chi/v5
go: downloading github.com/go-chi/chi/v5 v5.1.0
go: added github.com/go-chi/chi/v5 v5.1.0
[mbp2022 2:44:17:131] 

[mbp2022 5:06:44:131] go get github.com/go-chi/cors
go: added github.com/go-chi/cors v1.2.1
[mbp2022 5:06:46:131]

[mbp2022 20:06:46:91] go get github.com/jackc/pgconn
go: downloading github.com/jackc/pgconn v1.14.3
go: downloading github.com/jackc/pgproto3/v2 v2.3.3
go: added github.com/jackc/chunkreader/v2 v2.0.1
go: added github.com/jackc/pgconn v1.14.3
go: added github.com/jackc/pgio v1.0.0
go: added github.com/jackc/pgpassfile v1.0.0
go: added github.com/jackc/pgproto3/v2 v2.3.3
go: added github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a
go: added golang.org/x/crypto v0.20.0
go: added golang.org/x/text v0.14.0
[mbp2022 20:06:59:91] 

[mbp2022 20:07:01:91] go get github.com/jackc/pgx/v4
go: downloading github.com/jackc/pgx/v4 v4.18.3
go: added github.com/jackc/pgtype v1.14.0
go: added github.com/jackc/pgx/v4 v4.18.3
[mbp2022 20:07:11:91] 

[mbp2022 20:07:13:91] go get github.com/jackc/pgx/v4/stdlib
[mbp2022 20:07:24:91]  
```

### Frontend

- Install [Vue Router](https://router.vuejs.org/installation.html) if you didn't select it while setting up the project.

```
npm install vue-router@4
```