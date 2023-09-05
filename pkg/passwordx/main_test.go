package passwordx

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thangpham4/self-project/pkg/randomx"
)

func TestPassword(t *testing.T) {
	r := require.New(t)
	passwordString := randomx.RandString(10)
	password := Password{
		Password: passwordString,
	}

	hashedPassword, err := password.HasingPassword()
	r.NoError(err)
	r.NotEmpty(hashedPassword)

	ok := password.CheckPassword(hashedPassword)
	r.True(ok)
}
