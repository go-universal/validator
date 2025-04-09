package funcs_test

import (
	"testing"

	"github.com/go-universal/validator/funcs"
	"github.com/stretchr/testify/assert"
)

func TestIsValidUsername(t *testing.T) {
	assert.True(t, funcs.IsValidUsername("user_123"))
	assert.False(t, funcs.IsValidUsername("invalid-user!"))
}

func TestIsAlphaNumeric(t *testing.T) {
	assert.True(t, funcs.IsAlphaNumeric("abc123"))
	assert.True(t, funcs.IsAlphaNumeric("abc123@", "@"))
	assert.False(t, funcs.IsAlphaNumeric("abc123@", "#"))
}

func TestIsAlphaNumericWithPersian(t *testing.T) {
	assert.True(t, funcs.IsAlphaNumericWithPersian("سلام123"))
	assert.True(t, funcs.IsAlphaNumericWithPersian("سلام123@", "@"))
	assert.False(t, funcs.IsAlphaNumericWithPersian("سلام123#", "@"))
}

func TestIsValidIranianPhone(t *testing.T) {
	assert.True(t, funcs.IsValidIranianPhone("02123456789"))
	assert.False(t, funcs.IsValidIranianPhone("00123456789"))
}

func TestIsValidIranianMobile(t *testing.T) {
	assert.True(t, funcs.IsValidIranianMobile("09123456789"))
	assert.False(t, funcs.IsValidIranianMobile("08123456789"))
}

func TestIsValidIranianPostalCode(t *testing.T) {
	assert.True(t, funcs.IsValidIranianPostalCode("1234567890"))
	assert.False(t, funcs.IsValidIranianPostalCode("12345"))
}

func TestIsValidIranianIdNumber(t *testing.T) {
	assert.True(t, funcs.IsValidIranianIdNumber("1234567890"))
	assert.False(t, funcs.IsValidIranianIdNumber("12a4567890"))
}

func TestIsValidIranianNationalCode(t *testing.T) {
	assert.True(t, funcs.IsValidIranianNationalCode("0010350829"))
	assert.False(t, funcs.IsValidIranianNationalCode("1234567890"))
	assert.False(t, funcs.IsValidIranianNationalCode("abcdefghij"))
}

func TestIsValidIranianBankCard(t *testing.T) {
	assert.True(t, funcs.IsValidIranianBankCard("6362147010005732"))  // Valid Luhn
	assert.True(t, funcs.IsValidIranianBankCard("5029087000550593"))  // Valid Luhn
	assert.True(t, funcs.IsValidIranianBankCard("6037991199500590"))  // Valid Luhn
	assert.False(t, funcs.IsValidIranianBankCard("1234567890123456")) // Invalid Luhn
}

func TestIsValidIranianIBAN(t *testing.T) {
	assert.True(t, funcs.IsValidIranianIBAN("IR510550011775005110110001"))  // Valid
	assert.True(t, funcs.IsValidIranianIBAN("820540102680020817909002"))    // Without IR prefix
	assert.False(t, funcs.IsValidIranianIBAN("IR123456789000000000000000")) // Too short
}

func TestIsValidIP(t *testing.T) {
	assert.True(t, funcs.IsValidIP("192.168.1.1"))
	assert.True(t, funcs.IsValidIP("::1"))
	assert.False(t, funcs.IsValidIP("999.999.999.999"))
}

func TestIsValidIPPort(t *testing.T) {
	assert.True(t, funcs.IsValidIPPort("127.0.0.1:8080"))
	assert.False(t, funcs.IsValidIPPort("127.0.0.1"))
	assert.False(t, funcs.IsValidIPPort("127.0.0.1:99999"))
	assert.False(t, funcs.IsValidIPPort("invalid:port"))
}
