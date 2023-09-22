package db

import (
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const DBName = "red-ginger.db"

var DB *gorm.DB

// レースのポイントです
type Race struct {
	ID    string `gorm:"primaryKey"`
	Point int    `gorm:"column:point"`
}

func init() {
	d, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	DB = d
}

// DBに接続します
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DBName))
	if err != nil {
		return db, errors.NewError("DBをOPENできません", err)
	}

	// Raceテーブル
	if err = db.AutoMigrate(&Race{}); err != nil {
		return db, errors.NewError("テーブルを作成できません", err)
	}

	return db, nil
}

// Upsertします
//
// 既存のポイントに新しいポイントを加算します。
func (r *Race) Upsert() error {
	tx := DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"point": gorm.Expr("point + ?", r.Point),
		}),
	}).Create(&r)

	if tx.Error != nil {
		return errors.NewError("Upsertに失敗しました", tx.Error)
	}

	return nil
}

// 指定のIDのレコードを削除します
func (r Race) DeleteByID() error {
	if r.ID == "" {
		return errors.NewError("IDが空です")
	}

	// 削除操作
	if err := DB.Delete(&Race{}, "id = ?", r.ID).Error; err != nil {
		return errors.NewError("削除に失敗しました", err)
	}

	return nil
}

// 全てのレコードを削除します
func (r Race) DeleteAll() error {
	// 削除操作
	if err := DB.Exec("DELETE FROM races").Error; err != nil {
		return errors.NewError("全てのレコードの削除に失敗しました", err)
	}

	return nil
}

// 取得します
func (r Race) FindByID(id string) (Race, error) {
	var race Race
	if err := DB.First(&race, "id = ?", id).Error; err != nil {
		return race, errors.NewError("Failed to find record", err)
	}

	return race, nil
}

// 全てのデータを取得します
func (r Race) FindAll() ([]Race, error) {
	var races []Race
	if err := DB.Order("point DESC").Find(&races).Error; err != nil {
		return nil, errors.NewError("データを取得できません", err)
	}

	return races, nil
}
