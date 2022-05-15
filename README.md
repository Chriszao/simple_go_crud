<div align="center">
  <img style="height: 200px; border-radius: 10px;" src="/assets/golang.gif" />
</div>

This is a simple crud rest api that aims to apply [GO](https://go.dev/learn/) language fundamentals.



## What is inside?

This project uses lot of stuff as:

- [GO](https://go.dev/learn/)
- [Gorilla/mux](https://github.com/gorilla/mux)
- [MySQL](https://dev.mysql.com/doc/mysql-getting-started/en/)
- [Docker](https://docs.docker.com/compose/gettingstarted/)


## Getting Started

To run this api, you may need to install [Docker desktop](https://www.docker.com/products/docker-desktop/)

```bash
# MySQL container
docker run --name go-mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:latest

# Enter MySQL container
docker exec -it go-mysql bash

# access mysql with root user 
mysql -u root -p
```

```sql
-- Create the api database
CREATE DATABASE go_course;

-- Create user table
CREATE TABLE IF NOT EXISTS user (
  id int auto_increment primary key,
  name varchar(50) NOT NULL,
  email varchar(50) NOT NULL
);

-- Create user to not access the database with root user
CREATE USER 'golang'@'%' IDENTIFIED WITH authentication_plugin BY 'golang';

-- Grant all privileges to created user
GRANT ALL PRIVILEGES ON go_course.* TO 'golang'@'%';
```

And then run the development server:

```bash
go run main.go
# or
./crud # Builded API 
```

Open [http://localhost:5000](http://localhost:5000) with your browser (Or http client) to see the result.


### [Insomnia collection](/go_api_insomnia_collection.json)


<div align="center">
  <p>Developed with 💙 by <a href="https://github.com/Chriszao">Chriszao</a></p>
</div>