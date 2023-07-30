package supabase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"io"
	"net/http"
	"os"
)

type supabase struct {
	URL     string
	API_KEY string
}

func SupabaseENV() supabase {
	return supabase{
		URL:     os.Getenv("SUPABASE_URL"),
		API_KEY: os.Getenv("SUPABASE_SERVICE_ROLE"),
	}
}

// ポイントを取得します
func FindPointByID(discordID string) (int, error) {
	client := &http.Client{}

	url := fmt.Sprintf(
		"%s/rest/v1/red-ginger-quiz?id=eq.%s&select=*",
		SupabaseENV().URL,
		discordID,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, errors.NewError("リクエストを作成できません", err)
	}

	req.Header.Add("apikey", SupabaseENV().API_KEY)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", SupabaseENV().API_KEY))

	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.NewError("リクエストを実行できません", err)
	}

	defer resp.Body.Close()

	type Res struct {
		ID    string `json:"id"`
		Point int    `json:"point"`
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.NewError("レスポンスを読み取れません", err)
	}

	data := make([]Res, 0)
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return 0, errors.NewError("Unmarshalに失敗しました", err)
	}

	if len(data) == 0 {
		return 0, errors.NewError("データを取得できません", err)
	}

	return data[0].Point, nil
}

// ポイントを初期化します
func InitPoint(discordID string) error {
	client := &http.Client{}

	data := struct {
		ID    string `json:"id"`
		Point int    `json:"point"`
	}{
		ID:    discordID,
		Point: 0,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.NewError("Marshalに失敗しました", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/rest/v1/red-ginger-quiz", SupabaseENV().URL),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return errors.NewError("リクエストが作成できません", err)
	}

	req.Header.Add("apikey", SupabaseENV().API_KEY)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", SupabaseENV().API_KEY))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "resolution=merge-duplicates")

	resp, err := client.Do(req)
	if err != nil {
		return errors.NewError("リクエストを実行できません", err)
	}
	defer resp.Body.Close()

	return nil
}

// ポイントを+1します
func AddPoint(discordID string) error {
	currentPoint, err := FindPointByID(discordID)
	if err != nil {
		return errors.NewError("IDでポイントを取得できません", err)
	}

	client := &http.Client{}

	data := struct {
		Point int `json:"point"`
	}{
		Point: currentPoint + 1,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.NewError("Marshalに失敗しました", err)
	}

	// ここで "id=eq.your_id" の部分を更新するidに置き換えてください
	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf(
			"%s/rest/v1/red-ginger-quiz?id=eq.%s",
			SupabaseENV().URL,
			discordID,
		),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return errors.NewError("リクエストを作成できません", err)
	}

	req.Header.Add("apikey", SupabaseENV().API_KEY)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", SupabaseENV().API_KEY))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "return=minimal")

	resp, err := client.Do(req)
	if err != nil {
		return errors.NewError("リクエストを実行できません", err)
	}
	defer resp.Body.Close()

	return nil
}
