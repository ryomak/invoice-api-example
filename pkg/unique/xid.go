package unique

import "github.com/rs/xid"

var GenerateID = func() string {
	return xid.New().String()
}

func SetFakeGenerateID(id string) {
	GenerateID = func() string {
		return id
	}
}

func ResetGenerateID() {
	GenerateID = func() string {
		return xid.New().String()
	}
}
