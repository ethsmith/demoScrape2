package demoscrape2

type Game struct {
	//winnerID         int
	CoreID           string                  `json:"coreID,omitempty"`
	MapNum           int                     `json:"mapNum,omitempty"`
	WinnerClanName   string                  `json:"winnerClanName,omitempty"`
	Result           string                  `json:"result,omitempty"`
	Rounds           []*round                `json:"rounds,omitempty"`
	PotentialRound   *round                  `json:"potentialRound,omitempty"`
	Teams            map[string]*team        `json:"teams,omitempty"`
	Flags            flag                    `json:"flags"`
	MapName          string                  `json:"mapName,omitempty"`
	TickRate         int                     `json:"tickRate,omitempty"`
	TickLength       int                     `json:"tickLength,omitempty"`
	RoundsToWin      int                     `json:"roundsToWin,omitempty"` //30 or 16
	TotalPlayerStats map[uint64]*playerStats `json:"totalPlayerStats,omitempty"`
	CtPlayerStats    map[uint64]*playerStats `json:"ctPlayerStats,omitempty"`
	TPlayerStats     map[uint64]*playerStats `json:"TPlayerStats,omitempty"`
	TotalTeamStats   map[string]*teamStats   `json:"totalTeamStats,omitempty"`
	PlayerOrder      []uint64                `json:"playerOrder,omitempty"`
	TeamOrder        []string                `json:"teamOrder,omitempty"`
	TotalRounds      int                     `json:"totalRounds,omitempty"`
	TotalWPAlog      []*wpalog               `json:"totalWPAlog,omitempty"`
}

type flag struct {
	//all our sentinals and shit
	HasGameStarted            bool `json:"hasGameStarted,omitempty"`
	IsGameLive                bool `json:"isGameLive,omitempty"`
	IsGameOver                bool `json:"isGameOver,omitempty"`
	InRound                   bool `json:"inRound,omitempty"`
	PrePlant                  bool `json:"prePlant,omitempty"`
	PostPlant                 bool `json:"postPlant,omitempty"`
	PostWinCon                bool `json:"postWinCon,omitempty"`
	RoundIntegrityStart       int  `json:"roundIntegrityStart,omitempty"`
	RoundIntegrityEnd         int  `json:"roundIntegrityEnd,omitempty"`
	RoundIntegrityEndOfficial int  `json:"roundIntegrityEndOfficial,omitempty"`

	//for the round (gets reset on a new round) maybe should be in a new struct
	TAlive            int    `json:"TAlive,omitempty"`
	CtAlive           int    `json:"ctAlive,omitempty"`
	TMoney            bool   `json:"TMoney,omitempty"`
	TClutchVal        int    `json:"TClutchVal,omitempty"`
	CtClutchVal       int    `json:"ctClutchVal,omitempty"`
	TClutchSteam      uint64 `json:"TClutchSteam,omitempty"`
	CtClutchSteam     uint64 `json:"ctClutchSteam,omitempty"`
	OpeningKill       bool   `json:"openingKill,omitempty"`
	LastTickProcessed int    `json:"lastTickProcessed,omitempty"`
	TicksProcessed    int    `json:"ticksProcessed,omitempty"`
	DidRoundEndFire   bool   `json:"didRoundEndFire,omitempty"`
	RoundStartedAt    int    `json:"roundStartedAt,omitempty"`
	HaveInitRound     bool   `json:"haveInitRound,omitempty"`
}

type team struct {
	//id    int //meaningless?
	Name          string `json:"name,omitempty"`
	Score         int    `json:"score,omitempty"`
	ScoreAdjusted int    `json:"scoreAdjusted,omitempty"`
}

