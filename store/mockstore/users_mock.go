// generated by gen-mocks; DO NOT EDIT

package mockstore

import (
	"golang.org/x/net/context"
	"sourcegraph.com/sourcegraph/sourcegraph/go-sourcegraph/sourcegraph"
	"sourcegraph.com/sourcegraph/sourcegraph/store"
)

type Users struct {
	Get_          func(ctx context.Context, user sourcegraph.UserSpec) (*sourcegraph.User, error)
	GetWithEmail_ func(ctx context.Context, emailAddr sourcegraph.EmailAddr) (*sourcegraph.User, error)
	List_         func(ctx context.Context, opt *sourcegraph.UsersListOptions) ([]*sourcegraph.User, error)
	ListEmails_   func(v0 context.Context, v1 sourcegraph.UserSpec) ([]*sourcegraph.EmailAddr, error)
	Count_        func(v0 context.Context) (int32, error)
}

func (s *Users) Get(ctx context.Context, user sourcegraph.UserSpec) (*sourcegraph.User, error) {
	return s.Get_(ctx, user)
}

func (s *Users) GetWithEmail(ctx context.Context, emailAddr sourcegraph.EmailAddr) (*sourcegraph.User, error) {
	return s.GetWithEmail_(ctx, emailAddr)
}

func (s *Users) List(ctx context.Context, opt *sourcegraph.UsersListOptions) ([]*sourcegraph.User, error) {
	return s.List_(ctx, opt)
}

func (s *Users) ListEmails(v0 context.Context, v1 sourcegraph.UserSpec) ([]*sourcegraph.EmailAddr, error) {
	return s.ListEmails_(v0, v1)
}

func (s *Users) Count(v0 context.Context) (int32, error) { return s.Count_(v0) }

var _ store.Users = (*Users)(nil)

type Accounts struct {
	Create_               func(ctx context.Context, newUser *sourcegraph.User, email *sourcegraph.EmailAddr) (*sourcegraph.User, error)
	GetByGitHubID_        func(ctx context.Context, id int) (*sourcegraph.User, error)
	Update_               func(v0 context.Context, v1 *sourcegraph.User) error
	UpdateEmails_         func(v0 context.Context, v1 sourcegraph.UserSpec, v2 []*sourcegraph.EmailAddr) error
	RequestPasswordReset_ func(v0 context.Context, v1 *sourcegraph.User) (*sourcegraph.PasswordResetToken, error)
	ResetPassword_        func(v0 context.Context, v1 *sourcegraph.NewPassword) error
	Delete_               func(v0 context.Context, v1 int32) error
}

func (s *Accounts) Create(ctx context.Context, newUser *sourcegraph.User, email *sourcegraph.EmailAddr) (*sourcegraph.User, error) {
	return s.Create_(ctx, newUser, email)
}

func (s *Accounts) GetByGitHubID(ctx context.Context, id int) (*sourcegraph.User, error) {
	return s.GetByGitHubID_(ctx, id)
}

func (s *Accounts) Update(v0 context.Context, v1 *sourcegraph.User) error { return s.Update_(v0, v1) }

func (s *Accounts) UpdateEmails(v0 context.Context, v1 sourcegraph.UserSpec, v2 []*sourcegraph.EmailAddr) error {
	return s.UpdateEmails_(v0, v1, v2)
}

func (s *Accounts) RequestPasswordReset(v0 context.Context, v1 *sourcegraph.User) (*sourcegraph.PasswordResetToken, error) {
	return s.RequestPasswordReset_(v0, v1)
}

func (s *Accounts) ResetPassword(v0 context.Context, v1 *sourcegraph.NewPassword) error {
	return s.ResetPassword_(v0, v1)
}

func (s *Accounts) Delete(v0 context.Context, v1 int32) error { return s.Delete_(v0, v1) }

var _ store.Accounts = (*Accounts)(nil)

type Directory struct {
	GetUserByEmail_ func(ctx context.Context, email string) (*sourcegraph.UserSpec, error)
}

func (s *Directory) GetUserByEmail(ctx context.Context, email string) (*sourcegraph.UserSpec, error) {
	return s.GetUserByEmail_(ctx, email)
}

var _ store.Directory = (*Directory)(nil)

type ExternalAuthTokens struct {
	GetUserToken_      func(ctx context.Context, user int, host, clientID string) (*store.ExternalAuthToken, error)
	SetUserToken_      func(ctx context.Context, tok *store.ExternalAuthToken) error
	DeleteToken_       func(ctx context.Context, tok *sourcegraph.ExternalTokenSpec) error
	ListExternalUsers_ func(ctx context.Context, extUIDs []int, host, clientID string) ([]*store.ExternalAuthToken, error)
}

func (s *ExternalAuthTokens) GetUserToken(ctx context.Context, user int, host, clientID string) (*store.ExternalAuthToken, error) {
	return s.GetUserToken_(ctx, user, host, clientID)
}

func (s *ExternalAuthTokens) SetUserToken(ctx context.Context, tok *store.ExternalAuthToken) error {
	return s.SetUserToken_(ctx, tok)
}

func (s *ExternalAuthTokens) DeleteToken(ctx context.Context, tok *sourcegraph.ExternalTokenSpec) error {
	return s.DeleteToken_(ctx, tok)
}

func (s *ExternalAuthTokens) ListExternalUsers(ctx context.Context, extUIDs []int, host, clientID string) ([]*store.ExternalAuthToken, error) {
	return s.ListExternalUsers_(ctx, extUIDs, host, clientID)
}

var _ store.ExternalAuthTokens = (*ExternalAuthTokens)(nil)
