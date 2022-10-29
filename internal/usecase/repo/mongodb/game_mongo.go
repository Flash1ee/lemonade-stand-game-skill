package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/google/uuid"
)

type GameRepository struct {
	clt               *mongo.Client
	collectUsers      *mongo.Collection
	collectStatistics *mongo.Collection
	ctx               context.Context
}

func NewGameRepository(clt *mongo.Client) *GameRepository {
	return &GameRepository{
		clt:               clt,
		collectUsers:      clt.Database("game").Collection("users"),
		collectStatistics: clt.Database("game").Collection("statistics"),
		ctx:               context.Background(),
	}
}

func (repo *GameRepository) CreateUser(username string) (string, error) {
	sessionId := uuid.New().String()

	_, err := repo.collectUsers.InsertOne(repo.ctx, entity.NewSession(username, sessionId))

	return sessionId, err
}

func (repo *GameRepository) GetBalance(sessionID string) (int64, error) {
	session, err := repo.GetSession(sessionID)
	if err != nil {
		return 0, err
	}

	return session.Balance, nil
}

func (repo *GameRepository) GetSession(sessionID string) (entity.Session, error) {
	session := entity.Session{}
	err := repo.collectUsers.FindOne(repo.ctx, bson.D{{"session_id", sessionID}}).Decode(&session)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Session{}, errors.New("not found user")
		}
		return entity.Session{}, err
	}

	return session, nil
}

func (repo *GameRepository) SetDayWeather(sessionID string, weather entity.Weather) error {
	session, err := repo.GetSession(sessionID)
	if err != nil {
		return err
	}
	curDay := session.CurDay
	session.Weather[curDay] = weather

	filter := bson.D{{"session_id", sessionID}}
	update := bson.D{{"$set", bson.D{{"weather", session.Weather}}}}

	_, err = repo.collectUsers.UpdateOne(repo.ctx, filter, update)

	return err
}

func (repo *GameRepository) NextDay(sessionID string, newBalance int64) error {
	session, err := repo.GetSession(sessionID)
	if err != nil {
		return err
	}
	session.CurDay += 1
	session.Balance = newBalance
	filter := bson.D{{"session_id", sessionID}}
	update := bson.D{{"$set", bson.D{{"balance", session.Balance}, {"cur_day", session.CurDay}}}}

	_, err = repo.collectUsers.UpdateOne(repo.ctx, filter, update)

	return err
}

func (repo *GameRepository) SaveResult(sessionID string, result int64) error {
	session, err := repo.GetSession(sessionID)
	if err != nil {
		return err
	}

	_, err = repo.collectStatistics.InsertOne(repo.ctx,
		entity.Statistics{VKUserId: session.VKUserId, Result: result},
	)

	return err
}

func (repo *GameRepository) GetResult(userID string) ([]entity.Statistics, error) {
	userRes := entity.Statistics{}
	err := repo.collectStatistics.FindOne(repo.ctx,
		bson.D{{"vk_user_id", userID}}).
		Decode(&userRes)
	if err != nil {
		return nil, err
	}
	results := make([]entity.Statistics, 0, 5)
	opts := options.Find().SetSort(bson.D{{"rating", 1}})
	cur, err := repo.collectStatistics.Find(repo.ctx, opts)
	if err != nil {
		return nil, err
	}
	err = cur.All(repo.ctx, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}

	results[len(results)-1] = userRes
	return results[:5], nil
}
