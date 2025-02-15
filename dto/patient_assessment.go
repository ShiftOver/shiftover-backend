package dto

import "time"

type PatientAssessmentEntity struct {
	PatientID              string           `bson:"patientId" json:"patientId"`
	Spiritual              Spiritual        `bson:"spiritual" json:"spiritual"`
	Nutrition              Nutrition        `bson:"nutrition" json:"nutrition"`
	Pulmonary              Pulmonary        `bson:"pulmonary" json:"pulmonary"`
	Cardiovascular         Cardiovascular   `bson:"cardiovascular" json:"cardiovascular"`
	CardioCurrentTreatment string           `bson:"cardioCurrentTreatment" json:"cardioCurrentTreatment"`
	Neurosensory           Neurosensory     `bson:"neurosensory" json:"neurosensory"`
	Musculoskeletal        Musculoskeletal  `bson:"musculoskeletal" json:"musculoskeletal"`
	Mobility               Mobility         `bson:"mobility" json:"mobility"`
	Teaching               []string         `bson:"teaching" json:"teaching"`
	Gastrointestinal       Gastrointestinal `bson:"gastrointestinal" json:"gastrointestinal"`
	Genitourinary          Genitourinary    `bson:"genitourinary" json:"genitourinary"`
	PainManagement         PainManagement   `bson:"painManagement" json:"painManagement"`
	Discharge              Discharge        `bson:"discharge" json:"discharge"`
	CreatedAt              time.Time        `bson:"createdAt" json:"createdAt"`
	UpdatedAt              time.Time        `bson:"updatedAt" json:"updatedAt"`
}

type Spiritual struct {
	Religion             string `bson:"religion" json:"religion"`
	SpecialConsideration string `bson:"specialConsideration" json:"specialConsideration"`
	Anxiety              string `bson:"anxiety" json:"anxiety"`
	SupportSystem        string `bson:"supportSystem" json:"supportSystem"`
}

type Nutrition struct {
	DietType     string             `bson:"dietType" json:"dietType"`
	Appetite     string             `bson:"appetite" json:"appetite"`
	SpecialDiet  string             `bson:"specialDiet" json:"specialDiet"`
	Feeding      string             `bson:"feeding" json:"feeding"`
	Swallowing   string             `bson:"swallowing" json:"swallowing"`
	GlProblem    string             `bson:"glProblem" json:"glProblem"`
	WeightChange WeightChangeStatus `bson:"weightChange" json:"weightChange"`
}

type WeightChangeStatus struct {
	Status string       `bson:"status" json:"status"`
	Change WeightChange `bson:"change" json:"change"`
}

type WeightChange struct {
	Type        string `bson:"type" json:"type"`
	Weight      string `bson:"weight" json:"weight"`
	PeriodValue int    `bson:"periodValue" json:"periodValue"`
	PeriodType  string `bson:"periodType" json:"periodType"`
}

type Pulmonary struct {
	Rate   string `bson:"rate" json:"rate"`
	Rhythm string `bson:"rhythm" json:"rhythm"`
	Effort string `bson:"effort" json:"effort"`
	Cough  string `bson:"cough" json:"cough"`
	Sputum string `bson:"sputum" json:"sputum"`
}

type Cardiovascular struct {
	PulseRhythm      string             `bson:"pulseRhythm" json:"pulseRhythm"`
	PulseAmplitude   string             `bson:"pulseAmplitude" json:"pulseAmplitude"`
	PulseRate        string             `bson:"pulseRate" json:"pulseRate"`
	Edema            string             `bson:"edema" json:"edema"`
	NeckVeinEngorged string             `bson:"neckVeinEngorged" json:"neckVeinEngorged"`
	ChestPain        CardiovascularPain `bson:"chestPain" json:"chestPain"`
}

type CardiovascularPain struct {
	Yes          bool   `bson:"yes" json:"yes"`
	Location     string `bson:"location" json:"location"`
	ReferredPain string `bson:"referredPain" json:"referredPain"`
	Duration     string `bson:"duration" json:"duration"`
	Frequency    string `bson:"frequency" json:"frequency"`
}

type Neurosensory struct {
	LevelOfConsciousness string  `bson:"levelOfConsciousness" json:"levelOfConsciousness"`
	Vision               Vision  `bson:"vision" json:"vision"`
	Hearing              Hearing `bson:"hearing" json:"hearing"`
	Speech               Speech  `bson:"speech" json:"speech"`
	Smell                Smell   `bson:"smell" json:"smell"`
	Sensation            string  `bson:"sensation" json:"sensation"`
}

type Vision struct {
	Lt     bool   `bson:"lt" json:"lt"`
	Rt     bool   `bson:"rt" json:"rt"`
	Device string `bson:"device" json:"device"`
}

type Hearing struct {
	Lt     bool   `bson:"lt" json:"lt"`
	Rt     bool   `bson:"rt" json:"rt"`
	Device string `bson:"device" json:"device"`
}

type Speech struct {
	Normal bool   `bson:"normal" json:"normal"`
	Device string `bson:"device" json:"device"`
}

type Smell struct {
	Normal      bool   `bson:"normal" json:"normal"`
	Description string `bson:"description" json:"description"`
}

