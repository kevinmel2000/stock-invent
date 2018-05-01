# STOCKINVENT
Simple REST API that demonstrate Simple Inventory Management (CRUD) using Go Programming Language and SQLite3

### Running in development
1. `dep init` (first time only)
2. Start the server `go run main.go`

### Running in Linux Environment
1. `make all`

#### TODO
- [ ] Use Client to access the API
- [ ] Import data from CSV
- [ ] Export data to CSV
- [ ] Example of API Call

#### API Endpoint
- `POST` `/items/blouse` to insert single blouse
- `GET` `/items/blouse` to get all blouses
- `GET` `/items/blouse:id` to get single blouse
- `PUT` `/items/blouse:id` to update single blouse
- `DELETE` `/items/blouse:id` to remove single blouse



