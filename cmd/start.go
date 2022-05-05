/*
Copyright © 2022 Dackota Johnson dackota.j@gmail.com

*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Long:  `Start the server to begin listening for requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		router()
	},
}

var characters = []Character{}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func router() {
	router := gin.Default()
	router.GET("/characters", getCharacters)
	router.POST("/characters", addCharacter)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters)
}

// addCharacter adds a character from JSON received in the request body.
func addCharacter(c *gin.Context) {
	var newCharacter Character

	// Call BindJSON to bind the received JSON to
	// newCharacter.
	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	// Add the new character to the slice.
	characters = append(characters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)
}

type Character struct {
	ID      int    `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID                         int         `json:"id"`
		UserID                     int         `json:"userId"`
		ReadonlyURL                string      `json:"readonlyUrl"`
		AvatarURL                  string      `json:"avatarUrl"`
		FrameAvatarURL             interface{} `json:"frameAvatarUrl"`
		BackdropAvatarURL          interface{} `json:"backdropAvatarUrl"`
		SmallBackdropAvatarURL     interface{} `json:"smallBackdropAvatarUrl"`
		LargeBackdropAvatarURL     interface{} `json:"largeBackdropAvatarUrl"`
		ThumbnailBackdropAvatarURL interface{} `json:"thumbnailBackdropAvatarUrl"`
		DefaultBackdrop            struct {
			BackdropAvatarURL          string `json:"backdropAvatarUrl"`
			SmallBackdropAvatarURL     string `json:"smallBackdropAvatarUrl"`
			LargeBackdropAvatarURL     string `json:"largeBackdropAvatarUrl"`
			ThumbnailBackdropAvatarURL string `json:"thumbnailBackdropAvatarUrl"`
		} `json:"defaultBackdrop"`
		AvatarID                  int         `json:"avatarId"`
		FrameAvatarID             interface{} `json:"frameAvatarId"`
		BackdropAvatarID          interface{} `json:"backdropAvatarId"`
		SmallBackdropAvatarID     interface{} `json:"smallBackdropAvatarId"`
		LargeBackdropAvatarID     interface{} `json:"largeBackdropAvatarId"`
		ThumbnailBackdropAvatarID interface{} `json:"thumbnailBackdropAvatarId"`
		ThemeColorID              interface{} `json:"themeColorId"`
		ThemeColor                interface{} `json:"themeColor"`
		Decorations               struct {
			AvatarURL                  string      `json:"avatarUrl"`
			FrameAvatarURL             interface{} `json:"frameAvatarUrl"`
			BackdropAvatarURL          interface{} `json:"backdropAvatarUrl"`
			SmallBackdropAvatarURL     interface{} `json:"smallBackdropAvatarUrl"`
			LargeBackdropAvatarURL     interface{} `json:"largeBackdropAvatarUrl"`
			ThumbnailBackdropAvatarURL interface{} `json:"thumbnailBackdropAvatarUrl"`
			DefaultBackdrop            struct {
				BackdropAvatarURL          string `json:"backdropAvatarUrl"`
				SmallBackdropAvatarURL     string `json:"smallBackdropAvatarUrl"`
				LargeBackdropAvatarURL     string `json:"largeBackdropAvatarUrl"`
				ThumbnailBackdropAvatarURL string `json:"thumbnailBackdropAvatarUrl"`
			} `json:"defaultBackdrop"`
			AvatarID                             int         `json:"avatarId"`
			PortraitDecorationKey                interface{} `json:"portraitDecorationKey"`
			FrameAvatarDecorationKey             interface{} `json:"frameAvatarDecorationKey"`
			FrameAvatarID                        interface{} `json:"frameAvatarId"`
			BackdropAvatarDecorationKey          interface{} `json:"backdropAvatarDecorationKey"`
			BackdropAvatarID                     interface{} `json:"backdropAvatarId"`
			SmallBackdropAvatarDecorationKey     string      `json:"smallBackdropAvatarDecorationKey"`
			SmallBackdropAvatarID                interface{} `json:"smallBackdropAvatarId"`
			LargeBackdropAvatarDecorationKey     string      `json:"largeBackdropAvatarDecorationKey"`
			LargeBackdropAvatarID                interface{} `json:"largeBackdropAvatarId"`
			ThumbnailBackdropAvatarDecorationKey string      `json:"thumbnailBackdropAvatarDecorationKey"`
			ThumbnailBackdropAvatarID            interface{} `json:"thumbnailBackdropAvatarId"`
			ThemeColor                           interface{} `json:"themeColor"`
		} `json:"decorations"`
		Name               string      `json:"name"`
		SocialName         interface{} `json:"socialName"`
		Gender             string      `json:"gender"`
		Faith              interface{} `json:"faith"`
		Age                int         `json:"age"`
		Hair               string      `json:"hair"`
		Eyes               string      `json:"eyes"`
		Skin               string      `json:"skin"`
		Height             string      `json:"height"`
		Weight             int         `json:"weight"`
		Inspiration        bool        `json:"inspiration"`
		BaseHitPoints      int         `json:"baseHitPoints"`
		BonusHitPoints     interface{} `json:"bonusHitPoints"`
		OverrideHitPoints  int         `json:"overrideHitPoints"`
		RemovedHitPoints   int         `json:"removedHitPoints"`
		TemporaryHitPoints int         `json:"temporaryHitPoints"`
		CurrentXp          int         `json:"currentXp"`
		AlignmentID        int         `json:"alignmentId"`
		LifestyleID        int         `json:"lifestyleId"`
		Stats              []struct {
			ID    int         `json:"id"`
			Name  interface{} `json:"name"`
			Value int         `json:"value"`
		} `json:"stats"`
		BonusStats []struct {
			ID    int         `json:"id"`
			Name  interface{} `json:"name"`
			Value interface{} `json:"value"`
		} `json:"bonusStats"`
		OverrideStats []struct {
			ID    int         `json:"id"`
			Name  interface{} `json:"name"`
			Value interface{} `json:"value"`
		} `json:"overrideStats"`
		Background struct {
			HasCustomBackground bool `json:"hasCustomBackground"`
			Definition          struct {
				ID                                  int         `json:"id"`
				EntityTypeID                        int         `json:"entityTypeId"`
				Name                                string      `json:"name"`
				Description                         string      `json:"description"`
				Snippet                             string      `json:"snippet"`
				ShortDescription                    string      `json:"shortDescription"`
				SkillProficienciesDescription       string      `json:"skillProficienciesDescription"`
				ToolProficienciesDescription        string      `json:"toolProficienciesDescription"`
				LanguagesDescription                string      `json:"languagesDescription"`
				EquipmentDescription                string      `json:"equipmentDescription"`
				FeatureName                         string      `json:"featureName"`
				FeatureDescription                  string      `json:"featureDescription"`
				AvatarURL                           interface{} `json:"avatarUrl"`
				LargeAvatarURL                      interface{} `json:"largeAvatarUrl"`
				SuggestedCharacteristicsDescription string      `json:"suggestedCharacteristicsDescription"`
				SuggestedProficiencies              interface{} `json:"suggestedProficiencies"`
				SuggestedLanguages                  interface{} `json:"suggestedLanguages"`
				Organization                        interface{} `json:"organization"`
				ContractsDescription                string      `json:"contractsDescription"`
				SpellsPreDescription                string      `json:"spellsPreDescription"`
				SpellsPostDescription               string      `json:"spellsPostDescription"`
				PersonalityTraits                   []struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					DiceRoll    int    `json:"diceRoll"`
				} `json:"personalityTraits"`
				Ideals []struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					DiceRoll    int    `json:"diceRoll"`
				} `json:"ideals"`
				Bonds []struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					DiceRoll    int    `json:"diceRoll"`
				} `json:"bonds"`
				Flaws []struct {
					ID          int    `json:"id"`
					Description string `json:"description"`
					DiceRoll    int    `json:"diceRoll"`
				} `json:"flaws"`
				IsHomebrew bool `json:"isHomebrew"`
				Sources    []struct {
					SourceID   int         `json:"sourceId"`
					PageNumber interface{} `json:"pageNumber"`
					SourceType int         `json:"sourceType"`
				} `json:"sources"`
				SpellListIds []interface{} `json:"spellListIds"`
			} `json:"definition"`
			DefinitionID     interface{} `json:"definitionId"`
			CustomBackground struct {
				ID                                    int         `json:"id"`
				EntityTypeID                          int         `json:"entityTypeId"`
				Name                                  interface{} `json:"name"`
				Description                           interface{} `json:"description"`
				FeaturesBackground                    interface{} `json:"featuresBackground"`
				CharacteristicsBackground             interface{} `json:"characteristicsBackground"`
				FeaturesBackgroundDefinitionID        interface{} `json:"featuresBackgroundDefinitionId"`
				CharacteristicsBackgroundDefinitionID interface{} `json:"characteristicsBackgroundDefinitionId"`
				BackgroundType                        interface{} `json:"backgroundType"`
			} `json:"customBackground"`
		} `json:"background"`
		Race struct {
			IsSubRace         bool          `json:"isSubRace"`
			BaseRaceName      string        `json:"baseRaceName"`
			EntityRaceID      int           `json:"entityRaceId"`
			EntityRaceTypeID  int           `json:"entityRaceTypeId"`
			FullName          string        `json:"fullName"`
			BaseRaceID        int           `json:"baseRaceId"`
			BaseRaceTypeID    int           `json:"baseRaceTypeId"`
			Description       string        `json:"description"`
			AvatarURL         string        `json:"avatarUrl"`
			LargeAvatarURL    string        `json:"largeAvatarUrl"`
			PortraitAvatarURL string        `json:"portraitAvatarUrl"`
			MoreDetailsURL    string        `json:"moreDetailsUrl"`
			IsHomebrew        bool          `json:"isHomebrew"`
			SourceIds         []interface{} `json:"sourceIds"`
			GroupIds          []int         `json:"groupIds"`
			Type              int           `json:"type"`
			SupportsSubrace   interface{}   `json:"supportsSubrace"`
			SubRaceShortName  interface{}   `json:"subRaceShortName"`
			BaseName          string        `json:"baseName"`
			RacialTraits      []struct {
				Definition struct {
					ID               int           `json:"id"`
					DefinitionKey    string        `json:"definitionKey"`
					EntityTypeID     int           `json:"entityTypeId"`
					DisplayOrder     int           `json:"displayOrder"`
					Name             string        `json:"name"`
					Description      string        `json:"description"`
					Snippet          string        `json:"snippet"`
					HideInBuilder    bool          `json:"hideInBuilder"`
					HideInSheet      bool          `json:"hideInSheet"`
					Activation       interface{}   `json:"activation"`
					SourceID         int           `json:"sourceId"`
					SourcePageNumber int           `json:"sourcePageNumber"`
					CreatureRules    []interface{} `json:"creatureRules"`
					SpellListIds     []interface{} `json:"spellListIds"`
					FeatureType      int           `json:"featureType"`
					Sources          []struct {
						SourceID   int `json:"sourceId"`
						PageNumber int `json:"pageNumber"`
						SourceType int `json:"sourceType"`
					} `json:"sources"`
					AffectedFeatureDefinitionKeys []interface{} `json:"affectedFeatureDefinitionKeys"`
					IsCalledOut                   bool          `json:"isCalledOut"`
					EntityType                    string        `json:"entityType"`
					EntityID                      string        `json:"entityID"`
					EntityRaceID                  int           `json:"entityRaceId"`
					EntityRaceTypeID              int           `json:"entityRaceTypeId"`
					DisplayConfiguration          struct {
						Racialtrait  int `json:"RACIALTRAIT"`
						Abilityscore int `json:"ABILITYSCORE"`
						Language     int `json:"LANGUAGE"`
						Classfeature int `json:"CLASSFEATURE"`
					} `json:"displayConfiguration"`
					RequiredLevel interface{} `json:"requiredLevel"`
				} `json:"definition"`
			} `json:"racialTraits"`
			WeightSpeeds struct {
				Normal struct {
					Walk   int `json:"walk"`
					Fly    int `json:"fly"`
					Burrow int `json:"burrow"`
					Swim   int `json:"swim"`
					Climb  int `json:"climb"`
				} `json:"normal"`
				Encumbered        interface{} `json:"encumbered"`
				HeavilyEncumbered interface{} `json:"heavilyEncumbered"`
				PushDragLift      interface{} `json:"pushDragLift"`
				Override          interface{} `json:"override"`
			} `json:"weightSpeeds"`
			FeatIds []interface{} `json:"featIds"`
			Size    interface{}   `json:"size"`
			SizeID  int           `json:"sizeId"`
			Sources []struct {
				SourceID   int         `json:"sourceId"`
				PageNumber interface{} `json:"pageNumber"`
				SourceType int         `json:"sourceType"`
			} `json:"sources"`
		} `json:"race"`
		RaceDefinitionID     interface{} `json:"raceDefinitionId"`
		RaceDefinitionTypeID interface{} `json:"raceDefinitionTypeId"`
		Notes                struct {
			Allies              string      `json:"allies"`
			PersonalPossessions string      `json:"personalPossessions"`
			OtherHoldings       interface{} `json:"otherHoldings"`
			Organizations       string      `json:"organizations"`
			Enemies             string      `json:"enemies"`
			Backstory           string      `json:"backstory"`
			OtherNotes          string      `json:"otherNotes"`
		} `json:"notes"`
		Traits struct {
			PersonalityTraits string      `json:"personalityTraits"`
			Ideals            string      `json:"ideals"`
			Bonds             string      `json:"bonds"`
			Flaws             string      `json:"flaws"`
			Appearance        interface{} `json:"appearance"`
		} `json:"traits"`
		Preferences struct {
			UseHomebrewContent          bool `json:"useHomebrewContent"`
			ProgressionType             int  `json:"progressionType"`
			EncumbranceType             int  `json:"encumbranceType"`
			IgnoreCoinWeight            bool `json:"ignoreCoinWeight"`
			HitPointType                int  `json:"hitPointType"`
			ShowUnarmedStrike           bool `json:"showUnarmedStrike"`
			ShowScaledSpells            bool `json:"showScaledSpells"`
			PrimarySense                int  `json:"primarySense"`
			PrimaryMovement             int  `json:"primaryMovement"`
			PrivacyType                 int  `json:"privacyType"`
			SharingType                 int  `json:"sharingType"`
			AbilityScoreDisplayType     int  `json:"abilityScoreDisplayType"`
			EnforceFeatRules            bool `json:"enforceFeatRules"`
			EnforceMulticlassRules      bool `json:"enforceMulticlassRules"`
			EnableOptionalClassFeatures bool `json:"enableOptionalClassFeatures"`
			EnableOptionalOrigins       bool `json:"enableOptionalOrigins"`
			EnableDarkMode              bool `json:"enableDarkMode"`
			EnableContainerCurrency     bool `json:"enableContainerCurrency"`
		} `json:"preferences"`
		Configuration struct {
			StartingEquipmentType int  `json:"startingEquipmentType"`
			AbilityScoreType      int  `json:"abilityScoreType"`
			ShowHelpText          bool `json:"showHelpText"`
		} `json:"configuration"`
		Lifestyle interface{} `json:"lifestyle"`
		Inventory []struct {
			ID           int `json:"id"`
			EntityTypeID int `json:"entityTypeId"`
			Definition   struct {
				ID                    int         `json:"id"`
				BaseTypeID            int         `json:"baseTypeId"`
				EntityTypeID          int         `json:"entityTypeId"`
				CanEquip              bool        `json:"canEquip"`
				Magic                 bool        `json:"magic"`
				Name                  string      `json:"name"`
				Snippet               string      `json:"snippet"`
				Weight                float64     `json:"weight"`
				WeightMultiplier      float64     `json:"weightMultiplier"`
				Capacity              interface{} `json:"capacity"`
				CapacityWeight        float32     `json:"capacityWeight"`
				Type                  string      `json:"type"`
				Description           string      `json:"description"`
				CanAttune             bool        `json:"canAttune"`
				AttunementDescription string      `json:"attunementDescription"`
				Rarity                string      `json:"rarity"`
				IsHomebrew            bool        `json:"isHomebrew"`
				Version               interface{} `json:"version"`
				SourceID              interface{} `json:"sourceId"`
				SourcePageNumber      interface{} `json:"sourcePageNumber"`
				Stackable             bool        `json:"stackable"`
				BundleSize            int         `json:"bundleSize"`
				AvatarURL             interface{} `json:"avatarUrl"`
				LargeAvatarURL        interface{} `json:"largeAvatarUrl"`
				FilterType            string      `json:"filterType"`
				Cost                  interface{} `json:"cost"`
				IsPack                bool        `json:"isPack"`
				Tags                  []string    `json:"tags"`
				GrantedModifiers      []struct {
					FixedValue   int         `json:"fixedValue"`
					ID           string      `json:"id"`
					EntityID     interface{} `json:"entityId"`
					EntityTypeID interface{} `json:"entityTypeId"`
					Type         string      `json:"type"`
					SubType      string      `json:"subType"`
					Dice         struct {
						DiceCount      int         `json:"diceCount"`
						DiceValue      int         `json:"diceValue"`
						DiceMultiplier interface{} `json:"diceMultiplier"`
						FixedValue     int         `json:"fixedValue"`
						DiceString     string      `json:"diceString"`
					} `json:"dice"`
					Restriction           string        `json:"restriction"`
					StatID                interface{}   `json:"statId"`
					RequiresAttunement    bool          `json:"requiresAttunement"`
					Duration              interface{}   `json:"duration"`
					FriendlyTypeName      string        `json:"friendlyTypeName"`
					FriendlySubtypeName   string        `json:"friendlySubtypeName"`
					IsGranted             bool          `json:"isGranted"`
					BonusTypes            []interface{} `json:"bonusTypes"`
					Value                 interface{}   `json:"value"`
					AvailableToMulticlass bool          `json:"availableToMulticlass"`
					ModifierTypeID        int           `json:"modifierTypeId"`
					ModifierSubTypeID     int           `json:"modifierSubTypeId"`
					ComponentID           int           `json:"componentId"`
					ComponentTypeID       int           `json:"componentTypeId"`
				} `json:"grantedModifiers"`
				SubType              interface{}   `json:"subType"`
				IsConsumable         bool          `json:"isConsumable"`
				WeaponBehaviors      []interface{} `json:"weaponBehaviors"`
				BaseItemID           interface{}   `json:"baseItemId"`
				BaseArmorName        interface{}   `json:"baseArmorName"`
				StrengthRequirement  interface{}   `json:"strengthRequirement"`
				ArmorClass           interface{}   `json:"armorClass"`
				StealthCheck         interface{}   `json:"stealthCheck"`
				Damage               interface{}   `json:"damage"`
				DamageType           interface{}   `json:"damageType"`
				FixedDamage          interface{}   `json:"fixedDamage"`
				Properties           interface{}   `json:"properties"`
				AttackType           interface{}   `json:"attackType"`
				CategoryID           interface{}   `json:"categoryId"`
				Range                interface{}   `json:"range"`
				LongRange            interface{}   `json:"longRange"`
				IsMonkWeapon         bool          `json:"isMonkWeapon"`
				LevelInfusionGranted interface{}   `json:"levelInfusionGranted"`
				Sources              []struct {
					SourceID   int         `json:"sourceId"`
					PageNumber interface{} `json:"pageNumber"`
					SourceType int         `json:"sourceType"`
				} `json:"sources"`
				ArmorTypeID           interface{} `json:"armorTypeId"`
				GearTypeID            int         `json:"gearTypeId"`
				GroupedID             int         `json:"groupedId"`
				CanBeAddedToInventory bool        `json:"canBeAddedToInventory"`
				IsContainer           bool        `json:"isContainer"`
				IsCustomItem          bool        `json:"isCustomItem"`
			} `json:"definition"`
			DefinitionID         int         `json:"definitionId"`
			DefinitionTypeID     int         `json:"definitionTypeId"`
			DisplayAsAttack      interface{} `json:"displayAsAttack"`
			Quantity             int         `json:"quantity"`
			IsAttuned            bool        `json:"isAttuned"`
			Equipped             bool        `json:"equipped"`
			EquippedEntityTypeID interface{} `json:"equippedEntityTypeId"`
			EquippedEntityID     interface{} `json:"equippedEntityId"`
			ChargesUsed          int         `json:"chargesUsed"`
			LimitedUse           struct {
				MaxUses              int    `json:"maxUses"`
				NumberUsed           int    `json:"numberUsed"`
				ResetType            string `json:"resetType"`
				ResetTypeDescription string `json:"resetTypeDescription"`
			} `json:"limitedUse"`
			ContainerEntityID      int         `json:"containerEntityId"`
			ContainerEntityTypeID  int         `json:"containerEntityTypeId"`
			ContainerDefinitionKey string      `json:"containerDefinitionKey"`
			Currency               interface{} `json:"currency"`
		} `json:"inventory"`
		Currencies struct {
			Cp int `json:"cp"`
			Sp int `json:"sp"`
			Gp int `json:"gp"`
			Ep int `json:"ep"`
			Pp int `json:"pp"`
		} `json:"currencies"`
		Classes []struct {
			ID                   int         `json:"id"`
			EntityTypeID         int         `json:"entityTypeId"`
			Level                int         `json:"level"`
			IsStartingClass      bool        `json:"isStartingClass"`
			HitDiceUsed          int         `json:"hitDiceUsed"`
			DefinitionID         int         `json:"definitionId"`
			SubclassDefinitionID interface{} `json:"subclassDefinitionId"`
			Definition           struct {
				ID                    int           `json:"id"`
				Name                  string        `json:"name"`
				Description           string        `json:"description"`
				EquipmentDescription  string        `json:"equipmentDescription"`
				ParentClassID         interface{}   `json:"parentClassId"`
				AvatarURL             string        `json:"avatarUrl"`
				LargeAvatarURL        string        `json:"largeAvatarUrl"`
				PortraitAvatarURL     string        `json:"portraitAvatarUrl"`
				MoreDetailsURL        string        `json:"moreDetailsUrl"`
				SpellCastingAbilityID int           `json:"spellCastingAbilityId"`
				SourceIds             []interface{} `json:"sourceIds"`
				Sources               []struct {
					SourceID   int `json:"sourceId"`
					PageNumber int `json:"pageNumber"`
					SourceType int `json:"sourceType"`
				} `json:"sources"`
				HitDice       int `json:"hitDice"`
				ClassFeatures []struct {
					ID            int         `json:"id"`
					Name          string      `json:"name"`
					Prerequisite  interface{} `json:"prerequisite"`
					Description   string      `json:"description"`
					RequiredLevel int         `json:"requiredLevel"`
					DisplayOrder  int         `json:"displayOrder"`
				} `json:"classFeatures"`
				ClassFeatureDefinitions interface{} `json:"classFeatureDefinitions"`
				WealthDice              struct {
					DiceCount      int         `json:"diceCount"`
					DiceValue      int         `json:"diceValue"`
					DiceMultiplier int         `json:"diceMultiplier"`
					FixedValue     interface{} `json:"fixedValue"`
					DiceString     string      `json:"diceString"`
				} `json:"wealthDice"`
				CanCastSpells      bool        `json:"canCastSpells"`
				KnowsAllSpells     bool        `json:"knowsAllSpells"`
				SpellPrepareType   interface{} `json:"spellPrepareType"`
				SpellContainerName interface{} `json:"spellContainerName"`
				SourceID           interface{} `json:"sourceId"`
				SourcePageNumber   int         `json:"sourcePageNumber"`
				SubclassDefinition interface{} `json:"subclassDefinition"`
				IsHomebrew         bool        `json:"isHomebrew"`
				PrimaryAbilities   []int       `json:"primaryAbilities"`
				SpellRules         struct {
					MultiClassSpellSlotDivisor  int     `json:"multiClassSpellSlotDivisor"`
					IsRitualSpellCaster         bool    `json:"isRitualSpellCaster"`
					LevelCantripsKnownMaxes     []int   `json:"levelCantripsKnownMaxes"`
					LevelSpellKnownMaxes        []int   `json:"levelSpellKnownMaxes"`
					LevelSpellSlots             [][]int `json:"levelSpellSlots"`
					MultiClassSpellSlotRounding int     `json:"multiClassSpellSlotRounding"`
				} `json:"spellRules"`
				Prerequisites []struct {
					Description          string `json:"description"`
					PrerequisiteMappings []struct {
						ID                  int    `json:"id"`
						EntityID            int    `json:"entityId"`
						EntityTypeID        int    `json:"entityTypeId"`
						Type                string `json:"type"`
						SubType             string `json:"subType"`
						Value               int    `json:"value"`
						FriendlyTypeName    string `json:"friendlyTypeName"`
						FriendlySubTypeName string `json:"friendlySubTypeName"`
					} `json:"prerequisiteMappings"`
				} `json:"prerequisites"`
			} `json:"definition"`
			SubclassDefinition struct {
				ID                    int           `json:"id"`
				Name                  string        `json:"name"`
				Description           string        `json:"description"`
				EquipmentDescription  interface{}   `json:"equipmentDescription"`
				ParentClassID         int           `json:"parentClassId"`
				AvatarURL             interface{}   `json:"avatarUrl"`
				LargeAvatarURL        interface{}   `json:"largeAvatarUrl"`
				PortraitAvatarURL     interface{}   `json:"portraitAvatarUrl"`
				MoreDetailsURL        string        `json:"moreDetailsUrl"`
				SpellCastingAbilityID int           `json:"spellCastingAbilityId"`
				SourceIds             []interface{} `json:"sourceIds"`
				Sources               []struct {
					SourceID   int `json:"sourceId"`
					PageNumber int `json:"pageNumber"`
					SourceType int `json:"sourceType"`
				} `json:"sources"`
				HitDice       int `json:"hitDice"`
				ClassFeatures []struct {
					ID            int         `json:"id"`
					Name          string      `json:"name"`
					Prerequisite  interface{} `json:"prerequisite"`
					Description   string      `json:"description"`
					RequiredLevel int         `json:"requiredLevel"`
					DisplayOrder  int         `json:"displayOrder"`
				} `json:"classFeatures"`
				ClassFeatureDefinitions interface{} `json:"classFeatureDefinitions"`
				WealthDice              interface{} `json:"wealthDice"`
				CanCastSpells           bool        `json:"canCastSpells"`
				KnowsAllSpells          bool        `json:"knowsAllSpells"`
				SpellPrepareType        interface{} `json:"spellPrepareType"`
				SpellContainerName      interface{} `json:"spellContainerName"`
				SourceID                interface{} `json:"sourceId"`
				SourcePageNumber        int         `json:"sourcePageNumber"`
				SubclassDefinition      interface{} `json:"subclassDefinition"`
				IsHomebrew              bool        `json:"isHomebrew"`
				PrimaryAbilities        interface{} `json:"primaryAbilities"`
				SpellRules              interface{} `json:"spellRules"`
				Prerequisites           interface{} `json:"prerequisites"`
			} `json:"subclassDefinition"`
			ClassFeatures []struct {
				Definition struct {
					ID                    int           `json:"id"`
					DefinitionKey         string        `json:"definitionKey"`
					EntityTypeID          int           `json:"entityTypeId"`
					DisplayOrder          int           `json:"displayOrder"`
					Name                  string        `json:"name"`
					Description           string        `json:"description"`
					Snippet               string        `json:"snippet"`
					Activation            interface{}   `json:"activation"`
					MultiClassDescription string        `json:"multiClassDescription"`
					RequiredLevel         int           `json:"requiredLevel"`
					IsSubClassFeature     bool          `json:"isSubClassFeature"`
					LimitedUse            []interface{} `json:"limitedUse"`
					HideInBuilder         bool          `json:"hideInBuilder"`
					HideInSheet           bool          `json:"hideInSheet"`
					SourceID              int           `json:"sourceId"`
					SourcePageNumber      int           `json:"sourcePageNumber"`
					CreatureRules         []interface{} `json:"creatureRules"`
					LevelScales           []interface{} `json:"levelScales"`
					InfusionRules         []interface{} `json:"infusionRules"`
					SpellListIds          []interface{} `json:"spellListIds"`
					ClassID               int           `json:"classId"`
					FeatureType           int           `json:"featureType"`
					Sources               []struct {
						SourceID   int `json:"sourceId"`
						PageNumber int `json:"pageNumber"`
						SourceType int `json:"sourceType"`
					} `json:"sources"`
					AffectedFeatureDefinitionKeys []interface{} `json:"affectedFeatureDefinitionKeys"`
					EntityType                    string        `json:"entityType"`
					EntityID                      string        `json:"entityID"`
				} `json:"definition"`
				LevelScale interface{} `json:"levelScale"`
			} `json:"classFeatures"`
		} `json:"classes"`
		Feats []struct {
			ComponentTypeID interface{} `json:"componentTypeId"`
			ComponentID     interface{} `json:"componentId"`
			Definition      struct {
				ID           int    `json:"id"`
				EntityTypeID int    `json:"entityTypeId"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Snippet      string `json:"snippet"`
				Activation   struct {
					ActivationTime interface{} `json:"activationTime"`
					ActivationType interface{} `json:"activationType"`
				} `json:"activation"`
				SourceID         interface{}   `json:"sourceId"`
				SourcePageNumber interface{}   `json:"sourcePageNumber"`
				CreatureRules    []interface{} `json:"creatureRules"`
				Prerequisites    []interface{} `json:"prerequisites"`
				IsHomebrew       bool          `json:"isHomebrew"`
				Sources          []struct {
					SourceID   int         `json:"sourceId"`
					PageNumber interface{} `json:"pageNumber"`
					SourceType int         `json:"sourceType"`
				} `json:"sources"`
				SpellListIds []interface{} `json:"spellListIds"`
			} `json:"definition"`
			DefinitionID int `json:"definitionId"`
		} `json:"feats"`
		CustomDefenseAdjustments []interface{} `json:"customDefenseAdjustments"`
		CustomSenses             []interface{} `json:"customSenses"`
		CustomSpeeds             []interface{} `json:"customSpeeds"`
		CustomProficiencies      []interface{} `json:"customProficiencies"`
		SpellDefenses            interface{}   `json:"spellDefenses"`
		CustomActions            []interface{} `json:"customActions"`
		CharacterValues          []interface{} `json:"characterValues"`
		Conditions               []struct {
			ID    int `json:"id"`
			Level int `json:"level"`
		} `json:"conditions"`
		DeathSaves struct {
			FailCount    int  `json:"failCount"`
			SuccessCount int  `json:"successCount"`
			IsStabilized bool `json:"isStabilized"`
		} `json:"deathSaves"`
		AdjustmentXp int `json:"adjustmentXp"`
		SpellSlots   []struct {
			Level     int `json:"level"`
			Used      int `json:"used"`
			Available int `json:"available"`
		} `json:"spellSlots"`
		PactMagic []struct {
			Level     int `json:"level"`
			Used      int `json:"used"`
			Available int `json:"available"`
		} `json:"pactMagic"`
		ActiveSourceCategories []int `json:"activeSourceCategories"`
		Spells                 struct {
			Race []struct {
				OverrideSaveDc interface{} `json:"overrideSaveDc"`
				LimitedUse     interface{} `json:"limitedUse"`
				ID             int         `json:"id"`
				EntityTypeID   int         `json:"entityTypeId"`
				Definition     struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Level    int    `json:"level"`
					School   string `json:"school"`
					Duration struct {
						DurationInterval int    `json:"durationInterval"`
						DurationUnit     string `json:"durationUnit"`
						DurationType     string `json:"durationType"`
					} `json:"duration"`
					Activation struct {
						ActivationTime int `json:"activationTime"`
						ActivationType int `json:"activationType"`
					} `json:"activation"`
					Range struct {
						Origin     string      `json:"origin"`
						RangeValue int         `json:"rangeValue"`
						AoeType    interface{} `json:"aoeType"`
						AoeValue   interface{} `json:"aoeValue"`
					} `json:"range"`
					AsPartOfWeaponAttack  bool          `json:"asPartOfWeaponAttack"`
					Description           string        `json:"description"`
					Snippet               string        `json:"snippet"`
					Concentration         bool          `json:"concentration"`
					Ritual                bool          `json:"ritual"`
					RangeArea             interface{}   `json:"rangeArea"`
					DamageEffect          interface{}   `json:"damageEffect"`
					Components            []int         `json:"components"`
					ComponentsDescription string        `json:"componentsDescription"`
					SaveDcAbilityID       interface{}   `json:"saveDcAbilityId"`
					Healing               interface{}   `json:"healing"`
					HealingDice           []interface{} `json:"healingDice"`
					TempHpDice            []interface{} `json:"tempHpDice"`
					AttackType            interface{}   `json:"attackType"`
					CanCastAtHigherLevel  bool          `json:"canCastAtHigherLevel"`
					IsHomebrew            bool          `json:"isHomebrew"`
					Version               interface{}   `json:"version"`
					SourceID              interface{}   `json:"sourceId"`
					SourcePageNumber      int           `json:"sourcePageNumber"`
					RequiresSavingThrow   bool          `json:"requiresSavingThrow"`
					RequiresAttackRoll    bool          `json:"requiresAttackRoll"`
					AtHigherLevels        struct {
						ScaleType              string        `json:"scaleType"`
						HigherLevelDefinitions []interface{} `json:"higherLevelDefinitions"`
						AdditionalAttacks      []interface{} `json:"additionalAttacks"`
						AdditionalTargets      []interface{} `json:"additionalTargets"`
						AreaOfEffect           []interface{} `json:"areaOfEffect"`
						Duration               []interface{} `json:"duration"`
						Creatures              []interface{} `json:"creatures"`
						Special                []interface{} `json:"special"`
						Points                 []interface{} `json:"points"`
					} `json:"atHigherLevels"`
					Modifiers              []interface{} `json:"modifiers"`
					Conditions             []interface{} `json:"conditions"`
					Tags                   []string      `json:"tags"`
					CastingTimeDescription string        `json:"castingTimeDescription"`
					ScaleType              string        `json:"scaleType"`
					Sources                []struct {
						SourceID   int         `json:"sourceId"`
						PageNumber interface{} `json:"pageNumber"`
						SourceType int         `json:"sourceType"`
					} `json:"sources"`
					SpellGroups []interface{} `json:"spellGroups"`
				} `json:"definition"`
				DefinitionID          int         `json:"definitionId"`
				Prepared              bool        `json:"prepared"`
				CountsAsKnownSpell    bool        `json:"countsAsKnownSpell"`
				UsesSpellSlot         bool        `json:"usesSpellSlot"`
				CastAtLevel           interface{} `json:"castAtLevel"`
				AlwaysPrepared        bool        `json:"alwaysPrepared"`
				Restriction           string      `json:"restriction"`
				SpellCastingAbilityID int         `json:"spellCastingAbilityId"`
				DisplayAsAttack       interface{} `json:"displayAsAttack"`
				AdditionalDescription interface{} `json:"additionalDescription"`
				CastOnlyAsRitual      bool        `json:"castOnlyAsRitual"`
				RitualCastingType     interface{} `json:"ritualCastingType"`
				Range                 struct {
					Origin     string      `json:"origin"`
					RangeValue int         `json:"rangeValue"`
					AoeType    interface{} `json:"aoeType"`
					AoeValue   interface{} `json:"aoeValue"`
				} `json:"range"`
				Activation struct {
					ActivationTime int `json:"activationTime"`
					ActivationType int `json:"activationType"`
				} `json:"activation"`
				BaseLevelAtWill       bool        `json:"baseLevelAtWill"`
				AtWillLimitedUseLevel interface{} `json:"atWillLimitedUseLevel"`
				IsSignatureSpell      interface{} `json:"isSignatureSpell"`
				ComponentID           int         `json:"componentId"`
				ComponentTypeID       int         `json:"componentTypeId"`
				SpellListID           interface{} `json:"spellListId"`
			} `json:"race"`
			Class []struct {
				OverrideSaveDc interface{} `json:"overrideSaveDc"`
				LimitedUse     interface{} `json:"limitedUse"`
				ID             int         `json:"id"`
				EntityTypeID   int         `json:"entityTypeId"`
				Definition     struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Level    int    `json:"level"`
					School   string `json:"school"`
					Duration struct {
						DurationInterval int    `json:"durationInterval"`
						DurationUnit     string `json:"durationUnit"`
						DurationType     string `json:"durationType"`
					} `json:"duration"`
					Activation struct {
						ActivationTime int `json:"activationTime"`
						ActivationType int `json:"activationType"`
					} `json:"activation"`
					Range struct {
						Origin     string      `json:"origin"`
						RangeValue int         `json:"rangeValue"`
						AoeType    interface{} `json:"aoeType"`
						AoeValue   interface{} `json:"aoeValue"`
					} `json:"range"`
					AsPartOfWeaponAttack  bool          `json:"asPartOfWeaponAttack"`
					Description           string        `json:"description"`
					Snippet               string        `json:"snippet"`
					Concentration         bool          `json:"concentration"`
					Ritual                bool          `json:"ritual"`
					RangeArea             interface{}   `json:"rangeArea"`
					DamageEffect          interface{}   `json:"damageEffect"`
					Components            []int         `json:"components"`
					ComponentsDescription string        `json:"componentsDescription"`
					SaveDcAbilityID       interface{}   `json:"saveDcAbilityId"`
					Healing               interface{}   `json:"healing"`
					HealingDice           []interface{} `json:"healingDice"`
					TempHpDice            []interface{} `json:"tempHpDice"`
					AttackType            interface{}   `json:"attackType"`
					CanCastAtHigherLevel  bool          `json:"canCastAtHigherLevel"`
					IsHomebrew            bool          `json:"isHomebrew"`
					Version               interface{}   `json:"version"`
					SourceID              interface{}   `json:"sourceId"`
					SourcePageNumber      int           `json:"sourcePageNumber"`
					RequiresSavingThrow   bool          `json:"requiresSavingThrow"`
					RequiresAttackRoll    bool          `json:"requiresAttackRoll"`
					AtHigherLevels        struct {
						ScaleType              string `json:"scaleType"`
						HigherLevelDefinitions []struct {
							Level   int         `json:"level"`
							TypeID  int         `json:"typeId"`
							Dice    interface{} `json:"dice"`
							Value   int         `json:"value"`
							Details string      `json:"details"`
						} `json:"higherLevelDefinitions"`
						AdditionalAttacks []interface{} `json:"additionalAttacks"`
						AdditionalTargets []interface{} `json:"additionalTargets"`
						AreaOfEffect      []interface{} `json:"areaOfEffect"`
						Duration          []interface{} `json:"duration"`
						Creatures         []interface{} `json:"creatures"`
						Special           []interface{} `json:"special"`
						Points            []interface{} `json:"points"`
					} `json:"atHigherLevels"`
					Modifiers []struct {
						FixedValue            interface{}   `json:"fixedValue"`
						ID                    string        `json:"id"`
						EntityID              interface{}   `json:"entityId"`
						EntityTypeID          interface{}   `json:"entityTypeId"`
						Type                  string        `json:"type"`
						SubType               string        `json:"subType"`
						Dice                  interface{}   `json:"dice"`
						Restriction           string        `json:"restriction"`
						StatID                interface{}   `json:"statId"`
						RequiresAttunement    bool          `json:"requiresAttunement"`
						Duration              interface{}   `json:"duration"`
						FriendlyTypeName      string        `json:"friendlyTypeName"`
						FriendlySubtypeName   string        `json:"friendlySubtypeName"`
						IsGranted             bool          `json:"isGranted"`
						BonusTypes            []interface{} `json:"bonusTypes"`
						Value                 interface{}   `json:"value"`
						AvailableToMulticlass interface{}   `json:"availableToMulticlass"`
						ModifierTypeID        int           `json:"modifierTypeId"`
						ModifierSubTypeID     int           `json:"modifierSubTypeId"`
						ComponentID           int           `json:"componentId"`
						ComponentTypeID       int           `json:"componentTypeId"`
						Die                   struct {
							DiceCount      int         `json:"diceCount"`
							DiceValue      int         `json:"diceValue"`
							DiceMultiplier interface{} `json:"diceMultiplier"`
							FixedValue     interface{} `json:"fixedValue"`
							DiceString     string      `json:"diceString"`
						} `json:"die"`
						Count          int         `json:"count"`
						DurationUnit   interface{} `json:"durationUnit"`
						UsePrimaryStat bool        `json:"usePrimaryStat"`
						AtHigherLevels struct {
							ScaleType              string        `json:"scaleType"`
							HigherLevelDefinitions []interface{} `json:"higherLevelDefinitions"`
							AdditionalAttacks      []interface{} `json:"additionalAttacks"`
							AdditionalTargets      []interface{} `json:"additionalTargets"`
							AreaOfEffect           []interface{} `json:"areaOfEffect"`
							Duration               []interface{} `json:"duration"`
							Creatures              []interface{} `json:"creatures"`
							Special                []interface{} `json:"special"`
							Points                 []interface{} `json:"points"`
						} `json:"atHigherLevels"`
					} `json:"modifiers"`
					Conditions             []interface{} `json:"conditions"`
					Tags                   []string      `json:"tags"`
					CastingTimeDescription string        `json:"castingTimeDescription"`
					ScaleType              string        `json:"scaleType"`
					Sources                []struct {
						SourceID   int         `json:"sourceId"`
						PageNumber interface{} `json:"pageNumber"`
						SourceType int         `json:"sourceType"`
					} `json:"sources"`
					SpellGroups []interface{} `json:"spellGroups"`
				} `json:"definition"`
				DefinitionID          int         `json:"definitionId"`
				Prepared              bool        `json:"prepared"`
				CountsAsKnownSpell    bool        `json:"countsAsKnownSpell"`
				UsesSpellSlot         bool        `json:"usesSpellSlot"`
				CastAtLevel           interface{} `json:"castAtLevel"`
				AlwaysPrepared        bool        `json:"alwaysPrepared"`
				Restriction           string      `json:"restriction"`
				SpellCastingAbilityID interface{} `json:"spellCastingAbilityId"`
				DisplayAsAttack       interface{} `json:"displayAsAttack"`
				AdditionalDescription string      `json:"additionalDescription"`
				CastOnlyAsRitual      bool        `json:"castOnlyAsRitual"`
				RitualCastingType     interface{} `json:"ritualCastingType"`
				Range                 struct {
					Origin     string      `json:"origin"`
					RangeValue int         `json:"rangeValue"`
					AoeType    interface{} `json:"aoeType"`
					AoeValue   interface{} `json:"aoeValue"`
				} `json:"range"`
				Activation struct {
					ActivationTime int `json:"activationTime"`
					ActivationType int `json:"activationType"`
				} `json:"activation"`
				BaseLevelAtWill       bool        `json:"baseLevelAtWill"`
				AtWillLimitedUseLevel interface{} `json:"atWillLimitedUseLevel"`
				IsSignatureSpell      interface{} `json:"isSignatureSpell"`
				ComponentID           int         `json:"componentId"`
				ComponentTypeID       int         `json:"componentTypeId"`
				SpellListID           interface{} `json:"spellListId"`
			} `json:"class"`
			Background interface{}   `json:"background"`
			Item       []interface{} `json:"item"`
			Feat       []struct {
				OverrideSaveDc        interface{} `json:"overrideSaveDc"`
				LimitedUse            interface{} `json:"limitedUse"`
				ID                    interface{} `json:"id"`
				EntityTypeID          interface{} `json:"entityTypeId"`
				Definition            interface{} `json:"definition"`
				DefinitionID          int         `json:"definitionId"`
				Prepared              bool        `json:"prepared"`
				CountsAsKnownSpell    interface{} `json:"countsAsKnownSpell"`
				UsesSpellSlot         bool        `json:"usesSpellSlot"`
				CastAtLevel           interface{} `json:"castAtLevel"`
				AlwaysPrepared        bool        `json:"alwaysPrepared"`
				Restriction           interface{} `json:"restriction"`
				SpellCastingAbilityID interface{} `json:"spellCastingAbilityId"`
				DisplayAsAttack       interface{} `json:"displayAsAttack"`
				AdditionalDescription interface{} `json:"additionalDescription"`
				CastOnlyAsRitual      bool        `json:"castOnlyAsRitual"`
				RitualCastingType     interface{} `json:"ritualCastingType"`
				Range                 interface{} `json:"range"`
				Activation            interface{} `json:"activation"`
				BaseLevelAtWill       bool        `json:"baseLevelAtWill"`
				AtWillLimitedUseLevel interface{} `json:"atWillLimitedUseLevel"`
				IsSignatureSpell      interface{} `json:"isSignatureSpell"`
				ComponentID           int         `json:"componentId"`
				ComponentTypeID       int         `json:"componentTypeId"`
				SpellListID           interface{} `json:"spellListId"`
			} `json:"feat"`
		} `json:"spells"`
		Options struct {
			Race  []interface{} `json:"race"`
			Class []struct {
				ComponentID     int `json:"componentId"`
				ComponentTypeID int `json:"componentTypeId"`
				Definition      struct {
					ID               int           `json:"id"`
					EntityTypeID     int           `json:"entityTypeId"`
					Name             string        `json:"name"`
					Description      string        `json:"description"`
					Snippet          string        `json:"snippet"`
					Activation       interface{}   `json:"activation"`
					SourceID         int           `json:"sourceId"`
					SourcePageNumber interface{}   `json:"sourcePageNumber"`
					CreatureRules    []interface{} `json:"creatureRules"`
					SpellListIds     []interface{} `json:"spellListIds"`
				} `json:"definition"`
			} `json:"class"`
			Background interface{}   `json:"background"`
			Item       interface{}   `json:"item"`
			Feat       []interface{} `json:"feat"`
		} `json:"options"`
		Choices struct {
			Race  []interface{} `json:"race"`
			Class []struct {
				ComponentID     int           `json:"componentId"`
				ComponentTypeID int           `json:"componentTypeId"`
				ID              string        `json:"id"`
				ParentChoiceID  interface{}   `json:"parentChoiceId"`
				Type            int           `json:"type"`
				SubType         interface{}   `json:"subType"`
				OptionValue     int           `json:"optionValue"`
				Label           string        `json:"label"`
				IsOptional      bool          `json:"isOptional"`
				IsInfinite      bool          `json:"isInfinite"`
				DefaultSubtypes []interface{} `json:"defaultSubtypes"`
				DisplayOrder    int           `json:"displayOrder"`
				Options         []struct {
					ID          int    `json:"id"`
					Label       string `json:"label"`
					Description string `json:"description"`
				} `json:"options"`
				OptionIds []interface{} `json:"optionIds"`
			} `json:"class"`
			Background []struct {
				ComponentID     int         `json:"componentId"`
				ComponentTypeID int         `json:"componentTypeId"`
				ID              string      `json:"id"`
				ParentChoiceID  interface{} `json:"parentChoiceId"`
				Type            int         `json:"type"`
				SubType         int         `json:"subType"`
				OptionValue     int         `json:"optionValue"`
				Label           string      `json:"label"`
				IsOptional      bool        `json:"isOptional"`
				IsInfinite      bool        `json:"isInfinite"`
				DefaultSubtypes []string    `json:"defaultSubtypes"`
				DisplayOrder    interface{} `json:"displayOrder"`
				Options         []struct {
					ID          int         `json:"id"`
					Label       string      `json:"label"`
					Description interface{} `json:"description"`
				} `json:"options"`
				OptionIds []interface{} `json:"optionIds"`
			} `json:"background"`
			Item interface{} `json:"item"`
			Feat []struct {
				ComponentID     int           `json:"componentId"`
				ComponentTypeID int           `json:"componentTypeId"`
				ID              string        `json:"id"`
				ParentChoiceID  interface{}   `json:"parentChoiceId"`
				Type            int           `json:"type"`
				SubType         interface{}   `json:"subType"`
				OptionValue     interface{}   `json:"optionValue"`
				Label           string        `json:"label"`
				IsOptional      bool          `json:"isOptional"`
				IsInfinite      bool          `json:"isInfinite"`
				DefaultSubtypes []interface{} `json:"defaultSubtypes"`
				DisplayOrder    interface{}   `json:"displayOrder"`
				Options         []struct {
					ID          int    `json:"id"`
					Label       string `json:"label"`
					Description string `json:"description"`
				} `json:"options"`
				OptionIds []interface{} `json:"optionIds"`
			} `json:"feat"`
			ChoiceDefinitions []interface{} `json:"choiceDefinitions"`
		} `json:"choices"`
		Actions struct {
			Race  []interface{} `json:"race"`
			Class []struct {
				ComponentID     int    `json:"componentId"`
				ComponentTypeID int    `json:"componentTypeId"`
				ID              string `json:"id"`
				EntityTypeID    string `json:"entityTypeId"`
				LimitedUse      struct {
					Name                     interface{} `json:"name"`
					StatModifierUsesID       interface{} `json:"statModifierUsesId"`
					ResetType                int         `json:"resetType"`
					NumberUsed               int         `json:"numberUsed"`
					MinNumberConsumed        int         `json:"minNumberConsumed"`
					MaxNumberConsumed        int         `json:"maxNumberConsumed"`
					MaxUses                  int         `json:"maxUses"`
					Operator                 int         `json:"operator"`
					UseProficiencyBonus      bool        `json:"useProficiencyBonus"`
					ProficiencyBonusOperator int         `json:"proficiencyBonusOperator"`
					ResetDice                interface{} `json:"resetDice"`
				} `json:"limitedUse"`
				Name                   string      `json:"name"`
				Description            string      `json:"description"`
				Snippet                string      `json:"snippet"`
				AbilityModifierStatID  interface{} `json:"abilityModifierStatId"`
				OnMissDescription      string      `json:"onMissDescription"`
				SaveFailDescription    string      `json:"saveFailDescription"`
				SaveSuccessDescription string      `json:"saveSuccessDescription"`
				SaveStatID             interface{} `json:"saveStatId"`
				FixedSaveDc            interface{} `json:"fixedSaveDc"`
				AttackTypeRange        interface{} `json:"attackTypeRange"`
				ActionType             int         `json:"actionType"`
				AttackSubtype          interface{} `json:"attackSubtype"`
				Dice                   interface{} `json:"dice"`
				Value                  interface{} `json:"value"`
				DamageTypeID           interface{} `json:"damageTypeId"`
				IsMartialArts          bool        `json:"isMartialArts"`
				IsProficient           bool        `json:"isProficient"`
				SpellRangeType         interface{} `json:"spellRangeType"`
				DisplayAsAttack        interface{} `json:"displayAsAttack"`
				Range                  struct {
					Range                    interface{} `json:"range"`
					LongRange                interface{} `json:"longRange"`
					AoeType                  interface{} `json:"aoeType"`
					AoeSize                  interface{} `json:"aoeSize"`
					HasAoeSpecialDescription bool        `json:"hasAoeSpecialDescription"`
					MinimumRange             interface{} `json:"minimumRange"`
				} `json:"range"`
				Activation struct {
					ActivationTime int `json:"activationTime"`
					ActivationType int `json:"activationType"`
				} `json:"activation"`
				NumberOfTargets interface{} `json:"numberOfTargets"`
				FixedToHit      interface{} `json:"fixedToHit"`
				Ammunition      interface{} `json:"ammunition"`
			} `json:"class"`
			Background interface{}   `json:"background"`
			Item       interface{}   `json:"item"`
			Feat       []interface{} `json:"feat"`
		} `json:"actions"`
		Modifiers struct {
			Race []struct {
				FixedValue            int           `json:"fixedValue"`
				ID                    string        `json:"id"`
				EntityID              int           `json:"entityId"`
				EntityTypeID          int           `json:"entityTypeId"`
				Type                  string        `json:"type"`
				SubType               string        `json:"subType"`
				Dice                  interface{}   `json:"dice"`
				Restriction           string        `json:"restriction"`
				StatID                interface{}   `json:"statId"`
				RequiresAttunement    bool          `json:"requiresAttunement"`
				Duration              interface{}   `json:"duration"`
				FriendlyTypeName      string        `json:"friendlyTypeName"`
				FriendlySubtypeName   string        `json:"friendlySubtypeName"`
				IsGranted             bool          `json:"isGranted"`
				BonusTypes            []interface{} `json:"bonusTypes"`
				Value                 int           `json:"value"`
				AvailableToMulticlass bool          `json:"availableToMulticlass"`
				ModifierTypeID        int           `json:"modifierTypeId"`
				ModifierSubTypeID     int           `json:"modifierSubTypeId"`
				ComponentID           int           `json:"componentId"`
				ComponentTypeID       int           `json:"componentTypeId"`
			} `json:"race"`
			Class []struct {
				FixedValue            interface{}   `json:"fixedValue"`
				ID                    string        `json:"id"`
				EntityID              interface{}   `json:"entityId"`
				EntityTypeID          interface{}   `json:"entityTypeId"`
				Type                  string        `json:"type"`
				SubType               string        `json:"subType"`
				Dice                  interface{}   `json:"dice"`
				Restriction           string        `json:"restriction"`
				StatID                interface{}   `json:"statId"`
				RequiresAttunement    bool          `json:"requiresAttunement"`
				Duration              interface{}   `json:"duration"`
				FriendlyTypeName      string        `json:"friendlyTypeName"`
				FriendlySubtypeName   string        `json:"friendlySubtypeName"`
				IsGranted             bool          `json:"isGranted"`
				BonusTypes            []interface{} `json:"bonusTypes"`
				Value                 interface{}   `json:"value"`
				AvailableToMulticlass bool          `json:"availableToMulticlass"`
				ModifierTypeID        int           `json:"modifierTypeId"`
				ModifierSubTypeID     int           `json:"modifierSubTypeId"`
				ComponentID           int           `json:"componentId"`
				ComponentTypeID       int           `json:"componentTypeId"`
			} `json:"class"`
			Background []struct {
				FixedValue            interface{}   `json:"fixedValue"`
				ID                    string        `json:"id"`
				EntityID              int           `json:"entityId"`
				EntityTypeID          int           `json:"entityTypeId"`
				Type                  string        `json:"type"`
				SubType               string        `json:"subType"`
				Dice                  interface{}   `json:"dice"`
				Restriction           string        `json:"restriction"`
				StatID                interface{}   `json:"statId"`
				RequiresAttunement    bool          `json:"requiresAttunement"`
				Duration              interface{}   `json:"duration"`
				FriendlyTypeName      string        `json:"friendlyTypeName"`
				FriendlySubtypeName   string        `json:"friendlySubtypeName"`
				IsGranted             bool          `json:"isGranted"`
				BonusTypes            []interface{} `json:"bonusTypes"`
				Value                 interface{}   `json:"value"`
				AvailableToMulticlass bool          `json:"availableToMulticlass"`
				ModifierTypeID        int           `json:"modifierTypeId"`
				ModifierSubTypeID     int           `json:"modifierSubTypeId"`
				ComponentID           int           `json:"componentId"`
				ComponentTypeID       int           `json:"componentTypeId"`
			} `json:"background"`
			Item []struct {
				FixedValue   int         `json:"fixedValue"`
				ID           string      `json:"id"`
				EntityID     interface{} `json:"entityId"`
				EntityTypeID interface{} `json:"entityTypeId"`
				Type         string      `json:"type"`
				SubType      string      `json:"subType"`
				Dice         struct {
					DiceCount      int         `json:"diceCount"`
					DiceValue      int         `json:"diceValue"`
					DiceMultiplier interface{} `json:"diceMultiplier"`
					FixedValue     int         `json:"fixedValue"`
					DiceString     string      `json:"diceString"`
				} `json:"dice"`
				Restriction           string        `json:"restriction"`
				StatID                interface{}   `json:"statId"`
				RequiresAttunement    bool          `json:"requiresAttunement"`
				Duration              interface{}   `json:"duration"`
				FriendlyTypeName      string        `json:"friendlyTypeName"`
				FriendlySubtypeName   string        `json:"friendlySubtypeName"`
				IsGranted             bool          `json:"isGranted"`
				BonusTypes            []interface{} `json:"bonusTypes"`
				Value                 interface{}   `json:"value"`
				AvailableToMulticlass bool          `json:"availableToMulticlass"`
				ModifierTypeID        int           `json:"modifierTypeId"`
				ModifierSubTypeID     int           `json:"modifierSubTypeId"`
				ComponentID           int           `json:"componentId"`
				ComponentTypeID       int           `json:"componentTypeId"`
			} `json:"item"`
			Feat      []interface{} `json:"feat"`
			Condition []interface{} `json:"condition"`
		} `json:"modifiers"`
		ClassSpells []struct {
			EntityTypeID     int `json:"entityTypeId"`
			CharacterClassID int `json:"characterClassId"`
			Spells           []struct {
				OverrideSaveDc interface{} `json:"overrideSaveDc"`
				LimitedUse     interface{} `json:"limitedUse"`
				ID             int         `json:"id"`
				EntityTypeID   int         `json:"entityTypeId"`
				Definition     struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Level    int    `json:"level"`
					School   string `json:"school"`
					Duration struct {
						DurationInterval int         `json:"durationInterval"`
						DurationUnit     interface{} `json:"durationUnit"`
						DurationType     string      `json:"durationType"`
					} `json:"duration"`
					Activation struct {
						ActivationTime int `json:"activationTime"`
						ActivationType int `json:"activationType"`
					} `json:"activation"`
					Range struct {
						Origin     string      `json:"origin"`
						RangeValue int         `json:"rangeValue"`
						AoeType    interface{} `json:"aoeType"`
						AoeValue   interface{} `json:"aoeValue"`
					} `json:"range"`
					AsPartOfWeaponAttack  bool          `json:"asPartOfWeaponAttack"`
					Description           string        `json:"description"`
					Snippet               string        `json:"snippet"`
					Concentration         bool          `json:"concentration"`
					Ritual                bool          `json:"ritual"`
					RangeArea             interface{}   `json:"rangeArea"`
					DamageEffect          interface{}   `json:"damageEffect"`
					Components            []int         `json:"components"`
					ComponentsDescription string        `json:"componentsDescription"`
					SaveDcAbilityID       int           `json:"saveDcAbilityId"`
					Healing               interface{}   `json:"healing"`
					HealingDice           []interface{} `json:"healingDice"`
					TempHpDice            []interface{} `json:"tempHpDice"`
					AttackType            interface{}   `json:"attackType"`
					CanCastAtHigherLevel  bool          `json:"canCastAtHigherLevel"`
					IsHomebrew            bool          `json:"isHomebrew"`
					Version               interface{}   `json:"version"`
					SourceID              interface{}   `json:"sourceId"`
					SourcePageNumber      int           `json:"sourcePageNumber"`
					RequiresSavingThrow   bool          `json:"requiresSavingThrow"`
					RequiresAttackRoll    bool          `json:"requiresAttackRoll"`
					AtHigherLevels        struct {
						ScaleType              string        `json:"scaleType"`
						HigherLevelDefinitions []interface{} `json:"higherLevelDefinitions"`
						AdditionalAttacks      []interface{} `json:"additionalAttacks"`
						AdditionalTargets      []interface{} `json:"additionalTargets"`
						AreaOfEffect           []interface{} `json:"areaOfEffect"`
						Duration               []interface{} `json:"duration"`
						Creatures              []interface{} `json:"creatures"`
						Special                []interface{} `json:"special"`
						Points                 []interface{} `json:"points"`
					} `json:"atHigherLevels"`
					Modifiers []struct {
						FixedValue            interface{}   `json:"fixedValue"`
						ID                    string        `json:"id"`
						EntityID              interface{}   `json:"entityId"`
						EntityTypeID          interface{}   `json:"entityTypeId"`
						Type                  string        `json:"type"`
						SubType               string        `json:"subType"`
						Dice                  interface{}   `json:"dice"`
						Restriction           string        `json:"restriction"`
						StatID                interface{}   `json:"statId"`
						RequiresAttunement    bool          `json:"requiresAttunement"`
						Duration              interface{}   `json:"duration"`
						FriendlyTypeName      string        `json:"friendlyTypeName"`
						FriendlySubtypeName   string        `json:"friendlySubtypeName"`
						IsGranted             bool          `json:"isGranted"`
						BonusTypes            []interface{} `json:"bonusTypes"`
						Value                 interface{}   `json:"value"`
						AvailableToMulticlass interface{}   `json:"availableToMulticlass"`
						ModifierTypeID        int           `json:"modifierTypeId"`
						ModifierSubTypeID     int           `json:"modifierSubTypeId"`
						ComponentID           int           `json:"componentId"`
						ComponentTypeID       int           `json:"componentTypeId"`
						Die                   struct {
							DiceCount      int         `json:"diceCount"`
							DiceValue      int         `json:"diceValue"`
							DiceMultiplier interface{} `json:"diceMultiplier"`
							FixedValue     interface{} `json:"fixedValue"`
							DiceString     string      `json:"diceString"`
						} `json:"die"`
						Count          int         `json:"count"`
						DurationUnit   interface{} `json:"durationUnit"`
						UsePrimaryStat bool        `json:"usePrimaryStat"`
						AtHigherLevels struct {
							ScaleType              string `json:"scaleType"`
							HigherLevelDefinitions []struct {
								Level  int `json:"level"`
								TypeID int `json:"typeId"`
								Dice   struct {
									DiceCount      int         `json:"diceCount"`
									DiceValue      int         `json:"diceValue"`
									DiceMultiplier interface{} `json:"diceMultiplier"`
									FixedValue     int         `json:"fixedValue"`
									DiceString     string      `json:"diceString"`
								} `json:"dice"`
								Value   interface{} `json:"value"`
								Details string      `json:"details"`
							} `json:"higherLevelDefinitions"`
							AdditionalAttacks []interface{} `json:"additionalAttacks"`
							AdditionalTargets []interface{} `json:"additionalTargets"`
							AreaOfEffect      []interface{} `json:"areaOfEffect"`
							Duration          []interface{} `json:"duration"`
							Creatures         []interface{} `json:"creatures"`
							Special           []interface{} `json:"special"`
							Points            []interface{} `json:"points"`
						} `json:"atHigherLevels"`
					} `json:"modifiers"`
					Conditions             []interface{} `json:"conditions"`
					Tags                   []string      `json:"tags"`
					CastingTimeDescription string        `json:"castingTimeDescription"`
					ScaleType              string        `json:"scaleType"`
					Sources                []struct {
						SourceID   int `json:"sourceId"`
						PageNumber int `json:"pageNumber"`
						SourceType int `json:"sourceType"`
					} `json:"sources"`
					SpellGroups []interface{} `json:"spellGroups"`
				} `json:"definition"`
				DefinitionID          int         `json:"definitionId"`
				Prepared              bool        `json:"prepared"`
				CountsAsKnownSpell    bool        `json:"countsAsKnownSpell"`
				UsesSpellSlot         bool        `json:"usesSpellSlot"`
				CastAtLevel           interface{} `json:"castAtLevel"`
				AlwaysPrepared        bool        `json:"alwaysPrepared"`
				Restriction           interface{} `json:"restriction"`
				SpellCastingAbilityID interface{} `json:"spellCastingAbilityId"`
				DisplayAsAttack       interface{} `json:"displayAsAttack"`
				AdditionalDescription interface{} `json:"additionalDescription"`
				CastOnlyAsRitual      bool        `json:"castOnlyAsRitual"`
				RitualCastingType     interface{} `json:"ritualCastingType"`
				Range                 struct {
					Origin     string      `json:"origin"`
					RangeValue int         `json:"rangeValue"`
					AoeType    interface{} `json:"aoeType"`
					AoeValue   interface{} `json:"aoeValue"`
				} `json:"range"`
				Activation struct {
					ActivationTime int `json:"activationTime"`
					ActivationType int `json:"activationType"`
				} `json:"activation"`
				BaseLevelAtWill       bool        `json:"baseLevelAtWill"`
				AtWillLimitedUseLevel interface{} `json:"atWillLimitedUseLevel"`
				IsSignatureSpell      interface{} `json:"isSignatureSpell"`
				ComponentID           int         `json:"componentId"`
				ComponentTypeID       int         `json:"componentTypeId"`
				SpellListID           interface{} `json:"spellListId"`
			} `json:"spells"`
		} `json:"classSpells"`
		CustomItems []interface{} `json:"customItems"`
		Campaign    struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Link        string `json:"link"`
			PublicNotes string `json:"publicNotes"`
			DmUserID    int    `json:"dmUserId"`
			DmUsername  string `json:"dmUsername"`
			Characters  []struct {
				UserID        int         `json:"userId"`
				Username      string      `json:"username"`
				CharacterID   int         `json:"characterId"`
				CharacterName string      `json:"characterName"`
				CharacterURL  string      `json:"characterUrl"`
				AvatarURL     string      `json:"avatarUrl"`
				PrivacyType   int         `json:"privacyType"`
				CampaignID    interface{} `json:"campaignId"`
				IsAssigned    bool        `json:"isAssigned"`
			} `json:"characters"`
		} `json:"campaign"`
		Creatures []struct {
			ID                 int         `json:"id"`
			EntityTypeID       int         `json:"entityTypeId"`
			Name               interface{} `json:"name"`
			Description        interface{} `json:"description"`
			IsActive           bool        `json:"isActive"`
			RemovedHitPoints   int         `json:"removedHitPoints"`
			TemporaryHitPoints interface{} `json:"temporaryHitPoints"`
			GroupID            int         `json:"groupId"`
			Definition         struct {
				ID                    int    `json:"id"`
				EntityTypeID          int    `json:"entityTypeId"`
				Name                  string `json:"name"`
				AlignmentID           int    `json:"alignmentId"`
				SizeID                int    `json:"sizeId"`
				TypeID                int    `json:"typeId"`
				ArmorClass            int    `json:"armorClass"`
				ArmorClassDescription string `json:"armorClassDescription"`
				AverageHitPoints      int    `json:"averageHitPoints"`
				HitPointDice          struct {
					DiceCount      int         `json:"diceCount"`
					DiceValue      int         `json:"diceValue"`
					DiceMultiplier interface{} `json:"diceMultiplier"`
					FixedValue     int         `json:"fixedValue"`
					DiceString     string      `json:"diceString"`
				} `json:"hitPointDice"`
				Movements []struct {
					MovementID int    `json:"movementId"`
					Speed      int    `json:"speed"`
					Notes      string `json:"notes"`
				} `json:"movements"`
				PassivePerception int           `json:"passivePerception"`
				IsHomebrew        bool          `json:"isHomebrew"`
				ChallengeRatingID int           `json:"challengeRatingId"`
				SourceID          int           `json:"sourceId"`
				SourcePageNumber  int           `json:"sourcePageNumber"`
				IsLegendary       bool          `json:"isLegendary"`
				IsMythic          bool          `json:"isMythic"`
				HasLair           bool          `json:"hasLair"`
				AvatarURL         interface{}   `json:"avatarUrl"`
				LargeAvatarURL    interface{}   `json:"largeAvatarUrl"`
				BasicAvatarURL    interface{}   `json:"basicAvatarUrl"`
				Version           interface{}   `json:"version"`
				Swarm             interface{}   `json:"swarm"`
				SubTypes          []interface{} `json:"subTypes"`
				Environments      []int         `json:"environments"`
				Tags              []string      `json:"tags"`
				Sources           []struct {
					SourceID   int         `json:"sourceId"`
					PageNumber interface{} `json:"pageNumber"`
					SourceType int         `json:"sourceType"`
				} `json:"sources"`
				Stats []struct {
					StatID int         `json:"statId"`
					Name   interface{} `json:"name"`
					Value  int         `json:"value"`
				} `json:"stats"`
				Senses              []interface{} `json:"senses"`
				DamageAdjustments   []interface{} `json:"damageAdjustments"`
				ConditionImmunities []interface{} `json:"conditionImmunities"`
				SavingThrows        []interface{} `json:"savingThrows"`
				Skills              []struct {
					SkillID         int         `json:"skillId"`
					Value           int         `json:"value"`
					AdditionalBonus interface{} `json:"additionalBonus"`
				} `json:"skills"`
				Languages                   []interface{} `json:"languages"`
				SpecialTraitsDescription    string        `json:"specialTraitsDescription"`
				ActionsDescription          string        `json:"actionsDescription"`
				ReactionsDescription        string        `json:"reactionsDescription"`
				LegendaryActionsDescription string        `json:"legendaryActionsDescription"`
				MythicActionsDescription    interface{}   `json:"mythicActionsDescription"`
				BonusActionsDescription     interface{}   `json:"bonusActionsDescription"`
				CharacteristicsDescription  string        `json:"characteristicsDescription"`
				LairDescription             string        `json:"lairDescription"`
				LanguageDescription         string        `json:"languageDescription"`
				LanguageNote                string        `json:"languageNote"`
				HideCr                      bool          `json:"hideCr"`
				IsLegacy                    bool          `json:"isLegacy"`
			} `json:"definition"`
		} `json:"creatures"`
		OptionalOrigins       []interface{} `json:"optionalOrigins"`
		OptionalClassFeatures []interface{} `json:"optionalClassFeatures"`
		DateModified          time.Time     `json:"dateModified"`
		ProvidedFrom          string        `json:"providedFrom"`
	} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
