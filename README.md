# Blockchain-Based Library Management System

A decentralized library management system built with Go, implementing blockchain technology to maintain an immutable record of book checkouts.

## Project Overview

This project demonstrates the implementation of blockchain concepts in a practical application using Go. The system tracks library book checkouts using a chain of blocks, where each block contains transaction data about book checkouts.

## System Architecture

### Core Components
```mermaid
graph TD
    A[Book Management System] --> B[Blockchain]
    B --> C[Block Structure]
    C --> D[Contains BookCheckout Data]
    
    E[Book] --> |Has Properties| F[ID]
    E --> |Has Properties| G[Title]
    E --> |Has Properties| H[Author]
    E --> |Has Properties| I[PublishDate]
    E --> |Has Properties| J[ISBN]
    
    K[BookCheckout] --> |Records| L[BookID]
    K --> |Records| M[User]

```

### Transaction flow
```mermaid
sequenceDiagram
    participant User
    participant API
    participant Blockchain
    
    User->>API: Request Book Checkout
    API->>Blockchain: Create New Block
    Blockchain->>Blockchain: Add Block with Checkout Data
    Blockchain-->>API: Confirm Addition
    API-->>User: Checkout Confirmation

```
## Technical Stack
- Language: Go 1.x
- Web Framework: Gorilla Mux
- API: RESTful endpoints
- Data Structure: Blockchain

## Features
- Blockchain-based transaction recording
- Book management (add, query)
- Checkout tracking
- Immutable transaction history
- RESTful API endpoints

## Getting Started
- Prerequisites
- Go 1.x
- Git

## Installation
### Clone the repository
git clone https://github.com/kedarvartak/Go-BLock

## Navigate to project directory
cd golang-blockchain

## Install dependencies
go mod download

## Run the application
go run main.go

## Contributing
- Fork the repository
- Create your feature branch (git checkout -b feature/AmazingFeature)
- Commit your changes (git commit -m 'Add some AmazingFeature')
- Push to the branch (git push origin feature/AmazingFeature)
- Open a Pull Request

## Learning Objectives
- Understanding blockchain fundamentals
- Go programming practices
- RESTful API development
- Distributed systems concepts
