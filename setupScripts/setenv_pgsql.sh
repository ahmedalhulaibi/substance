#!bin/bash
psql -c "CREATE USER travis_test WITH PASSWORD 'password';" -U postgres
psql -c "CREATE TABLE Persons (ID int PRIMARY KEY,LastName varchar(255),FirstName varchar(255),Address varchar(255),City varchar(255));CREATE TABLE Orders (OrderID int UNIQUE NOT NULL,OrderNumber int NOT NULL,PersonID int DEFAULT NULL,PRIMARY KEY (OrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));CREATE TABLE AntiOrders (AntiOrderID int UNIQUE NOT NULL,AntiOrderNumber int NOT NULL,PersonID int UNIQUE DEFAULT NULL,PRIMARY KEY (AntiOrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));" -U postgres
psql -c "GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO travis_test;" -U postgres
