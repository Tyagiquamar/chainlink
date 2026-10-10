package solutils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSafeArchiveExtractPath(t *testing.T) {
	t.Parallel()

	targetDir := t.TempDir()

	outPath, err := safeArchiveExtractPath(targetDir, "program.so")
	require.NoError(t, err)
	require.Equal(t, filepath.Join(targetDir, "program.so"), outPath)

	_, err = safeArchiveExtractPath(targetDir, "../../etc/passwd")
	require.ErrorContains(t, err, "escapes target directory")
}
