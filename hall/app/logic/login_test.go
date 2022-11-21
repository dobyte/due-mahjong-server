package logic

import "testing"

func TestLogin_doGoogleLogin(t *testing.T) {
	logic := NewLogin(nil)

	_, err := logic.doGoogleLogin("dfjslfjsdlfjlsdf", "")
	if err != nil {
		t.Fatal(err)
	}
}
