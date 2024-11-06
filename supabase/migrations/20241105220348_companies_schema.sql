

CREATE SCHEMA companies_schema;
create table companies_schema.companies(
	company_id serial PRIMARY KEY,
	company_name varchar(200) NOT NULL UNIQUE,
	company_name_ar varchar(200) ,
	company_phone varchar(200) UNIQUE,
	company_address TEXT,
	company_address_ar TEXT,
	company_description varchar(200),
	company_description_ar varchar(200),
	company_email varchar(200) UNIQUE,
	company_logo TEXT ,
	company_logo_vertical TEXT,
	company_logo_dark TEXT,
	company_logo_dark_vertical TEXT,
	created_at timestamp not null default now(),
	updated_at timestamp,
	deleted_at timestamp 
);