type teamStats struct {
	WinPoints      float64 `json:"winPoints,omitempty"`
	ImpactPoints   float64 `json:"impactPoints,omitempty"`
	TWinPoints     float64 `json:"TWinPoints,omitempty"`
	CtWinPoints    float64 `json:"ctWinPoints,omitempty"`
	TImpactPoints  float64 `json:"TImpactPoints,omitempty"`
	CtImpactPoints float64 `json:"ctImpactPoints,omitempty"`
	_4v5w          int     `json:"_4V5W,omitempty"`
	_4v5s          int     `json:"_4V5S,omitempty"`
	_5v4w          int     `json:"_5V4W,omitempty"`
	_5v4s          int     `json:"_5V4S,omitempty"`
	Pistols        int     `json:"pistols,omitempty"`
	PistolsW       int     `json:"pistolsW,omitempty"`
	Saves          int     `json:"saves,omitempty"`
	Clutches       int     `json:"clutches,omitempty"`
	Traded         int     `json:"traded,omitempty"`
	Fass           int     `json:"fass,omitempty"`
	Ef             int     `json:"ef,omitempty"`
	Ud             int     `json:"ud,omitempty"`
	Util           int     `json:"util,omitempty"`
	CtR            int     `json:"ctR,omitempty"`
	CtRW           int     `json:"ctRW,omitempty"`
	TR             int     `json:"TR,omitempty"`
	TRW            int     `json:"TRW,omitempty"`
	Deaths         int     `json:"deaths,omitempty"`

	//kinda garbo
	Normalizer int `json:"normalizer,omitempty"`
}

type round struct {
	//round value
	RoundNum            int8                    `json:"roundNum,omitempty"`
	StartingTick        int                     `json:"startingTick,omitempty"`
	EndingTick          int                     `json:"endingTick,omitempty"`
	PlayerStats         map[uint64]*playerStats `json:"playerStats,omitempty"`
	TeamStats           map[string]*teamStats   `json:"teamStats,omitempty"`
	InitTerroristCount  int                     `json:"initTerroristCount,omitempty"`
	InitCTerroristCount int                     `json:"initCTerroristCount,omitempty"`
	WinnerClanName      string                  `json:"winnerClanName,omitempty"`
	//winnerID            int //this is the unique ID which should not change BUT IT DOES
	WinnerENUM         int     `json:"winnerENUM,omitempty"` //this effectively represents the side that won: 2 (T) or 3 (CT)
	IntegrityCheck     bool    `json:"integrityCheck,omitempty"`
	Planter            uint64  `json:"planter,omitempty"`
	Defuser            uint64  `json:"defuser,omitempty"`
	EndDueToBombEvent  bool    `json:"endDueToBombEvent,omitempty"`
	WinTeamDmg         int     `json:"winTeamDmg,omitempty"`
	ServerNormalizer   int     `json:"serverNormalizer,omitempty"`
	ServerImpactPoints float64 `json:"serverImpactPoints,omitempty"`
	KnifeRound         bool    `json:"knifeRound,omitempty"`

	WPAlog        []*wpalog `json:"WPAlog,omitempty"`
	BombStartTick int       `json:"bombStartTick,omitempty"`
}

type wpalog struct {
	Round               int `json:"round,omitempty"`
	Tick                int `json:"tick,omitempty"`
	Clock               int `json:"clock,omitempty"`
	Planted             int `json:"planted,omitempty"`
	CtAlive             int `json:"ctAlive,omitempty"`
	TAlive              int `json:"TAlive,omitempty"`
	CtEquipVal          int `json:"ctEquipVal,omitempty"`
	TEquipVal           int `json:"TEquipVal,omitempty"`
	CtFlashes           int `json:"ctFlashes,omitempty"`
	CtSmokes            int `json:"ctSmokes,omitempty"`
	CtMolys             int `json:"ctMolys,omitempty"`
	CtFrags             int `json:"ctFrags,omitempty"`
	TFlashes            int `json:"TFlashes,omitempty"`
	TSmokes             int `json:"TSmokes,omitempty"`
	TMolys              int `json:"TMolys,omitempty"`
	TFrags              int `json:"TFrags,omitempty"`
	ClosestCTDisttoBomb int `json:"closestCTDisttoBomb,omitempty"`
	Kits                int `json:"kits,omitempty"`
	CtArmor             int `json:"ctArmor,omitempty"`
	TArmor              int `json:"TArmor,omitempty"`
	Winner              int `json:"winner,omitempty"`
}

