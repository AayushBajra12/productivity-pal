package models

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	Preferences   Preferences   `json:"preferences"`
	HealthDetails HealthDetails `json:"health_details"`
}

type Preferences struct {
	HealthAdvice  bool `json:"health_advice"`
	FinanceAdvice bool `json:"finance_advice"`
	DailyTasks    bool `json:"daily_tasks"`
}

type HealthDetails struct {
	HealthID      int64   `json:"health_id"`
	Age           int     `json:"age"`
	Height        float64 `json:"height"`
	Weight        float64 `json:"weight"`
	ActivityLevel string  `json:"activity_level"`
	HeartRate     int     `json:"heart_rate"`
	GoalWeight    float64 `json:"goal_weight"`
	GeneralGoal   string  `json:"general_goal"`
}
