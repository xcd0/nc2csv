package util

/*

func Save(path string, in interface{}) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	enc := gob.NewEncoder(f)

	if err := enc.Encode(in); err != nil {
		log.Fatal(err)
	}
}

func Load(path string, out *[]token.Token) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	if err := dec.Decode(&out); err != nil {
		log.Fatal("decode error:", err)
	}
}
*/
