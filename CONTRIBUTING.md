# Contributing to Akash Chain SDK

Thank you for your interest in contributing to the Akash Chain SDK! This document provides guidelines and instructions for contributing to this repository.

## Code of Conduct

Please note that this project is released with a Contributor Code of Conduct. By participating in this project you agree to abide by its terms.

## Getting Started

### Prerequisites

- **direnv**: For environment management

Direnv will prompt for any other missing dependency.

### Development Environment Setup

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/chain-sdk.git
   cd chain-sdk
   ```

3. **Set up direnv** (required for development):
   ```bash
   # Install direnv (macOS example)
   brew install direnv
   # Install direnv (Ubuntu-based)
   apt install direnv
   
   # Hook direnv to your shell
   # For bash: echo 'eval "$(direnv hook bash)"' >> ~/.bashrc
   # For zsh: echo 'eval "$(direnv hook zsh)"' >> ~/.zshrc
   
   # Allow direnv in the project directory
   direnv allow
   ```

4. **Set up vendor dependencies**:
   ```bash
   make modvendor
   ```

## Development Workflow

### Building and Testing

#### Protobuf Code Generation
If you modify any `.proto` files, you need to regenerate the generated code:

```bash
# Generate code for all languages
make proto-gen

# Generate code for specific language
make proto-gen-go    # For Go
make proto-gen-ts    # For TypeScript
```

## Code Style and Guidelines

### Go Code

- Follow standard Go conventions
- Use `gofmt` for formatting
- Write comprehensive tests
- Use interfaces appropriately
- Document exported functions and types

### TypeScript Code

- Follow TypeScript best practices
- Use ESLint for code quality
- Write unit tests with Jest
- Use proper typing and interfaces

### Protobuf Definitions

- Follow the [Protocol Buffers Style Guide](https://developers.google.com/protocol-buffers/docs/style)
- Update documentation when modifying `.proto` files
- Regenerate all language bindings after changes

### Commit Messages

Use conventional commit format:
```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Codestyle changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Example:
```
feat(node): add new query endpoint for deployment status

- Add new gRPC endpoint
- Update client libraries
- Add comprehensive tests

closes/fixes/refs #123
```

## Release and Versioning

### Version Tags

This repository uses semantic versioning with module-specific tagging. When tagging Go modules, use the `repo-tools` utility instead of creating tags manually:

```bash
# Command structure:
repo-tools bump-go <major/minor/patch/prerel> <module>
```


```bash
# Main Go module pre release
repo-tools bump-go prerel go

# CLI module pre-release
repo-tools bump-go prerel go/cli

# SDL module pre-release
repo-tools bump-go prerel go/sdl
```

The `repo-tools` utility is set up by direnv and available in your `PATH` after running `direnv allow`.

### Release Process

1. **Version bump**: Use `repo-tools` to create new version tags
2. **CI/CD**: Automated release pipelines will build and publish packages
3. **Documentation**: Update changelog and documentation as needed

## Issue Reporting

### Where to Report Issues

All issues must be reported via the [Akash Network Support Repository](https://github.com/akash-network/support/issues) with the `repo/chain-sdk` tag.

## Getting Help

- **Documentation**: Check the [README.md](./README.md) for project overview
- **Discussions**: Use GitHub Discussions for questions and ideas
- **Community**: Join the Akash Network community channels

---

Thank you for contributing to the Akash Network ecosystem!
