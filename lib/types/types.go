package tailslideTypes

type Flag struct {
	FlagId int `json:"id"`
	AppId int `json:"app_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
	RolloutPercentage int `json:"rollout_percentage"`
	WhiteListedUsers string `json:"white_listed_users"`
	ErrorThresholdPercentage int `json:"error_threshold_percentage"`
	CircuitStatus string `json:"circuit_status"`
	IsRecoverable bool `json:"is_recoverable"`
	CircuitRecoveryPercentage int `json:"circuit_recovery_percentage"`
	CircuitRecoveryDelay int `json:"circuit_recovery_delay"`
	CircuitInitialRecoveryPercentage int `json:"circuit_initial_recovery_percentage"`
	CircuitRecoveryRate int `json:"circuit_recovery_rate"`
	CircuitRecoveryIncrementPercentage int `json:"circuit_recovery_increment_percentage"`
	CircuitRecoveryProfile string `json:"circuit_recovery_profile"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}



