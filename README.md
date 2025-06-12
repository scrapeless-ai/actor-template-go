# Actor Template (Go)

## Project Description
This is a Go-based Actor pattern template project for quickly bootstrapping service development that supports the Actor model. It includes basic configurations, Docker support, and input/output specifications.

## Prerequisites
- Go 1.20+ ([Installation Guide](https://go.dev/doc/install))

## Installation & Running
### Local Execution
1. Clone the repository:
```bash
git clone https://github.com/scrapeless-ai/actor-template-go.git
cd actor-template-go
```
2. Install dependencies:
```bash
go mod tidy
```
3. Start the service (using example input):
```bash
go run main.go
```

## Configuration Explanation
- `.env.example`: Environment variable template (rename to `.env` and modify as needed)
- `.actor/actor.json`: Actor metadata configuration (name, version, etc.)
- `.actor/input_schema.json`: Input data validation schema

## Contribution & Feedback
Welcome to submit Issues or Pull Requests to contribute to improvements!

## License
This project is licensed under the [MIT License](LICENSE).
        