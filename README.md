# Billing Engine

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.21
2. [GNU Make](https://www.gnu.org/software/make/)
3. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

    Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ````

4. [Docker](https://docs.docker.com/get-docker/) version 20

5. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29

## Running

You should be able to run using the script `run.sh`:

```bash
./run.sh
```

You should be able to access the API at http://localhost:8080

## Document Erd and Sequence

[ERD And Sequence Diagram](./documentation/billing-engine-doc.md)

## What Next

1. Penalty for Missed Payments
2. Grace Period Handling
3. Create Loan Product 
4. Add Log
5. Use Partition to loan transaction
6. Update Logic to Use Event driven approach
7. Update Logic Balance Principal to EOD Balance Principal
8. Add Unit Test
9. Add repayment_allocation_order in loans