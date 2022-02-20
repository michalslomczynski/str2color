# str2color
Package for colorizing strings with ANSI 256 byte colors for terminal outputs (without white and black).
### Why
Can be useful for logger if you do not want to manually apply colors for many tags. Collisions might happen since range is narrow.
### Usage
```go
fmt.Println(str2color.Format("test string"))
```
$ <span style="color:MediumSeaGreen">test string</span>