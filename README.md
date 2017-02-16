A simple library that tries to mimic the functionality of the Python csv classes
DictReader and DictWriter.

#### Example 1:

> // file.csv:<br>
>field1,field2,field3<br>
>aaa,bbb,ccc

```go
fd, _ := os.Open("/path/to/file.csv")
defer fd.Close()
reader := mapcsv.NewMapReader(fd, nil, 0)
fmt.Println(reader.Fields())   // ["field1", "field2", "field3"]
record, _ := reader.AsMap()   // map[string]string {"field1": "aaa", "field2": "bbb", "field3": "ccc"}
```


#### Example 2:

```go
fd, _ := os.OpenFile("/path/to/file.csv", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
defer fd.Close()
var fields = []string{"field1", "field2", "field3"}
w := mapcsv.NewMapWriter(fd, fields, 0)
w.WriteHeader()
var row = map[string]string{"field1": "aaa", "field2": "bbb", "field3": "ccc"}
w.WriteRow(row)
w.Flush()
```
