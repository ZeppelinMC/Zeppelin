package registry

import (
	"bytes"
	_ "embed"
	"reflect"
	"sync"

	"github.com/dynamitemc/aether/nbt"
)

var biome_id_mu sync.Mutex
var BiomeId = biomeIdMap{m: make(map[string]int32)}

type biomeIdMap struct {
	mu sync.Mutex
	m  map[string]int32
}

func (b *biomeIdMap) SetMap(m map[string]int32) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.m = m
}

func (b *biomeIdMap) GetMap() map[string]int32 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.m
}

var Registries registries

type registries_t map[string]any

var RegistryMap = make(map[string]any)

type Dimension1 struct {
	FixedTime                   int64   `nbt:"fixed_time"`
	AmbientLight                float32 `nbt:"ambient_light"`
	BedWorks                    bool    `nbt:"bed_works"`
	CoordinateScale             float64 `nbt:"coordinate_scale"`
	Effects                     string  `nbt:"effects"`
	HasCeiling                  bool    `nbt:"has_ceiling"`
	HasRaids                    bool    `nbt:"has_raids"`
	HasSkylight                 bool    `nbt:"has_skylight"`
	Height                      int32   `nbt:"height"`
	Infiniburn                  string  `nbt:"infiniburn"`
	LogicalHeight               int32   `nbt:"logical_height"`
	MinY                        int32   `nbt:"min_y"`
	MonsterSpawnBlockLightLimit int32   `nbt:"monster_spawn_block_light_limit"`
	Natural                     bool    `nbt:"natural"`
	Ultrawarm                   bool    `nbt:"ultrawarm"`
	PiglinSafe                  bool    `nbt:"piglin_safe"`
	RespawnAnchorWorks          bool    `nbt:"respawn_anchor_works"`
	MonsterSpawnLightLevel      int32   `nbt:"monster_spawn_light_level"`
}

type Dimension struct {
	FixedTime                   int64   `nbt:"fixed_time"`
	AmbientLight                float32 `nbt:"ambient_light"`
	BedWorks                    bool    `nbt:"bed_works"`
	CoordinateScale             float64 `nbt:"coordinate_scale"`
	Effects                     string  `nbt:"effects"`
	HasCeiling                  bool    `nbt:"has_ceiling"`
	HasRaids                    bool    `nbt:"has_raids"`
	HasSkylight                 bool    `nbt:"has_skylight"`
	Height                      int32   `nbt:"height"`
	Infiniburn                  string  `nbt:"infiniburn"`
	LogicalHeight               int32   `nbt:"logical_height"`
	MinY                        int32   `nbt:"min_y"`
	MonsterSpawnBlockLightLimit int32   `nbt:"monster_spawn_block_light_limit"`
	Natural                     bool    `nbt:"natural"`
	Ultrawarm                   bool    `nbt:"ultrawarm"`
	PiglinSafe                  bool    `nbt:"piglin_safe"`
	RespawnAnchorWorks          bool    `nbt:"respawn_anchor_works"`
	MonsterSpawnLightLevel      struct {
		MaxInclusive int32  `nbt:"max_inclusive"`
		MinInclusive int32  `nbt:"min_inclusive"`
		Type         string `nbt:"type"`
	} `nbt:"monster_spawn_light_level"`
}

type ChatType struct {
	Chat struct {
		Parameters     []string `nbt:"parameters"`
		TranslationKey string   `nbt:"translation_key"`

		Style struct {
			Color  string `nbt:"color"`
			Italic bool   `nbt:"italic"`
		} `nbt:"style,omitempty"`
	} `nbt:"chat"`
	Narration struct {
		Parameters     []string `nbt:"parameters"`
		TranslationKey string   `nbt:"translation_key"`
	} `nbt:"narration"`
}