type playerStats struct {
	Name    string `json:"name,omitempty"`
	SteamID uint64 `json:"steamID,omitempty"`
	IsBot   bool   `json:"isBot,omitempty"`
	//teamID  int
	TeamENUM     int    `json:"teamENUM,omitempty"`
	TeamClanName string `json:"teamClanName,omitempty"`
	Side         int    `json:"side,omitempty"`
	Rounds       int    `json:"rounds,omitempty"`
	//playerPoints float32
	//teamPoints float32
	Damage              int     `json:"damage,omitempty"`
	Kills               uint8   `json:"kills,omitempty"`
	Assists             uint8   `json:"assists,omitempty"`
	Deaths              uint8   `json:"deaths,omitempty"`
	DeathTick           int     `json:"deathTick,omitempty"`
	DeathPlacement      float64 `json:"deathPlacement,omitempty"`
	TicksAlive          int     `json:"ticksAlive,omitempty"`
	Trades              int     `json:"trades,omitempty"`
	Traded              int     `json:"traded,omitempty"`
	Ok                  int     `json:"ok,omitempty"`
	Ol                  int     `json:"ol,omitempty"`
	Cl_1                int     `json:"cl_1,omitempty"`
	Cl_2                int     `json:"cl_2,omitempty"`
	Cl_3                int     `json:"cl_3,omitempty"`
	Cl_4                int     `json:"cl_4,omitempty"`
	Cl_5                int     `json:"cl_5,omitempty"`
	_2k                 int     `json:"_2K,omitempty"`
	_3k                 int     `json:"_3K,omitempty"`
	_4k                 int     `json:"_4K,omitempty"`
	_5k                 int     `json:"_5K,omitempty"`
	NadeDmg             int     `json:"nadeDmg,omitempty"`
	InfernoDmg          int     `json:"infernoDmg,omitempty"`
	UtilDmg             int     `json:"utilDmg,omitempty"`
	Ef                  int     `json:"ef,omitempty"`
	FAss                int     `json:"FAss,omitempty"`
	EnemyFlashTime      float64 `json:"enemyFlashTime,omitempty"`
	Hs                  int     `json:"hs,omitempty"`
	KastRounds          float64 `json:"kastRounds,omitempty"`
	Saves               int     `json:"saves,omitempty"`
	Entries             int     `json:"entries,omitempty"`
	KillPoints          float64 `json:"killPoints,omitempty"`
	ImpactPoints        float64 `json:"impactPoints,omitempty"`
	WinPoints           float64 `json:"winPoints,omitempty"`
	AwpKills            int     `json:"awpKills,omitempty"`
	RF                  int     `json:"RF,omitempty"`
	RA                  int     `json:"RA,omitempty"`
	NadesThrown         int     `json:"nadesThrown,omitempty"`
	FiresThrown         int     `json:"firesThrown,omitempty"`
	FlashThrown         int     `json:"flashThrown,omitempty"`
	SmokeThrown         int     `json:"smokeThrown,omitempty"`
	DamageTaken         int     `json:"damageTaken,omitempty"`
	SuppRounds          int     `json:"suppRounds,omitempty"`
	SuppDamage          int     `json:"suppDamage,omitempty"`
	LurkerBlips         int     `json:"lurkerBlips,omitempty"`
	DistanceToTeammates int     `json:"distanceToTeammates,omitempty"`
	LurkRounds          int     `json:"lurkRounds,omitempty"`
	Wlp                 float64 `json:"wlp,omitempty"`
	Mip                 float64 `json:"mip,omitempty"`
	Rws                 float64 `json:"rws,omitempty"` //round win shares
	Eac                 int     `json:"eac,omitempty"` //effective assist contributions

	Rwk int `json:"rwk,omitempty"` //rounds with Kills

	//derived
	UtilThrown   int     `json:"utilThrown,omitempty"`
	Atd          int     `json:"atd,omitempty"`
	Kast         float64 `json:"kast,omitempty"`
	KillPointAvg float64 `json:"killPointAvg,omitempty"`
	Iiwr         float64 `json:"iiwr,omitempty"`
	Adr          float64 `json:"adr,omitempty"`
	DrDiff       float64 `json:"drDiff,omitempty"`
	KR           float64 `json:"KR,omitempty"`
	Tr           float64 `json:"tr,omitempty"` //trade ratio
	ImpactRating float64 `json:"impactRating,omitempty"`
	Rating       float64 `json:"rating,omitempty"`

	//side specific
	TDamage               int     `json:"TDamage,omitempty"`
	CtDamage              int     `json:"ctDamage,omitempty"`
	TImpactPoints         float64 `json:"TImpactPoints,omitempty"`
	TWinPoints            float64 `json:"TWinPoints,omitempty"`
	TOK                   int     `json:"TOK,omitempty"`
	TOL                   int     `json:"TOL,omitempty"`
	CtImpactPoints        float64 `json:"ctImpactPoints,omitempty"`
	CtWinPoints           float64 `json:"ctWinPoints,omitempty"`
	CtOK                  int     `json:"ctOK,omitempty"`
	CtOL                  int     `json:"ctOL,omitempty"`
	TKills                uint8   `json:"TKills,omitempty"`
	TDeaths               uint8   `json:"TDeaths,omitempty"`
	TKAST                 float64 `json:"TKAST,omitempty"`
	TKASTRounds           float64 `json:"TKASTRounds,omitempty"`
	TADR                  float64 `json:"TADR,omitempty"`
	CtKills               uint8   `json:"ctKills,omitempty"`
	CtDeaths              uint8   `json:"ctDeaths,omitempty"`
	CtKAST                float64 `json:"ctKAST,omitempty"`
	CtKASTRounds          float64 `json:"ctKASTRounds,omitempty"`
	CtADR                 float64 `json:"ctADR,omitempty"`
	TTeamsWinPoints       float64 `json:"TTeamsWinPoints,omitempty"`
	CtTeamsWinPoints      float64 `json:"ctTeamsWinPoints,omitempty"`
	TWinPointsNormalizer  int     `json:"TWinPointsNormalizer,omitempty"`
	CtWinPointsNormalizer int     `json:"ctWinPointsNormalizer,omitempty"`
	TRounds               int     `json:"TRounds,omitempty"`
	CtRounds              int     `json:"ctRounds,omitempty"`
	CtRating              float64 `json:"ctRating,omitempty"`
	CtImpactRating        float64 `json:"ctImpactRating,omitempty"`
	TRating               float64 `json:"TRating,omitempty"`
	TImpactRating         float64 `json:"TImpactRating,omitempty"`
	TADP                  float64 `json:"TADP,omitempty"`
	CtADP                 float64 `json:"ctADP,omitempty"`

	TRF   int `json:"TRF,omitempty"`
	CtAWP int `json:"ctAWP,omitempty"`

	//kinda garbo
	TeamsWinPoints      float64 `json:"teamsWinPoints,omitempty"`
	WinPointsNormalizer int     `json:"winPointsNormalizer,omitempty"`

	//"flags"
	Health             int            `json:"health,omitempty"`
	TradeList          map[uint64]int `json:"tradeList,omitempty"`
	MostRecentFlasher  uint64         `json:"mostRecentFlasher,omitempty"`
	MostRecentFlashVal float64        `json:"mostRecentFlashVal,omitempty"`
	DamageList         map[uint64]int `json:"damageList,omitempty"`
}

