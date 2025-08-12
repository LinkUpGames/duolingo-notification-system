# Notifcation System

This is a small scale architecture design for Duolingo's notification system. Here is a link to the [paper](https://research.duolingo.com/papers/yancey.kdd20.pdf). This project is divided into 4 main components: **server**, **batch job**, **database**, **client**.

## Getting Started

You can test the application by running `docker compose up -d` and wait for all containers to initialize. Once the containers are up and running you can enter the **client** service and run `npm run app` to test the application from the client's perspective. Additionally, you can go into the sub directories of the project for more information on how the component can be run separately from the rest.

### Environment File

Before testing the project, make sure you have the following saved to a `.env` file. This is where you can update parameters to tweak the score algorithm and ports.

```bash
# .env
# Database
export POSTGRES_DB=database
export POSTGRES_USER=dev
export POSTGRES_PASSWORD=0000
export POSTGRES_PORT=8000

# Server Vars
export SERVER_PORT=3000
export SERVER_NAME=server
export SCORER_NAME=scorer

# Algorithm

# Defaults
export DEFAULT_SCORE="0.5"

# History Recency Penalty
export PENALTY="0.02"
export FACTOR=10
export CUTOFF=10
export EXPLORE="1.2"
```

### Server

The server is a simple REST http server with various routes for updating and fetching database values.

### Scorer

A simple batch script that runs every 5 minutes updating the notifications scores.

### Client

A simple client that fetches and shows a chosen notification by the system.

### Database

A PostgresSQL server for database purposes.
