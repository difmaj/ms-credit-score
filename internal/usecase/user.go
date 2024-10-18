package usecase

import (
	"context"

	"github.com/difmaj/ms-credit-score/internal/domain"
	"github.com/difmaj/ms-credit-score/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

// Login logs in a user.
func (uc *Usecase) Login(ctx context.Context, input *dto.LoginHTTPInput) (*dto.LoginHTTPOutput, error) {
	user, err := uc.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, err
	}

	privileges, err := uc.repo.GetPrivilegesByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	claims, err := uc.GenerateToken(user, getPrivileges(privileges))
	if err != nil {
		return nil, err
	}

	return &dto.LoginHTTPOutput{
		ID:           user.ID,
		RefreshToken: claims.Ref,
		AccessToken:  claims.Token,
	}, nil
}

func getPrivileges(privileges []*domain.Privilege) domain.Privileges {
	privs := make(domain.Privileges, 0)
	for _, privilege := range privileges {
		if _, ok := privs[privilege.Context]; !ok {
			privs[privilege.Context] = []string{}
		}
		privs[privilege.Context] = append(privs[privilege.Context], privilege.Action)
	}
	return privs
}
