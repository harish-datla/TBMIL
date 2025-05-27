# Part1
# Storage Engines
## Prologue
  - DBMS’s core role is to reliably store data and making it available for users.
  - Using databases lets developers focus on application logic instead of inventing a new way to organzie data every time we create a new app.
  - Databases are modular and composed of the following parts:
    - Transport layer (handles incoming requests)
    - Query processor (chooses most efficient way to run queries)
    - Execution engine (performs operations)
    - Storage engine (manages persistent data in memory and on disk)

The storage engine exposes a simple CRUD API for creating, reading, updating, and deleting records

DBMSes build on storage engines by providing schemas, query languages, indexing, transactions, and other high-level features







