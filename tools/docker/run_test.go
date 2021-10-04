package docker

import (
	"fmt"
	"github.com/fizyk/magex/file"
	"github.com/stretchr/testify/suite"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"testing"
)

type dockerSuite struct {
	suite.Suite
}

// Docker suite
func TestDockerSuite(t *testing.T) {
	suite.Run(t, new(dockerSuite))
}

// Check that run runs docker with proper user permissions
func (s *dockerSuite) TestRunCreateFile() {
	fileName := "test.txt"
	s.False(file.Exists(fileName))
	currentUser, err := user.Current()
	s.NoError(err)
	cwd, err := os.Getwd()
	s.NoError(err)
	err = Run("--rm",
		"-v",
		fmt.Sprintf("%s/:/testCreate", cwd),
		"-w",
		"/testCreate",
		"alpine",
		"touch",
		fileName,
	)
	s.NoError(err)
	defer os.Remove("test.txt")
	info, err := os.Stat(fileName)
	s.NoError(err)
	statInfo := info.Sys().(*syscall.Stat_t)
	s.NoError(err)
	// Compare the owner
	s.Equal(currentUser.Gid, strconv.FormatUint(uint64(statInfo.Gid), 10))
	s.Equal(currentUser.Uid, strconv.FormatUint(uint64(statInfo.Uid), 10))
}
