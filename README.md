# go-recuitls
Go wrapper for interacting with [recutils](https://www.gnu.org/software/recutils/manual/recutils.html) with some niceties for the modern world like providing structured return values.

## Usage

See the [example folder](example) to run sample code that tests all currently-implemented functions.

### csv2rec

TODO 

### mdb2rec
TODO; may not be implemented.

### rec2csv

TODO

### recdel

TODO

### recfix

**Recfix(filename string, operation RecfixOperation, useExternalDesc bool, force bool)** runs the specified `recfix` operation and returns an error if warranted. `RecfixOperation` may have the value `Auto`, `Check`, or `Sort`. Decryption/encryption operations are not currently supported.

### recfmt

TODO

### recinf 

**Recinf(filename string)** returns an array of objects for each record definition (`%rec`) in the provided file and an error if warranted. Objects follow the structure shown below.
```go
type RecinfResponse struct {
	Record      string
	Count       int
	Rectypedefs []Rectype 
	Doc         []string
	Rectypes    []Rectype
	Key         []string
	Mandatory   []string
	Singular    []string
	Allowed     []string
	Prohibited  []string
	Unique      []string
	Auto        []string
	Sort        []string
	Comments    []string
}

type Rectype struct {
	Name  string
	Value string
	Enum  []string `json:",omitempty"`
}
```
As this object returns both descriptors and counts, there is no need to use `recinf` flags `-d` and `-n`.

### recins

TODO

### recsel

TODO

### recset