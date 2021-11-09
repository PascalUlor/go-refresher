# GO Lang code challenge

**ABOUT** This is a very basic REST API implementation with Golang. It's bee over a year since wrote Golang code thus I really enjoyed working on this project. 

# Table Of Content

1. [Getting Started](#getting-started 'Getting Started')
   .._[prerequisites](#prerequisites 'Prerequisites')
   .._[Installation](#installation 'Installation')
2. [Technology Stack](#technology-stack)
3. [Built With](#built-with 'Built With')
4. [Useful Links](#useful-links 'Useful Links')
5. [Improvements](#improvements 'Improvements')

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

To get started with this project you need a basic knowledge of :

```
Golang
Version Control (Git)
```

### Installation

The following instructions will install the project on your local machine

1. Clone the [**repository here**](https://github.com/NuriCareers/Pascal-Ulor-coding-challenge)
2. [**cd**] into the root directory of the project folder `pascal-ulor-coding-challenge`.
3. Run `go mod tidy` on the terminal to install Dependencies/Modules

### How to run

1. use the `.env.example` file to create your environment variables.
2. Now open terminal to build then run the application

```
go build -o bin/main
```

```
./bin/main
```

If all goes well, you should see something similar to this on the console:

```
Application started on http://localhost:<port>
```

## Integration tests

To test the endpoints, create a test database and run the following:

```
go test
```

## Technology Stack

### Built With

- [GO](https://golang.org/) - Golang

### API Architecture

The API is built for easy use and understanding. It includes the following:

1. The GET endpoint to get the details of a block.
2. The GET TO GET details of a transaction.
## API END POINTS AND FUNCTIONALITY

| EndPoint                                                       | Functionality                                                                        |
| -------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| POST `/block/:network/:ref`                               | Gets details of a _block_ in a blockchain                                |
| POST `/trx/network/:ref` | Gets details of a _transaction_ in a block using _specified transaction reference_ |

Now you can test the endpoints with a client e.g POSTMAN

## Useful Links

1. [Dependency Management With Go Modules](https://dev.to/pc_codes/dependency-management-with-go-modules-434c)
2. [Github Repo](https://github.com/NuriCareers/Pascal-Ulor-coding-challenge)

## Improvements
1. Use go channels to improve application runtime.
2. Improve test coverage
## Authors

- **Ulor Pascal** - [PascalUlor](https://github.com/PascalUlor)