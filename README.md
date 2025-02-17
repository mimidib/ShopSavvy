# ShopSavvy
This project is a Go application that scrapes product prices from specified websites and sends email notifications if the price is lower than a specified threshold.

## Getting Started

To run this application, you need to have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

Create a .env file in the project directory with the following content:
```
EMAIL_FROM=your_email@gmail.com
EMAIL_PASSWORD=your_app_specific_password
EMAIL_TO=recipient_email@example.com
```
Install the required Go packages:  
```
go get -u github.com/gocolly/colly/v2
go get github.com/joho/godotenv
```


### Running the Application

1. Clone the repository or download the project files.
2. Navigate to the project directory:
   ```sh
   cd my-go-project
   ```
3. Run the application using the following command:
   ```
   go run cmd/main.go
   ```

### Purpose

This project serves as a basic introduction to Go programming, showcasing how to structure a Go application with separate packages and an entry point. The application simply prints "Hello, World!" to the console.

