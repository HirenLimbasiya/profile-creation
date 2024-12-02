# Profile Creation API

This project provides a simple API for creating user profiles. It validates user input for the PAN number, mobile number, and email using the Fiber framework and the validator.v10 package.

## Table of Contents
- Project Setup
- API Endpoints
- Validation
- Testing
- Folder Structure

## Project Setup

### Installation
1. Clone the repository:
```
https://github.com/HirenLimbasiya/profile-creation.git
cd profile-creation
```
2. Install dependencies:
```
go mod tidy
```
3. Add .env file [Optional]
```
PORT=3000
```
4. Run the application:
```
go run main.go
```

## API Endpoints
### POST /user
Creates a new user profile.
Request Body:
```
{
    "name": "John Doe",
    "pan": "ABCDE1234F",
    "mobile": 9876543210,
    "email": "test@example.com"
}
```
Example Request:
```
curl -X POST http://localhost:3000/user -d '{"name": "John Doe", "pan": "ABCDE1234F", "mobile": 9876543210, "email": "test@example.com"}' -H "Content-Type: application/json"
```
Response:
```
{
    "status_code": 200,
    "message": "User created successfully",
    "data": {
        "name": "John Doe",
        "pan": "ABCDE1234F",
        "mobile": 9876543210,
        "email": "test@example.com"
    }
}

```
## Validation
- PAN: Required, exactly 10 characters (5 letters, 4 digits, 1 letter at the end).
- Mobile: Required, exactly 10 digits.
- Email: Required, valid email format.

## Testing
```
go test -count=1 ./test -v   
```
## Author

This project is created and maintained by **Hiren Limbasiya**.
You can explore more of my work on my [Portfolio](https://www.hirenlimbasiya.com/).