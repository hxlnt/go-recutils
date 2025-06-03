# go-recutils
Go wrapper for interacting with [recutils](https://www.gnu.org/software/recutils/manual/recutils.html) with some niceties for the modern world like providing structured return values and errors.

## Usage

See the [example folder](example) to run sample code that tests all currently-implemented functions.

### csv2rec, mdb2rec, rec2csv
TODO; mdb2rec may not be implemented.

### recinf

```go
recinfo, err = recfile.Inf()
```

### recsel

```go
// type SelectionParams struct {
// 	Type       string
// 	Expression string
// 	Quick      string
// 	Number     []int
// 	Random     int
// 	Join       string
// }

params = rec.SelectionParams{Type: "books", Expression: "Title='Junkyard Jam Band'"}
options = rec.DefaultOptions
recordSet = recfile.Sel(params, options)
recordSeterror = recordSet.Error
```

### recfix

```go
// type OptionFlags struct {
//	Force           bool
//	NoExternal      bool
//	NoAuto          bool
//	CaseInsensitive bool
//	Unique          bool
// }

options = rec.OptionFlags{CaseInsensitive: true}
newRecordSet = recordSet.Fix(rec.Check, rec.DefaultOptions)
newRecfile = recfile.Fix(rec.Check, rec.DefaultOptions)
recfileError = newRecfile.Error
```

### recdel

```go
newRecordSet = recordSet.Del(params, rec.DefaultOptions, rec.Remove)
newRecfile = recfile.Del(params, options, rec.Comment) 
```

### recins

```go
anotherRecordSet = recordSet.Ins(newRecordSet, params, options)
newRecfile = recfile.Ins(newRecordSet, params, options)
```

### recset

```go
fields = []string{"Status"}
action = rec.FieldAction{ActionType: "SetAdd", ActionValue: "Read"}
newRecfile := recFile.Set(fields, action, params, options)
```

### recfmt

```go
templateOutput, err = recordSet.Fmt("{{Title}}: {{Subtitle}}", false)
template2Output, err = recordSet.Fmt("template.txt", true)
```
