package models

// Profile - User profile sruct
type Profile struct {
	Role     *int      `json:"role,omitempty" bson:"role,omitempty" valid:"required"`
	Phone    *string   `json:"phone,omitempty" bson:"phone,omitempty" valid:"required"`
	Password *string   `json:"pass,omitempty" bson:"pass" valid:"required"`
	Location *Location `json:"location,omitempty" bson:"location" valid:"required"`
}

// Login - User login sruct
type Login struct {
	Role     *int    `json:"role,omitempty" bson:"role,omitempty" valid:"required"`
	Phone    *string `json:"phone,omitempty" bson:"phone,omitempty" valid:"required"`
	Password *string `json:"pass,omitempty" bson:"pass" valid:"required"`
}
