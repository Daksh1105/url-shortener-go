package encoder_test

import (
	"URL-Shortner/internal/encoder"
	"testing"
)

func TestBase62EncodeDecode(t *testing.T) {
	// Test ID 0
	id := uint64(0)
	encoded := encoder.Base62Encode(id)
	if encoded != "0" {
		t.Errorf("Base62Encode(0) = %s, want '0'", encoded)
	}
	decoded := encoder.Base62Decode(encoded)
	if decoded != id {
		t.Errorf("Base62Decode(%s) = %d, want %d", encoded, decoded, id)
	}

	// Test ID 1
	id = 1
	encoded = encoder.Base62Encode(id)
	if encoded != "1" {
		t.Errorf("Base62Encode(1) = %s, want '1'", encoded)
	}
	decoded = encoder.Base62Decode(encoded)
	if decoded != id {
		t.Errorf("Base62Decode(%s) = %d, want %d", encoded, decoded, id)
	}

	// Test ID 62 (should be "10")
	id = 62
	encoded = encoder.Base62Encode(id)
	if encoded != "10" {
		t.Errorf("Base62Encode(62) = %s, want '10'", encoded)
	}
	decoded = encoder.Base62Decode(encoded)
	if decoded != id {
		t.Errorf("Base62Decode(%s) = %d, want %d", encoded, decoded, id)
	}

	// Test large ID (max uint64)
	id = 18446744073709551615 // max uint64
	encoded = encoder.Base62Encode(id)
	decoded = encoder.Base62Decode(encoded)
	if decoded != id {
		t.Errorf("Round-trip for max uint64 failed: got %d, want %d", decoded, id)
	}
}
