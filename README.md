# wdog
Windows process watch dog, make your program never die.

### Install

Use `go get github.com/songjiayang/wdog`  or download the binary [release](https://github.com/songjiayang/wdog/releases).

### Config 

A example: 

```
{
  "processes": [
     // die check, halt check, auto reload.
    {
      "name": "java.exe",
      "rcmd": "E:/example/start.bat",
      "endpoint": "http://localhost:8080",
      
      "checkInterval": 15,
      "reloadInterval": 10800
    },
     // die check, halt check.
    {
      "name": "gateway.exe",
      "rcmd": "E:/example/gateway.exe",
      "endpoint": "http://localhost:9090",
      "checkInterval": 15
    },
     // die check.
    {
      "name": "anylysis.exe",
      "rcmd": "E:/example/anylysis.exe",
      "checkInterval": 15
    }
  ]
}
```

### Features

- die check
- hangs check  (http can't respoonse)
- auto reload with interval
