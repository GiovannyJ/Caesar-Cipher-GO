package caesarciphergo
import (
	"testing"
)

func Test_Enc(t *testing.T){
	org := "GIO"
	key := 1
	sol := "HJP"

	if encrypted := Enc(org, key); encrypted != sol{
		t.Error("TEST FAILED input:", org, ", expected: ", sol, "received: ", encrypted)
	}
}

func Test_Dec(t *testing.T){
	org := "HJP"
	key := 1
	sol := "GIO"

	if decrypted := Dec(org, key); decrypted != sol{
		t.Error("TEST FAILED input:", org, ", expected: ", sol, "received: ", decrypted)
	}
}