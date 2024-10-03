## Docker

To start the PostgreSQL container:

```
[mbp2022 12:35:37:89] docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
[mbp2022 12:37:59:89] docker ps -a
CONTAINER ID   IMAGE                     COMMAND                  CREATED        STATUS                      PORTS                                            NAMES
72e7a4cbdd5e   tecnotree/tomweb          "/usr/bin/java -jar …"   5 weeks ago    Exited (143) 5 weeks ago                                                     tomweb
e445f7a40490   php:7.2-apache            "docker-php-entrypoi…"   6 weeks ago    Exited (0) 6 weeks ago                                                       apache-php
7199b6a2bc03   ubuntu:22.04              "/bin/bash"              5 months ago   Exited (137) 5 months ago                                                    ubuntu-script-playground
482dfc1cd8e1   alpine-go                 "/bin/sh"                5 months ago   Exited (0) 5 months ago                                                      alpine-go
53c7af4ca3f4   tomweb-load               "/usr/bin/java -jar …"   5 months ago   Exited (143) 5 months ago                                                    tomweb-load
fb36b12452f0   python-hello              "python /src/server.…"   6 months ago   Exited (137) 6 months ago                                                    hellopython
46ffcd9605b0   nodejs-hello              "docker-entrypoint.s…"   6 months ago   Exited (137) 6 months ago                                                    hellonodejs
2160abd443dd   9c4ed427e83b              "docker-php-entrypoi…"   6 months ago   Exited (255) 4 months ago   0.0.0.0:8081->80/tcp                             hellophp
3b64712fc252   bitnami/openldap:latest   "/opt/bitnami/script…"   7 months ago   Exited (0) 5 weeks ago                                                       openldap
ef9cea777c5e   grafana/grafana           "/run.sh"                7 months ago   Exited (255) 7 months ago   0.0.0.0:3000->3000/tcp                           grafana
d55abc4b4df5   prom/prometheus           "/bin/prometheus --c…"   7 months ago   Exited (255) 7 months ago   0.0.0.0:9090->9090/tcp                           prometheus
8a222511cf59   cassandra:latest          "docker-entrypoint.s…"   7 months ago   Exited (143) 7 months ago                                                    cassandra4
eb756328a645   mongo:latest              "docker-entrypoint.s…"   8 months ago   Exited (0) 7 months ago                                                      mongodb
fa5ac54da016   jcalonso/mailhog          "/MailHog"               8 months ago   Exited (255) 5 months ago   0.0.0.0:1025->1025/tcp, 0.0.0.0:8025->8025/tcp   mailhog
cbf0f20dbfd0   mysql:8                   "docker-entrypoint.s…"   8 months ago   Exited (0) 4 weeks ago                                                       mysql8
6e23dff22940   postgres:14               "docker-entrypoint.s…"   8 months ago   Exited (0) 16 minutes ago                                                    postgres14
[mbp2022 12:38:00:89] docker start postgres14 
postgres14
[mbp2022 12:38:04:89] 
```

To login into the container and access PostgreSQL database:

```
[mbp2022 12:38:05:89] docker exec -it postgres14 bash
root@6e23dff22940:/# 
root@6e23dff22940:/# 
root@6e23dff22940:/# psql
psql: error: connection to server on socket "/var/run/postgresql/.s.PGSQL.5432" failed: FATAL:  role "root" does not exist
root@6e23dff22940:/# psql -U postgres
psql (14.10 (Debian 14.10-1.pgdg120+1))
Type "help" for help.

postgres=# 
postgres=# \l
                                  List of databases
     Name     |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   
--------------+----------+----------+------------+------------+-----------------------
 bookappdb    | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 bookstoredb  | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 postgres     | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 simplebankdb | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
 template0    | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
              |          |          |            |            | postgres=CTc/postgres
 template1    | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
              |          |          |            |            | postgres=CTc/postgres
 tokens       | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
(7 rows)

postgres=# 
```