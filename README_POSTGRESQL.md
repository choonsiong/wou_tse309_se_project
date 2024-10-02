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