type Accolades struct {
	Awp        int `json:"awp,omitempty"`
	Deagle     int `json:"deagle,omitempty"`
	Knife      int `json:"knife,omitempty"`
	Dinks      int `json:"dinks,omitempty"`
	BlindKills int `json:"blindKills,omitempty"`
	BombPlants int `json:"bombPlants,omitempty"`
	Jumps      int `json:"jumps,omitempty"`
	TeamDMG    int `json:"teamDMG,omitempty"`
	SelfDMG    int `json:"selfDMG,omitempty"`
	Ping       int `json:"ping,omitempty"`
	PingPoints int `json:"pingPoints,omitempty"`
	//footsteps         int //unnecessary processing?
	BombTaps          int `json:"bombTaps,omitempty"`
	KillsThroughSmoke int `json:"killsThroughSmoke,omitempty"`
	Penetrations      int `json:"penetrations,omitempty"`
	NoScopes          int `json:"noScopes,omitempty"`
	MidairKills       int `json:"midairKills,omitempty"`
	CrouchedKills     int `json:"crouchedKills,omitempty"`
	BombzoneKills     int `json:"bombzoneKills,omitempty"`
	KillsWhileMoving  int `json:"killsWhileMoving,omitempty"`
	MostMoneySpent    int `json:"mostMoneySpent,omitempty"`
	MostShotsOnLegs   int `json:"mostShotsOnLegs,omitempty"`
	ShotsFired        int `json:"shotsFired,omitempty"`
	Ak                int `json:"ak,omitempty"`
	M4                int `json:"m4,omitempty"`
	Pistol            int `json:"pistol,omitempty"`
	Scout             int `json:"scout,omitempty"`
}
