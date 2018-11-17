package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexdriaguine/toggl/models"
	"github.com/joho/godotenv"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

const basePath = "https://www.toggl.com/api/v8/"
const reportsPath = "https://toggl.com/reports/api/v2/"
const apiTokenKey = "API_TOKEN"

var apiToken string

func makeRequest(path string, params map[string]string, data interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	token := getToken()
	auth := basicAuth(token, "api_token")

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return err
	}

	requestQuery := req.URL.Query()
	requestQuery.Add("user_agent", "alexdriagin12@gmail.com")

	for key, value := range params {
		requestQuery.Add(key, value)
	}

	req.URL.RawQuery = requestQuery.Encode()
	req.Header.Add("Authorization", "Basic "+auth)

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func getToken() string {
	if len(apiToken) > 0 {
		return apiToken
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv(apiTokenKey)

	if len(token) == 0 {
		log.Fatal("No API_TOKEN found in .env")
	}

	apiToken = token
	return token
}

type flexTags struct {
	Plus  *models.Tag
	Minus *models.Tag
}

func extractFlexTags(tags []*models.Tag) *flexTags {
	var flexPlusTag *models.Tag
	var flexMinusTag *models.Tag

	for i := 0; i < len(tags); i++ {
		if tags[i].Name == "flex-plus" {
			flexPlusTag = tags[i]
		}
		if tags[i].Name == "flex-minus" {
			flexMinusTag = tags[i]
		}
	}

	return &flexTags{Plus: flexPlusTag, Minus: flexMinusTag}
}

type flexTime struct {
	Hours    int
	Minutes  int
	Seconds  int
	Positive bool
}

func (f *flexTime) ToString() string {
	secondsString := strconv.Itoa(f.Seconds)
	hoursString := strconv.Itoa(f.Hours)
	minutesString := strconv.Itoa(f.Minutes)

	if f.Seconds < 10 {
		secondsString = "0" + secondsString
	}
	if f.Minutes < 10 {
		minutesString = "0" + minutesString
	}
	if f.Hours < 10 {
		hoursString = "0" + hoursString
	}

	sign := "+"
	if !f.Positive {
		sign = "-"
	}

	return sign + hoursString + ":" + minutesString + ":" + secondsString
}

func calculateTime(summary *models.Summary) *flexTime {
	plusTime := 0
	minusTime := 0

	if len(summary.Data) > 0 {
		plusTime = summary.Data[0].Time
	}

	if len(summary.Data) > 1 {
		minusTime = summary.Data[1].Time
	}

	totalTime := int(math.Abs(float64(plusTime - minusTime)))

	seconds := (totalTime / 1000) % 60
	minutes := (totalTime / (1000 * 60)) % 60
	hours := (totalTime / (1000 * 60 * 60)) % 24

	return &flexTime{
		Hours:    hours,
		Minutes:  minutes,
		Seconds:  seconds,
		Positive: plusTime >= minusTime,
	}

}

func main() {
	profile := &models.Profile{}
	profilePath := basePath + "/me"
	if err := makeRequest(profilePath, nil, &profile); err != nil {
		log.Fatal(err.Error())
	}

	workspaceID := profile.Data.Workspaces[0].ID
	userID := profile.Data.ID

	tags := make([]*models.Tag, 40)
	tagsPath := basePath + "/workspaces/" + strconv.Itoa(workspaceID) + "/tags"
	if err := makeRequest(tagsPath, nil, &tags); err != nil {
		log.Fatal(err.Error())
	}

	flexTags := extractFlexTags(tags)

	now := time.Now().Local()
	lastYear := time.Date(now.Year()-1, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	summary := models.Summary{}
	query := map[string]string{
		"workspace_id": strconv.Itoa(workspaceID),
		"since":        lastYear.Format("2006-01-02"),
		"until":        now.Format("2006-01-02"),
		"user_ids":     strconv.Itoa(userID),
		"tag_ids":      strconv.Itoa(flexTags.Plus.ID) + "," + strconv.Itoa(flexTags.Minus.ID),
	}

	if err := makeRequest(reportsPath+"summary.json", query, &summary); err != nil {
		log.Fatal(err.Error())
	}

	flexTime := calculateTime(&summary)

	fmt.Println(flexTime.ToString())
}
