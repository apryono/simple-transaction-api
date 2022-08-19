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

CREATE TABLE transactions (
	id serial PRIMARY KEY,
	transaction_number TEXT CHECK (char_length(transaction_number) <= 255),
	customer_id int4 NOT NULL REFERENCES customers(id),
	pic_name TEXT NOT NULL CHECK (char_length(pic_name) <= 255),
	pic_phone TEXT NULL CHECK (char_length(pic_phone) <= 255),
	pic_email TEXT NULL CHECK (char_length(pic_email) <= 255),
	total_price numeric(20,3) DEFAULT 0,
	type_of_payment TEXT NULL CHECK (char_length(type_of_payment) <= 50),
	note TEXT,
	status text,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER transactions BEFORE UPDATE ON transactions FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE transaction_details (
	id serial PRIMARY KEY,
	transaction_id int4 NOT NULL REFERENCES transactions(id),
	product_id int4 NOT NULL REFERENCES products(id),
	product_name TEXT NOT NULL CHECK (char_length(product_name) <= 255),
	product_price numeric(20,3) DEFAULT 0,
	product_quantity int4 DEFAULT 0,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP WITH TIME ZONE 
);
CREATE TRIGGER transaction_details BEFORE UPDATE ON transaction_details FOR EACH ROW EXECUTE PROCEDURE update_modified_column();