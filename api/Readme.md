Go Projects API

A simple Go-based API that manages user authentication and coin balances using a mock database.

🚀 Features

User authentication with username and Authorization token

Fetch user coin balance

Uses middleware for authorization

⚙️ Installation

1️⃣ Clone the Repository

git clone https://github.com/Bisruxa/go_projects.git
cd go_projects

2️⃣ Install Dependencies

go mod tidy

3️⃣ Run the API

go run cmd/api/main.go

📌 API Endpoints

🔐 Authorization Middleware

The API requires a valid Authorization token to access endpoints.

🔹 Get User Coin Balance

Request:

GET /account/coins/?username={username}

Headers:

Authorization: 123ABC

Example:

curl -H "Authorization: 123ABC" "http://localhost:8000/account/coins/?username=alex"

Response:

{
  "balance": 10,
  "code": 200
}

🏗️ Project Structure

/go_projects
│── /cmd
│   └── api
│       └── main.go  # API entry point
│── /api
│   ├── api.go  # API handlers
│   ├── internal
│   │   ├── middleware  # Auth Middleware
│   │   ├── handlers  # Request handlers
│   │   └── tools  # Database interface & mock data
└── go.mod  # Go dependencies

🛠️ Tech Stack

Golang for backend

Logrus for logging

Gorilla/schema for request parsing


