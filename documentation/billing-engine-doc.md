---
label: Billing Engine
title: Billing Engine
geometry: "left=3cm,right=3cm"
order: 99
---

# Billing Engine

## Overview

# Version History

| Version | Description           | PIC           | Date       |
|---------|-----------------------|---------------|------------|
| 1.0     | Initial Documentation | Jafar Rohmadi | 2025-01-31 |

## Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    loans {
        uuid id PK
        string user_id
        float total_amount
        float balance_principal
        float due_principal
        float due_interest
        float paid_principal
        float paid_interest
        int interest_rate
        string status
        int number_of_installment
        timestamp start_date
        timestamp end_date
        timestamp created_at
        timestamp updated_at
    }

    loan_transactions {
        uuid id PK
        uuid loan_id FK
        float amount_paid
        timestamp payment_date
        string type
        timestamp created_at
        timestamp updated_at
    }

    loan_schedules {
        uuid id PK
        uuid loan_id FK
        timestamp start_date
        timestamp due_date
        float amount
        string status
        timestamp created_at
        timestamp updated_at
    }

    loans ||--o| loan_transactions : has
    loans ||--o| loan_schedules : has

```

## Sequence Diagrams

### Simple Billing

```mermaid
sequenceDiagram
    participant Client
    participant Product
    participant BillingEngine
    participant LoanRepository
    participant LoanScheduleRepository
    participant LoanTransactionRepository
    participant Database

    Client ->> Product: Request To Create Loan
    Product ->> BillingEngine: Create Loan
    BillingEngine ->> LoanRepository: Create Loan Record
    LoanRepository -->> BillingEngine: Loan Created
    BillingEngine ->> LoanScheduleRepository: Generate Loan Schedule
    LoanScheduleRepository -->> BillingEngine: Loan Schedule Generated
    BillingEngine ->> Database: Save Loan Details
    BillingEngine ->> Database: Save Loan Schedule
    Database -->> BillingEngine: Loan and Schedule Saved

    loop Weekly Payment Cycle
        Client ->> Product: Make Payment
        Product ->> BillingEngine: Make Payment
        BillingEngine ->> LoanRepository: Get Loan by ID
        LoanRepository -->> BillingEngine: Loan Found
        BillingEngine ->> LoanScheduleRepository: Get Loan Schedules
        LoanScheduleRepository -->> BillingEngine: Loan Schedules Found
        alt Payment Validation
            BillingEngine ->> LoanTransactionRepository: Record Transaction
            LoanTransactionRepository -->> BillingEngine: Transaction Recorded
            BillingEngine ->> LoanScheduleRepository: Update Loan Schedule
            LoanScheduleRepository -->> BillingEngine: Loan Schedule Updated
            BillingEngine ->> Database: Save Transaction and Schedule
            Database -->> BillingEngine: Transaction and Schedule Saved
        else Payment Rejected
            BillingEngine -->> Product  : Payment Invalid
            Product -->> Client: Payment Invalid
        end
    end
```
