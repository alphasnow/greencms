// @author AlphaSnow

package utils

import "testing"

func TestPassword(t *testing.T) {
	pass, err := PasswordHash("admins")
	// guests
	// $2a$10$sNfxGrG4PXikrFn/Eb2YXOj9.lhCo7pTQZXhHrP/AMA7.Va0QxLra
	// admins
	// $2a$10$BBA/jGFggvv0qy7vtnFxfOa7Q.LAzMLSvYRfdZsktY/JNj3mdCdEO
	pl := len(pass)
	// 60
	t.Log(pass, err, pl)

	ok := PasswordVerify("admin", pass)
	t.Log(ok)

	t.Log("ok")
}
