package server

var errJson = struct {
	MethodNotAllowed string
}{
	MethodNotAllowed: `{"error": "StatusMethodNotAllowed", "code": 405}`,
}
