# Gola/Gobar
This repo is part of one of my interview assignment.
###### Start date: 17/07/2020
###### End date: 18/07/2020

This is just a minimalistic sample application to demonstrate my coding practice in Golang and MongoDB.
This application is a subset of the needs of a cab booking platform like Gobar/Gola.

- [x] A user should be able to request a booking for a cab from pickup location A to pickup location B
- [x] A user should be able to view his past bookings
- [x] A user should be able to get cabs that are nearby.
- [x] Writing middleware for proper http error statuses
- [x] Using gin

P.S. an actual cab booking platform has way more detail than you could think of.
That usually includes of plenty of other necesary microservices, a messaging queue, socket connections, VRP engine and ofcourse some statistics to generate the price estimate.


## How to run
### Step 1
###### Initialize your mongo database with commands in the file below

```sh
dbinit.js
```

### Step 2
###### On your terminal
```sh
git clone github.com/Glitchfix/golagobar
cd golagobar
go get ./...
```

### Update your config file
```json
{
  "database": {
    "database": "golagobar",
    "host": "DBHOST",
    "port": "DBHOST",
    "username": "YOUR_USERNAME",
    "password": "YOUR_PASSWORD",
    "authsource": "YOUR_AUTHENTICATION_SOURCE_DB"
  },
  "server": {
    "host": "HOST_IP",
    "port": "PORT"
  }
}

```

### Click on the postman icon to see API documentation
[![N|Click on me](https://lh4.googleusercontent.com/Dfqo9J42K7-xRvHW3GVpTU7YCa_zpy3kEDSIlKjpd2RAvVlNfZe5pn8Swaa4TgCWNTuOJOAfwWY=s128-h128-e365)](https://documenter.getpostman.com/view/12089646/T1DjjzRK)