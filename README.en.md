<h1 align="center">📔 note-gin</h1>

<p align="center">
<a target="_blank" href="https://gitee.com/zhouboyi/note-requests">
<img src="https://img.shields.io/badge/license-MIT-red"> 
<img src="https://img.shields.io/badge/Go-1.18-darkturquoise">
<img src="https://img.shields.io/badge/Gin-1.8.0-dodgerblue">
<img src="https://img.shields.io/badge/MongoDB Go Driver-1.9.1-seagreen">
</a>
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
