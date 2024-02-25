-- PostgreSQL database init script

-- Create the Customer table
CREATE TABLE IF NOT EXISTS customers (
    customer_id SERIAL PRIMARY KEY,
    customer_limit INT NOT NULL,
    customer_balance INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the Transaction table
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id SERIAL PRIMARY KEY,
    transaction_amount INT NOT NULL,
    transaction_type VARCHAR(255) NOT NULL,
    transaction_description VARCHAR(255) NOT NULL,
    transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    transaction_customer INT NOT NULL,
    FOREIGN KEY (transaction_customer) REFERENCES customers(customer_id)
);

-- Create the Indexes
CREATE INDEX IF NOT EXISTS customer_id_index ON customers(customer_id);
CREATE INDEX IF NOT EXISTS transaction_id_index ON transactions(transaction_id);
CREATE INDEX IF NOT EXISTS transaction_customer_index ON transactions(transaction_customer);

