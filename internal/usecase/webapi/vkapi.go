package webapi

import (
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/go-vk-api/vk"
)

// VKWebAPI - work with VK API.
type VKWebAPI struct {
	clt *vk.Client
}

// New - VKWebAPI.
func New(accessToken string) *VKWebAPI {
	clt, _ := vk.NewClientWithOptions(vk.WithToken(accessToken))

	return &VKWebAPI{
		clt: clt,
	}
}

// GetVKUser - return info about vk user
func (v *VKWebAPI) GetVKUser(userIds ...string) ([]entity.VKUser, error) {
	results := make([]entity.VKUser, 0)
	err := v.clt.CallMethod("users.get", vk.RequestParams{
		"user_ids": userIds,
	}, &results)

	return results, err
}

// GetVKUserFriends - return user friend
func (v *VKWebAPI) GetVKUserFriends(userId string) (entity.VKFriends, error) {
	results := entity.VKFriends{}
	err := v.clt.CallMethod("friends.get", vk.RequestParams{
		"user_id": userId,
		"order":   "random",
	}, &results)

	return results, err
}
