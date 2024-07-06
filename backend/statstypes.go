package main

type Stat struct {
	Season     int         `json:"season"`
	Team       string      `json:"team"`
	Conference string      `json:"conference"`
	StatName   string      `json:"statName"`
	StatValue  interface{} `json:"statValue"`
}

type TeamStats struct {
	KickReturns           int `json:"kickReturns"`
	Games                 int `json:"games"`
	PassingTDs            int `json:"passingTDs"`
	RushingTDs            int `json:"rushingTDs"`
	InterceptionYards     int `json:"interceptionYards"`
	TacklesForLoss        int `json:"tacklesForLoss"`
	KickReturnTDs         int `json:"kickReturnTDs"`
	RushingYards          int `json:"rushingYards"`
	FourthDownConversions int `json:"fourthDownConversions"`
	PossessionTime        int `json:"possessionTime"`
	Penalties             int `json:"penalties"`
	PuntReturnYards       int `json:"puntReturnYards"`
	TotalYards            int `json:"totalYards"`
	InterceptionTDs       int `json:"interceptionTDs"`
	PuntReturnTDs         int `json:"puntReturnTDs"`
	KickReturnYards       int `json:"kickReturnYards"`
	FirstDowns            int `json:"firstDowns"`
	Sacks                 int `json:"sacks"`
	PassesIntercepted     int `json:"passesIntercepted"`
	PuntReturns           int `json:"puntReturns"`
	FumblesLost           int `json:"fumblesLost"`
	PassCompletions       int `json:"passCompletions"`
	NetPassingYards       int `json:"netPassingYards"`
	FourthDowns           int `json:"fourthDowns"`
	Turnovers             int `json:"turnovers"`
	PassAttempts          int `json:"passAttempts"`
	Interceptions         int `json:"interceptions"`
	FumblesRecovered      int `json:"fumblesRecovered"`
	ThirdDowns            int `json:"thirdDowns"`
	PenaltyYards          int `json:"penaltyYards"`
	ThirdDownConversions  int `json:"thirdDownConversions"`
	RushingAttempts       int `json:"rushingAttempts"`
}

type AdvTeamStats struct {
	Season     int    `json:"season"`
	Team       string `json:"team"`
	Conference string `json:"conference"`
	Offense    struct {
		Plays                 int     `json:"plays"`
		Drives                int     `json:"drives"`
		PPA                   float64 `json:"ppa"`
		TotalPPA              float64 `json:"totalPPA"`
		SuccessRate           float64 `json:"successRate"`
		Explosiveness         float64 `json:"explosiveness"`
		PowerSuccess          float64 `json:"powerSuccess"`
		StuffRate             float64 `json:"stuffRate"`
		LineYards             float64 `json:"lineYards"`
		LineYardsTotal        int     `json:"lineYardsTotal"`
		SecondLevelYards      float64 `json:"secondLevelYards"`
		SecondLevelYardsTotal int     `json:"secondLevelYardsTotal"`
		OpenFieldYards        float64 `json:"openFieldYards"`
		OpenFieldYardsTotal   int     `json:"openFieldYardsTotal"`
		TotalOpportunies      int     `json:"totalOpportunies"`
		PointsPerOpportunity  float64 `json:"pointsPerOpportunity"`
		FieldPosition         struct {
			AverageStart           float64 `json:"averageStart"`
			AveragePredictedPoints float64 `json:"averagePredictedPoints"`
		} `json:"fieldPosition"`
		Havoc struct {
			Total      float64 `json:"total"`
			FrontSeven float64 `json:"frontSeven"`
			DB         float64 `json:"db"`
		} `json:"havoc"`
		StandardDowns struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"standardDowns"`
		PassingDowns struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"passingDowns"`
		RushingPlays struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			TotalPPA      float64 `json:"totalPPA"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"rushingPlays"`
		PassingPlays struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			TotalPPA      float64 `json:"totalPPA"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"passingPlays"`
	} `json:"offense"`
	Defense struct {
		Plays                 int     `json:"plays"`
		Drives                int     `json:"drives"`
		PPA                   float64 `json:"ppa"`
		TotalPPA              float64 `json:"totalPPA"`
		SuccessRate           float64 `json:"successRate"`
		Explosiveness         float64 `json:"explosiveness"`
		PowerSuccess          float64 `json:"powerSuccess"`
		StuffRate             float64 `json:"stuffRate"`
		LineYards             float64 `json:"lineYards"`
		LineYardsTotal        int     `json:"lineYardsTotal"`
		SecondLevelYards      float64 `json:"secondLevelYards"`
		SecondLevelYardsTotal int     `json:"secondLevelYardsTotal"`
		OpenFieldYards        float64 `json:"openFieldYards"`
		OpenFieldYardsTotal   int     `json:"openFieldYardsTotal"`
		TotalOpportunies      int     `json:"totalOpportunies"`
		PointsPerOpportunity  float64 `json:"pointsPerOpportunity"`
		FieldPosition         struct {
			AverageStart           float64 `json:"averageStart"`
			AveragePredictedPoints float64 `json:"averagePredictedPoints"`
		} `json:"fieldPosition"`
		Havoc struct {
			Total      float64 `json:"total"`
			FrontSeven float64 `json:"frontSeven"`
			DB         float64 `json:"db"`
		} `json:"havoc"`
		StandardDowns struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"standardDowns"`
		PassingDowns struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"passingDowns"`
		RushingPlays struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			TotalPPA      float64 `json:"totalPPA"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"rushingPlays"`
		PassingPlays struct {
			Rate          float64 `json:"rate"`
			PPA           float64 `json:"ppa"`
			TotalPPA      float64 `json:"totalPPA"`
			SuccessRate   float64 `json:"successRate"`
			Explosiveness float64 `json:"explosiveness"`
		} `json:"passingPlays"`
	} `json:"defense"`
}

type Team struct {
	ID             int      `json:"id"`
	School         string   `json:"school"`
	Mascot         string   `json:"mascot"`
	Abbreviation   string   `json:"abbreviation"`
	AltName1       string   `json:"alt_name_1"`
	AltName2       string   `json:"alt_name_2"`
	AltName3       string   `json:"alt_name_3"`
	Classification string   `json:"classification"`
	Conference     string   `json:"conference"`
	Division       string   `json:"division"`
	Color          string   `json:"color"`
	AltColor       string   `json:"alt_color"`
	Logos          []string `json:"logos"`
	Twitter        string   `json:"twitter"`
	Location       struct {
		VenueID         int     `json:"venue_id"`
		Name            string  `json:"name"`
		City            string  `json:"city"`
		State           string  `json:"state"`
		Zip             string  `json:"zip"`
		CountryCode     string  `json:"country_code"`
		Timezone        string  `json:"timezone"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		Elevation       string  `json:"elevation"`
		Capacity        int     `json:"capacity"`
		YearConstructed int     `json:"year_constructed"`
		Grass           bool    `json:"grass"`
		Dome            bool    `json:"dome"`
	} `json:"location"`
}

type Ranking struct {
	Team  string  `json:team`
	Score float64 `json:score`
}
