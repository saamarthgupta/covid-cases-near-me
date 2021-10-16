package repository

import (
	"context"
	"covid_cases_near_me/config"
	"covid_cases_near_me/constants"
	"covid_cases_near_me/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CaseRepository struct {
}

func NewCaseRepository(config *config.Config) *CaseRepository {
	return &CaseRepository{}
}

func (caseRepository *CaseRepository) Save(countryCaseData model.CaseData, stateVsCaseDataMap map[string]model.CaseData) error {
	ctx, cancel := getMongoContext()
	countryCollection := getMongoClient(ctx).Database(constants.DATABASE_NAME).Collection(constants.COUNTRY_COLLECTION)
	countryCollection.InsertOne(ctx, countryCaseData)
	stateCollection := getMongoClient(ctx).Database(constants.DATABASE_NAME).Collection(constants.STATE_COLLECTION)
	if len(stateVsCaseDataMap) > 0 {
		countriesToUpdate := make([]interface{}, len(stateVsCaseDataMap))
		i := 0
		for _, value := range stateVsCaseDataMap {
			countriesToUpdate[i] = value
			i++
		}
		opts := options.InsertMany().SetOrdered(false)
		stateCollection.InsertMany(ctx, countriesToUpdate, opts)
	}
	defer cancel()
	return nil
}

func (caseRepository *CaseRepository) GetCasesByCountry(country string) (model.CaseData, error) {
	ctx, cancel := getMongoContext()
	countryCollection := getMongoClient(ctx).Database(constants.DATABASE_NAME).Collection(constants.COUNTRY_COLLECTION)
	opts := options.FindOne().SetSort(bson.M{"date": -1})
	countryQuery := bson.M{"country": country}
	countryData := model.CaseData{}
	err := countryCollection.FindOne(ctx, countryQuery, opts).Decode(&countryData)
	defer cancel()
	return countryData, err
}

func (caseRepository *CaseRepository) GetCasesByState(state string) (model.CaseData, error) {
	ctx, cancel := getMongoContext()
	stateCollection := getMongoClient(ctx).Database(constants.DATABASE_NAME).Collection(constants.STATE_COLLECTION)
	opts := options.FindOne().SetSort(bson.M{"date": -1})
	stateQuery := bson.M{"province": state}
	stateData := model.CaseData{}
	err := stateCollection.FindOne(ctx, stateQuery, opts).Decode(&stateData)
	defer cancel()
	return stateData, err
}

func getMongoContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	return ctx, cancel
}

func getMongoClient(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://root:1234@covid-cases.5kaaz.mongodb.net/covid_cases?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal("Unable to Connect to mongo!!")
		return nil
	}
	return client
}
