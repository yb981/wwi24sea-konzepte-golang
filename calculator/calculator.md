# RPN Calculator

=========================================
Welcome to the RPN Calculator  
=========================================

## Overview

The RPN Calculator is a command-line tool that allows users to perform mathematical calculations using Reverse Polish Notation (RPN).

## Functionality

The RPN Calculator supports the following operations:

- **Addition (`+`)**: Adds the top two numbers on the stack.
- **Subtraction (`-`)**: Subtracts the top number from the second top number on the stack.
- **Multiplication (`*`)**: Multiplies the top two numbers on the stack.
- **Division (`/`)**: Divides the second top number by the top number on the stack.
- **Exponentiation (`^`)**: Raises the second top number to the power of the top number on the stack.
- **Square Root (`sqrt`)**: Computes the square root of the top number on the stack.
- **Logarithm (`log`)**: Calculates the logarithm (base 10) of the top number on the stack.
- **Factorial (`!`)**: Computes the factorial of the top integer on the stack.
- **Absolute Value (`abs`)**: Computes the absolute value of the top number on the stack.
- **Sum All Numbers (`++`)**: Sums all numbers currently on the stack.
- **Multiply All Numbers (`**`)\*\*: Multiplies all numbers currently on the stack.

## Available Commands

- Type `help` for assistance.
- Type `latex` for LaTeX formatted output.
- Type `exit` to close the application.

## Getting Started

### Prerequisites

To run the RPN Calculator, ensure you have Go installed on your machine. You can download Go from [golang.org](https://golang.org/dl/).

### Run

1. **Clone the repository**
2. **run in ./caclulator/**: go run .