type registries struct {
	BannerPattern map[string]struct {
		AssetId        string `nbt:"asset_id"`
		TranslationKey string `nbt:"translation_key"`
	} `nbt:"minecraft:banner_pattern"`
	ChatType   map[string]ChatType `nbt:"minecraft:chat_type"`
	DamageType map[string]struct {
		Exhaustion       float32 `nbt:"exhaustion"`
		MessageID        string  `nbt:"message_id"`
		Scaling          string  `nbt:"scaling"`
		DeathMessageType string  `nbt:"death_message_type,omitempty"`
		Effects          string  `nbt:"effects,omitempty"`
	} `nbt:"minecraft:damage_type"`
	DimensionType struct {
		Overworld      Dimension  `nbt:"minecraft:overworld"`
		OverworldCaves Dimension  `nbt:"minecraft:overworld_caves"`
		TheEnd         Dimension  `nbt:"minecraft:the_end"`
		TheNether      Dimension1 `nbt:"minecraft:the_nether"`
	} `nbt:"minecraft:dimension_type"`
	/*Enchantment map[string]struct {
		AnvilCost   int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects struct {
			SmashDamagePerFallenBlock []struct {
				Effect struct {
					Type  string `nbt:"type"`
					Value struct {
						Base               float32 `nbt:"base"`
						PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
						Type               string  `nbt:"type"`
					} `nbt:"value"`
				} `nbt:"effect"`
			} `nbt:"minecraft:smash_damage_per_fallen_block"`
			PreventArmorChange struct{} `nbt:"minecraft:prevent_armor_change"`
			HitBlock           []struct {
				Effect struct {
					Effects []struct {
						Type string `nbt:"type"`

						Entity string `nbt:"entity"`

						Pitch  float32 `nbt:"pitch"`
						Sound  string  `nbt:"sound"`
						Volume float32 `nbt:"volume"`
					} `nbt:"effects"`
					Type string `nbt:"type"`
				} `nbt:"effect"`
				Requirements struct {
					Condition string `nbt:"condition"`
					Terms     []struct {
						Block string `nbt:"block"`

						Condition string `nbt:"condition"`

						Thundering bool `nbt:"thundering"`

						Entity string `nbt:"entity"`

						Predicate struct {
							Type      string `nbt:"type"`
							CanSeeSky bool   `nbt:"can_see_sky"`
						} `nbt:"predicate"`
					} `nbt:"terms"`
				} `nbt:"requirements"`
			} `nbt:"minecraft:hit_block"`
			ArmorEffectiveness []struct {
				Effect struct {
					Type  string `nbt:"type"`
					Value struct {
						Base               float32 `nbt:"base"`
						PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
						Type               string  `nbt:"type"`
					} `nbt:"value"`
				} `nbt:"effect"`
			} `nbt:"minecraft:armor_effectiveness"`
			Attributes []struct {
				Amount struct {
					Added              float32 `nbt:"added"`
					Base               float32 `nbt:"base"`
					PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
					Type               string  `nbt:"type"`
				} `nbt:"amount"`
				Atrribute string `nbt:"attribute"`
				Id        string `nbt:"id"`
				Operation string `nbt:"operation"`
			} `nbt:"minecraft:attributes"`
			AmmoUse []struct {
				Effect struct {
					Type  string  `nbt:"type"`
					Value float32 `nbt:"value"`
				} `nbt:"effect"`
				Requirements struct {
					Condition string `nbt:"condition"`
					Predicate struct {
						Items string `nbt:"items"`
					} `nbt:"predicate"`
				} `nbt:"requirements"`
			} `nbt:"minecraft:ammo_use"`
			ProjectileSpawned []struct {
				Effect struct {
					Duration float32 `nbt:"duration"`
					Type     string  `nbt:"type"`
				} `nbt:"effect"`
			} `nbt:"minecraft:projectile_spawned"`
			LocationChanged []struct {
				Effect struct {
					BlockState struct {
						State struct {
							Name       string
							Properties map[string]any
						} `nbt:"state"`
						Type string `nbt:"type"`
					} `nbt:"block_state"`
					Height    float32 `nbt:"height"`
					Offset    []int32 `nbt:"offset"`
					Predicate struct {
						Predicates []struct {
							Offset []int32 `nbt:"offset"`
							Tag    string  `nbt:"tag"`
							Type   string  `nbt:"type"`
							Blocks string  `nbt:"blocks"`
							Fluids string  `nbt:"fluids"`
						} `nbt:"predicates"`
						Radius struct {
							Max   float32 `nbt:"max"`
							Min   float32 `nbt:"min"`
							Type  string  `nbt:"type"`
							Value struct {
								Base               float32 `nbt:"base"`
								PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
								Type               string  `nbt:"type"`
							} `nbt:"value"`
						} `nbt:"radius"`
						Type string `nbt:"type"`
					} `nbt:"predicate"`
					TriggerGameEvent string `nbt:"trigger_game_event"`
					Type             string `nbt:"type"`
				} `nbt:"effect"`
				Requirements struct {
					Condition string `nbt:"condition"`
					Entity    string `nbt:"entity"`
					Predicate struct {
						Flags struct {
							IsOnGround bool `nbt:"is_on_ground"`
						} `nbt:"flags"`
					} `nbt:"predicate"`
				} `nbt:"requirements"`
			} `nbt:"minecraft:location_changed"`
			DamageImmunity []struct {
				Effect       struct{}
				Requirements struct {
					Condition string `nbt:"condition"`
					Predicate struct {
						Tags []struct {
							Expected bool   `nbt:"expected"`
							Id       string `nbt:"id"`
						} `nbt:"tags"`
					} `nbt:"predicate"`
				} `nbt:"requirements"`
			} `nbt:"minecraft:damage_immunity"`
			DamageProtection []struct {
				Effect struct {
					Type  string `nbt:"type"`
					Value struct {
						Base               float32 `nbt:"base"`
						PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
						Type               string  `nbt:"type"`
					} `nbt:"value"`
				} `nbt:"effect"`
				Requirements struct {
					Condition string `nbt:"condition"`
					Predicate struct {
						Tags []struct {
							Expected bool   `nbt:"expected"`
							Id       string `nbt:"id"`
						} `nbt:"tags"`
					} `nbt:"predicate"`
					Terms []struct {
						Condition string `nbt:"condition"`
						Predicate struct {
							Tags []struct {
								Expected bool   `nbt:"expected"`
								Id       string `nbt:"id"`
							} `nbt:"tags"`
						} `nbt:"predicate"`
					} `nbt:"terms"`
				} `nbt:"requirements"`
			} `nbt:"minecraft:damage_protection"`
			Damage []struct {
				Effect struct {
					Type  string `nbt:"type"`
					Value struct {
						Base               float32 `nbt:"base"`
						PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
						Type               string  `nbt:"type"`
					} `nbt:"value"`
				} `nbt:"effect"`
				Requirements struct {
					Condition string `nbt:"condition"`
					Entity    string `nbt:"entity"`
					Predicate struct {
						Type string `nbt:"type"`
					} `nbt:"predicate"`
				} `nbt:"requirements"`
			} `nbt:"minecraft:damage"`
			PostAttack []struct {
				Affected string `nbt:"affected"`
				Effect   struct {
					MaxAmplifier float32 `nbt:"max_amplifier"`
					MaxDuration  struct {
						Base               float32 `nbt:"base"`
						PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
						Type               string  `nbt:"type"`
					} `nbt:"max_duration"`
					Duration struct {
						Base               float32 `nbt:"base"`
						PerLevelAboveFirst float32 `nbt:"per_level_above_first"`
						Type               string  `nbt:"type"`
					} `nbt:"duration"`
					MinAmplifier float32 `nbt:"min_amplifier"`
					MinDuration  float32 `nbt:"min_duration"`
					ToApply      string  `nbt:"to_apply"`
					Type         string  `nbt:"type"`

					Effects []struct {
						Type string `nbt:"type"`

						Entity string `nbt:"entity"`

						Pitch  float32 `nbt:"pitch"`
						Sound  string  `nbt:"sound"`
						Volume float32 `nbt:"volume"`
					} `nbt:"effects"`
				} `nbt:"effect"`
				Enchanted    string `nbt:"enchanted"`
				Requirements struct {
					Condition string `nbt:"condition"`
					Predicate struct {
						IsDirect bool `nbt:"is_direct"`
					} `nbt:"predicate"`
					Terms []struct {
						Condition  string `nbt:"condition"`
						Thundering bool   `nbt:"thundering"`
						Entity     string `nbt:"entity"`
						Predicate  struct {
							IsDirect bool   `nbt:"is_direct"`
							Type     string `nbt:"type"`
							Location struct {
								CanSeeSky bool `nbt:"can_see_sky"`
							} `nbt:"location"`
						} `nbt:"predicate"`
					} `nbt:"terms"`
				} `nbt:"requirements"`
			} `nbt:"minecraft:post_attack"`
		} `nbt:"effects"`
		MaxCost struct {
			Base               int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MinCost struct {
			Base               int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		MaxLevel       int32    `nbt:"max_level"`
		Slots          []string `nbt:"slots"`
		SupportedItems string   `nbt:"supported_items"`
		Weight         int32    `nbt:"weight"`
		ExclusiveSet   string   `nbt:"exclusive_set"`
		PrimaryItems   string   `nbt:"primary_items"`
	} `nbt:"minecraft:enchantment"`*/
	JukeboxSong map[string]struct {
		ComparatorOutput int32 `nbt:"comparator_output"`
		Description      struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		LengthInSeconds float32 `nbt:"length_in_seconds"`
		SoundEvent      string  `nbt:"sound_event"`
	} `nbt:"minecraft:jukebox_song"`
	PaintingVariant map[string]struct {
		AssetId string `nbt:"asset_id"`
		Height  int32  `nbt:"height"`
		Weight  int32  `nbt:"width"`
	} `nbt:"minecraft:painting_variant"`
	TrimMaterial map[string]struct {
		AssetName   string `nbt:"asset_name"`
		Description struct {
			Color     string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Ingredient             string  `nbt:"ingredient"`
		ItemModelIndex         float32 `nbt:"item_model_index"`
		OverrideArmorMaterials struct {
			Diamond string `nbt:"minecraft:diamond,omitempty"`
			Gold    string `nbt:"minecraft:gold,omitempty"`
			Iron    string `nbt:"minecraft:iron,omitempty"`
		} `nbt:"override_armor_materials,omitempty"`
	} `nbt:"minecraft:trim_material"`
	TrimPattern map[string]struct {
		AssetId     string `nbt:"asset_id"`
		Decal       bool   `nbt:"decal"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		TemplateItem string `nbt:"template_item"`
	} `nbt:"minecraft:trim_pattern"`
	WolfVariant map[string]struct {
		AngryTexture string `nbt:"angry_texture"`
		Biomes       string `nbt:"biomes"`
		TameTexture  string `nbt:"tame_texture"`
		WildTexture  string `nbt:"wild_texture"`
	} `nbt:"minecraft:wolf_variant"`
	WorldgenBiome map[string]struct {
		Downfall float32 `nbt:"downfall"`
		Effects  struct {
			FogColor     int32 `nbt:"fog_color"`
			FoliageColor int32 `nbt:"foliage_color,omitempty"`
			GrassColor   int32 `nbt:"grass_color,omitempty"`
			MoodSound    struct {
				BlockSearchExtent int32   `nbt:"block_search_extent"`
				Offset            float64 `nbt:"offset"`
				Sound             string  `nbt:"sound"`
				TickDelay         int32   `nbt:"tick_delay"`
			} `nbt:"mood_sound"`
			Music struct {
				MaxDelay            int32  `nbt:"max_delay"`
				MinDelay            int32  `nbt:"min_delay"`
				ReplaceCurrentMusic bool   `nbt:"replace_current_music"`
				Sound               string `nbt:"sound"`
			} `nbt:"music,omitempty"`
			SkyColor      int32 `nbt:"sky_color"`
			WaterColor    int32 `nbt:"water_color"`
			WaterFogColor int32 `nbt:"water_fog_color"`
		} `nbt:"effects"`
		HasPrecipitation bool    `nbt:"has_precipitation"`
		Temperature      float32 `nbt:"temperature"`
	} `nbt:"minecraft:worldgen/biome"`
}

//go:embed registries.nbt
var registriesFile []byte

func LoadRegistry() error {
	_, err := nbt.NewDecoder(bytes.NewReader(registriesFile)).Decode(&Registries)

	if err != nil {
		return err
	}
	v := reflect.ValueOf(Registries)
	for i := 0; i < v.NumField(); i++ {
		RegistryMap[v.Type().Field(i).Tag.Get("nbt")] = v.Field(i).Interface()
	}

	return err
}
