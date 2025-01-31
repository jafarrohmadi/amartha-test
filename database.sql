-- Create the Loan table
CREATE TABLE IF NOT EXISTS loans (
                                     id UUID PRIMARY KEY,
                                     user_id VARCHAR(255) NOT NULL,
                                     total_amount FLOAT NOT NULL,
                                     balance_principal FLOAT NOT NULL,
                                     due_principal FLOAT NOT NULL,
                                     due_interest FLOAT NOT NULL,
                                     paid_principal FLOAT NOT NULL,
                                     paid_interest FLOAT NOT NULL,
                                     interest_rate INT NOT NULL,
                                     status VARCHAR(50) NOT NULL,
                                     number_of_installment INT NOT NULL,
                                     start_date TIMESTAMP NOT NULL,
                                     end_date TIMESTAMP NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the LoanTransaction table
CREATE TABLE IF NOT EXISTS loan_transactions (
                                                 id UUID PRIMARY KEY,
                                                 loan_id UUID REFERENCES loans(id) ON DELETE CASCADE,
                                                 amount_paid FLOAT NOT NULL,
                                                 payment_date TIMESTAMP NOT NULL,
                                                 type VARCHAR(50) NOT NULL,
                                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the LoanSchedule table
CREATE TABLE IF NOT EXISTS loan_schedules (
                                              id UUID PRIMARY KEY,
                                              loan_id UUID REFERENCES loans(id) ON DELETE CASCADE,
                                              start_date TIMESTAMP NOT NULL,
                                              due_date TIMESTAMP NOT NULL,
                                              amount FLOAT NOT NULL,
                                              status VARCHAR(50) NOT NULL,
                                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for faster lookups
CREATE INDEX idx_loan_transactions_loan_id ON loan_transactions (loan_id);
CREATE INDEX idx_loan_schedules_loan_id ON loan_schedules (loan_id);
