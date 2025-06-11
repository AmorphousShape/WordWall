# WordWall

[![Go Report Card](https://goreportcard.com/badge/github.com/AmorphousShape/wordwall/pkg/wordwall)](https://goreportcard.com/report/github.com/AmorphousShape/wordwall/pkg/wordwall)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

**WordWall** is a Go package and CLI tool for detecting and filtering zero-tolerance or banned language in strings. It includes support for obfuscated text (e.g., leetspeak, Greek letters, and other special characters).

---

## Features

- Define banned words with customizable filtering rules
- Detect obfuscated words (e.g., `h@t3`, `sp@m`, Greek or Unicode lookalikes)
- Regex-based matching with case insensitivity and spacing tolerance
- Support for different rule types:
  - Censor: Replace offending words
  - Filter: Flag content for review
  - Ban: Trigger zero-tolerance logic

---

## Installation

### Go package:

```bash
go get github.com/AmorphousShape/wordwall/pkg/wordwall
```

---

## Usage

### As a Go package:

```go
package main

import (
    "fmt"
    "github.com/AmorphousShape/wordwall/pkg/wordwall"
)

func main() {
    // Set banned words and the rule to apply
    wordwall.SetBannedWords([]string{"hate", "spam"}, wordwall.RuleCensor)

    input := "I h@t3 sp@m"
    filtered, hitCensor, hitFilter, hitBan := wordwall.FilterString(input)

    fmt.Println("Filtered:", filtered)
    fmt.Println("Hit censor?", hitCensor)
    fmt.Println("Hit filter?", hitFilter)
    fmt.Println("Hit ban?", hitBan)
}
```

#### Output:
```
Filtered: I **** ****
Hit censor? true
Hit filter? false
Hit ban? false
```

---

## API

#### `SetBannedWords(words []string, response Rule)`

Sets the banned words and the default filtering rule to apply when they’re detected.

#### `FilterString(message string) (newMessage string, hitCensor bool, hitFilter bool, hitBan bool)`

Filters a message and returns the cleaned version, along with booleans indicating which rule(s) were triggered.
