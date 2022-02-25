

CREATE TABLE schools (
  id bigserial NOT NULL PRIMARY KEY ,
  school_name varchar(255) NOT NULL,
  street_address varchar(255) NOT NULL,
  city varchar(255) NOT NULL,
  providence varchar(255) NOT NULL,
  zip varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  fax varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  website varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);