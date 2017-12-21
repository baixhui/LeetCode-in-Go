package Problem0677

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapSum(t *testing.T) {
	ast := assert.New(t)

	ms := Constructor()

	ms.Insert("apple", 3)

	ast.Equal(3, ms.Sum("ap"))

	ms.Insert("app", 2)

	ast.Equal(5, ms.Sum("ap"))
}
