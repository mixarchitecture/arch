package app

import "github.com/mixarchitecture/arch/auth/src/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	Login        command.LoginHandler
	Register     command.RegisterHandler
	RefreshToken command.RefreshTokenHandler
	Logout       command.LogoutHandler
}
