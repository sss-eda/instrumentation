package instrumentation

var connectTests = []struct{ int }{
	1, 2, 3,
}

// TestConnect TODO
func TestConnect(t *testing.T) {
	for _, test := range connectTests {
		if _, err := Connect(test); err != nil {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
