#Create a users table
```
create table users (
  id serial primary key,
  email text not null unique,
  password text not null
);
```

#Create a books table

```
create table books (
  id serial primary key,
  user_id integer,
  title text,
  author text,
  year text
);
```

#create user book relational table

```
select Books.id, Users.email, Books.Author, 
Books.Year, Books.Title FROM Books 
INNER JOIN Users ON Books.id=Users.id;
```