# Restify

Convert any JSON file into a REST API supporting CRUD operations.

Currently loads whole file into memory
Goals:

- Provide GET, POST, PUT, DELETE routes for JSON file
- Index based on optional key
- Optional authentication endpoint
  - Use redis or bolt for caching tokens
- Provide stats endpoint
