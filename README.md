# goenvlib

A Go library for managing environment variables with type safety and reloading capabilities for testing purpose mostly

## Features

- Type-safe environment variable access
- Support for various data types:
  - String
  - Integer
  - Float64
  - Boolean
  - Slices of all above types
- Automatic environment variable reloading
- Thread-safe operations
- Default value support
- Error handling with fallback to default values

## Installation

```bash
go get github.com/bennyscetbun/goenvlib
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/bennyscetbun/goenvlib"
)

func main() {
    // Get a string environment variable
    port := goenvlib.GetenvString("PORT", "8080")
    fmt.Println("Port:", *port)

    // Get an integer environment variable
    timeout := goenvlib.GetenvInt("TIMEOUT", 30)
    fmt.Println("Timeout:", *timeout)

    // Get a boolean environment variable
    debug := goenvlib.GetenvBool("DEBUG", false)
    fmt.Println("Debug mode:", *debug)

    // Get a slice of integers
    ports := goenvlib.GetenvIntSlice("PORTS", []int{80, 443})
    fmt.Println("Ports:", *ports)

    // Reload all environment variables
    // mostly usable for
    goenvlib.ReloadEnv()
}
```

## Environment Variable Format

- For single values, use standard environment variable format
- For slices, use comma-separated values:
  - Example: `PORTS=80,443,8080`
  - Example: `DEBUG_LEVELS=true,false,true`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 