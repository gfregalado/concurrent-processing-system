# Overview

In this exercise, you are expected to build a simple system that executes multiple tasks concurrently and efficiently using Golang's concurrency primitives. The system will receive several tasks that simulate work by sleeping for a random duration and then return the result of a computation. Your goal is to implement a worker pool that manages these tasks concurrently while limiting the number of tasks running in parallel.

## Detailed Requirements

### 1. Task Generation

Each task is a function that sleeps for a random duration between 1 and 5 seconds to simulate doing work.
After "work" is done, each task returns a unique result. This could be a simple message, a computed number, etc.
Generate a list of N tasks (N can be a large number, like 100 or more) that the system will execute. Each task must have unique characteristics (e.g., different sleep time or result).

### 2. Worker Pool

Implement a worker pool using goroutines. The number of workers should be configurable (e.g., a variable at the beginning of the main function or passed as an argument).
Each worker is responsible for executing a task from a queue of tasks.
The system must limit the number of concurrent tasks to the number of workers in the worker pool.
Workers should receive tasks via a channel.

### 3. Results Collection

The main function should wait for all tasks to complete and then terminate.
As tasks complete, the system collects their results.
After all tasks are complete, the program should output the results, along with some information about each task's execution (e.g., execution time).

### 4. Error Handling

If a task encounters an error, it should be sent to a separate error channel. The main function will print out the error.
The program should handle potential panics in tasks using recover, to ensure that the entire system doesn't crash due to a failure in one task.
### 5. Testing and Documentation

Include unit tests demonstrating the functionality of your concurrency model and error handling.
Provide a README file that includes:
Instructions on how to run your application.
An overview of your approach and design choices.
Challenges you faced and how you addressed them.
Possible improvements if given more time.

## Evaluation Criteria

Correct implementation of concurrency: proper use of goroutines, channels, and synchronization mechanisms.
Efficiency: the program should not use more goroutines than the configured number of workers.
Robustness: effective error handling and ability to recover from panics.
Code quality: code should be well-structured, properly formatted, and clearly commented.
Testing: presence of meaningful unit tests.
Documentation: clarity of explanations and instructions in the README.