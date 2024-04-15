# GOLANG BACKEND

Golang backend app with fiber framework

## 1. Setup

### 1.1 Clone the repository

```bash
git clone https://github.com/gotired/golang_backend.git
cd golang_backend
go mod tidy
```

### 1.2 Create `.env` file

.env variable

- `DB_USERNAME` : database username
- `DB_PASSWORD` : database password
- `DB_HOST` : database host server
- `DB_PORT` : database port
- `DB_NAME` : database name
- `DB_SSLMODE` : database ssl mode
- `APP_PORT` : app port
- `JWT_SECRET_KEY` : jwt secret key

### 1.3 Start app

#### 1.3.1 Start app by go

```bash
go run cmd/main.go
```

#### 1.3.1 Start app by docker

```bash
docker-compose up --build
```

Certainly! Here's how you can format the tables in a README file using Markdown:

## 2. Provide APIs

### 2.1 Roles group

#### 2.1.1 List roles

- `Method` : `GET`
- `Endpoint` : /roles

#### 2.1.1 Register roles

- `Method` : `POST`
- `Endpoint` : /roles/register
- `Header` :
  - `Content-Type` : application/json
- `Body` :
  | Key | Type | Detail |
  | ------ | ------ | ------ |
  | role | String | role name |

### 2.2 Users group

#### 2.2.1 List users

- `Method` : `GET`
- `Endpoint` : /users

#### 2.2.2 Register user

- `Method` : `POST`
- `Endpoint` : /users/register
- `Header` :
  - `Content-Type` : application/json
- `Body` :
  | Key | Type | Detail |
  | ------ | ------ | ------ |
  | first_name | String | account first name |
  | last_name | String | account last name |
  | email | String | account email |
  | password | String | account password |
  | confirm | String | account confirm password |
  | num | String | account phone number |
  | user_name | String | account user name |
  | role | UUID | role uuid|

#### 2.2.3 Login

- `Method` : `POST`
- `Endpoint` : /users/login
- `Header` :
  - `Content-Type` : application/x-www-form-urlencoded
- `Body` :
  | Key | Type | Detail |
  | ------ | ------ | -------------------- |
  | identifier | String | account user naeme or email |
  | password | String | account password |
