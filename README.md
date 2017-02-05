A simple library that tries to mimic the functionality of the Python csv classes
DictReader and DictWriter.

#### Example 1:

With csv file:

>field1,field2,field3<br>
>aaa,bbb,ccc

import "os"<br>
import "github.com/wayeast/mapcsv"<br><br>
fd, _ := os.Open("/path/to/file.csv")<br>
defer fd.Close()<br><br>
reader := mapcsv.NewMapReader(fd, nil, ',')<br>
fmt.Println(reader.Fields())&nbsp;&nbsp;&nbsp;*// ["field1", "field2", "field3"]*<br>
record, err := reader.Read()&nbsp;&nbsp;&nbsp;*// map[string]string {"field1": "aaa", "field2": "bbb", "field3": "ccc"}*<br><br>


#### Example 2:

import "os"<br>
import "github.com/wayeast/mapcsv"<br><br>
fd, _ := os.OpenFile("/path/to/file.csv", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)<br>
var fields = []string{"field1", "field2", "field3"}<br>
w := mapcsv.NewMapWriter(fd, fields, ',')<br>
w.WriteHeader()<br><br>
var record = map[string]string{"field1": "aaa", "field2": "bbb", "field3": "ccc"}<br>
w.Write(record)<br>
w.Flush()<br>
