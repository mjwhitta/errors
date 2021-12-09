# errors

<a href="https://www.buymeacoffee.com/mjwhitta">🍪 Buy me a cookie</a>

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/mjwhitta/errors)](https://goreportcard.com/report/gitlab.com/mjwhitta/errors)

## What is this?

A simple errors package that dynamically prepends the package name.

## How to install

Open a terminal and run the following:

```
$ go get --ldflags="-s -w" --trimpath -u gitlab.com/mjwhitta/errors
```

## Usage

Below is a sample usage:

```
package mypackage

import (
    "os"

    "gitlab.com/mjwhitta/errors"
)

func OpenFileByName(name string) (*os.File, error) {
    var e error
    var f *os.File

    if name == "" {
        return nil, errors.New("no name was provided")
    }

    if f, e = os.Open(name); e != nil {
        return nil, errors.Newf("failed to open %s: %w", name, e)
    }

    return f, nil
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/errors)
