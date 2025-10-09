# karina v0.1.2a

## Overview

karina is a competitive programming test generator/judge, built with wails and svelte.

The project is under active development, expect bugs and missing features.

## Features (so far)

- Run many participants at the same time.
- Generate tests with a simple python script.

## Supported runtimes

- Cpython 3.13.7
- More coming soon.

## Installation

1. Download the release archive (`karina_0.1.2a_windows_amd64.zip`)
2. Unzip it to a folder
3. Run `karina.exe`

## Building and Developing

You don't want to do this.

Prerequisites:

- Go 1.23
- Wails CLI
- NodeJS and pnpm
- upx for compressing the binary (optional)

Steps:

```sh
git clone https://github.com/shellawa/karina.git
cd karina
cd frontend && pnpm install
cd ..
wails dev -s / wails build (-upx)
```
