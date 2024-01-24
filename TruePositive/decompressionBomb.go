func decompressionBombNoncompliant() {
	b := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	bb := bytes.NewReader(b)
	r, err := zlib.NewReader(bb)
	if err != nil {
		panic(err)
	}
	// Noncompliant: bytes read is not limited in `io.Copy()`.
	out, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("out: %v", out)
}
