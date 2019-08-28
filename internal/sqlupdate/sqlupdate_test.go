package sqlupdate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	ov, cv string
	b      bool
}{
	{"0.0.1", "0.2.0", true},
	{"v1.0.2", "v3.0.1", true},
	{"v1.3.2", "v1.2.4", false},
	{"2.3.4", "2.4", true},
	{"2.3.4", "2.3.2.20180809", false},
	{"2.3.4", "2.3.4.20180809", true},
}

func TestSqlUpdate_Decode(t *testing.T) {
	s := new(SqlUpdate)
	err := s.decode("./testdata/record.json")
	assert.NoError(t, err)
	assert.Equal(t, "test", s.Project)
	assert.Equal(t, "0.0.1", s.Update[0].Version)
	// 比较版本号
	res1 := s.compareResult("0.0.0", "1.1.1")
	assert.Equal(t, 4, len(res1))
	res2 := s.compareResult("0.1.0", "1.1.1")
	assert.Equal(t, 3, len(res2))
	assert.Equal(t, "1.0.0", res2[0].Version)
}

func TestCompare(t *testing.T) {
	for _, v := range testData {
		assert.Equal(t, v.b, compare(v.ov, v.cv))
	}
}

func TestSqlUpdate_GetSqls(t *testing.T) {
	s := new(SqlUpdate)
	str, err := s.GetSqls("./testdata/record.json", "0.0.0", "3.0.2")
	assert.NoError(t, err)
	assert.Contains(t, str, "COMMENT ON COLUMN public.account.balance IS")
}
