package installconfig

import (
	"testing"
)

const testDataDir = "../../test/testdata"

func TestLoad(t *testing.T) {
	Load(testDataDir)
}
