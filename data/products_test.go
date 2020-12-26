package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "tea",
		Price: 1.00,
		SKU:   "abs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
