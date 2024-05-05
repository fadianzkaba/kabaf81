type Reader struct {
	Read([]byte) (int, error)
} 

type File struct { ... }
func (f File) Read(b []byte) (n int, e error)

type TCPConn struct { ... }
func (t TCPConn) Read(b []byte) (n int, e error) 

var f File
var t TCPConn

var r Reader
r = f
r.Read(...)

r = t
r.Reader(...)