type Musculoskeletal struct {
	HandGrasps HandGrasps `bson:"handGrasps" json:"handGrasps"`
	Joint      Joint      `bson:"joint" json:"joint"`
	Weakness   StatusInfo `bson:"weakness" json:"weakness"`
	Paralysis  StatusInfo `bson:"paralysis" json:"paralysis"`
	Seizure    StatusInfo `bson:"seizure" json:"seizure"`
	Movement   StatusInfo `bson:"movement" json:"movement"`
}

type HandGrasps struct {
	Status string `bson:"status" json:"status"`
	Rt     string `bson:"rt" json:"rt"`
	Lt     string `bson:"lt" json:"lt"`
}

type Joint struct {
	Status      string `bson:"status" json:"status"`
	Description string `bson:"description" json:"description"`
}

type StatusInfo struct {
	Yes         bool   `bson:"yes" json:"yes"`
	Description string `bson:"description" json:"description"`
}

type Mobility struct {
	Activity          MobilityActivity  `bson:"activity" json:"activity"`
	ProstheticDevices ProstheticDevices `bson:"prostheticDevices" json:"prostheticDevices"`
	Injury            StatusInfo        `bson:"injury" json:"injury"`
	DominantHand      string            `bson:"dominantHand" json:"dominantHand"`
}

type MobilityActivity struct {
	Independent bool     `bson:"independent" json:"independent"`
	Assistance  []string `bson:"assistance" json:"assistance"`
}

type ProstheticDevices struct {
	Cane           bool   `bson:"cane" json:"cane"`
	Walker         bool   `bson:"walker" json:"walker"`
	WheelChair     bool   `bson:"wheelChair" json:"wheelChair"`
	ArtificialLimb string `bson:"artificialLimb" json:"artificialLimb"`
	Other          string `bson:"other" json:"other"`
}

type Gastrointestinal struct {
	OralCavity         string       `bson:"oralCavity" json:"oralCavity"`
	Abdomen            string       `bson:"abdomen" json:"abdomen"`
	BowelPattern       BowelPattern `bson:"bowelPattern" json:"bowelPattern"`
	EliminationProblem string       `bson:"eliminationProblem" json:"eliminationProblem"`
}

type BowelPattern struct {
	Frequency int `bson:"frequency" json:"frequency"`
	Days      int `bson:"days" json:"days"`
}

type Genitourinary struct {
	Bladder string  `bson:"bladder" json:"bladder"`
	Voiding Voiding `bson:"voiding" json:"voiding"`
	Urine   string  `bson:"urine" json:"urine"`
}

type Voiding struct {
	Day    int    `bson:"day" json:"day"`
	Night  int    `bson:"night" json:"night"`
	Status string `bson:"status" json:"status"`
}

type PainManagement struct {
	Pain            PainInfo `bson:"pain" json:"pain"`
	Pattern         string   `bson:"pattern" json:"pattern"`
	PatientDescribe string   `bson:"patientDescribe" json:"patientDescribe"`
	Intensity       int      `bson:"intensity" json:"intensity"`
	AffectedAbility string   `bson:"affectedAbility" json:"affectedAbility"`
	Relieves        string   `bson:"relieves" json:"relieves"`
}

type PainInfo struct {
	Yes         bool   `bson:"yes" json:"yes"`
	Description string `bson:"description" json:"description"`
	When        string `bson:"when" json:"when"`
	Cause       string `bson:"cause" json:"cause"`
}

type Discharge struct {
	ScreeningCriteria ScreeningCriteria `bson:"screeningCriteria" json:"screeningCriteria"`
	HomeEnvironment   HomeEnvironment   `bson:"homeEnvironment" json:"homeEnvironment"`
	PlanningNeeds     PlanningNeeds     `bson:"planningNeeds" json:"planningNeeds"`
	ReferralNeeds     ReferralNeeds     `bson:"referralNeeds" json:"referralNeeds"`
}

type ScreeningCriteria struct {
	PostDischargeAssistance PostDischargeAssistance `bson:"postDischargeAssistance" json:"postDischargeAssistance"`
	FinancialConcern        StatusInfo              `bson:"financialConcern" json:"financialConcern"`
}

type PostDischargeAssistance struct {
	Yes                 bool       `bson:"yes" json:"yes"`
	FamilyCapable       StatusInfo `bson:"familyCapable" json:"familyCapable"`
	FamilyCannotProvide StatusInfo `bson:"familyCannotProvide" json:"familyCannotProvide"`
}

type HomeEnvironment struct {
	LiveWith  string `bson:"liveWith" json:"liveWith"`
	LiveWhere string `bson:"liveWhere" json:"liveWhere"`
}

type PlanningNeeds struct {
	Medication         StatusInfo `bson:"medication" json:"medication"`
	Environment        StatusInfo `bson:"environment" json:"environment"`
	Treatment          StatusInfo `bson:"treatment" json:"treatment"`
	Health             StatusInfo `bson:"health" json:"health"`
	OutpatientReferral StatusInfo `bson:"outpatientReferral" json:"outpatientReferral"`
	Diet               StatusInfo `bson:"diet" json:"diet"`
}

type ReferralNeeds struct {
	PossibleNeeds string `bson:"possibleNeeds" json:"possibleNeeds"`
	RnAssessment  string `bson:"rnAssessment" json:"rnAssessment"`
	Date          string `bson:"date" json:"date"`
	Time          string `bson:"time" json:"time"`
}
