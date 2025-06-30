# SMOL Virtual Machine 🚀

A stack-based virtual machine implementation in Go that executes programs written in ABM (Abstract Machine) language. This project was built for fun and educational purposes to explore virtual machine concepts and stack-based computation.

## 🎯 What is ABM?

ABM is a simple, stack-based programming language designed for this virtual machine. It features:

- Stack-based computation model
- Variable assignment and manipulation
- Arithmetic and logical operations
- Label-based control flow (GOTO/RETURN)
- Function-like procedures with labels

## 🏗️ Architecture

The VM consists of several key components:

- **Stack**: Generic stack implementation for operand management
- **Memory**: Variable storage with pointer-based memory addressing
- **Instruction Processor**: Executes ABM instructions sequentially
- **Label System**: Supports jumps and procedure calls

## 📚 Language Reference

### Basic Operations

| Instruction      | Description             | Example              |
| ---------------- | ----------------------- | -------------------- |
| `show <text>`  | Display text            | `show Hello World` |
| `push <value>` | Push integer onto stack | `push 42`          |
| `pop`          | Remove top stack item   | `pop`              |
| `print`        | Print top stack value   | `print`            |
| `halt`         | Stop program execution  | `halt`             |

### Variables

| Instruction      | Description                        | Example      |
| ---------------- | ---------------------------------- | ------------ |
| `lvalue <var>` | Create variable reference          | `lvalue x` |
| `rvalue <var>` | Push variable value to stack       | `rvalue x` |
| `:=`           | Assign top stack value to variable | `:=`       |

### Arithmetic Operations

| Operator | Description    | Usage                         |
| -------- | -------------- | ----------------------------- |
| `+`    | Addition       | Push two values, then `+`   |
| `-`    | Subtraction    | Push two values, then `-`   |
| `*`    | Multiplication | Push two values, then `*`   |
| `/`    | Division       | Push two values, then `/`   |
| `div`  | Modulo         | Push two values, then `div` |

### Logical Operations

| Operator | Description | Result                                |
| -------- | ----------- | ------------------------------------- |
| `&`    | Bitwise AND | `1` if both true, `0` otherwise   |
| `\|`    | Bitwise OR  | `1` if either true, `0` otherwise |
| `!`    | Logical NOT | `1` if false, `0` if true         |

### Comparison Operations

| Operator | Description           | Result                            |
| -------- | --------------------- | --------------------------------- |
| `<>`   | Not equal             | `1` if different, `0` if same |
| `<=`   | Less than or equal    | `1` if true, `0` if false     |
| `>=`   | Greater than or equal | `1` if true, `0` if false     |
| `<`    | Less than             | `1` if true, `0` if false     |
| `>`    | Greater than          | `1` if true, `0` if false     |

### Control Flow

| Instruction      | Description                 | Example        |
| ---------------- | --------------------------- | -------------- |
| `label <name>` | Define a label              | `label loop` |
| `goto <name>`  | Jump to label               | `goto loop`  |
| `return`       | Return to previous position | `return`     |

## 🚀 Getting Started

### Prerequisites

- Go 1.20 or higher

### Installation

1. Clone the repository:

```bash
git clone https://github.com/vmskonakanchi/smol-vm.git
cd abm-vm
```

2. Build the virtual machine:

```bash
go build -o abm-vm
```

### Running Programs

Execute an ABM program:

```bash
./smol-vm <filename.abm>
```

Example:

```bash
./smol-vm test/first.abm
```

## 📝 Example Programs

### Hello World (`first.abm`)

```abm
show This is my first program
lvalue f
push 5
:=
rvalue f
print
halt
```

### Arithmetic Operations (`operatorsTest.abm`)

```abm
push 5
push 4
show 5 - 4 = 1
-
print
pop
halt
```

### Function-like Procedures (`demo.abm`)

```abm
lvalue x
push 0
:=
goto work

label work
  lvalue xx
  rvalue x
  push 1
  +
  :=
return
```

## 🎮 Sample Programs

The `test/` directory contains several example programs:

- **`first.abm`** - Basic variable assignment and printing
- **`demo.abm`** - Parameter passing demonstration
- **`operatorsTest.abm`** - Arithmetic and logical operations showcase
- **`factProc.abm`** - Factorial calculation using loops
- **`recursiveFactorial.abm`** - Recursive factorial implementation

## 🔧 Technical Details

### Stack Implementation

- Generic stack with type safety
- Automatic memory management
- Stack overflow protection
- Items store both name and value for debugging

### Memory Model

- Variables stored as pointers for efficient memory access
- Automatic variable creation on first use
- Support for both values and memory addresses

### Instruction Processing

- Two-pass execution: label collection, then instruction execution
- Line-by-line parsing with whitespace normalization
- Error handling for invalid operations

## 🎯 Why This Project?

This project was created for fun and learning! It explores:

- Virtual machine design principles
- Stack-based computation models
- Language interpreter implementation
- Go programming techniques (generics, pointers, interfaces)

## 🤝 Contributing

This is a fun project! Feel free to:

- Add new ABM language features
- Improve error handling
- Create more example programs
- Optimize performance

## 📄 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

Inspired by classic stack-based virtual machines and educational programming language implementations.
