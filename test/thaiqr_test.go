package test

import (
	"testing"

	thaiqr "github.com/ThaiQR/ThaiQR-Go"
)

func TestMerchantPromptpayQRGenerate(t *testing.T) {
	receiveId := "0912345678"
	amount := "100.00"
	onetime := true

	expectedResult := "00020101021229370016A0000006770101110113006691234567853037645406100.005802TH6304D803"
	actualResult := thaiqr.MerchantPromptpayQRGenerate(receiveId, amount, onetime)

	if actualResult != expectedResult {
		t.Errorf("Unexpected result. Expected: %s, got: %s", expectedResult, actualResult)
	}
}

func TestMerchantBillpaymentQRGenerate(t *testing.T) {
	billerId := "011234567891233"
	merchantName := "THAIQR"
	reference1 := "REF123456789"
	reference2 := "REF987654321"
	amount := "200.00"
	onetime := true

	expectedResult := "00020101021230710016A00000067701011201150112345678912330212REF1234567890312REF98765432153037645406200.005802TH5906THAIQR6304F528"
	actualResult := thaiqr.MerchantBillpaymentQRGenerate(billerId, merchantName, reference1, reference2, amount, onetime)

	if actualResult != expectedResult {
		t.Errorf("Unexpected result. Expected: %s, got: %s", expectedResult, actualResult)
	}
}
