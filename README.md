# Reddit Clone

This project has been a passion project of mine for a while and I wanted to create something I would enjoy as my first real project.
Here is the [link](https://172-235-29-203.ip.linodeusercontent.com/) to this project which has been deployed.

## Tech Stack Used

For the backend, I used **golang** and just the standard library and a router package called **"chi-router"**. It's a highly performant language and easy to learn, hence my language of choice for the backend.
<br>

For the frontend,(which I am particulary not skilled in nor am I fond of) I used basic **HTML**, **CSS**, and very little javascript. The framework (I don't know if it can be called as such) I used is **HTMX**. It's a relatively new framework and if you haven't tried it out yet, I highly recommend it if you are tired of writing React code.

For the database,I used **Postgres** for storing all the user data.

## Getting Started

To run this application in your local machine, have a postgres database set up and create a database.yml and add the neccesary configurations. An example of the database.yml is there in the database.yml.example file.
<br>
Also run the following code to add the tables and indices to the database.
<br>
`soda migrate up`
<br>

Also add .env file with the key "DATABASE_URL" and add your database connection string. Below is an example of a connection string

`host=localhost port=5432 dbname=reddit user=reddit password=reddit sslmode=disable`



## Prerequisites

1. Have the latest version of go installed.
2. Latest version of Postgres.

## Installing

First build a binary of the application.

`go build -o reddit cmd/web/*.go`

In this case, I built the binary and named it "reddit". Run this binary as an executable and it will run in your machine.
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
