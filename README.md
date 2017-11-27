Lazydog
======

### Introduction

Just one thing , inject golang source code and print where is the code now !

Sometimes, l need to read some opensource golang project like tidb, jaeger, consul or other things. However, l'm too stupid to follow the code, especially when interface{} flies !

### Impl way

l use golang standard lib to parse itself (A language can parse itself, amazing!), and inject what l want to do to the ast file.This is simple actually.

### Install

```shell
make install
```

### run test

```shell
make jumptest
make overtest
```

### how to use

Inject the source code in dir recursively; never mind what the code will look like, just run the fucking opensource project and read the log.

```shell
lazydog jump -d src
```

And of course, lt can be restored

```shell
lazydog over -d src
```

### what the log contains

```text
[lazydog][{{goroutine id}}] {{pkgname}}:{{lineNo}} caller={{function}}
```