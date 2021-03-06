package mastodon

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Account hold information for mastodon account.
type Account struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username"`
	Acct           string    `json:"acct"`
	DisplayName    string    `json:"display_name"`
	Locked         bool      `json:"locked"`
	CreatedAt      time.Time `json:"created_at"`
	FollowersCount int64     `json:"followers_count"`
	FollowingCount int64     `json:"following_count"`
	StatusesCount  int64     `json:"statuses_count"`
	Note           string    `json:"note"`
	URL            string    `json:"url"`
	Avatar         string    `json:"avatar"`
	AvatarStatic   string    `json:"avatar_static"`
	Header         string    `json:"header"`
	HeaderStatic   string    `json:"header_static"`
}

// GetAccount return Account.
func (c *Client) GetAccount(ctx context.Context, id int) (*Account, error) {
	var account Account
	err := c.doAPI(ctx, http.MethodGet, fmt.Sprintf("/api/v1/accounts/%d", id), nil, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetAccountCurrentUser return Account of current user.
func (c *Client) GetAccountCurrentUser(ctx context.Context) (*Account, error) {
	var account Account
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/accounts/verify_credentials", nil, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Profile is a struct for updating profiles.
type Profile struct {
	// If it is nil it will not be updated.
	// If it is empty, update it with empty.
	DisplayName *string
	Note        *string

	// Set the base64 encoded character string of the image.
	Avatar string
	Header string
}

// AccountUpdate updates the information of the current user.
func (c *Client) AccountUpdate(ctx context.Context, profile *Profile) (*Account, error) {
	params := url.Values{}
	if profile.DisplayName != nil {
		params.Set("display_name", *profile.DisplayName)
	}
	if profile.Note != nil {
		params.Set("note", *profile.Note)
	}
	if profile.Avatar != "" {
		params.Set("avatar", profile.Avatar)
	}
	if profile.Header != "" {
		params.Set("header", profile.Header)
	}

	var account Account
	err := c.doAPI(ctx, http.MethodPatch, "/api/v1/accounts/update_credentials", params, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetAccountStatuses return statuses by specified accuont.
func (c *Client) GetAccountStatuses(ctx context.Context, id int64) ([]*Status, error) {
	var statuses []*Status
	err := c.doAPI(ctx, http.MethodGet, fmt.Sprintf("/api/v1/accounts/%d/statuses", id), nil, &statuses)
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

// GetAccountFollowers return followers list.
func (c *Client) GetAccountFollowers(ctx context.Context, id int64) ([]*Account, error) {
	var accounts []*Account
	err := c.doAPI(ctx, http.MethodGet, fmt.Sprintf("/api/v1/accounts/%d/followers", id), nil, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccountFollowing return following list.
func (c *Client) GetAccountFollowing(ctx context.Context, id int64) ([]*Account, error) {
	var accounts []*Account
	err := c.doAPI(ctx, http.MethodGet, fmt.Sprintf("/api/v1/accounts/%d/following", id), nil, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetBlocks return block list.
func (c *Client) GetBlocks(ctx context.Context) ([]*Account, error) {
	var accounts []*Account
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/blocks", nil, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Relationship hold information for relation-ship to the account.
type Relationship struct {
	ID         int64 `json:"id"`
	Following  bool  `json:"following"`
	FollowedBy bool  `json:"followed_by"`
	Blocking   bool  `json:"blocking"`
	Muting     bool  `json:"muting"`
	Requested  bool  `json:"requested"`
}

// AccountFollow follow the account.
func (c *Client) AccountFollow(ctx context.Context, id int64) (*Relationship, error) {
	var relationship Relationship
	err := c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/accounts/%d/follow", id), nil, &relationship)
	if err != nil {
		return nil, err
	}
	return &relationship, nil
}

// AccountUnfollow unfollow the account.
func (c *Client) AccountUnfollow(ctx context.Context, id int64) (*Relationship, error) {
	var relationship Relationship
	err := c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/accounts/%d/unfollow", id), nil, &relationship)
	if err != nil {
		return nil, err
	}
	return &relationship, nil
}

// AccountBlock block the account.
func (c *Client) AccountBlock(ctx context.Context, id int64) (*Relationship, error) {
	var relationship Relationship
	err := c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/accounts/%d/block", id), nil, &relationship)
	if err != nil {
		return nil, err
	}
	return &relationship, nil
}

// AccountUnblock unblock the account.
func (c *Client) AccountUnblock(ctx context.Context, id int64) (*Relationship, error) {
	var relationship Relationship
	err := c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/accounts/%d/unblock", id), nil, &relationship)
	if err != nil {
		return nil, err
	}
	return &relationship, nil
}

// AccountMute mute the account.
func (c *Client) AccountMute(ctx context.Context, id int64) (*Relationship, error) {
	var relationship Relationship
	err := c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/accounts/%d/mute", id), nil, &relationship)
	if err != nil {
		return nil, err
	}
	return &relationship, nil
}

// AccountUnmute unmute the account.
func (c *Client) AccountUnmute(ctx context.Context, id int64) (*Relationship, error) {
	var relationship Relationship
	err := c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/accounts/%d/unmute", id), nil, &relationship)
	if err != nil {
		return nil, err
	}
	return &relationship, nil
}

// GetAccountRelationship return relationship for the account.
func (c *Client) GetAccountRelationships(ctx context.Context, ids []int64) ([]*Relationship, error) {
	params := url.Values{}
	for _, id := range ids {
		params.Add("id[]", fmt.Sprint(id))
	}

	var relationships []*Relationship
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/accounts/relationships", params, &relationships)
	if err != nil {
		return nil, err
	}
	return relationships, nil
}

// AccountsSearch search accounts by query.
func (c *Client) AccountsSearch(ctx context.Context, q string, limit int64) ([]*Account, error) {
	params := url.Values{}
	params.Set("q", q)
	params.Set("limit", fmt.Sprint(limit))

	var accounts []*Account
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/accounts/search", params, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// FollowRemoteUser send follow-request.
func (c *Client) FollowRemoteUser(ctx context.Context, uri string) (*Account, error) {
	params := url.Values{}
	params.Set("uri", uri)

	var account Account
	err := c.doAPI(ctx, http.MethodPost, "/api/v1/follows", params, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetFollowRequests return follow-requests.
func (c *Client) GetFollowRequests(ctx context.Context) ([]*Account, error) {
	var accounts []*Account
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/follow_requests", nil, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// FollowRequestAuthorize is authorize the follow request of user with id.
func (c *Client) FollowRequestAuthorize(ctx context.Context, id int64) error {
	return c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/follow_requests/%d/authorize", id), nil, nil)
}

// FollowRequestReject is rejects the follow request of user with id.
func (c *Client) FollowRequestReject(ctx context.Context, id int64) error {
	return c.doAPI(ctx, http.MethodPost, fmt.Sprintf("/api/v1/follow_requests/%d/reject", id), nil, nil)
}

// GetMutes returns the list of users muted by the current user.
func (c *Client) GetMutes(ctx context.Context) ([]*Account, error) {
	var accounts []*Account
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/mutes", nil, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
