package encoder

import (
	"fmt"
	"testing"
)

// Unit Tests

func TestEncodeReturnsErrorOnNegativeNumber(t *testing.T) {
	answer, err := Encode(-1)

	if answer != "" {
		t.Errorf("Expected '', but got %s", answer)
	}

	if err == nil {
		t.Errorf("Expected error, but got %v", err)
	}
}

func TestEncodeReturnsCorrectRomanString(t *testing.T) {
	data := []struct {
		got  int
		want string
	}{
		{0, "nulla"},
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{13, "XIII"},
		{14, "XIV"},
		{15, "XV"},
		{16, "XVI"},
		{17, "XVII"},
		{18, "XVIII"},
		{19, "XIX"},
		{20, "XX"},
		{59, "LIX"},
		{1990, "MCMXC"},
		{2008, "MMVIII"},
		{2020, "MMXX"},
		{1026, "MXXVI"},
		{1508, "MDVIII"},
		{1141, "MCXLI"},
		{1482, "MCDLXXXII"},
		{2207, "MMCCVII"},
		{2752, "MMDCCLII"},
		{2185, "MMCLXXXV"},
		{357, "CCCLVII"},
		{510, "DX"},
		{916, "CMXVI"},
		{2555, "MMDLV"},
		{2709, "MMDCCIX"},
		{652, "DCLII"},
		{1336, "MCCCXXXVI"},
		{2476, "MMCDLXXVI"},
		{2045, "MMXLV"},
		{1744, "MDCCXLIV"},
		{441, "CDXLI"},
		{1823, "MDCCCXXIII"},
		{2018, "MMXVIII"},
		{1421, "MCDXXI"},
		{462, "CDLXII"},
		{1344, "MCCCXLIV"},
		{2352, "MMCCCLII"},
		{792, "DCCXCII"},
		{761, "DCCLXI"},
		{68, "LXVIII"},
		{2143, "MMCXLIII"},
		{2317, "MMCCCXVII"},
		{2485, "MMCDLXXXV"},
		{1400, "MCD"},
		{1331, "MCCCXXXI"},
		{2802, "MMDCCCII"},
		{94, "XCIV"},
		{1467, "MCDLXVII"},
		{2700, "MMDCC"},
		{588, "DLXXXVIII"},
		{2793, "MMDCCXCIII"},
		{1137, "MCXXXVII"},
		{434, "CDXXXIV"},
		{394, "CCCXCIV"},
		{2612, "MMDCXII"},
		{941, "CMXLI"},
		{1433, "MCDXXXIII"},
		{2068, "MMLXVIII"},
		{800, "DCCC"},
		{281, "CCLXXXI"},
		{921, "CMXXI"},
		{2178, "MMCLXXVIII"},
		{2847, "MMDCCCXLVII"},
		{896, "DCCCXCVI"},
		{1863, "MDCCCLXIII"},
		{2824, "MMDCCCXXIV"},
		{1617, "MDCXVII"},
		{2531, "MMDXXXI"},
		{962, "CMLXII"},
		{2016, "MMXVI"},
		{430, "CDXXX"},
		{578, "DLXXVIII"},
		{1818, "MDCCCXVIII"},
		{1373, "MCCCLXXIII"},
		{136, "CXXXVI"},
		{1808, "MDCCCVIII"},
		{1240, "MCCXL"},
		{2704, "MMDCCIV"},
		{1467, "MCDLXVII"},
		{1165, "MCLXV"},
		{1207, "MCCVII"},
		{714, "DCCXIV"},
		{1563, "MDLXIII"},
		{391, "CCCXCI"},
		{2644, "MMDCXLIV"},
		{761, "DCCLXI"},
		{1244, "MCCXLIV"},
		{2139, "MMCXXXIX"},
		{2812, "MMDCCCXII"},
		{2155, "MMCLV"},
		{1885, "MDCCCLXXXV"},
	}

	for _, tt := range data {
		test := fmt.Sprintf("%d, %s", tt.got, tt.want)
		t.Run(test, func(t *testing.T) {
			answer, _ := Encode(tt.got)
			if answer != tt.want {
				t.Errorf("Got %s but wanted %s", answer, tt.want)
			}
		})
	}
}

// Benchmarks

const (
	MaxInt8  = 1<<7 - 1
	MaxInt16 = 1<<15 - 1
	MaxInt32 = 1<<31 - 1
)

func benchmark(i int, b *testing.B) string {
	var a string

	for n := 0; n < b.N; n++ {
		a, _ = Encode(i)
	}

	return a
}

func BenchmarkEncodeMaxInt8(b *testing.B)  { benchmark(MaxInt8, b) }
func BenchmarkEncodeMaxInt16(b *testing.B) { benchmark(MaxInt16, b) }
func BenchmarkEncodeMaxInt32(b *testing.B) { benchmark(MaxInt32, b) }
