# go-colony
PLF::Colony in Go (with type safety)

## Usage

```bash
# Install genny
go get github.com/cheekybits/genny
# Install go-colony
go get github.com/20zinnm/go-colony
# Generate a colony for a type
genny -in=/Path/to/go/src/github.com/20zinnm/go-colony/colony.go -out=[output_file_name] -pkg=[package] gen "Type=Position"
```