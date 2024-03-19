package main

import "testing"

// go test -v ./cmd/web/
// go test -cover ./cmd/web/
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out ./cmd/web/
func TestRun(t *testing.T) {
	err := run()
		if err != nil{
			t.Error("failed run()")
		
	}

}
