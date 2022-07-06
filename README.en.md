<h1 align="center">📔 note-gin</h1>

<p align="center">
<a target="_blank" href="https://github.com/zhouboyi1998/note-gin"> 
<img src="https://img.shields.io/github/stars/zhouboyi1998/note-gin?logo=github">
</a>
<a target="_blank" href="https://opensource.org/licenses/MIT"> 
<img src="https://img.shields.io/badge/license-MIT-red"> 
</a>
<img src="https://img.shields.io/badge/Go-1.18-darkturquoise">
<img src="https://img.shields.io/badge/Gin-1.8.0-dodgerblue">
<img src="https://img.shields.io/badge/MongoDB Go Driver-1.9.1-seagreen">
</p>

### 📖 Language

[简体中文](./README.md) | English

### ⌛ Start

#### Project configuration

* 1：Configure `Global GOPATH` & `Project GOPATH`
* 2：Configure `Environment`
    * `GOPROXY=https://goproxy.cn,direct`
    * `GOFLAGS=-buildvcs=false`

#### Install dependencies

* Run the command in the root path of the project：`go mod tidy`

### 🐳 Docker

* Run the command in the project root directory

#### Compile the Golang code to Linux executable file

```
set GOOS=linux

set GOARCH=amd64

go build main.go
```

#### Docker Build

```
docker build -t note-gin .
```

#### Docker Run

```
docker run -d -p 18091:18091 --name note-gin note-gin
```

### 📜 Licence

[MIT License](https://opensource.org/licenses/MIT) Copyright (c) 2022 周博义
