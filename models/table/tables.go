package table

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Role string    `gorm:"type:string" json:"role"`
}

type UserInfo struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	FirstName string    `gorm:"type:string" json:"first_name"`
	LastName  string    `gorm:"type:string" json:"last_name"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type UserCredential struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email    string    `gorm:"type:string;uniqueIndex"  json:"email"`
	Password string    `gorm:"type:string" json:"password"`
	UserName string    `gorm:"type:string;uniqueIndex" json:"user_name"`
	Role     string    `gorm:"foreignKey:RoleID" json:"role_id"`
}
