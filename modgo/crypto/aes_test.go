package crypto

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAES(m *testing.T) {
	raw := `{
    "batch_len": 1000,
    "batch_size": 2097152, 
    "batch_interval": 60,
    "batch_try_times": 3
}
`
	en, _ := AesEncode(raw)
	fmt.Println(en)
	de, _ := AesDecode(en)
	fmt.Println(de)

	fmt.Println(raw == de)
}

func TestFile(t *testing.T) {
	bs, _ := ioutil.ReadFile("testfile")
	de, err := AesDecode(string(bs))
	if err != nil {
		panic(err)
	}
	fmt.Println(de)
}
