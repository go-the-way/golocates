package locatevar

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	starts = 1000000
	seq    = 2000
)

func TestInit(t *testing.T) {
	generateGoFiles(seq)
}

func generateGoFiles(seq int) {
	var (
		arrConst  []string
		arrDefine []string
		arrVar    []string
	)
	for i := 0; i < seq; i++ {
		curSeq := starts + i
		arrConst = append(arrConst, fmt.Sprintf("\tId%d = cst.Id%d", curSeq, curSeq))
		curDefine := ""
		if i == seq-1 {
			curDefine = fmt.Sprintf("\tId%d string", curSeq)
		} else {
			curDefine = fmt.Sprintf("\tId%d,", curSeq)
		}
		arrDefine = append(arrDefine, curDefine)
		arrVar = append(arrVar, fmt.Sprintf("\tId%d: \"Id%d\",", curSeq, curSeq))
	}
	arrConstStr := strings.Join(arrConst, "\n")
	arrDefineStr := strings.Join(arrDefine, "\n")
	arrVarStr := strings.Join(arrVar, "\n")

	_ = os.WriteFile("const.go", []byte(`package locatevar

var (
`+arrConstStr+`
)
`), 0700)

	_ = os.WriteFile("define.go", []byte(`package locatevar

type constant struct {
`+arrDefineStr+`
}
`), 0700)

	_ = os.WriteFile("var.go", []byte(`package locatevar

var cst = &constant{
`+arrVarStr+`
}
`), 0700)

}
