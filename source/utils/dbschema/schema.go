package dbschema

import "time"

//Patient struct for the sql table schema for Patient object
type Patient struct {
	ID                           int64     `gorm:"primary_key" json:"id"`
	CreatedAt                    time.Time `json:"createdAt"`
	UpdatedAt                    time.Time `json:"updatedAt"`
	Name                         string
	Dob                          string
	AdmissionDate                string
	DatofMeasureMent             string
	TimeofMeasurement            string
	GlucoseLevel                 string
	PulseRate                    string
	BodyTemperature              string
	SystolicPressure             string
	DiastolicPressure            string
	RespirationRate              string
	IntracranialPressure         string
	OxygenSaturation             string
	CentralVenousPressure        string
	CardiacOutput                string
	MixedVenousOxygenSaturation  string
	RightAtriumPressure          string
	RightVentriclePeakSystolic   string
	RightVentricleEndDiastolic   string
	PulmonaryArteryMean          string
	PulmonaryArteryPeakSystolic  string
	PulmonaryArteryEndDiastolic  string
	PulmonaryArteryOcclusionMean string
	LeftAtriumMean               string
	LeftAtriumAWave              string
	LeftAtriumVWave              string
	LeftVentriclePeakSystolic    string
	LeftVentricleEndDiastolic    string
	BrachialArteryMean           string
	BrachialArteryPeakSystolic   string
	BrachialArteryEndDiastolic   string
}
