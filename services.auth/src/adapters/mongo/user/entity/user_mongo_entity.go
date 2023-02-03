package entity

import "github.com/mixarchitecture/arch/auth/domain/user"

type MongoUser struct {
	UUID      string   `bson:"uuid"`
	Email     string   `bson:"email"`
	Roles     []string `bson:"roles"`
	Password  []byte   `bson:"password"`
	IsActive  bool     `bson:"is_active"`
	CreatedAt string   `bson:"created_at"`
	UpdatedAt string   `bson:"updated_at"`
}

func (m *MongoUser) ToUser() *user.User {
	return &user.User{
		UUID:     m.UUID,
		Email:    m.Email,
		Roles:    m.Roles,
		Password: m.Password,
		IsActive: m.IsActive,
	}
}

func (m *MongoUser) FromUser(user *user.User) *MongoUser {
	m.UUID = user.UUID
	m.Roles = user.Roles
	m.Email = user.Email
	m.Password = user.Password
	m.IsActive = user.IsActive
	return m
}
