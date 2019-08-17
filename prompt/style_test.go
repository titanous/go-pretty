package prompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimestampOptions_Generate(t *testing.T) {
	tso := TimestampOptionsOff
	assert.Equal(t, "", tso.Generate())

	tso = TimestampOptionsSimple
	assert.NotEmpty(t, tso.Generate())
}
