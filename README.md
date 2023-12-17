This project is a simple HTTP (REST) API built using Go and the GoFr framework for managing a car garage. 
The API supports CRUD (Create, Read, Update, Delete) operations for car entries in the garage and integrates with a MariaDB database (https://www.teoalida.com/cardatabase/samples/United-Kingdom-Car-Database-by-Teoalida-full-specs-SAMPLE.sql) for data persistence.
The project is organized into the following directories:
api: Contains API handler functions for managing car and garage entries.
db: Manages the database connection and implements functions for interacting with the database.
model: Defines the data models used in the project (e.g., Car, GarageEntry).
test: Contains unit tests for the API handlers.
