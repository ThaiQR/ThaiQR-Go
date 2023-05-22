package thaiqr

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/howeyc/crc16"
)

func MerchantPromptpayQRGenerate(recieveId, amount string, onetime bool) string {
	PFI := preprocessValue("00", "01")

	var pim_val string
	if onetime {
		pim_val = "12"
	} else {
		pim_val = "11"
	}
	PIM := preprocessValue("01", pim_val)

	/** Merchant Identifier */
	AID := preprocessValue("00", "A000000677010111")
	recieveId = preprocessRecieveID(recieveId)

	merchantSum := AID + recieveId
	merchantIdentifier := preprocessValue("29", merchantSum)
	/* */

	Currency := preprocessValue("53", "764")

	amount = preprocessAmount(amount)

	CountryCode := preprocessValue("58", "TH")

	crc := "6304"

	data := PFI + PIM + merchantIdentifier + Currency + amount + CountryCode + crc
	dataBuffer := []byte(data)
	crcResult := crc16.ChecksumCCITTFalse(dataBuffer)

	data += strings.ToUpper(strconv.FormatInt(int64(crcResult), 16))

	return data
}

func MerchantBillpaymentQRGenerate(billerId, merchantName, reference1, reference2, amount string, onetime bool) string {
	reference1 = strings.ToUpper(reference1)
	reference2 = strings.ToUpper(reference2)

	PFI := preprocessValue("00", "01")

	var pim_val string
	if onetime {
		pim_val = "12"
	} else {
		pim_val = "11"
	}
	PIM := preprocessValue("01", pim_val)

	/** Merchant Identifier */
	AID := preprocessValue("00", "A000000677010112")
	billerId = preprocessValue("01", billerId)
	reference1 = preprocessValue("02", reference1)
	reference2 = preprocessValue("03", reference2)

	merchantSum := AID + billerId + reference1 + reference2
	merchantIdentifier := preprocessValue("30", merchantSum)
	/* */

	Currency := preprocessValue("53", "764")

	amount = preprocessAmount(amount)

	CountryCode := preprocessValue("58", "TH")

	merchantName = preprocessValue("59", merchantName)

	TerminalID := preprocessValue("07", "")

	crc := "6304"

	data := PFI + PIM + merchantIdentifier + Currency + amount + CountryCode + merchantName + TerminalID + crc
	dataBuffer := []byte(data)
	crcResult := crc16.ChecksumCCITTFalse(dataBuffer)

	data += strings.ToUpper(strconv.FormatInt(int64(crcResult), 16))

	return data
}

func preprocessValue(prefix, value string) string {
	if len(value) != 0 {
		if len(value) < 10 {
			value = prefix + "0" + strconv.Itoa(len(value)) + value
		} else {
			value = prefix + strconv.Itoa(len(value)) + value
		}
	} else {
		value = ""
	}
	return value
}

func preprocessAmount(value string) string {
	checkAmount := strings.Split(value, ".")

	if len(checkAmount) > 1 {
		if checkAmount[1] == "" || len(checkAmount[1]) == 0 {
			checkAmount[1] = "00"
		} else if len(checkAmount[1]) == 1 {
			checkAmount[1] += "0"
		} else if len(checkAmount[1]) > 2 {
			checkAmount[1] = checkAmount[1][:2]
		}

		value = checkAmount[0] + "." + checkAmount[1]
	} else if len(checkAmount) == 1 {
		value = value + "." + "00"
	}

	value = preprocessValue("54", value)

	return value
}

func preprocessRecieveID(value string) string {
	if len(value) == 10 && value[0] == '0' { // Phone Number
		value = "0066" + trimFirstRune(value)
		value = preprocessValue("01", value)
	} else if len(value) == 13 { // National ID or Tax ID
		value = preprocessValue("02", value)
	} else if len(value) == 15 { // E-Wallet ID
		value = preprocessValue("03", value)
	} else { // Bank Account
		value = preprocessValue("04", value)
	}

	return value
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
