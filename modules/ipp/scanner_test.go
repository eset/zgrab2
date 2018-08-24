package ipp

import (
	"testing"
)

// TODO: Use testing harness to represent pass/failure conditions
func TestReadAllAttributes(t *testing.T) {
	var scanner Scanner
	scanner.config = &Flags{}
	// Makes truncation occur at a manageable 1024 bytes, which can be reached by just copy-paste
	scanner.config.MaxSize = 1

	body := []byte{2, 1, 4, 6, 0, 0, 0, 1, 1, 71, 0, 18, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 99, 104, 97, 114, 115, 101, 116, 0, 5, 117, 116, 102, 45, 56, 72, 0, 27, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 110, 97, 116, 117, 114, 97, 108, 45, 108, 97, 110, 103, 117, 97, 103, 101, 0, 5, 101, 110, 45, 117, 115, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 3}
	t.Error(readAllAttributes(body, &scanner))
	// Total slice is 130 bytes

	// Should have no attributes and a "Reported field length runs out of bounds." error
	tooLongName := []byte{2, 1, 4, 6, 0, 0, 0, 1, 1, 71, 0, 180, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 99, 104, 97, 114, 115, 101, 116, 0, 5, 117, 116, 102, 45, 56, 72, 0, 27, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 110, 97, 116, 117, 114, 97, 108, 45, 108, 97, 110, 103, 117, 97, 103, 101, 0, 5, 101, 110, 45, 117, 115, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 3}
	t.Error(readAllAttributes(tooLongName, &scanner))

	// Should have one attribute with no values and "attributes-charset" as its name; and a "Reported field length..." error
	tooLongValue := []byte{2, 1, 4, 6, 0, 0, 0, 1, 1, 71, 0, 18, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 99, 104, 97, 114, 115, 101, 116, 0, 150, 117, 116, 102, 45, 56, 72, 0, 27, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 110, 97, 116, 117, 114, 97, 108, 45, 108, 97, 110, 103, 117, 97, 103, 101, 0, 5, 101, 110, 45, 117, 115, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 3}
	t.Error(readAllAttributes(tooLongValue, &scanner))

	// 19 attributes
	// Should have no error, since the attributes can all be read. It's a final value of length 50.
	fullLength := []byte{2, 1, 4, 6, 0, 0, 0, 1, 1, 71, 0, 18, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 99, 104, 97, 114, 115, 101, 116, 0, 5, 117, 116, 102, 45, 56, 72, 0, 27, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 110, 97, 116, 117, 114, 97, 108, 45, 108, 97, 110, 103, 117, 97, 103, 101, 0, 5, 101, 110, 45, 117, 115, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 50, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 3}
	t.Error(readAllAttributes(fullLength, &scanner))
	attrs, _ := readAllAttributes(fullLength, &scanner)
	t.Errorf("%v\n", len(attrs[len(attrs)-1].Values[0].Bytes))

	// 19 attributes
	// Should have no error, even though not all can be read. The name of a new value is incompletely read, but then discarded without error, since truncation is detected.
	truncated := []byte{2, 1, 4, 6, 0, 0, 0, 1, 1, 71, 0, 18, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 99, 104, 97, 114, 115, 101, 116, 0, 5, 117, 116, 102, 45, 56, 72, 0, 27, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 110, 97, 116, 117, 114, 97, 108, 45, 108, 97, 110, 103, 117, 97, 103, 101, 0, 5, 101, 110, 45, 117, 115, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97}
	t.Error(readAllAttributes(truncated, &scanner))
	atrs, _ := readAllAttributes(truncated, &scanner)
	t.Errorf("%v\n", len(atrs[len(atrs)-1].Values[0].Bytes))

	// Should have no attributes and no error.
	noGroups := []byte{2, 1, 4, 6, 0, 0, 0, 1, 3}
	t.Error(readAllAttributes(noGroups, &scanner))

	// Should have usual 3 attributes and no error.
	emptyGroups := []byte{2, 1, 4, 6, 0, 0, 0, 1, 1, 71, 0, 18, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 99, 104, 97, 114, 115, 101, 116, 0, 5, 117, 116, 102, 45, 56, 72, 0, 27, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 110, 97, 116, 117, 114, 97, 108, 45, 108, 97, 110, 103, 117, 97, 103, 101, 0, 5, 101, 110, 45, 117, 115, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 0, 1, 2, 4, 5, 3}
	t.Error(readAllAttributes(emptyGroups, &scanner))

	// Should heave no attribute and no error.
	dataAfterEndOfAttrs := []byte{2, 1, 4, 6, 0, 0, 0, 1, 3, 1, 71, 0, 18, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 99, 104, 97, 114, 115, 101, 116, 0, 5, 117, 116, 102, 45, 56, 72, 0, 27, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 45, 110, 97, 116, 117, 114, 97, 108, 45, 108, 97, 110, 103, 117, 97, 103, 101, 0, 5, 101, 110, 45, 117, 115, 65, 0, 14, 115, 116, 97, 116, 117, 115, 45, 109, 101, 115, 115, 97, 103, 101, 0, 36, 84, 104, 101, 32, 112, 114, 105, 110, 116, 101, 114, 32, 111, 114, 32, 99, 108, 97, 115, 115, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 101, 120, 105, 115, 116, 46, 3}
	t.Error(readAllAttributes(dataAfterEndOfAttrs, &scanner))

	// We're never expecting to read 0 bytes unless it's into a 0-byte field, so io.EOF means something went wrong.
	// Should have 0 attributes and "Fewer body bytes read than expected." error, because we expected to read at least one delimiter-tag.
	noTagToRead := []byte{2, 1, 4, 6, 0, 0, 0, 1}
	t.Error(readAllAttributes(noTagToRead, &scanner))

	// We're never expecting to read some but not all bytes and then hit io.ErrUnexpectedEOF, so that would indicate an issue (one case is too-short body).
	tooShortBody := []byte{2, 1, 4, 6}
	t.Error(readAllAttributes(tooShortBody, &scanner))

	//things that should actually make this fail:
	//just blatantly not IPP: probably fail with wrong field-length error or reported field length error eventually?
}
