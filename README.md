# Currency Converter

This is simply a RESTful web service which convert currencies in golang. 

**Table of Contents**
- [About the project](#about-the-project)
- [Implementation](#implementation)
    - [Dependency](#dependency)
- [Layout](#layout)
- [Building](#building)
- [Running](#running)
- [Testing](#testing)
- [Improvement](#improvement)

## About the project

The currency converter is capable of converting between 3 currencies namely Nigerian Naira(NGN), Ghanaian Cedis(GHS)
and Kenyan Shillings (KSH). Any other currency results in a failure with appropriate response. 

## Implementation

The conversions between currencies are housed in a sqlite db and has been preloaded.

![](/db.png?raw=true "Sqlite conversion table")

To convert between currencies, you provide the source(base) and target(quote) currency codes. 
Application first santiates the input, verify the currencies and retrieves the rate from the database.

### Dependency

I tried to limit the number of dependencies for this service. Since I used an sqlite database, I added the following

* `sqlite driver` used for establish connection to the database and other functionalities
* `gorm` used as ORM tool to safely abstract database operates through sanitization of parameters

## Layout

Below is the layout of the application

```tree
├── README.md
├── cmd
│   └── main.go
├── data
│   └── test.db
├── db.png
├── go.mod
├── go.sum
├── pkg
│   ├── controllers
│   │   └── currency_controller.go
│   ├── model
│   │   ├── currency_model.go
│   │   └── setup.go
│   └── services
│       └── currency_service.go
├── run.sh
└── tests
    └── currency_test.go
```

A brief directories of the program layout

* `README.md` is a detailed description of the project.
* `cmd` contains main package for application startup.
* `data` contains the sqlite db file containing currency conversion.
* `go.mod` golang modules used.
* `go.sum` module checksum.
* `pkg` hold project business logic.
* `test` hold all application test cases.
* `run.sh` a simple bash script for running the application.

## Building

To build this project, run the command below

```bash
$ go build -o currency-converter cmd/main.go
```
This generate a executable in the root directory.
Note that you will require golang to be installed before building

## Running

After the project has been successfully built, you can run the compiled executable by running the below 

```bash
$ ./currency-converter --domain=0.0.0.0:9000
Starting HTTP server @ 0.0.0.0:9000
```
Note that this start the application on port 9000. The port can be modified in the command line argument
If the argument is absent, the application default the port to 8080

```bash
$ ./currency-converter                      
Starting HTTP server @ localhost:8080
```

Alternatively, running the run.sh script also starts the application
```bash
$ ./run.sh                                  
Starting HTTP server @ 0.0.0.0:9000
```

## Testing

You can test using command line utilities such as curl, httpie. or any other client or web browser.

```bash
$ http http://localhost:9000/convert-ccy\?source\=NGN\&target\=KSH
HTTP/1.1 200 OK
Content-Length: 11
Content-Type: text/plain; charset=utf-8
Date: Thu, 06 Jan 2022 14:52:14 GMT

rate: 0.270

$ http http://localhost:9000/convert-ccy\?source\=GHS\&target\=GHS
HTTP/1.1 200 OK
Content-Length: 11
Content-Type: text/plain; charset=utf-8
Date: Thu, 06 Jan 2022 14:51:47 GMT

rate: 1.000

$ http http://localhost:9000/convert-ccy\?source\=GHS\&target\=USD
HTTP/1.1 400 Bad Request
Content-Length: 21
Content-Type: text/plain; charset=utf-8
Date: Thu, 06 Jan 2022 14:50:57 GMT

invalid currency: USD
```

## Improvement

Ideally, if this solution is dockerized, it eliminates the need to have golang installed on host system
and provides a consistent environment for both testing and production.