package SigninRequestDto

import (
	"fmt"
	"strings"

	authv1 "github.com/MoreWiktor/go.sso.proto/auth"
	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmailLoginSigninVo struct {
	Login string `json:"login" validate:"required,email"`
}

type PhoneLoginSigninVo struct {
	Login string `json:"login" validate:"required,e164"`
}

type PasswordSigninVo struct {
	Password string `json:"password" validate:"required"`
}

type ServiceIdSigninVo struct {
	ServiceId string `json:"serviceId" validate:"required,uuid"`
}

type CompanyIdSigninVo struct {
	CompanyId string `json:"companyId" validate:"required,uuid"`
}

const (
	EMAIL_LOGIN_TYPE = "email"
	PHONE_LOGIN_TYPE = "phone"
)

type Result struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	ServiceId string `json:"serviceId"`
	CompanyId string `json:"companyId"`
	LoginType string `json:"loginType"`
}

const (
	LOGIN_VALID_ERROR      = "login must be email or phone"
	PASSWORD_VALID_ERROR   = "password is not allowed"
	SERVICE_ID_VALID_ERROR = "serviceId must be uuid"
	COMPANY_ID_VALID_ERROR = "companyId must be uuid"
)

func RequestValidator(req *authv1.SigninRequest) (*Result, error) {
	validate := validator.New()

	Login := req.GetLogin()
	Password := req.GetPassword()
	ServiceId := req.GetServiceId()
	// TODO replace GetServiceId to GetCompanyId when updates contract
	CompanyId := req.GetServiceId()

	errors := []string{}

	LoginType := ""

	// LOGIN MUST BE EMAIL OR PHONE
	if err := validate.Struct(&EmailLoginSigninVo{Login}); err != nil {
		if err := validate.Struct(&PhoneLoginSigninVo{Login}); err != nil {
			errors = append(errors, LOGIN_VALID_ERROR)
		} else {
			LoginType = PHONE_LOGIN_TYPE
		}
	} else {
		LoginType = EMAIL_LOGIN_TYPE
	}

	// PASSWORD VALIDATION
	if err := validate.Struct(&PasswordSigninVo{Password}); err != nil {
		errors = append(errors, PASSWORD_VALID_ERROR)
	}

	// SERVICE_ID VALIDATION
	if err := validate.Struct(&ServiceIdSigninVo{ServiceId}); err != nil {
		errors = append(errors, SERVICE_ID_VALID_ERROR)
	}

	// COMPANY_ID VALIDATION
	if err := validate.Struct(&CompanyIdSigninVo{CompanyId}); err != nil {
		errors = append(errors, COMPANY_ID_VALID_ERROR)
	}

	if len(errors) == 0 {
		return &Result{Login, Password, CompanyId, ServiceId, LoginType}, nil
	}

	return nil, status.Error(codes.InvalidArgument, fmt.Sprint(strings.Join(errors[:], "; ")))
}
