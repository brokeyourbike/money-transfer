package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionCanGetReferenceForHumans(t *testing.T) {
	trx1 := Transaction{ReferenceID: 1}
	trx2 := Transaction{ReferenceID: 4294967295}

	assert.Equal(t, "0000000001", trx1.GetReferenceForHumans())
	assert.Equal(t, "4294967295", trx2.GetReferenceForHumans())
}
