# Console-based Task List

Suggested learning project generated with the [Code Learning Assistant GPT](https://chat.openai.com/g/g-ji16Atlkg-code-learning-assistant).

## General Requirements

### Environment Setup
- Install Go (ensure you have the latest version).
- Set up your Go workspace and familiarize yourself with basic Go commands (`go run`, `go build`).

### Application Structure
- Organize your application using packages to separate concerns (e.g., data storage, task management, and CLI interface).

## Feature Requirements

### Task Management
- The application must support the following basic operations:
  1. **Add a Task**: Users can add a new task with a description.
  2. **List Tasks**: Users can view all tasks, displaying each task's ID, description, and status (Pending/Completed).
  3. **Complete a Task**: Users can mark a task as completed by specifying the task's ID.
  4. **Delete a Task**: Users can delete a task by specifying the task's ID.

### Data Storage
- Implement a simple file-based storage system to persist tasks between application runs.
- Tasks should be stored in a format that allows easy reading and writing (JSON, CSV, etc.).

### CLI Interface
- Use Go's `flag` or `cobra` package to parse command-line arguments.
- The application should support different commands corresponding to its operations (add, list, complete, delete).
- Provide a help command or option to display usage instructions.

## Technical Requirements

### Error Handling
- Properly handle potential errors, especially for file operations and input parsing.
- Display user-friendly error messages for common issues (e.g., task not found, invalid input).

### Code Quality
- Write clean, readable, and well-structured code.
- Follow Go's coding conventions and best practices.

### Documentation
- Include a README file with:
  - Instructions on how to build and run your application.
  - A brief description of the application and its features.
  - Any necessary setup steps (e.g., if a specific file directory is required for the data storage).

## Stretch Goals (Optional)

### Enhanced Task Details
- Allow users to specify due dates for tasks and include these in the task listings.
- Support for task priorities.

### Interactive Mode
- Implement an interactive mode that keeps the application running and allows users to execute multiple commands without restarting the app.

### Unit Tests
- Write unit tests for your application's core functionality to ensure reliability and facilitate future development.
