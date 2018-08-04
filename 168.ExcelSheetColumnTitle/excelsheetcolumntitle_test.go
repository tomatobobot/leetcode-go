package ExcelSheetColumnTitle

import "testing"

func TestConvertToTitle(t *testing.T) {
	str := convertToTitle(701)
	if str != "ZY" {
		t.Fatal(str)
	}
}
