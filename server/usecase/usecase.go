package usecase

import (
	"context"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"github.com/AbeTetsuya20/ddd_challenge/server/domain/repository"
)

type Usecase interface {

	// ユーザーは、【チャット】の【所属メンバー】に【メッセージ】を送信する
	SendMessage(ctx context.Context, message *domain.Message) error
	
	CreateUser(ctx context.Context, user *domain.User) error

	// チャンネルを作成する
	CreateChannel(ctx context.Context, channel *domain.Channel) error

	// ユーザーは、【チャット】の【名前】,【所属メンバー】などを設定する。つまり、チャンネルの名前を編集する？
	EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error

	// チャンネルを削除する
	DeleteChannel(ctx context.Context, channel *domain.Channel) error

	// チャンネルの全メッセージを取得する。これが画面に表示される
	// フロントエンドから 60 秒に一回のリクエストを想定
	MessageList(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error)

	// ユーザーは、【送信予定のメッセージ】を全て確認できる
	// header に ユーザー ID を入れているものとする
	// isSend が false のものを取得する
	GetScheduleToSendMessage(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error)

	// ユーザーは、【送信予定のメッセージ】を編集・削除できる
	EditScheduleToSendMessage(ctx context.Context, afterMessage *domain.Message) error
}

type ChatToolUsecase struct {
	UserRepo    repository.UserRepository
	ChannelRepo repository.ChannelRepository
	MessageRepo repository.MessageRepository
}

func (c ChatToolUsecase) SendMessage(ctx context.Context, message *domain.Message) error {
	err := c.MessageRepo.CreateMessage(ctx, message)
	return err
}

func (c ChatToolUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	err := c.UserRepo.CreateUser(ctx, user)
	return err
}

func (c ChatToolUsecase) CreateChannel(ctx context.Context, channel *domain.Channel) error {
	err := c.ChannelRepo.CreateChannel(ctx, channel)
	return err
}

func (c ChatToolUsecase) EditChannelConfig(ctx context.Context, beforeChannel *domain.Channel, afterChannel *domain.Channel) error {
	err := c.ChannelRepo.UpdateChannel(ctx, beforeChannel, afterChannel)
	return err
}

func (c ChatToolUsecase) DeleteChannel(ctx context.Context, channel *domain.Channel) error {
	err := c.ChannelRepo.DeleteChannel(ctx, channel)
	return err
}

func (c ChatToolUsecase) MessageList(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error) {
	messages, err := c.MessageRepo.GetAllSendMessage(ctx, channelID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (c ChatToolUsecase) GetScheduleToSendMessage(ctx context.Context, channelID domain.ChannelID) ([]*domain.Message, error) {
	// userID は header から取得
	messages, err := c.MessageRepo.GetMessageByChannelIDByIsNotSendAndUserID(ctx, channelID)
	if err != nil {
		return nil, err
	}
	return messages, nil

}

func (c ChatToolUsecase) EditScheduleToSendMessage(ctx context.Context, afterMessage *domain.Message) error {
	err := c.MessageRepo.UpdateMessage(ctx, afterMessage)
	return err
}
