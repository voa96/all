package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Camera struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Status      string    `json:"status" gorm:"default:активна"`
	InstalledAt time.Time `json:"installed_at" gorm:"autoCreateTime"`
	Events      []Event   `json:"events,omitempty" gorm:"foreignKey:CameraID"`
}

type Event struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CameraID   uint      `json:"camera_id"`
	Camera     Camera    `json:"-" gorm:"foreignKey:CameraID"`
	Timestamp  time.Time `json:"timestamp" gorm:"autoCreateTime"`
	ArmedCount int       `json:"armed_count"`
	TotalCount int       `json:"total_count"`
	RiskLevel  string    `json:"risk_level" gorm:"default:низкий"`
}

type DashboardStats struct {
	TotalCameras  int64   `json:"total_cameras"`
	ActiveCameras int64   `json:"active_cameras"`
	TotalEvents   int64   `json:"total_events"`
	TotalArmed    int64   `json:"total_armed"`
	AvgRisk       float64 `json:"avg_risk"`
}

var db *gorm.DB

// ===== MAIN =====

func main() {
	initDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// ===== СНАЧАЛА API =====
	http.HandleFunc("/api/cameras", getCameras)
	http.HandleFunc("/api/events", getEvents)
	http.HandleFunc("/api/dashboard", getDashboardStats)

	// ===== ПОТОМ СТАТИКА =====
	fs := http.FileServer(http.Dir("./сайт"))
	http.Handle("/", fs)

	log.Println("🚀 Сервер запущен на http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// ===== БД =====

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	db.AutoMigrate(&Camera{}, &Event{})
	fillTestData()
}

func fillTestData() {
	var count int64
	db.Model(&Camera{}).Count(&count)
	if count > 0 {
		return
	}

	cameras := []Camera{
		{Name: "Вход в ТЦ", Location: "ул. Ленина, 15", Status: "активна"},
		{Name: "Парковка", Location: "ул. Ленина, 15", Status: "активна"},
		{Name: "Вход в офис", Location: "пр. Мира, 42", Status: "неактивна"},
	}
	db.Create(&cameras)

	now := time.Now()
	events := []Event{
		{CameraID: 1, Timestamp: now.Add(-2 * time.Hour), ArmedCount: 2, TotalCount: 14, RiskLevel: "высокий"},
		{CameraID: 2, Timestamp: now.Add(-3 * time.Hour), ArmedCount: 0, TotalCount: 8, RiskLevel: "низкий"},
		{CameraID: 1, Timestamp: now.Add(-4 * time.Hour), ArmedCount: 1, TotalCount: 11, RiskLevel: "средний"},
		{CameraID: 2, Timestamp: now.Add(-5 * time.Hour), ArmedCount: 0, TotalCount: 6, RiskLevel: "низкий"},
		{CameraID: 1, Timestamp: now.Add(-6 * time.Hour), ArmedCount: 3, TotalCount: 9, RiskLevel: "высокий"},
		{CameraID: 1, Timestamp: now.Add(-7 * time.Hour), ArmedCount: 0, TotalCount: 15, RiskLevel: "низкий"},
		{CameraID: 2, Timestamp: now.Add(-8 * time.Hour), ArmedCount: 1, TotalCount: 5, RiskLevel: "средний"},
		{CameraID: 1, Timestamp: now.Add(-24 * time.Hour), ArmedCount: 2, TotalCount: 20, RiskLevel: "высокий"},
		{CameraID: 1, Timestamp: now.Add(-26 * time.Hour), ArmedCount: 0, TotalCount: 12, RiskLevel: "низкий"},
	}
	db.Create(&events)
}

// ===== API =====

func getCameras(w http.ResponseWriter, r *http.Request) {
	var cameras []Camera
	db.Find(&cameras)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cameras)
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	var events []Event
	db.Preload("Camera").Order("timestamp DESC").Limit(20).Find(&events)

	type EventWithCameraName struct {
		ID         uint      `json:"id"`
		CameraID   uint      `json:"camera_id"`
		CameraName string    `json:"camera_name"`
		Timestamp  time.Time `json:"timestamp"`
		ArmedCount int       `json:"armed_count"`
		TotalCount int       `json:"total_count"`
		RiskLevel  string    `json:"risk_level"`
	}

	result := make([]EventWithCameraName, len(events))
	for i, e := range events {
		result[i] = EventWithCameraName{
			ID:         e.ID,
			CameraID:   e.CameraID,
			CameraName: e.Camera.Name,
			Timestamp:  e.Timestamp,
			ArmedCount: e.ArmedCount,
			TotalCount: e.TotalCount,
			RiskLevel:  e.RiskLevel,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getDashboardStats(w http.ResponseWriter, r *http.Request) {
	var stats DashboardStats

	db.Model(&Camera{}).Count(&stats.TotalCameras)
	db.Model(&Camera{}).Where("status = ?", "активна").Count(&stats.ActiveCameras)
	db.Model(&Event{}).Count(&stats.TotalEvents)
	db.Model(&Event{}).Select("COALESCE(SUM(armed_count), 0)").Scan(&stats.TotalArmed)

	var riskResult struct{ Avg float64 }
	db.Model(&Event{}).Select("COALESCE(AVG(CASE risk_level WHEN 'низкий' THEN 1 WHEN 'средний' THEN 2 WHEN 'высокий' THEN 3 ELSE 0 END), 0) as avg").Scan(&riskResult)
	stats.AvgRisk = riskResult.Avg

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
