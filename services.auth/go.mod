module github.com/mixarchitecture/arch/auth

replace github.com/mixarchitecture/arch/shared => ../services.shared

go 1.19

require (
	github.com/mixarchitecture/arch/shared v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.5.0
)

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.2.1 // indirect
	golang.org/x/text v0.6.0 // indirect
)
