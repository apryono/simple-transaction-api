CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
END;
$$ language 'plpgsql';

create table brands (
	id serial PRIMARY KEY,
	name TEXT NOT NULL CHECK (char_length(name) <= 255),
	description TEXT CHECK (char_length(description) <= 5000),
	made_in TEXT CHECK (char_length(made_in) <= 100),
    status bool NOT NULL default true,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER brands BEFORE UPDATE ON brands FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

create table products (
	id serial PRIMARY KEY,
	brand_id int4 NOT NULL REFERENCES brands(id),
	name TEXT NOT NULL CHECK (char_length(name) <= 255),
	overview_description TEXT CHECK (char_length(overview_description) <= 5000),
    price NUMERIC(20,3) NOT NULL,
    sku TEXT CHECK (char_length(sku) <= 255),
    status bool NOT NULL default true,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER products BEFORE UPDATE ON products FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

create table customers (
	id serial PRIMARY KEY,
	name TEXT NOT NULL CHECK (char_length(name) <= 255),
    username TEXT CHECK (char_length(username) <= 255),
	password TEXT CHECK (char_length(password) <= 255),
	status TEXT NOT NULL CHECK (char_length(status) <= 20),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER users BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();