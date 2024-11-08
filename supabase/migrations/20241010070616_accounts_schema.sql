
CREATE SCHEMA accounts_schema;


create table accounts_schema.permissions(
	permission_id serial PRIMARY KEY,
	permission_function varchar(200) NOT NULL UNIQUE,
	permission_name varchar(200) NOT NULL,
	permission_description varchar(200),
	permission_group varchar(200) NOT NULL
);


create table accounts_schema.roles(
	role_id serial primary key,
	role_name varchar(200) not null unique,
	role_description varchar(200),
	created_at timestamp not null default now(),
	updated_at timestamp,
	deleted_at timestamp 
);


CREATE TABLE accounts_schema.role_permissions(
	role_id int NOT NULL,
	FOREIGN KEY (role_id) REFERENCES accounts_schema.roles(role_id),
	permission_id int NOT NULL,
	FOREIGN KEY (permission_id) REFERENCES accounts_schema.permissions(permission_id),
	PRIMARY KEY (role_id, permission_id)
);
create table accounts_schema.user_types(
	user_type_id serial primary key,
	user_type_name varchar(200) not null unique
);

create table accounts_schema.users(
	user_id serial primary key,
	user_name varchar(200) not null,
	user_security_level int NOT NULL,
	user_type_id int NOT NULL,
	FOREIGN KEY (user_type_id) REFERENCES accounts_schema.user_types  (user_type_id),
	user_phone varchar(200) unique,
	user_email varchar(200) not null unique,
	user_password varchar(200),
	created_at timestamp not null default now(),
	updated_at timestamp,
	deleted_at timestamp 
);

CREATE TABLE accounts_schema.user_roles(
	user_id int NOT NULL,
	FOREIGN KEY (user_id) REFERENCES accounts_schema.users(user_id),
	role_id int NOT NULL,
	FOREIGN KEY (role_id) REFERENCES accounts_schema.roles(role_id),
	PRIMARY KEY (user_id, role_id)
);

CREATE TABLE accounts_schema.navigation_bars(
    navigation_bar_id serial PRIMARY KEY,
    menu_key varchar(200) UNIQUE NOT NULL,
    label varchar(200) NOT NULL,
    label_ar varchar(200),
    icon varchar(200),
    "route" varchar(200) UNIQUE,
    parent_id int,
    FOREIGN KEY (parent_id) REFERENCES accounts_schema.navigation_bars(navigation_bar_id),
    permission_id int
);
