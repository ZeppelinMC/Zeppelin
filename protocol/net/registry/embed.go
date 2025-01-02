package registry

var Registries = map[string]interface{}{"minecraft:banner_pattern": map[string]struct {
	AssetId        string "nbt:\"asset_id\""
	TranslationKey string "nbt:\"translation_key\""
}{"minecraft:base": {AssetId: "minecraft:base", TranslationKey: "block.minecraft.banner.base"}, "minecraft:border": {AssetId: "minecraft:border", TranslationKey: "block.minecraft.banner.border"}, "minecraft:bricks": {AssetId: "minecraft:bricks", TranslationKey: "block.minecraft.banner.bricks"}, "minecraft:circle": {AssetId: "minecraft:circle", TranslationKey: "block.minecraft.banner.circle"}, "minecraft:creeper": {AssetId: "minecraft:creeper", TranslationKey: "block.minecraft.banner.creeper"}, "minecraft:cross": {AssetId: "minecraft:cross", TranslationKey: "block.minecraft.banner.cross"}, "minecraft:curly_border": {AssetId: "minecraft:curly_border", TranslationKey: "block.minecraft.banner.curly_border"}, "minecraft:diagonal_left": {AssetId: "minecraft:diagonal_left", TranslationKey: "block.minecraft.banner.diagonal_left"}, "minecraft:diagonal_right": {AssetId: "minecraft:diagonal_right", TranslationKey: "block.minecraft.banner.diagonal_right"}, "minecraft:diagonal_up_left": {AssetId: "minecraft:diagonal_up_left", TranslationKey: "block.minecraft.banner.diagonal_up_left"}, "minecraft:diagonal_up_right": {AssetId: "minecraft:diagonal_up_right", TranslationKey: "block.minecraft.banner.diagonal_up_right"}, "minecraft:flow": {AssetId: "minecraft:flow", TranslationKey: "block.minecraft.banner.flow"}, "minecraft:flower": {AssetId: "minecraft:flower", TranslationKey: "block.minecraft.banner.flower"}, "minecraft:globe": {AssetId: "minecraft:globe", TranslationKey: "block.minecraft.banner.globe"}, "minecraft:gradient": {AssetId: "minecraft:gradient", TranslationKey: "block.minecraft.banner.gradient"}, "minecraft:gradient_up": {AssetId: "minecraft:gradient_up", TranslationKey: "block.minecraft.banner.gradient_up"}, "minecraft:guster": {AssetId: "minecraft:guster", TranslationKey: "block.minecraft.banner.guster"}, "minecraft:half_horizontal": {AssetId: "minecraft:half_horizontal", TranslationKey: "block.minecraft.banner.half_horizontal"}, "minecraft:half_horizontal_bottom": {AssetId: "minecraft:half_horizontal_bottom", TranslationKey: "block.minecraft.banner.half_horizontal_bottom"}, "minecraft:half_vertical": {AssetId: "minecraft:half_vertical", TranslationKey: "block.minecraft.banner.half_vertical"}, "minecraft:half_vertical_right": {AssetId: "minecraft:half_vertical_right", TranslationKey: "block.minecraft.banner.half_vertical_right"}, "minecraft:mojang": {AssetId: "minecraft:mojang", TranslationKey: "block.minecraft.banner.mojang"}, "minecraft:piglin": {AssetId: "minecraft:piglin", TranslationKey: "block.minecraft.banner.piglin"}, "minecraft:rhombus": {AssetId: "minecraft:rhombus", TranslationKey: "block.minecraft.banner.rhombus"}, "minecraft:skull": {AssetId: "minecraft:skull", TranslationKey: "block.minecraft.banner.skull"}, "minecraft:small_stripes": {AssetId: "minecraft:small_stripes", TranslationKey: "block.minecraft.banner.small_stripes"}, "minecraft:square_bottom_left": {AssetId: "minecraft:square_bottom_left", TranslationKey: "block.minecraft.banner.square_bottom_left"}, "minecraft:square_bottom_right": {AssetId: "minecraft:square_bottom_right", TranslationKey: "block.minecraft.banner.square_bottom_right"}, "minecraft:square_top_left": {AssetId: "minecraft:square_top_left", TranslationKey: "block.minecraft.banner.square_top_left"}, "minecraft:square_top_right": {AssetId: "minecraft:square_top_right", TranslationKey: "block.minecraft.banner.square_top_right"}, "minecraft:straight_cross": {AssetId: "minecraft:straight_cross", TranslationKey: "block.minecraft.banner.straight_cross"}, "minecraft:stripe_bottom": {AssetId: "minecraft:stripe_bottom", TranslationKey: "block.minecraft.banner.stripe_bottom"}, "minecraft:stripe_center": {AssetId: "minecraft:stripe_center", TranslationKey: "block.minecraft.banner.stripe_center"}, "minecraft:stripe_downleft": {AssetId: "minecraft:stripe_downleft", TranslationKey: "block.minecraft.banner.stripe_downleft"}, "minecraft:stripe_downright": {AssetId: "minecraft:stripe_downright", TranslationKey: "block.minecraft.banner.stripe_downright"}, "minecraft:stripe_left": {AssetId: "minecraft:stripe_left", TranslationKey: "block.minecraft.banner.stripe_left"}, "minecraft:stripe_middle": {AssetId: "minecraft:stripe_middle", TranslationKey: "block.minecraft.banner.stripe_middle"}, "minecraft:stripe_right": {AssetId: "minecraft:stripe_right", TranslationKey: "block.minecraft.banner.stripe_right"}, "minecraft:stripe_top": {AssetId: "minecraft:stripe_top", TranslationKey: "block.minecraft.banner.stripe_top"}, "minecraft:triangle_bottom": {AssetId: "minecraft:triangle_bottom", TranslationKey: "block.minecraft.banner.triangle_bottom"}, "minecraft:triangle_top": {AssetId: "minecraft:triangle_top", TranslationKey: "block.minecraft.banner.triangle_top"}, "minecraft:triangles_bottom": {AssetId: "minecraft:triangles_bottom", TranslationKey: "block.minecraft.banner.triangles_bottom"}, "minecraft:triangles_top": {AssetId: "minecraft:triangles_top", TranslationKey: "block.minecraft.banner.triangles_top"}}, "minecraft:chat_type": map[string]ChatType{"minecraft:chat": {Chat: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
	Style          struct {
		Color  string "nbt:\"color\""
		Italic bool   "nbt:\"italic\""
	} "nbt:\"style,omitempty\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.text", Style: struct {
	Color  string "nbt:\"color\""
	Italic bool   "nbt:\"italic\""
}{Color: "", Italic: false}}, Narration: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.text.narrate"}}, "minecraft:emote_command": {Chat: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
	Style          struct {
		Color  string "nbt:\"color\""
		Italic bool   "nbt:\"italic\""
	} "nbt:\"style,omitempty\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.emote", Style: struct {
	Color  string "nbt:\"color\""
	Italic bool   "nbt:\"italic\""
}{Color: "", Italic: false}}, Narration: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.emote"}}, "minecraft:msg_command_incoming": {Chat: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
	Style          struct {
		Color  string "nbt:\"color\""
		Italic bool   "nbt:\"italic\""
	} "nbt:\"style,omitempty\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "commands.message.display.incoming", Style: struct {
	Color  string "nbt:\"color\""
	Italic bool   "nbt:\"italic\""
}{Color: "", Italic: false}}, Narration: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.text.narrate"}}, "minecraft:msg_command_outgoing": {Chat: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
	Style          struct {
		Color  string "nbt:\"color\""
		Italic bool   "nbt:\"italic\""
	} "nbt:\"style,omitempty\""
}{Parameters: []string{"target", "content"}, TranslationKey: "commands.message.display.outgoing", Style: struct {
	Color  string "nbt:\"color\""
	Italic bool   "nbt:\"italic\""
}{Color: "", Italic: false}}, Narration: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.text.narrate"}}, "minecraft:say_command": {Chat: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
	Style          struct {
		Color  string "nbt:\"color\""
		Italic bool   "nbt:\"italic\""
	} "nbt:\"style,omitempty\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.announcement", Style: struct {
	Color  string "nbt:\"color\""
	Italic bool   "nbt:\"italic\""
}{Color: "", Italic: false}}, Narration: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.text.narrate"}}, "minecraft:team_msg_command_incoming": {Chat: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
	Style          struct {
		Color  string "nbt:\"color\""
		Italic bool   "nbt:\"italic\""
	} "nbt:\"style,omitempty\""
}{Parameters: []string{"target", "sender", "content"}, TranslationKey: "chat.type.team.text", Style: struct {
	Color  string "nbt:\"color\""
	Italic bool   "nbt:\"italic\""
}{Color: "", Italic: false}}, Narration: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.text.narrate"}}, "minecraft:team_msg_command_outgoing": {Chat: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
	Style          struct {
		Color  string "nbt:\"color\""
		Italic bool   "nbt:\"italic\""
	} "nbt:\"style,omitempty\""
}{Parameters: []string{"target", "sender", "content"}, TranslationKey: "chat.type.team.sent", Style: struct {
	Color  string "nbt:\"color\""
	Italic bool   "nbt:\"italic\""
}{Color: "", Italic: false}}, Narration: struct {
	Parameters     []string "nbt:\"parameters\""
	TranslationKey string   "nbt:\"translation_key\""
}{Parameters: []string{"sender", "content"}, TranslationKey: "chat.type.text.narrate"}}}, "minecraft:damage_type": map[string]struct {
	Exhaustion       float32 "nbt:\"exhaustion\""
	MessageID        string  "nbt:\"message_id\""
	Scaling          string  "nbt:\"scaling\""
	DeathMessageType string  "nbt:\"death_message_type,omitempty\""
	Effects          string  "nbt:\"effects,omitempty\""
}{"minecraft:arrow": {Exhaustion: 0.1, MessageID: "arrow", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:bad_respawn_point": {Exhaustion: 0.1, MessageID: "badRespawnPoint", Scaling: "always", DeathMessageType: "", Effects: ""}, "minecraft:cactus": {Exhaustion: 0.1, MessageID: "cactus", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:campfire": {Exhaustion: 0.1, MessageID: "inFire", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:cramming": {Exhaustion: 0, MessageID: "cramming", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:dragon_breath": {Exhaustion: 0, MessageID: "dragonBreath", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:drown": {Exhaustion: 0, MessageID: "drown", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:dry_out": {Exhaustion: 0.1, MessageID: "dryout", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:explosion": {Exhaustion: 0.1, MessageID: "explosion", Scaling: "always", DeathMessageType: "", Effects: ""}, "minecraft:fall": {Exhaustion: 0, MessageID: "fall", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:falling_anvil": {Exhaustion: 0.1, MessageID: "anvil", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:falling_block": {Exhaustion: 0.1, MessageID: "fallingBlock", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:falling_stalactite": {Exhaustion: 0.1, MessageID: "fallingStalactite", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:fireball": {Exhaustion: 0.1, MessageID: "fireball", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:fireworks": {Exhaustion: 0.1, MessageID: "fireworks", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:fly_into_wall": {Exhaustion: 0, MessageID: "flyIntoWall", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:freeze": {Exhaustion: 0, MessageID: "freeze", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:generic": {Exhaustion: 0, MessageID: "generic", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:generic_kill": {Exhaustion: 0, MessageID: "genericKill", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:hot_floor": {Exhaustion: 0.1, MessageID: "hotFloor", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:in_fire": {Exhaustion: 0.1, MessageID: "inFire", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:in_wall": {Exhaustion: 0, MessageID: "inWall", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:indirect_magic": {Exhaustion: 0, MessageID: "indirectMagic", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:lava": {Exhaustion: 0.1, MessageID: "lava", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:lightning_bolt": {Exhaustion: 0.1, MessageID: "lightningBolt", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:magic": {Exhaustion: 0, MessageID: "magic", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:mob_attack": {Exhaustion: 0.1, MessageID: "mob", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:mob_attack_no_aggro": {Exhaustion: 0.1, MessageID: "mob", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:mob_projectile": {Exhaustion: 0.1, MessageID: "mob", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:on_fire": {Exhaustion: 0, MessageID: "onFire", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:out_of_world": {Exhaustion: 0, MessageID: "outOfWorld", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:outside_border": {Exhaustion: 0, MessageID: "outsideBorder", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:player_attack": {Exhaustion: 0.1, MessageID: "player", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:player_explosion": {Exhaustion: 0.1, MessageID: "explosion.player", Scaling: "always", DeathMessageType: "", Effects: ""}, "minecraft:sonic_boom": {Exhaustion: 0, MessageID: "sonic_boom", Scaling: "always", DeathMessageType: "", Effects: ""}, "minecraft:spit": {Exhaustion: 0.1, MessageID: "mob", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:stalagmite": {Exhaustion: 0, MessageID: "stalagmite", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:starve": {Exhaustion: 0, MessageID: "starve", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:sting": {Exhaustion: 0.1, MessageID: "sting", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:sweet_berry_bush": {Exhaustion: 0.1, MessageID: "sweetBerryBush", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:thorns": {Exhaustion: 0.1, MessageID: "thorns", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:thrown": {Exhaustion: 0.1, MessageID: "thrown", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:trident": {Exhaustion: 0.1, MessageID: "trident", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:unattributed_fireball": {Exhaustion: 0.1, MessageID: "onFire", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:wind_charge": {Exhaustion: 0.1, MessageID: "mob", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:wither": {Exhaustion: 0, MessageID: "wither", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}, "minecraft:wither_skull": {Exhaustion: 0.1, MessageID: "witherSkull", Scaling: "when_caused_by_living_non_player", DeathMessageType: "", Effects: ""}}, "minecraft:dimension_type": struct {
	Overworld      Dimension  "nbt:\"minecraft:overworld\""
	OverworldCaves Dimension  "nbt:\"minecraft:overworld_caves\""
	TheEnd         Dimension  "nbt:\"minecraft:the_end\""
	TheNether      Dimension1 "nbt:\"minecraft:the_nether\""
}{Overworld: Dimension{AmbientLight: 0, BedWorks: true, CoordinateScale: 1, Effects: "minecraft:overworld", HasCeiling: false, HasRaids: true, HasSkylight: true, Height: 384, Infiniburn: "#minecraft:infiniburn_overworld", LogicalHeight: 384, MinY: -64, MonsterSpawnBlockLightLimit: 0, Natural: true, Ultrawarm: false, PiglinSafe: false, RespawnAnchorWorks: false, MonsterSpawnLightLevel: struct {
	MaxInclusive int32  "nbt:\"max_inclusive\""
	MinInclusive int32  "nbt:\"min_inclusive\""
	Type         string "nbt:\"type\""
}{MaxInclusive: 7, MinInclusive: 0, Type: "minecraft:uniform"}}, OverworldCaves: Dimension{AmbientLight: 0, BedWorks: true, CoordinateScale: 1, Effects: "minecraft:overworld", HasCeiling: true, HasRaids: true, HasSkylight: true, Height: 384, Infiniburn: "#minecraft:infiniburn_overworld", LogicalHeight: 384, MinY: -64, MonsterSpawnBlockLightLimit: 0, Natural: true, Ultrawarm: false, PiglinSafe: false, RespawnAnchorWorks: false, MonsterSpawnLightLevel: struct {
	MaxInclusive int32  "nbt:\"max_inclusive\""
	MinInclusive int32  "nbt:\"min_inclusive\""
	Type         string "nbt:\"type\""
}{MaxInclusive: 7, MinInclusive: 0, Type: "minecraft:uniform"}}, TheEnd: Dimension{FixedTime: 6000, AmbientLight: 0, BedWorks: false, CoordinateScale: 1, Effects: "minecraft:the_end", HasCeiling: false, HasRaids: true, HasSkylight: false, Height: 256, Infiniburn: "#minecraft:infiniburn_end", LogicalHeight: 256, MinY: 0, MonsterSpawnBlockLightLimit: 0, Natural: false, Ultrawarm: false, PiglinSafe: false, RespawnAnchorWorks: false, MonsterSpawnLightLevel: struct {
	MaxInclusive int32  "nbt:\"max_inclusive\""
	MinInclusive int32  "nbt:\"min_inclusive\""
	Type         string "nbt:\"type\""
}{MaxInclusive: 7, MinInclusive: 0, Type: "minecraft:uniform"}}, TheNether: Dimension1{FixedTime: 18000, AmbientLight: 0.1, BedWorks: false, CoordinateScale: 8, Effects: "minecraft:the_nether", HasCeiling: true, HasRaids: false, HasSkylight: false, Height: 256, Infiniburn: "#minecraft:infiniburn_nether", LogicalHeight: 128, MinY: 0, MonsterSpawnBlockLightLimit: 15, Natural: false, Ultrawarm: true, PiglinSafe: true, RespawnAnchorWorks: true, MonsterSpawnLightLevel: 7}}, "minecraft:enchantment": map[string]struct {
	AnvilCost   int32 "nbt:\"anvil_cost\""
	Description struct {
		Translate string "nbt:\"translate\""
	} "nbt:\"description\""
	Effects struct {
		SmashDamagePerFallenBlock []struct {
			Effect struct {
				Type  string "nbt:\"type\""
				Value struct {
					Base               float32 "nbt:\"base\""
					PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
					Type               string  "nbt:\"type\""
				} "nbt:\"value\""
			} "nbt:\"effect\""
		} "nbt:\"minecraft:smash_damage_per_fallen_block\""
		PreventArmorChange struct{} "nbt:\"minecraft:prevent_armor_change\""
		HitBlock           []struct {
			Effect struct {
				Effects []struct {
					Type   string  "nbt:\"type\""
					Entity string  "nbt:\"state\""
					Pitch  float32 "nbt:\"pitch\""
					Sound  string  "nbt:\"sound\""
					Volume float32 "nbt:\"volume\""
				} "nbt:\"effects\""
				Type string "nbt:\"type\""
			} "nbt:\"effect\""
			Requirements struct {
				Condition string "nbt:\"condition\""
				Terms     []struct {
					Block      string "nbt:\"block\""
					Condition  string "nbt:\"condition\""
					Thundering bool   "nbt:\"thundering\""
					Entity     string "nbt:\"state\""
					Predicate  struct {
						Type      string "nbt:\"type\""
						CanSeeSky bool   "nbt:\"can_see_sky\""
					} "nbt:\"predicate\""
				} "nbt:\"terms\""
			} "nbt:\"requirements\""
		} "nbt:\"minecraft:hit_block\""
		ArmorEffectiveness []struct {
			Effect struct {
				Type  string "nbt:\"type\""
				Value struct {
					Base               float32 "nbt:\"base\""
					PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
					Type               string  "nbt:\"type\""
				} "nbt:\"value\""
			} "nbt:\"effect\""
		} "nbt:\"minecraft:armor_effectiveness\""
		Attributes []struct {
			Amount struct {
				Added              float32 "nbt:\"added\""
				Base               float32 "nbt:\"base\""
				PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
				Type               string  "nbt:\"type\""
			} "nbt:\"amount\""
			Atrribute string "nbt:\"attribute\""
			Id        string "nbt:\"id\""
			Operation string "nbt:\"operation\""
		} "nbt:\"minecraft:attributes\""
		AmmoUse []struct {
			Effect struct {
				Type  string  "nbt:\"type\""
				Value float32 "nbt:\"value\""
			} "nbt:\"effect\""
			Requirements struct {
				Condition string "nbt:\"condition\""
				Predicate struct {
					Items string "nbt:\"items\""
				} "nbt:\"predicate\""
			} "nbt:\"requirements\""
		} "nbt:\"minecraft:ammo_use\""
		ProjectileSpawned []struct {
			Effect struct {
				Duration float32 "nbt:\"duration\""
				Type     string  "nbt:\"type\""
			} "nbt:\"effect\""
		} "nbt:\"minecraft:projectile_spawned\""
		LocationChanged []struct {
			Effect struct {
				BlockState struct {
					State struct {
						Name       string
						Properties map[string]interface{}
					} "nbt:\"state\""
					Type string "nbt:\"type\""
				} "nbt:\"block_state\""
				Height    float32 "nbt:\"height\""
				Offset    []int32 "nbt:\"offset\""
				Predicate struct {
					Predicates []struct {
						Offset []int32 "nbt:\"offset\""
						Tag    string  "nbt:\"tag\""
						Type   string  "nbt:\"type\""
						Blocks string  "nbt:\"blocks\""
						Fluids string  "nbt:\"fluids\""
					} "nbt:\"predicates\""
					Radius struct {
						Max   float32 "nbt:\"max\""
						Min   float32 "nbt:\"min\""
						Type  string  "nbt:\"type\""
						Value struct {
							Base               float32 "nbt:\"base\""
							PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
							Type               string  "nbt:\"type\""
						} "nbt:\"value\""
					} "nbt:\"radius\""
					Type string "nbt:\"type\""
				} "nbt:\"predicate\""
				TriggerGameEvent string "nbt:\"trigger_game_event\""
				Type             string "nbt:\"type\""
			} "nbt:\"effect\""
			Requirements struct {
				Condition string "nbt:\"condition\""
				Entity    string "nbt:\"state\""
				Predicate struct {
					Flags struct {
						IsOnGround bool "nbt:\"is_on_ground\""
					} "nbt:\"flags\""
				} "nbt:\"predicate\""
			} "nbt:\"requirements\""
		} "nbt:\"minecraft:location_changed\""
		DamageImmunity []struct {
			Effect       struct{}
			Requirements struct {
				Condition string "nbt:\"condition\""
				Predicate struct {
					Tags []struct {
						Expected bool   "nbt:\"expected\""
						Id       string "nbt:\"id\""
					} "nbt:\"tags\""
				} "nbt:\"predicate\""
			} "nbt:\"requirements\""
		} "nbt:\"minecraft:damage_immunity\""
		DamageProtection []struct {
			Effect struct {
				Type  string "nbt:\"type\""
				Value struct {
					Base               float32 "nbt:\"base\""
					PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
					Type               string  "nbt:\"type\""
				} "nbt:\"value\""
			} "nbt:\"effect\""
			Requirements struct {
				Condition string "nbt:\"condition\""
				Predicate struct {
					Tags []struct {
						Expected bool   "nbt:\"expected\""
						Id       string "nbt:\"id\""
					} "nbt:\"tags\""
				} "nbt:\"predicate\""
				Terms []struct {
					Condition string "nbt:\"condition\""
					Predicate struct {
						Tags []struct {
							Expected bool   "nbt:\"expected\""
							Id       string "nbt:\"id\""
						} "nbt:\"tags\""
					} "nbt:\"predicate\""
				} "nbt:\"terms\""
			} "nbt:\"requirements\""
		} "nbt:\"minecraft:damage_protection\""
		Damage []struct {
			Effect struct {
				Type  string "nbt:\"type\""
				Value struct {
					Base               float32 "nbt:\"base\""
					PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
					Type               string  "nbt:\"type\""
				} "nbt:\"value\""
			} "nbt:\"effect\""
			Requirements struct {
				Condition string "nbt:\"condition\""
				Entity    string "nbt:\"state\""
				Predicate struct {
					Type string "nbt:\"type\""
				} "nbt:\"predicate\""
			} "nbt:\"requirements\""
		} "nbt:\"minecraft:damage\""
		PostAttack []struct {
			Affected string "nbt:\"affected\""
			Effect   struct {
				MaxAmplifier float32 "nbt:\"max_amplifier\""
				MaxDuration  struct {
					Base               float32 "nbt:\"base\""
					PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
					Type               string  "nbt:\"type\""
				} "nbt:\"max_duration\""
				Duration struct {
					Base               float32 "nbt:\"base\""
					PerLevelAboveFirst float32 "nbt:\"per_level_above_first\""
					Type               string  "nbt:\"type\""
				} "nbt:\"duration\""
				MinAmplifier float32 "nbt:\"min_amplifier\""
				MinDuration  float32 "nbt:\"min_duration\""
				ToApply      string  "nbt:\"to_apply\""
				Type         string  "nbt:\"type\""
				Effects      []struct {
					Type   string  "nbt:\"type\""
					Entity string  "nbt:\"state\""
					Pitch  float32 "nbt:\"pitch\""
					Sound  string  "nbt:\"sound\""
					Volume float32 "nbt:\"volume\""
				} "nbt:\"effects\""
			} "nbt:\"effect\""
			Enchanted    string "nbt:\"enchanted\""
			Requirements struct {
				Condition string "nbt:\"condition\""
				Predicate struct {
					IsDirect bool "nbt:\"is_direct\""
				} "nbt:\"predicate\""
				Terms []struct {
					Condition  string "nbt:\"condition\""
					Thundering bool   "nbt:\"thundering\""
					Entity     string "nbt:\"state\""
					Predicate  struct {
						IsDirect bool   "nbt:\"is_direct\""
						Type     string "nbt:\"type\""
						Location struct {
							CanSeeSky bool "nbt:\"can_see_sky\""
						} "nbt:\"location\""
					} "nbt:\"predicate\""
				} "nbt:\"terms\""
			} "nbt:\"requirements\""
		} "nbt:\"minecraft:post_attack\""
	} "nbt:\"effects\""
	MaxCost struct {
		Base               int32 "nbt:\"base\""
		PerLevelAboveFirst int32 "nbt:\"per_level_above_first\""
	} "nbt:\"max_cost\""
	MinCost struct {
		Base               int32 "nbt:\"base\""
		PerLevelAboveFirst int32 "nbt:\"per_level_above_first\""
	} "nbt:\"min_cost\""
	MaxLevel       int32    "nbt:\"max_level\""
	Slots          []string "nbt:\"slots\""
	SupportedItems string   "nbt:\"supported_items\""
	Weight         int32    "nbt:\"weight\""
	ExclusiveSet   string   "nbt:\"exclusive_set\""
	PrimaryItems   string   "nbt:\"primary_items\""
}(nil), "minecraft:jukebox_song": map[string]struct {
	ComparatorOutput int32 "nbt:\"comparator_output\""
	Description      struct {
		Translate string "nbt:\"translate\""
	} "nbt:\"description\""
	LengthInSeconds float32 "nbt:\"length_in_seconds\""
	SoundEvent      string  "nbt:\"sound_event\""
}{"minecraft:11": {ComparatorOutput: 11, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.11"}, LengthInSeconds: 71, SoundEvent: "minecraft:music_disc.11"}, "minecraft:13": {ComparatorOutput: 1, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.13"}, LengthInSeconds: 178, SoundEvent: "minecraft:music_disc.13"}, "minecraft:5": {ComparatorOutput: 15, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.5"}, LengthInSeconds: 178, SoundEvent: "minecraft:music_disc.5"}, "minecraft:blocks": {ComparatorOutput: 3, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.blocks"}, LengthInSeconds: 345, SoundEvent: "minecraft:music_disc.blocks"}, "minecraft:cat": {ComparatorOutput: 2, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.cat"}, LengthInSeconds: 185, SoundEvent: "minecraft:music_disc.cat"}, "minecraft:chirp": {ComparatorOutput: 4, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.chirp"}, LengthInSeconds: 185, SoundEvent: "minecraft:music_disc.chirp"}, "minecraft:creator": {ComparatorOutput: 12, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.creator"}, LengthInSeconds: 176, SoundEvent: "minecraft:music_disc.creator"}, "minecraft:creator_music_box": {ComparatorOutput: 11, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.creator_music_box"}, LengthInSeconds: 73, SoundEvent: "minecraft:music_disc.creator_music_box"}, "minecraft:far": {ComparatorOutput: 5, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.far"}, LengthInSeconds: 174, SoundEvent: "minecraft:music_disc.far"}, "minecraft:mall": {ComparatorOutput: 6, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.mall"}, LengthInSeconds: 197, SoundEvent: "minecraft:music_disc.mall"}, "minecraft:mellohi": {ComparatorOutput: 7, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.mellohi"}, LengthInSeconds: 96, SoundEvent: "minecraft:music_disc.mellohi"}, "minecraft:otherside": {ComparatorOutput: 14, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.otherside"}, LengthInSeconds: 195, SoundEvent: "minecraft:music_disc.otherside"}, "minecraft:pigstep": {ComparatorOutput: 13, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.pigstep"}, LengthInSeconds: 149, SoundEvent: "minecraft:music_disc.pigstep"}, "minecraft:precipice": {ComparatorOutput: 13, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.precipice"}, LengthInSeconds: 299, SoundEvent: "minecraft:music_disc.precipice"}, "minecraft:relic": {ComparatorOutput: 14, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.relic"}, LengthInSeconds: 218, SoundEvent: "minecraft:music_disc.relic"}, "minecraft:stal": {ComparatorOutput: 8, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.stal"}, LengthInSeconds: 150, SoundEvent: "minecraft:music_disc.stal"}, "minecraft:strad": {ComparatorOutput: 9, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.strad"}, LengthInSeconds: 188, SoundEvent: "minecraft:music_disc.strad"}, "minecraft:wait": {ComparatorOutput: 12, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.wait"}, LengthInSeconds: 238, SoundEvent: "minecraft:music_disc.wait"}, "minecraft:ward": {ComparatorOutput: 10, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "jukebox_song.minecraft.ward"}, LengthInSeconds: 251, SoundEvent: "minecraft:music_disc.ward"}}, "minecraft:painting_variant": map[string]struct {
	AssetId string "nbt:\"asset_id\""
	Height  int32  "nbt:\"height\""
	Weight  int32  "nbt:\"width\""
}{"minecraft:alban": {AssetId: "minecraft:alban", Height: 1, Weight: 1}, "minecraft:aztec": {AssetId: "minecraft:aztec", Height: 1, Weight: 1}, "minecraft:aztec2": {AssetId: "minecraft:aztec2", Height: 1, Weight: 1}, "minecraft:backyard": {AssetId: "minecraft:backyard", Height: 4, Weight: 3}, "minecraft:baroque": {AssetId: "minecraft:baroque", Height: 2, Weight: 2}, "minecraft:bomb": {AssetId: "minecraft:bomb", Height: 1, Weight: 1}, "minecraft:bouquet": {AssetId: "minecraft:bouquet", Height: 3, Weight: 3}, "minecraft:burning_skull": {AssetId: "minecraft:burning_skull", Height: 4, Weight: 4}, "minecraft:bust": {AssetId: "minecraft:bust", Height: 2, Weight: 2}, "minecraft:cavebird": {AssetId: "minecraft:cavebird", Height: 3, Weight: 3}, "minecraft:changing": {AssetId: "minecraft:changing", Height: 2, Weight: 4}, "minecraft:cotan": {AssetId: "minecraft:cotan", Height: 3, Weight: 3}, "minecraft:courbet": {AssetId: "minecraft:courbet", Height: 1, Weight: 2}, "minecraft:creebet": {AssetId: "minecraft:creebet", Height: 1, Weight: 2}, "minecraft:donkey_kong": {AssetId: "minecraft:donkey_kong", Height: 3, Weight: 4}, "minecraft:earth": {AssetId: "minecraft:earth", Height: 2, Weight: 2}, "minecraft:endboss": {AssetId: "minecraft:endboss", Height: 3, Weight: 3}, "minecraft:fern": {AssetId: "minecraft:fern", Height: 3, Weight: 3}, "minecraft:fighters": {AssetId: "minecraft:fighters", Height: 2, Weight: 4}, "minecraft:finding": {AssetId: "minecraft:finding", Height: 2, Weight: 4}, "minecraft:fire": {AssetId: "minecraft:fire", Height: 2, Weight: 2}, "minecraft:graham": {AssetId: "minecraft:graham", Height: 2, Weight: 1}, "minecraft:humble": {AssetId: "minecraft:humble", Height: 2, Weight: 2}, "minecraft:kebab": {AssetId: "minecraft:kebab", Height: 1, Weight: 1}, "minecraft:lowmist": {AssetId: "minecraft:lowmist", Height: 2, Weight: 4}, "minecraft:match": {AssetId: "minecraft:match", Height: 2, Weight: 2}, "minecraft:meditative": {AssetId: "minecraft:meditative", Height: 1, Weight: 1}, "minecraft:orb": {AssetId: "minecraft:orb", Height: 4, Weight: 4}, "minecraft:owlemons": {AssetId: "minecraft:owlemons", Height: 3, Weight: 3}, "minecraft:passage": {AssetId: "minecraft:passage", Height: 2, Weight: 4}, "minecraft:pigscene": {AssetId: "minecraft:pigscene", Height: 4, Weight: 4}, "minecraft:plant": {AssetId: "minecraft:plant", Height: 1, Weight: 1}, "minecraft:pointer": {AssetId: "minecraft:pointer", Height: 4, Weight: 4}, "minecraft:pond": {AssetId: "minecraft:pond", Height: 4, Weight: 3}, "minecraft:pool": {AssetId: "minecraft:pool", Height: 1, Weight: 2}, "minecraft:prairie_ride": {AssetId: "minecraft:prairie_ride", Height: 2, Weight: 1}, "minecraft:sea": {AssetId: "minecraft:sea", Height: 1, Weight: 2}, "minecraft:skeleton": {AssetId: "minecraft:skeleton", Height: 3, Weight: 4}, "minecraft:skull_and_roses": {AssetId: "minecraft:skull_and_roses", Height: 2, Weight: 2}, "minecraft:stage": {AssetId: "minecraft:stage", Height: 2, Weight: 2}, "minecraft:sunflowers": {AssetId: "minecraft:sunflowers", Height: 3, Weight: 3}, "minecraft:sunset": {AssetId: "minecraft:sunset", Height: 1, Weight: 2}, "minecraft:tides": {AssetId: "minecraft:tides", Height: 3, Weight: 3}, "minecraft:unpacked": {AssetId: "minecraft:unpacked", Height: 4, Weight: 4}, "minecraft:void": {AssetId: "minecraft:void", Height: 2, Weight: 2}, "minecraft:wanderer": {AssetId: "minecraft:wanderer", Height: 2, Weight: 1}, "minecraft:wasteland": {AssetId: "minecraft:wasteland", Height: 1, Weight: 1}, "minecraft:water": {AssetId: "minecraft:water", Height: 2, Weight: 2}, "minecraft:wind": {AssetId: "minecraft:wind", Height: 2, Weight: 2}, "minecraft:wither": {AssetId: "minecraft:wither", Height: 2, Weight: 2}}, "minecraft:trim_material": map[string]struct {
	AssetName   string "nbt:\"asset_name\""
	Description struct {
		Color     string "nbt:\"color\""
		Translate string "nbt:\"translate\""
	} "nbt:\"description\""
	Ingredient             string  "nbt:\"ingredient\""
	ItemModelIndex         float32 "nbt:\"item_model_index\""
	OverrideArmorMaterials struct {
		Diamond string "nbt:\"minecraft:diamond,omitempty\""
		Gold    string "nbt:\"minecraft:gold,omitempty\""
		Iron    string "nbt:\"minecraft:iron,omitempty\""
	} "nbt:\"override_armor_materials,omitempty\""
}{"minecraft:amethyst": {AssetName: "amethyst", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#9A5CC6", Translate: "trim_material.minecraft.amethyst"}, Ingredient: "minecraft:amethyst_shard", ItemModelIndex: 1, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:copper": {AssetName: "copper", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#B4684D", Translate: "trim_material.minecraft.copper"}, Ingredient: "minecraft:copper_ingot", ItemModelIndex: 0.5, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:diamond": {AssetName: "diamond", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#6EECD2", Translate: "trim_material.minecraft.diamond"}, Ingredient: "minecraft:diamond", ItemModelIndex: 0.8, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:emerald": {AssetName: "emerald", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#11A036", Translate: "trim_material.minecraft.emerald"}, Ingredient: "minecraft:emerald", ItemModelIndex: 0.7, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:gold": {AssetName: "gold", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#DEB12D", Translate: "trim_material.minecraft.gold"}, Ingredient: "minecraft:gold_ingot", ItemModelIndex: 0.6, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:iron": {AssetName: "iron", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#ECECEC", Translate: "trim_material.minecraft.iron"}, Ingredient: "minecraft:iron_ingot", ItemModelIndex: 0.2, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:lapis": {AssetName: "lapis", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#416E97", Translate: "trim_material.minecraft.lapis"}, Ingredient: "minecraft:lapis_lazuli", ItemModelIndex: 0.9, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:netherite": {AssetName: "netherite", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#625859", Translate: "trim_material.minecraft.netherite"}, Ingredient: "minecraft:netherite_ingot", ItemModelIndex: 0.3, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:quartz": {AssetName: "quartz", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#E3D4C4", Translate: "trim_material.minecraft.quartz"}, Ingredient: "minecraft:quartz", ItemModelIndex: 0.1, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}, "minecraft:redstone": {AssetName: "redstone", Description: struct {
	Color     string "nbt:\"color\""
	Translate string "nbt:\"translate\""
}{Color: "#971607", Translate: "trim_material.minecraft.redstone"}, Ingredient: "minecraft:redstone", ItemModelIndex: 0.4, OverrideArmorMaterials: struct {
	Diamond string "nbt:\"minecraft:diamond,omitempty\""
	Gold    string "nbt:\"minecraft:gold,omitempty\""
	Iron    string "nbt:\"minecraft:iron,omitempty\""
}{Diamond: "", Gold: "", Iron: ""}}}, "minecraft:trim_pattern": map[string]struct {
	AssetId     string "nbt:\"asset_id\""
	Decal       bool   "nbt:\"decal\""
	Description struct {
		Translate string "nbt:\"translate\""
	} "nbt:\"description\""
	TemplateItem string "nbt:\"template_item\""
}{"minecraft:bolt": {AssetId: "minecraft:bolt", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.bolt"}, TemplateItem: "minecraft:bolt_armor_trim_smithing_template"}, "minecraft:coast": {AssetId: "minecraft:coast", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.coast"}, TemplateItem: "minecraft:coast_armor_trim_smithing_template"}, "minecraft:dune": {AssetId: "minecraft:dune", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.dune"}, TemplateItem: "minecraft:dune_armor_trim_smithing_template"}, "minecraft:eye": {AssetId: "minecraft:eye", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.eye"}, TemplateItem: "minecraft:eye_armor_trim_smithing_template"}, "minecraft:flow": {AssetId: "minecraft:flow", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.flow"}, TemplateItem: "minecraft:flow_armor_trim_smithing_template"}, "minecraft:host": {AssetId: "minecraft:host", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.host"}, TemplateItem: "minecraft:host_armor_trim_smithing_template"}, "minecraft:raiser": {AssetId: "minecraft:raiser", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.raiser"}, TemplateItem: "minecraft:raiser_armor_trim_smithing_template"}, "minecraft:rib": {AssetId: "minecraft:rib", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.rib"}, TemplateItem: "minecraft:rib_armor_trim_smithing_template"}, "minecraft:sentry": {AssetId: "minecraft:sentry", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.sentry"}, TemplateItem: "minecraft:sentry_armor_trim_smithing_template"}, "minecraft:shaper": {AssetId: "minecraft:shaper", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.shaper"}, TemplateItem: "minecraft:shaper_armor_trim_smithing_template"}, "minecraft:silence": {AssetId: "minecraft:silence", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.silence"}, TemplateItem: "minecraft:silence_armor_trim_smithing_template"}, "minecraft:snout": {AssetId: "minecraft:snout", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.snout"}, TemplateItem: "minecraft:snout_armor_trim_smithing_template"}, "minecraft:spire": {AssetId: "minecraft:spire", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.spire"}, TemplateItem: "minecraft:spire_armor_trim_smithing_template"}, "minecraft:tide": {AssetId: "minecraft:tide", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.tide"}, TemplateItem: "minecraft:tide_armor_trim_smithing_template"}, "minecraft:vex": {AssetId: "minecraft:vex", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.vex"}, TemplateItem: "minecraft:vex_armor_trim_smithing_template"}, "minecraft:ward": {AssetId: "minecraft:ward", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.ward"}, TemplateItem: "minecraft:ward_armor_trim_smithing_template"}, "minecraft:wayfinder": {AssetId: "minecraft:wayfinder", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.wayfinder"}, TemplateItem: "minecraft:wayfinder_armor_trim_smithing_template"}, "minecraft:wild": {AssetId: "minecraft:wild", Decal: false, Description: struct {
	Translate string "nbt:\"translate\""
}{Translate: "trim_pattern.minecraft.wild"}, TemplateItem: "minecraft:wild_armor_trim_smithing_template"}}, "minecraft:wolf_variant": map[string]struct {
	AngryTexture string "nbt:\"angry_texture\""
	Biomes       string "nbt:\"biomes\""
	TameTexture  string "nbt:\"tame_texture\""
	WildTexture  string "nbt:\"wild_texture\""
}{"minecraft:ashen": {AngryTexture: "minecraft:state/wolf/wolf_ashen_angry", Biomes: "minecraft:snowy_taiga", TameTexture: "minecraft:state/wolf/wolf_ashen_tame", WildTexture: "minecraft:state/wolf/wolf_ashen"}, "minecraft:black": {AngryTexture: "minecraft:state/wolf/wolf_black_angry", Biomes: "minecraft:old_growth_pine_taiga", TameTexture: "minecraft:state/wolf/wolf_black_tame", WildTexture: "minecraft:state/wolf/wolf_black"}, "minecraft:chestnut": {AngryTexture: "minecraft:state/wolf/wolf_chestnut_angry", Biomes: "minecraft:old_growth_spruce_taiga", TameTexture: "minecraft:state/wolf/wolf_chestnut_tame", WildTexture: "minecraft:state/wolf/wolf_chestnut"}, "minecraft:pale": {AngryTexture: "minecraft:state/wolf/wolf_angry", Biomes: "minecraft:taiga", TameTexture: "minecraft:state/wolf/wolf_tame", WildTexture: "minecraft:state/wolf/wolf"}, "minecraft:rusty": {AngryTexture: "minecraft:state/wolf/wolf_rusty_angry", Biomes: "#minecraft:is_jungle", TameTexture: "minecraft:state/wolf/wolf_rusty_tame", WildTexture: "minecraft:state/wolf/wolf_rusty"}, "minecraft:snowy": {AngryTexture: "minecraft:state/wolf/wolf_snowy_angry", Biomes: "minecraft:grove", TameTexture: "minecraft:state/wolf/wolf_snowy_tame", WildTexture: "minecraft:state/wolf/wolf_snowy"}, "minecraft:spotted": {AngryTexture: "minecraft:state/wolf/wolf_spotted_angry", Biomes: "#minecraft:is_savanna", TameTexture: "minecraft:state/wolf/wolf_spotted_tame", WildTexture: "minecraft:state/wolf/wolf_spotted"}, "minecraft:striped": {AngryTexture: "minecraft:state/wolf/wolf_striped_angry", Biomes: "#minecraft:is_badlands", TameTexture: "minecraft:state/wolf/wolf_striped_tame", WildTexture: "minecraft:state/wolf/wolf_striped"}, "minecraft:woods": {AngryTexture: "minecraft:state/wolf/wolf_woods_angry", Biomes: "minecraft:forest", TameTexture: "minecraft:state/wolf/wolf_woods_tame", WildTexture: "minecraft:state/wolf/wolf_woods"}}, "minecraft:worldgen/biome": map[string]struct {
	Downfall float32 "nbt:\"downfall\""
	Effects  struct {
		FogColor     int32 "nbt:\"fog_color\""
		FoliageColor int32 "nbt:\"foliage_color,omitempty\""
		GrassColor   int32 "nbt:\"grass_color,omitempty\""
		MoodSound    struct {
			BlockSearchExtent int32   "nbt:\"block_search_extent\""
			Offset            float64 "nbt:\"offset\""
			Sound             string  "nbt:\"sound\""
			TickDelay         int32   "nbt:\"tick_delay\""
		} "nbt:\"mood_sound\""
		Music struct {
			MaxDelay            int32  "nbt:\"max_delay\""
			MinDelay            int32  "nbt:\"min_delay\""
			ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
			Sound               string "nbt:\"sound\""
		} "nbt:\"music,omitempty\""
		SkyColor      int32 "nbt:\"sky_color\""
		WaterColor    int32 "nbt:\"water_color\""
		WaterFogColor int32 "nbt:\"water_fog_color\""
	} "nbt:\"effects\""
	HasPrecipitation bool    "nbt:\"has_precipitation\""
	Temperature      float32 "nbt:\"temperature\""
}{"minecraft:badlands": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:bamboo_jungle": {Downfall: 0.9, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7842047, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.95}, "minecraft:basalt_deltas": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 6840176, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.basalt_deltas.mood", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:beach": {Downfall: 0.4, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7907327, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.8}, "minecraft:birch_forest": {Downfall: 0.6, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8037887, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.6}, "minecraft:cherry_grove": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 6141935, WaterFogColor: 6141935}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:cold_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4020182, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:crimson_forest": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 3343107, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.crimson_forest.mood", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:dark_forest": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7972607, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.7}, "minecraft:deep_cold_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4020182, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:deep_dark": {Downfall: 0.4, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7907327, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.8}, "minecraft:deep_frozen_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 3750089, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:deep_lukewarm_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4566514, WaterFogColor: 267827}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:deep_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:desert": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:dripstone_caves": {Downfall: 0.4, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7907327, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.8}, "minecraft:end_barrens": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 10518688, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 0, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 0.5}, "minecraft:end_highlands": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 10518688, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 0, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 0.5}, "minecraft:end_midlands": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 10518688, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 0, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 0.5}, "minecraft:eroded_badlands": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:flower_forest": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7972607, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.7}, "minecraft:forest": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7972607, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.7}, "minecraft:frozen_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8364543, WaterColor: 3750089, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0}, "minecraft:frozen_peaks": {Downfall: 0.9, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8756735, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: -0.7}, "minecraft:frozen_river": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8364543, WaterColor: 3750089, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0}, "minecraft:grove": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8495359, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: -0.2}, "minecraft:ice_spikes": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8364543, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0}, "minecraft:jagged_peaks": {Downfall: 0.9, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8756735, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: -0.7}, "minecraft:jungle": {Downfall: 0.9, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7842047, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.95}, "minecraft:lukewarm_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4566514, WaterFogColor: 267827}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:lush_caves": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:mangrove_swamp": {Downfall: 0.9, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7907327, WaterColor: 3832426, WaterFogColor: 5077600}, HasPrecipitation: true, Temperature: 0.8}, "minecraft:meadow": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 937679, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:mushroom_fields": {Downfall: 1, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7842047, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.9}, "minecraft:nether_wastes": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 3344392, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.nether_wastes.mood", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:old_growth_birch_forest": {Downfall: 0.6, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8037887, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.6}, "minecraft:old_growth_pine_taiga": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8168447, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.3}, "minecraft:old_growth_spruce_taiga": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8233983, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.25}, "minecraft:plains": {Downfall: 0.4, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7907327, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.8}, "minecraft:river": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:savanna": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:savanna_plateau": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:small_end_islands": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 10518688, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 0, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 0.5}, "minecraft:snowy_beach": {Downfall: 0.3, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8364543, WaterColor: 4020182, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.05}, "minecraft:snowy_plains": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8364543, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0}, "minecraft:snowy_slopes": {Downfall: 0.9, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8560639, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: -0.3}, "minecraft:snowy_taiga": {Downfall: 0.4, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8625919, WaterColor: 4020182, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: -0.5}, "minecraft:soul_sand_valley": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 1787717, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.soul_sand_valley.mood", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:sparse_jungle": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7842047, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.95}, "minecraft:stony_peaks": {Downfall: 0.3, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7776511, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 1}, "minecraft:stony_shore": {Downfall: 0.3, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8233727, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.2}, "minecraft:sunflower_plains": {Downfall: 0.4, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7907327, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.8}, "minecraft:swamp": {Downfall: 0.9, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7907327, WaterColor: 6388580, WaterFogColor: 2302743}, HasPrecipitation: true, Temperature: 0.8}, "minecraft:taiga": {Downfall: 0.8, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8233983, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.25}, "minecraft:the_end": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 10518688, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 0, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 0.5}, "minecraft:the_void": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 0.5}, "minecraft:warm_ocean": {Downfall: 0.5, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8103167, WaterColor: 4445678, WaterFogColor: 270131}, HasPrecipitation: true, Temperature: 0.5}, "minecraft:warped_forest": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 1705242, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.warped_forest.mood", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:windswept_forest": {Downfall: 0.3, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8233727, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.2}, "minecraft:windswept_gravelly_hills": {Downfall: 0.3, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8233727, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.2}, "minecraft:windswept_hills": {Downfall: 0.3, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 8233727, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: true, Temperature: 0.2}, "minecraft:windswept_savanna": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}, "minecraft:wooded_badlands": {Downfall: 0, Effects: struct {
	FogColor     int32 "nbt:\"fog_color\""
	FoliageColor int32 "nbt:\"foliage_color,omitempty\""
	GrassColor   int32 "nbt:\"grass_color,omitempty\""
	MoodSound    struct {
		BlockSearchExtent int32   "nbt:\"block_search_extent\""
		Offset            float64 "nbt:\"offset\""
		Sound             string  "nbt:\"sound\""
		TickDelay         int32   "nbt:\"tick_delay\""
	} "nbt:\"mood_sound\""
	Music struct {
		MaxDelay            int32  "nbt:\"max_delay\""
		MinDelay            int32  "nbt:\"min_delay\""
		ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
		Sound               string "nbt:\"sound\""
	} "nbt:\"music,omitempty\""
	SkyColor      int32 "nbt:\"sky_color\""
	WaterColor    int32 "nbt:\"water_color\""
	WaterFogColor int32 "nbt:\"water_fog_color\""
}{FogColor: 12638463, FoliageColor: 0, GrassColor: 0, MoodSound: struct {
	BlockSearchExtent int32   "nbt:\"block_search_extent\""
	Offset            float64 "nbt:\"offset\""
	Sound             string  "nbt:\"sound\""
	TickDelay         int32   "nbt:\"tick_delay\""
}{BlockSearchExtent: 8, Offset: 2, Sound: "minecraft:ambient.cave", TickDelay: 6000}, Music: struct {
	MaxDelay            int32  "nbt:\"max_delay\""
	MinDelay            int32  "nbt:\"min_delay\""
	ReplaceCurrentMusic bool   "nbt:\"replace_current_music\""
	Sound               string "nbt:\"sound\""
}{MaxDelay: 0, MinDelay: 0, ReplaceCurrentMusic: false, Sound: ""}, SkyColor: 7254527, WaterColor: 4159204, WaterFogColor: 329011}, HasPrecipitation: false, Temperature: 2}}}
