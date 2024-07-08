# Blog Project Showcase

## Purpose

The purpose of this project is to demonstrate the efficiency and speed of a human developer when leveraging tools like ChatGPT, combined with previous developer experience and online documentation. Through this showcase, we aim to illustrate how AI-powered tools can enhance development workflows and accelerate project delivery.

## Features

- **Blog Website**: Provides a platform where users can read blog entries in chronological order.
- **Administrator Access**: Password-protected access for administrators to create, edit, and delete blog posts.

## Technologies Used

- **Go (Golang)**: Backend programming language chosen for its performance and concurrency features.
- **PostgreSQL**: Database management system used to store blog posts and user data.
- **Docker Compose**: Containerization tool to manage the PostgreSQL database alongside the Go application.
- **Gorilla Mux**: Go package for handling HTTP requests and routing.
- **ChatGPT**: AI language model used for generating project ideas, refining code snippets, and providing insights into best practices.

## Project Structure

The project is structured as follows:

blog/
├── main.go
├── handlers/
│ ├── blog_handler.go
│ └── admin_handler.go
├── models/
│ └── post.go
└── db/
└── postgres.go

- **`main.go`**: Entry point of the Go application.
- **`handlers/`**: Contains HTTP handlers for blog and admin functionalities.
- **`models/`**: Defines data models like `Post`.
- **`db/`**: Includes database connection setup (`postgres.go`).

## Setup Instructions

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/blog.git
   cd blog

   ```

2. **Start PostgreSQL using Docker Compose:**
   `docker-compose up`

3. **Run the Go Application:**
   `go run main.go`

4. **Access the Application:**
   Open a web browser and go to http://localhost:8080 to view the blog website.

## Usage

- Regular users can view blog posts sorted by date.
- Administrators can log in to create new posts, edit existing ones, and delete posts using the password-protected admin interface.

## Contributors

- Developer: Robert Jones
- AI Assistant: ChatGPT by OpenAI
