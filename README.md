## Clone the Repo

### Run go.mod tidy in Terminal

### Run command main.go

### Open Postman & Choose for the GRPC Req

### Make request along with proper fields

```
Create Request For User
{
  "name": " srk",
  "salary": 50000,
  "department": "Engineering",
  "address": {
    "city": "New York",
    "state": "NY"
  }
}
```

```CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY, -- Use a VARCHAR to store UUIDs
    name VARCHAR(255) NOT NULL,
    salary BIGINT NOT NULL,
    department VARCHAR(255) NOT NULL,
    address_id VARCHAR(36) -- Assuming addresses also use UUIDs
);
```
