package googlemrf

import (
	"testing"
)

// Makes sure the client is compiled.
func TestClient(t *testing.T) {
	_ = testClient{}
	_ = mediaClient{}
}
