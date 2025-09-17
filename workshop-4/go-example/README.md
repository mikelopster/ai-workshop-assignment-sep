# Go Profile API

Simple Go HTTP API for the profile page used by the frontend UI.

Endpoints:
- GET /profile - returns the profile JSON
- PUT /profile - updates profile fields (partial updates supported)

Run tests:

```bash
cd workshop-4/go-example
go test ./...
```


## Example

```
curl -i http://localhost:3000/profile
HTTP/1.1 200 OK
Date: Thu, 21 Aug 2025 14:49:11 GMT
Content-Type: application/json
Content-Length: 210

{"membership_level":"Gold","membership_code":"LBK001234","first_name":"สมชาย","last_name":"ใจดี","phone":"081-234-5678","email":"somchai@example.com","joined_date":"2023-06-15","points":15420}
```