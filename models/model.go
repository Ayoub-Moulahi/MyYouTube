package models

import (
	"database/sql"
	"time"
)

type Channel struct {
	ID          int32          `json:"id"`
	CreatedAt   sql.NullTime   `json:"createdAt"`
	UpdatedAt   sql.NullTime   `json:"updatedAt"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Video       []string       `json:"video"`
	UserID      int32          `json:"userID"`
}

type Comment struct {
	ID        int32        `json:"id"`
	CreatedAt sql.NullTime `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
	Contents  string       `json:"contents"`
	UserID    int32        `json:"userID"`
	VideoID   int32        `json:"videoID"`
}

type Follower struct {
	ID         int32        `json:"id"`
	CreatedAt  sql.NullTime `json:"createdAt"`
	LeaderID   int32        `json:"leaderID"`
	FollowerID int32        `json:"followerID"`
}

type User struct {
	ID           int32        `json:"id"`
	CreatedAt    sql.NullTime `json:"createdAt"`
	UpdatedAt    sql.NullTime `json:"updatedAt"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	Birthdate    time.Time    `json:"birthdate"`
	Password     string       `json:"-"`
	PasswordHash string       `json:"passwordHash"`
	Remember     string       `json:"-"`
	RememberHash string       `json:"rememberHash"`
}

type Video struct {
	ID        int32        `json:"id"`
	CreatedAt sql.NullTime `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
	Title     string       `json:"title"`
	Url       string       `json:"url"`
	ChannelID int32        `json:"channelID"`
	UserID    int32        `json:"userID"`
}
