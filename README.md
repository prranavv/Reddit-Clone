# Reddit Clone

This project has been a passion project of mine for a while and I wanted to create something I would enjoy as my first real project.
Here is the [link](https://172-235-29-203.ip.linodeusercontent.com/) to this project which has been deployed.

## Getting Started

To run this application in your local machine, have a postgres database set up and create a database.yml and add the neccesary configurations. An example of the database.yml is there in the database.yml.example file.
Also add .env file with the key "DATABASE_URL" and add your database connection string. Below is an example of a connection string

`host=localhost port=5432 dbname=reddit user=reddit password=reddit sslmode=disable`

## Prerequisites

1. Have the latest version of go installed to run.
2. Latest version of Postgres.

## Installing

First build a binary of the application.

`go build -o reddit cmd/web/*.go`

In this case, I build the binary and named it "reddit". Run this binary as an executable and it will run in your machine.
If everything works you will get the following logs in your terminal

`{
  "time": "2024-01-30T12:39:57.558014061+05:30",
  "level": "INFO",
  "msg": "Connected to Database"
}
{
  "time": "2024-01-30T12:39:57.558167688+05:30",
  "level": "INFO",
  "msg": "Server is running on port 8080"
}
`
