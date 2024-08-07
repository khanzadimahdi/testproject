package register

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/khanzadimahdi/testproject/application/auth"
	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/user"
	"github.com/khanzadimahdi/testproject/infrastructure/jwt"
	"time"
)

type UseCase struct {
	userRepository user.Repository
	jwt            *jwt.JWT
	mailer         domain.Mailer
	mailFrom       string
}

func NewUseCase(
	userRepository user.Repository,
	JWT *jwt.JWT,
	mailer domain.Mailer,
	mailFrom string,
) *UseCase {
	return &UseCase{
		userRepository: userRepository,
		jwt:            JWT,
		mailer:         mailer,
		mailFrom:       mailFrom,
	}
}

func (uc *UseCase) Register(request Request) (*Response, error) {
	if ok, validation := request.Validate(); !ok {
		return &Response{
			ValidationErrors: validation,
		}, nil
	}

	exists, err := uc.IdentityExists(request.Identity)
	if err != nil {
		return nil, err
	}

	if exists {
		return &Response{
			ValidationErrors: map[string]string{
				"identity": "user with given email already exists",
			},
		}, nil
	}

	resetPasswordToken, err := uc.registrationToken(request.Identity)
	if err != nil {
		return nil, err
	}

	resetPasswordToken = base64.URLEncoding.EncodeToString([]byte(resetPasswordToken))

	var msg bytes.Buffer
	if _, err := msg.WriteString(resetPasswordToken); err != nil {
		return nil, err
	}

	if err := uc.mailer.SendMail(uc.mailFrom, request.Identity, "Registration", msg.Bytes()); err != nil {
		return nil, err
	}

	return &Response{}, nil
}

func (uc *UseCase) IdentityExists(identity string) (bool, error) {
	_, err := uc.userRepository.GetOneByIdentity(identity)
	if errors.Is(err, domain.ErrNotExists) {
		return true, nil
	}

	return false, err
}

func (uc *UseCase) registrationToken(identity string) (string, error) {
	b := jwt.NewClaimsBuilder()
	b.SetSubject(identity)
	b.SetNotBefore(time.Now())
	b.SetExpirationTime(time.Now().Add(24 * time.Hour))
	b.SetIssuedAt(time.Now())
	b.SetAudience([]string{auth.RegistrationToken})

	return uc.jwt.Generate(b.Build())
}
