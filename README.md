# Tiktak - A Video Sharing Platform

Tiktak is a video sharing platform, inspired by the popular social media app TikTok. It allows users to upload, view, comment on, and like videos. This project is designed to showcase a full-stack application using Go with the Fiber framework, along with other modern technologies for a robust and scalable service.

## Features

- User registration and authentication
- Video uploading and streaming
- Commenting on videos
- Liking videos

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (Version 1.21)
- Database setup (e.g., MySQL, PostgreSQL)
- Any other environment-specific setups (like external APIs if used)

### Installing

A step-by-step series of examples that tell you how to get a development environment running.

```bash
# Clone the repository
git clone https://github.com/BryceWayne/tiktak.git

# Navigate to the project directory
cd tiktak

# Install dependencies (if any)
go mod tidy

# Run the application
go run main.go
```

## API Overview

### User Endpoints

- `POST /register`: Register a new user.
- `POST /login`: Authenticate a user.

### Video Endpoints

- `POST /upload`: Upload a new video.
- `GET /videos`: Retrieve all videos.

### Comment Endpoints

- `POST /comment`: Post a comment on a video.
- `GET /comments/:videoId`: Retrieve all comments for a specific video.

### Like Endpoints

- `POST /like`: Like a video.
- `GET /likes/:videoId`: Get likes for a specific video.

_Note: Detailed API documentation with Swagger is currently in development._

## Built With

- [Go](https://golang.org/) - The Go Programming Language
- [Fiber](https://gofiber.io/) - An Express-inspired web framework for Go
- [MySQL/PostgreSQL/etc.] - Database
- [Docker](https://www.docker.com/) - Containerization (if used)
- [Any other major libraries or frameworks used]

## Contributing

Please read [CONTRIBUTING.md](link-to-your-contributing.md-file) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/BryceWayne/tiktak/tags).

## Authors

- **Bryce Chudomelka** - *Initial work* - [BryceWayne](https://github.com/BryceWayne)

See also the list of [contributors](https://github.com/BryceWayne/tiktak/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

- Hat tip to anyone whose code was used
- Inspiration
- etc.
