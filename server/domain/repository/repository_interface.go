package repository

import (
	"context"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context) (*domain.User, error)
	UpdateUser(ctx context.Context, beforeUser *domain.User, afterUser *domain.User) error
	DeleteUser(ctx context.Context, user *domain.User) error
}

type ChannelRepository interface {
	CreateChannel(ctx context.Context, channel *domain.Channel) error
	GetChannels(ctx context.Context) ([]*domain.Channel, error)
	UpdateChannel(ctx context.Context, channelID domain.ChannelID, updatedChannel *domain.Channel) error
	DeleteChannel(ctx context.Context, channelID domain.ChannelID) error
}

type MessageRepository interface {
	// 新しいメッセージを作成する
	CreateMessage(ctx context.Context, message *domain.Message) error

	// ChannelID を指定して送信済みすべてのメッセージ一覧を取得
	// フロントエンドから 1 分に 1 回のリクエストを想定
	GetAllSendMessage(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error)

	// ChannelID を指定して特定の user の未送信のメッセージ一覧を取得
	GetMessageByChannelIDByIsNotSendAndUserID(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error)

	// メッセージを更新する
	UpdateMessage(ctx context.Context, afterMessage *domain.Message) error
	// メッセージを削除する
	DeleteMessage(ctx context.Context, message *domain.Message) error
}

type JoinChannelToUserRepository interface {
	// userID を指定して channelID を GET
	GetChannelIDByUserID(ctx context.Context, userID domain.UserID) ([]domain.ChannelID, error)
	// channelID を指定して userID を GET
	GetUserIDByChannelID(ctx context.Context, userID domain.ChannelID) ([]domain.UserID, error)
	// チャンネルに入会したときに実行される
	CreateConnectionUserIDToChannelID(ctx context.Context, userid domain.UserID, channelID domain.ChannelID) error
	// チャンネルから脱退したときに実行される
	DeleteConnectionUserIDToChannelID(ctx context.Context, userid domain.UserID, channelID domain.ChannelID) error
}
