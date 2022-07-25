package tailslideTypes

type Flag struct {
	FlagId int `json:"id"`
	AppId int `json:"app_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
	RolloutPercentage float32 `json:"rollout_percentage"`
	WhiteListedUsers string `json:"white_listed_users"`
	ErrorThresholdPercentage float32 `json:"error_threshold_percentage"`
	CircuitStatus string `json:"circuit_status"`
	IsRecoverable bool `json:"is_recoverable"`
	CircuitRecoveryPercentage float32 `json:"circuit_recovery_percentage"`
	CircuitRecoveryDelay int `json:"circuit_recovery_delay"`
	CircuitInitialRecoveryPercentage float32 `json:"circuit_initial_recovery_percentage"`
	CircuitRecoveryRate int `json:"circuit_recovery_rate"`
	CircuitRecoveryIncrementPercentage float32 `json:"circuit_recovery_increment_percentage"`
	CircuitRecoveryProfile string `json:"circuit_recovery_profile"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type FlagManagerConfig struct {
	NatsServer string
	Stream string
	SdkKey string
	AppId int
	UserContext string
	RedisHost string
	RedisPort int
}

type GetFlags func() []Flag

