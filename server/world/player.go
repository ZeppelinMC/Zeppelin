package world

type playerData struct {
	Invulnerable        int8          `nbt:"Invulnerable"`
	FoodSaturationLevel float32       `nbt:"foodSaturationLevel"`
	UUID                []int32       `nbt:"UUID"`
	EnderItems          []interface{} `nbt:"EnderItems"`
	DataVersion         int32         `nbt:"DataVersion"`
	SelectedItemSlot    int32         `nbt:"SelectedItemSlot"`
	SleepTimer          int16         `nbt:"SleepTimer"`
	Abilities           struct {
		Mayfly       int8    `nbt:"mayfly"`
		Instabuild   int8    `nbt:"instabuild"`
		WalkSpeed    float32 `nbt:"walkSpeed"`
		MayBuild     int8    `nbt:"mayBuild"`
		Flying       int8    `nbt:"flying"`
		FlySpeed     float32 `nbt:"flySpeed"`
		Invulnerable int8    `nbt:"invulnerable"`
	} `nbt:"abilities"`
	RecipeBook struct {
		IsGuiOpen                           int8     `nbt:"isGuiOpen"`
		IsBlastingFurnaceGuiOpen            int8     `nbt:"isBlastingFurnaceGuiOpen"`
		IsSmokerGuiOpen                     int8     `nbt:"isSmokerGuiOpen"`
		IsBlastingFurnaceFilteringCraftable int8     `nbt:"isBlastingFurnaceFilteringCraftable"`
		IsFilteringCraftable                int8     `nbt:"isFilteringCraftable"`
		ToBeDisplayed                       []string `nbt:"toBeDisplayed"`
		IsFurnaceGuiOpen                    int8     `nbt:"isFurnaceGuiOpen"`
		IsFurnaceFilteringCraftable         int8     `nbt:"isFurnaceFilteringCraftable"`
		IsSmokerFilteringCraftable          int8     `nbt:"isSmokerFilteringCraftable"`
		Recipes                             []string `nbt:"recipes"`
	} `nbt:"recipeBook"`
	XpSeed          int32         `nbt:"XpSeed"`
	Inventory       []interface{} `nbt:"Inventory"`
	FoodLevel       int32         `nbt:"foodLevel"`
	HurtByTimestamp int32         `nbt:"HurtByTimestamp"`
	FallDistance    float32       `nbt:"FallDistance"`
	PlayerGameType  int32         `nbt:"playerGameType"`
	SeenCredits     int8          `nbt:"seenCredits"`
	Pos             []float64     `nbt:"Pos"`
	FoodTickTimer   int32         `nbt:"foodTickTimer"`
	Brain           struct {
		Memories struct {
		} `nbt:"memories"`
	} `nbt:"Brain"`
	AbsorptionAmount    float32   `nbt:"AbsorptionAmount"`
	DeathTime           int16     `nbt:"DeathTime"`
	XpLevel             int32     `nbt:"XpLevel"`
	XpP                 float32   `nbt:"XpP"`
	FallFlying          int8      `nbt:"FallFlying"`
	Motion              []float64 `nbt:"Motion"`
	OnGround            int8      `nbt:"OnGround"`
	Rotation            []float32 `nbt:"Rotation"`
	Score               int32     `nbt:"Score"`
	Fire                int16     `nbt:"Fire"`
	FoodExhaustionLevel float32   `nbt:"foodExhaustionLevel"`
	Attributes          []struct {
		Name string  `nbt:"Name"`
		Base float64 `nbt:"Base"`
	} `nbt:"Attributes"`
	PortalCooldown     int32   `nbt:"PortalCooldown"`
	Health             float32 `nbt:"Health"`
	Dimension          string  `nbt:"Dimension"`
	XpTotal            int32   `nbt:"XpTotal"`
	Air                int16   `nbt:"Air"`
	WardenSpawnTracker struct {
		WarningLevel          int32 `nbt:"warning_level"`
		TicksSinceLastWarning int32 `nbt:"ticks_since_last_warning"`
		CooldownTicks         int32 `nbt:"cooldown_ticks"`
	} `nbt:"warden_spawn_tracker"`
	HurtTime int16 `nbt:"HurtTime"`
}
