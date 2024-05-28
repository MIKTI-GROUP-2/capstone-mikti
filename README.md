# E-Ticketing Application

## Group 2

## Features

- Authentication (Login and Register)
- Filtering Event
- Wishlist Event
- Booking Ticket
- Payment Ticket using Midtrans
- Storage using Cloudinary

## 3rd Party App

1. Midtrans for Payment Gateway
2. Cloudinary for File and Image Storage
3. Database Using PostgreSQL

## Another Link

[Sprint Project]()
[API Documentation POSTMAN]()
[Database Design](https://miro.com/welcomeonboard/UEJyUDhVOHJlODBXRnBDTzV2eEVkS0ZKeHFQVWF2MnYwcUp4cERGUG0xQzZmdjFzS212QUhQakxtcWRvYXdQc3wzNDU4NzY0NTE4NjE5NzUxMTM5fDI=?share_link_id=489248718140)

## REST API Design

Request
GET {{BASE_URL}}/api/v1/event/category/1

Response

```json
{
  "data":[
    {
    "id":1,
    "category_name":"Music",
    "status":true
    }
  ],
  "message":"Success Get Data"
}
```

## Project Structure

Clean Architecture

## How To Use

- ```git clone https://github.com/irvanhau/capstone-mikti.git```
- ```cp .env.example .env```
- create the database first
- change .env with your config
- ```go mod tidy```
- ```go run main.go wire_gen.go```
- 