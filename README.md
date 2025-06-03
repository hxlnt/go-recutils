# go-recutils
Go wrapper for interacting with [recutils](https://www.gnu.org/software/recutils/manual/recutils.html) with some niceties for the modern world like providing structured return values and errors.

## Usage

See the [example folder](example) to run sample code that tests all currently-implemented functions.

### csv2rec, mdb2rec, rec2csv
TODO; mdb2rec may not be implemented.

### recdel

```go
func (recs RecordSet) del(removeOrComment DeleteStyle, options OptionFlags) RecordSet
func (file Recfile) del(removeOrComment DeleteStyle, options OptionFlags) Recfile
```

### recfix

```go
func (recs RecordSet) fix(action FixAction, options OptionsFlags) RecordSet
func (file Recfile) fix(action FixAction, options OptionsFlags) Recfile
```

### recfmt

```go
func (recs RecordSet) fmt(template string, isTemplateFilename bool) ([]string, error)
```

### recfix

```go
func (file Recfile).inf() Recinfo
```

### recins

```go
func (recs RecordSet) ins(newRecords RecordSet, options OptionsFlags) RecordSet
func (file Recfile) ins(newRecords RecordSet, options OptionsFlags) Recfile
```

### recsel

```go
func (file Recfile) sel(params SelectionParams, options OptionsFlags) RecordSet
```

### recset

```go
func (file Recfile) set(params SelectionParams, options OptionsFlags) Recfile
```