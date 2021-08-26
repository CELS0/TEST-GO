package services_test

import (
	"testing"

	"github.com/CELS0/TEST-GO/services"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	require.Equal(t, services.Sum(2, 3), 5)
}

func TestMultiply(t *testing.T) {
	require.Equal(t, services.Multiply(3, 7), 21)
}
