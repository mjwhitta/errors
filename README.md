# errors

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/errors?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/errors)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/mjwhitta/errors/ci.yaml?style=for-the-badge)](https://github.com/mjwhitta/errors/actions)
![License](https://img.shields.io/github/license/mjwhitta/errors?style=for-the-badge)

## What is this?

A simple errors package that dynamically prepends the package name.

## How to install

Open a terminal and run the following:

```
$ go get -u github.com/mjwhitta/errors
```

## Usage

Below is a sample usage:

```
package mypackage

import (
    "os"

    "github.com/mjwhitta/errors"
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

- [Source](https://github.com/mjwhitta/errors)
