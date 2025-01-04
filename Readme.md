# CMN Express API Documentation

Welcome to the **CMN Express API Documentation**. This project provides a robust API for managing logistics, tracking, and delivery services efficiently. Below, you'll find an overview of the features, endpoints, and usage guidelines.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [API Endpoints](#api-endpoints)
- [Admin Endpoints](#admin-endpoints)
- [Customer Endpoints](#customer-endpoints)
- [Delivery Person Endpoints](#delivery-person-endpoints)
- [Definitions](#definitions)
- [License](#license)

---

## Overview
**CMN Express API** simplifies the management of logistics operations through a set of well-designed endpoints.

### Key Features:
- Manage customers and delivery persons.
- Paginated retrieval of data.
- Comprehensive endpoint descriptions and schemas.

### API Details:
- **Host**: `203.145.47.225`
- **Base Path**: `/`
- **Version**: `1.0`

---

## Installation

### Prerequisites:
- Install Go (Golang) if needed.
- Ensure your system has internet access to interact with the host server.

### Set up:
1. Clone this repository.
2. Start the API server by following project-specific instructions.
3. Access API documentation at `{host}/swagger` (if available).

---

## API Endpoints

### Admin Endpoints

#### 1. List All Customers (Paginated)
**GET** `/admin/customers/all`
- **Description**: Retrieves a paginated list of all customers.
- **Parameters**:
  - `page` (integer): Page number (default: 1).
  - `page_size` (integer): Page size (default: 10).

#### 2. List Pending Customers
**GET** `/admin/customers/pending`
- **Description**: Retrieves a paginated list of customers with status "PENDING".
- **Parameters**:
  - `page` (integer): Page number (default: 1).
  - `page_size` (integer): Page size (default: 10).

#### 3. List All Delivery Persons (Paginated)
**GET** `/admin/delivery-persons/all`
- **Description**: Retrieves a paginated list of all delivery persons.
- **Parameters**:
  - `page` (integer): Page number (default: 1).
  - `page_size` (integer): Page size (default: 10).

#### 4. List Pending Delivery Persons
**GET** `/admin/delivery-persons/pending`
- **Description**: Retrieves a paginated list of delivery persons with status "PENDING".
- **Parameters**:
  - `page` (integer): Page number (default: 1).
  - `page_size` (integer): Page size (default: 10).

### Customer Endpoints

#### 1. Register New Customer
**POST** `/customers/register`
- **Description**: Registers a new customer.
- **Request Body**:
  - See [RegisterRequest](#registerrequest).

#### 2. Get Customer Information
**GET** `/customers/{id}`
- **Description**: Retrieves details of a specific customer.
- **Path Parameters**:
  - `id` (string): Customer ID.

#### 3. Update Customer Information
**PUT** `/customers/{id}`
- **Description**: Updates details of a specific customer.
- **Path Parameters**:
  - `id` (string): Customer ID.
- **Request Body**:
  - See [UpdateCustomerReq](#updatecustomerreq).

#### 4. Delete Customer Account
**DELETE** `/customers/{id}`
- **Description**: Deletes a customer account.
- **Path Parameters**:
  - `id` (string): Customer ID.

### Delivery Person Endpoints

#### 1. Register New Delivery Person
**POST** `/delivery-persons/register`
- **Description**: Registers a new delivery person.
- **Request Body**:
  - See [RegisterRequest](#registerrequest).

#### 2. Get Delivery Person Information
**GET** `/delivery-persons/{id}`
- **Description**: Retrieves details of a specific delivery person.
- **Path Parameters**:
  - `id` (string): Delivery Person ID.

#### 3. Update Delivery Person Information
**PUT** `/delivery-persons/{id}`
- **Description**: Updates details of a specific delivery person.
- **Path Parameters**:
  - `id` (string): Delivery Person ID.
- **Request Body**:
  - See [UpdateDeliveryPersonReq](#updatedeliverypersonreq).

#### 4. Delete Delivery Person Account
**DELETE** `/delivery-persons/{id}`
- **Description**: Deletes a delivery person account.
- **Path Parameters**:
  - `id` (string): Delivery Person ID.

---

## Definitions

### RegisterRequest
- **Required Fields**:
  - `phone` (string): Phone number.
  - `password` (string): Account password.
- **Optional Fields**:
  - `account_type` (string)
  - `current_address` (string)
  - `date_of_birth` (string)
  - `full_name` (string)
  - `gender` (string)
  - `identification_number` (string)
  - `latitude` (number)
  - `longitude` (number)
  - `nationality` (string)
  - `place_of_origin` (string)
  - `place_of_residence` (string)

### UpdateCustomerReq
- **Optional Fields**:
  - Same as [RegisterRequest](#registerrequest), except `phone` and `password` are not required.

### UpdateDeliveryPersonReq
- **Optional Fields**:
  - `current_address` (string)
  - `date_of_birth` (string)
  - `full_name` (string)
  - `gender` (string)
  - `identification_number` (string)
  - `nationality` (string)
  - `place_of_origin` (string)
  - `place_of_residence` (string)

---

## License
This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

For more information, contact **Lucas** at [nguyenmanh180102@gmail.com](mailto:nguyenmanh180102@gmail.com) or visit [Sharky Tech](https://sharkytech.vercel.app).

