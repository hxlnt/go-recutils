# go-recutils
Go wrapper for interacting with [recutils](https://www.gnu.org/software/recutils/manual/recutils.html) with some niceties for the modern world like providing structured return values and errors.

## Usage

See the [example folder](example) to run sample code that tests all currently-implemented functions.

### csv2rec, mdb2rec, rec2csv
TODO; mdb2rec may not be implemented.

### recdel

```go
func (recs RecordSet) Del(params SelectionParams, options OptionFlags, removeOrComment DeleteStyle) RecordSet
func (file Recfile) Del(params SelectionParams, options OptionFlags, removeOrComment DeleteStyle) Recfile
```

### recfix

```go
func (recs RecordSet) Fix(action FixAction, options OptionsFlags) RecordSet
func (file Recfile) Fix(action FixAction, options OptionsFlags) Recfile
```

### recfmt

```go
func (recs RecordSet) Fmt(template string, isTemplateFilename bool) (string, error)
```

### recinf

```go
func (file Recfile).Inf() ([]Recinfo, error)
```

### recins

```go
func (recs RecordSet) Ins(newRecords RecordSet, params SelectionParams, options OptionsFlags) RecordSet
func (file Recfile) Ins(newRecords RecordSet, params SelectionParams, options OptionsFlags) Recfile
```

### recsel

```go
func (file Recfile) sel(params SelectionParams, options OptionsFlags) RecordSet
```

### recset

```go
func (file Recfile) set(params SelectionParams, options OptionsFlags) Recfile
```