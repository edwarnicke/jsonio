json io provides two things:

**jsonio.Reader()** - returns an io.Reader with the same methods as json.Encoder, which read from what is written to the Writer.

**jsonio.Writer()** - returns an io.Writer with the same methods as json.Decode, output of which is read from the Reader.

Example:
```go
    r := jsonio.Reader()
    w := jsonio.Writer()
    r.Encode(&in)
    io.Copy(w, r)
    w.Decode(&out)
```
