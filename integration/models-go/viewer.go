package models

import "github.com/swiftcarrot/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
