A simple library that tries to mimic the functionality of the Python csv classes
DictReader and DictWriter.

#### Example 1:

> // file.csv:

>field1,field2,field3<br>
>aaa,bbb,ccc

```
fd,_ := os.Open("/path/to/file.csv")<br>
defer fd.Close()<br><br>
reader := mapcsv.NewMapReader(fd, nil, 0)<br>
fmt.Println(reader.Fields())&nbsp;&nbsp;&nbsp;*// ["field1", "field2", "field3"]*<br>
record, _ := reader.AsMap()&nbsp;&nbsp;&nbsp;*// map[string]string {"field1": "aaa", "field2": "bbb", "field3": "ccc"}*<br><br>
```


#### Example 2:

```
fd, _ := os.OpenFile("/path/to/file.csv", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)<br>
defer fd.Close()<br><br>
var fields = []string{"field1", "field2", "field3"}<br>
w := mapcsv.NewMapWriter(fd, fields, 0)<br>
w.WriteHeader()<br><br>
var row = map[string]string{"field1": "aaa", "field2": "bbb", "field3": "ccc"}<br>
w.WriteRow(row)<br>
w.Flush()<br>
```
