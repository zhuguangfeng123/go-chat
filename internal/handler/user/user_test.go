package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	dtov1 "github.com/zhuguangfeng123/go-chat/internal/dto/v1"
	iJwt "github.com/zhuguangfeng123/go-chat/internal/handler/jwt"
	"github.com/zhuguangfeng123/go-chat/internal/service/code"
	svcmocks "github.com/zhuguangfeng123/go-chat/internal/service/mocks"
	"github.com/zhuguangfeng123/go-chat/internal/service/user"
	"github.com/zhuguangfeng123/go-chat/pkg/ginx/result"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_Signup(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) (iJwt.JwtHandler, user.UserService, code.CodeService)
		reqBody  dtov1.UserSignupReq
		wantCode int
		wantBody result.Result
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) (iJwt.JwtHandler, user.UserService, code.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				userSvc.EXPECT().UserSignup(gomock.Any(), domain.User{
					Phone:    "18860313694",
					Password: "Zgf111111",
				}).Return("1111", nil)
				return nil, userSvc, codeSvc
			},
			reqBody: dtov1.UserSignupReq{
				Phone:           "18860313694",
				Password:        "Zgf111111",
				ConfirmPassword: "Zgf111111",
			},
			wantCode: http.StatusOK,
			wantBody: result.Result{
				Code: 200,
				Msg:  "SUCCESS",
				Data: nil,
			},
		},
		{
			name: "手机格式有误",
			mock: func(ctrl *gomock.Controller) (iJwt.JwtHandler, user.UserService, code.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return nil, userSvc, codeSvc
			},
			reqBody: dtov1.UserSignupReq{
				Phone:           "188603136942",
				Password:        "Zgf111111",
				ConfirmPassword: "Zgf111111",
			},
			wantCode: http.StatusOK,
			wantBody: result.Result{
				Code: -1,
				Msg:  "手机号码格式有误",
				Data: nil,
			},
		},
		{
			name: "两次密码不一致",
			mock: func(ctrl *gomock.Controller) (iJwt.JwtHandler, user.UserService, code.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return nil, userSvc, codeSvc
			},
			reqBody: dtov1.UserSignupReq{
				Phone:           "18860313694",
				Password:        "Zgf111111",
				ConfirmPassword: "111111",
			},
			wantCode: http.StatusOK,
			wantBody: result.Result{
				Code: -1,
				Msg:  "两次密码不一致",
				Data: nil,
			},
		},
		{
			name: "密码格式有误",
			mock: func(ctrl *gomock.Controller) (iJwt.JwtHandler, user.UserService, code.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return nil, userSvc, codeSvc
			},
			reqBody: dtov1.UserSignupReq{
				Phone:           "18860313694",
				Password:        "111111",
				ConfirmPassword: "111111",
			},
			wantCode: http.StatusOK,
			wantBody: result.Result{
				Code: -1,
				Msg:  "密码格式有误",
				Data: nil,
			},
		},
		{
			name: "手机号码冲突",
			mock: func(ctrl *gomock.Controller) (iJwt.JwtHandler, user.UserService, code.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				userSvc.EXPECT().UserSignup(gomock.Any(), domain.User{
					Phone:    "18860313695",
					Password: "Zgf111111",
				}).Return("", user.ErrUserDuplicatePhone)
				return nil, userSvc, codeSvc
			},
			reqBody: dtov1.UserSignupReq{
				Phone:           "18860313695",
				Password:        "Zgf111111",
				ConfirmPassword: "Zgf111111",
			},
			wantCode: http.StatusOK,
			wantBody: result.Result{
				Code: -1,
				Msg:  "手机号码已注册",
				Data: nil,
			},
		},
		{
			name: "系统异常",
			mock: func(ctrl *gomock.Controller) (iJwt.JwtHandler, user.UserService, code.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				userSvc.EXPECT().UserSignup(gomock.Any(), domain.User{
					Phone:    "18860313695",
					Password: "Zgf111111",
				}).Return("", errors.New("系统错误"))
				return nil, userSvc, codeSvc
			},
			reqBody: dtov1.UserSignupReq{
				Phone:           "18860313695",
				Password:        "Zgf111111",
				ConfirmPassword: "Zgf111111",
			},
			wantCode: http.StatusOK,
			wantBody: result.Result{
				Code: -1,
				Msg:  "系统错误",
				Data: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			server := gin.Default()
			hdl := NewUserHandler(tc.mock(ctrl))
			hdl.RegisterRouter(server)

			body, err := json.Marshal(tc.reqBody)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/user/signup", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			res := result.Result{}
			err = json.Unmarshal(resp.Body.Bytes(), &res)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, res)
		})
	}
}

func TestEncrypt(t *testing.T) {
	password := "helloword123"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}
