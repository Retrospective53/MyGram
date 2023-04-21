# MyGram Golang Backend

This is the backend for MyGram, an application that allows users to save photos and make comments on other people's photos. It is built using Golang and utilizes various libraries and technologies for authentication, validation, and documentation.

## Technologies Used

- Golang
- Postgres database with [GORM ORM](https://gorm.io/)
- [Gin Gonic](https://github.com/gin-gonic/gin) web framework
- [Go Validator](https://github.com/go-playground/validator) for field validations
- [JsonWebToken](https://github.com/dgrijalva/jwt-go) for authentication
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) for password hashing
- [Go Swagger](https://github.com/swaggo/swag) for API documentation using Go Swagger

## Installation

1. Clone the repository: `git clone <repository-url>`
2. Install the dependencies: `go mod tidy`
3. Create a Postgres database and update the database configuration in the `config` folder.
4. Migrates the database
5. Start the server: `go run main.go`

## Endpoints

### User Routes

- `POST("/users")`: Register a new user.
- `POST("/users/login")`: Login a user and obtain a JsonWebToken for authentication.
- `GET("/users")`: Get user data using the JsonWebToken obtained from login for authentication.

### Photo Routes

- `GET("/photos/all")`: Get all photos.
- `GET("/photos/:photoID")`: Get a photo by ID.
- `POST("/photos")`: Upload a new photo.
- `PUT("/photos/:photoID")`: Update a photo by ID.
- `DELETE("/photos/:photoID")`: Delete a photo by ID.

### Comment Routes

- `GET("/comments/all")`: Get all comments.
- `GET("/comments/:commentID")`: Get a comment by ID.
- `POST("/comments")`: Upload a new comment.
- `PUT("/comments/:commentID")`: Update a comment by ID.
- `DELETE("/comments/:commentID")`: Delete a comment by ID.

### Social Media Routes

- `GET("/socialmedias/all")`: Get all social media data.
- `GET("/socialmedias/:socialmediaID")`: Get social media data by ID.
- `POST("/socialmedias")`: Upload social media data.
- `PUT("/socialmedias/:socialmediaID")`: Update social media data by ID.
- `DELETE("/socialmedias/:socialmediatID")`: Delete social media data by ID.

## Authentication and Authorization

- Endpoints that access data in the SocialMedia, Photo, and Comment tables require authentication and must include a JsonWebToken obtained from the login endpoint.
- Endpoints that modify proprietary data such as updates or deletes require authorization to ensure that only authorized users can perform these actions.

## Documentation

This API is documented using Swagger and can be accessed at `<base-url>/swagger/index.html` after starting the server.
Base url by default is "/api/v1".
