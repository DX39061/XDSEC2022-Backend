package game_bind

type GameDataResponse struct {
	ID    uint    `mapstructure:"id" json:"id"`
	Score float64 `mapstructure:"score" json:"score"`
}
