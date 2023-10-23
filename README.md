# Concurrent Processing System

The Concurrent Processing System is a Go application designed to demonstrate the handling of multiple tasks in parallel using goroutines, channels, and the sync package. This system creates a pool of workers that execute tasks concurrently. Each task is simulated by a sleep operation, mimicking the behavior of time-consuming tasks such as I/O operations, computations, or web requests.

## Features

Concurrent task processing with a fixed number of workers.
Panic recovery in workers to handle task failures without crashing the entire system.
Dynamic task generation with random execution time.
Separate channels for successful task results and errors, ensuring clear and distinct handling of each.
Synchronization to wait for all tasks to complete before exiting the application.

## Prerequisites
Go 1.20 or later.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installation

Clone the repository to your Go workspace:
```shell
git clone https://github.com/gfregalado/concurrent-processing-system.git
```

Navigate to the project directory:

```shell
cd concurrent-processing-system
```

Build the application:
```shell
make build
```

Or Build and Run the executable generated from the build process:
Build the application:
```shell
make run
```
This will start the application, creating workers and tasks based on predefined constants. The output will display the results of the task executions and any errors encountered during the process.

### Configuration
The system's behavior can be adjusted by changing the constants in the main.go file:

**numWorkers:** The number of concurrent workers.
**numTasks:** The number of tasks to be generated and processed.

### Testing
To run the automated tests for this system, use the following command:
```shell
make test
```
This will run tests in all subdirectories of the project.