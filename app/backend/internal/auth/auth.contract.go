package auth

import "context"

// Auth Manager
type UsecaseContract interface {
	Login(ctx context.Context, payload LoginPayload) (TokenResponse, error)  //-> Credential & session manager    //compare email and hash password
	Refresh(ctx context.Context, refreshToken string) (TokenResponse, error) //-> token manager & session manager //refresh old token expired,
	Logout(ctx context.Context, refreshToken string) error                   //-> token manager & session manager //revoke token
	Register(ctx context.Context, payload RegisterPayload) error             //-> Credential repository             //register new Credential
	Verify(ctx context.Context, token string) (TokenEntity, error)           //-> token manager                   //verify valid token
}

type CredentialRepositoryContract interface {
	Authentication(ctx context.Context, email string) (AuthLogin, error) //authentication credential
	Registration(ctx context.Context, entity RegisterPayload) error      //registration new Credential
	GetRoleByUserId(ctx context.Context, userId string) (Role, error)
}

// Session Manager
type SessionRepositoryContract interface {
	Create(ctx context.Context, payload CreateSessionPayload) error                    //save token to session
	Rotate(ctx context.Context, payload UpdateSessionPayload) error                    //update token
	Delete(ctx context.Context, payload DeleteSessionPayload) error                    //delete token from session
	GetVersion(ctx context.Context, accessToken string) (int, error)                   //get user version
	GetByAccessToken(ctx context.Context, accessToken string) (SessionEntity, error)   //get token by access token
	GetByRefreshToken(ctx context.Context, refreshToken string) (SessionEntity, error) //get token by refresh token
}

// Token Manager
type TokenContract interface {
	CreateToken(ctx context.Context, payload TokenEntity) (string, error) //create token with entity
	VerifyToken(ctx context.Context, token string) (TokenEntity, error)   //verify token and return entity
}

const ContextAuthEntityKey string = "user-auth" //key to get user from context
