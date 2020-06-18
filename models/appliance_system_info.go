// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplianceSystemInfo Appliance:System Info
//
// The Intersight Appliance's system information. SystemInfo is a singleton managed object
// created during the Intersight Appliance setup. The Intersight Appliance updates the
// SystemInfo managed object with up to date cluster status information periodically.
//
// swagger:model applianceSystemInfo
type ApplianceSystemInfo struct {
	MoBaseMo

	// Connection state of the Intersight Appliance to the Intersight.
	// Read Only: true
	// Enum: [ Connected NotConnected ClaimInProgress Unclaimed]
	CloudConnStatus string `json:"CloudConnStatus,omitempty"`

	// Current running deployment size for the Intersight Appliance cluster. Eg. small, medium, large etc.
	// Read Only: true
	DeploymentSize string `json:"DeploymentSize,omitempty"`

	// Publicly accessible FQDN or IP address of the Intersight Appliance.
	// Read Only: true
	Hostname string `json:"Hostname,omitempty"`

	// Indicates that the setup initialization process has been completed.
	// Read Only: true
	InitDone *bool `json:"InitDone,omitempty"`

	// Operational status of the Intersight Appliance cluster.
	// Read Only: true
	// Enum: [Unknown Operational Impaired AttentionNeeded]
	OperationalStatus string `json:"OperationalStatus,omitempty"`

	// SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance.
	// Read Only: true
	SerialID string `json:"SerialId,omitempty"`

	// Timezone of the Intersight Appliance.
	// Enum: [Pacific/Niue Pacific/Pago_Pago Pacific/Honolulu Pacific/Rarotonga Pacific/Tahiti Pacific/Marquesas America/Anchorage Pacific/Gambier America/Los_Angeles America/Tijuana America/Vancouver America/Whitehorse Pacific/Pitcairn America/Dawson_Creek America/Denver America/Edmonton America/Hermosillo America/Mazatlan America/Phoenix America/Yellowknife America/Belize America/Chicago America/Costa_Rica America/El_Salvador America/Guatemala America/Managua America/Mexico_City America/Regina America/Tegucigalpa America/Winnipeg Pacific/Galapagos America/Bogota America/Cancun America/Cayman America/Guayaquil America/Havana America/Iqaluit America/Jamaica America/Lima America/Nassau America/New_York America/Panama America/Port-au-Prince America/Rio_Branco America/Toronto Pacific/Easter America/Caracas America/Asuncion America/Barbados America/Boa_Vista America/Campo_Grande America/Cuiaba America/Curacao America/Grand_Turk America/Guyana America/Halifax America/La_Paz America/Manaus America/Martinique America/Port_of_Spain America/Porto_Velho America/Puerto_Rico America/Santo_Domingo America/Thule Atlantic/Bermuda America/St_Johns America/Araguaina America/Argentina/Buenos_Aires America/Bahia America/Belem America/Cayenne America/Fortaleza America/Godthab America/Maceio America/Miquelon America/Montevideo America/Paramaribo America/Recife America/Santiago America/Sao_Paulo Antarctica/Palmer Antarctica/Rothera Atlantic/Stanley America/Noronha Atlantic/South_Georgia America/Scoresbysund Atlantic/Azores Atlantic/Cape_Verde Africa/Abidjan Africa/Accra Africa/Bissau Africa/Casablanca Africa/El_Aaiun Africa/Monrovia America/Danmarkshavn Atlantic/Canary Atlantic/Faroe Atlantic/Reykjavik Etc/GMT Europe/Dublin Europe/Lisbon Europe/London Africa/Algiers Africa/Ceuta Africa/Lagos Africa/Ndjamena Africa/Tunis Africa/Windhoek Europe/Amsterdam Europe/Andorra Europe/Belgrade Europe/Berlin Europe/Brussels Europe/Budapest Europe/Copenhagen Europe/Gibraltar Europe/Luxembourg Europe/Madrid Europe/Malta Europe/Monaco Europe/Oslo Europe/Paris Europe/Prague Europe/Rome Europe/Stockholm Europe/Tirane Europe/Vienna Europe/Warsaw Europe/Zurich Africa/Cairo Africa/Johannesburg Africa/Maputo Africa/Tripoli Asia/Amman Asia/Beirut Asia/Damascus Asia/Gaza Asia/Jerusalem Asia/Nicosia Europe/Athens Europe/Bucharest Europe/Chisinau Europe/Helsinki Europe/Istanbul Europe/Kaliningrad Europe/Kiev Europe/Riga Europe/Sofia Europe/Tallinn Europe/Vilnius Africa/Khartoum Africa/Nairobi Antarctica/Syowa Asia/Baghdad Asia/Qatar Asia/Riyadh Europe/Minsk Europe/Moscow Asia/Tehran Asia/Baku Asia/Dubai Asia/Tbilisi Asia/Yerevan Europe/Samara Indian/Mahe Indian/Mauritius Indian/Reunion Asia/Kabul Antarctica/Mawson Asia/Aqtau Asia/Aqtobe Asia/Ashgabat Asia/Dushanbe Asia/Karachi Asia/Tashkent Asia/Yekaterinburg Indian/Kerguelen Indian/Maldives Asia/Calcutta Asia/Kolkata Asia/Colombo Asia/Katmandu Antarctica/Vostok Asia/Almaty Asia/Bishkek Asia/Dhaka Asia/Omsk Asia/Thimphu Indian/Chagos Asia/Rangoon Indian/Cocos Antarctica/Davis Asia/Bangkok Asia/Hovd Asia/Jakarta Asia/Krasnoyarsk Asia/Saigon Indian/Christmas Antarctica/Casey Asia/Brunei Asia/Choibalsan Asia/Hong_Kong Asia/Irkutsk Asia/Kuala_Lumpur Asia/Macau Asia/Makassar Asia/Manila Asia/Shanghai Asia/Singapore Asia/Taipei Asia/Ulaanbaatar Australia/Perth Asia/Pyongyang Asia/Dili Asia/Jayapura Asia/Seoul Asia/Tokyo Asia/Yakutsk Pacific/Palau Australia/Adelaide Australia/Darwin Antarctica/DumontDUrville Asia/Magadan Asia/Vladivostok Australia/Brisbane Australia/Hobart Australia/Sydney Pacific/Chuuk Pacific/Guam Pacific/Port_Moresby Pacific/Efate Pacific/Guadalcanal Pacific/Kosrae Pacific/Norfolk Pacific/Noumea Pacific/Pohnpei Asia/Kamchatka Pacific/Auckland Pacific/Fiji Pacific/Funafuti Pacific/Kwajalein Pacific/Majuro Pacific/Nauru Pacific/Tarawa Pacific/Wake Pacific/Wallis Pacific/Apia Pacific/Enderbury Pacific/Fakaofo Pacific/Tongatapu Pacific/Kiritimati]
	TimeZone *string `json:"TimeZone,omitempty"`

	// Current software version of the Intersight Appliance.
	// Read Only: true
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplianceSystemInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		CloudConnStatus string `json:"CloudConnStatus,omitempty"`

		DeploymentSize string `json:"DeploymentSize,omitempty"`

		Hostname string `json:"Hostname,omitempty"`

		InitDone *bool `json:"InitDone,omitempty"`

		OperationalStatus string `json:"OperationalStatus,omitempty"`

		SerialID string `json:"SerialId,omitempty"`

		TimeZone *string `json:"TimeZone,omitempty"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CloudConnStatus = dataAO1.CloudConnStatus

	m.DeploymentSize = dataAO1.DeploymentSize

	m.Hostname = dataAO1.Hostname

	m.InitDone = dataAO1.InitDone

	m.OperationalStatus = dataAO1.OperationalStatus

	m.SerialID = dataAO1.SerialID

	m.TimeZone = dataAO1.TimeZone

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplianceSystemInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CloudConnStatus string `json:"CloudConnStatus,omitempty"`

		DeploymentSize string `json:"DeploymentSize,omitempty"`

		Hostname string `json:"Hostname,omitempty"`

		InitDone *bool `json:"InitDone,omitempty"`

		OperationalStatus string `json:"OperationalStatus,omitempty"`

		SerialID string `json:"SerialId,omitempty"`

		TimeZone *string `json:"TimeZone,omitempty"`

		Version string `json:"Version,omitempty"`
	}

	dataAO1.CloudConnStatus = m.CloudConnStatus

	dataAO1.DeploymentSize = m.DeploymentSize

	dataAO1.Hostname = m.Hostname

	dataAO1.InitDone = m.InitDone

	dataAO1.OperationalStatus = m.OperationalStatus

	dataAO1.SerialID = m.SerialID

	dataAO1.TimeZone = m.TimeZone

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this appliance system info
func (m *ApplianceSystemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudConnStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationalStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applianceSystemInfoTypeCloudConnStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","Connected","NotConnected","ClaimInProgress","Unclaimed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applianceSystemInfoTypeCloudConnStatusPropEnum = append(applianceSystemInfoTypeCloudConnStatusPropEnum, v)
	}
}

// property enum
func (m *ApplianceSystemInfo) validateCloudConnStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applianceSystemInfoTypeCloudConnStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplianceSystemInfo) validateCloudConnStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudConnStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudConnStatusEnum("CloudConnStatus", "body", m.CloudConnStatus); err != nil {
		return err
	}

	return nil
}

var applianceSystemInfoTypeOperationalStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Operational","Impaired","AttentionNeeded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applianceSystemInfoTypeOperationalStatusPropEnum = append(applianceSystemInfoTypeOperationalStatusPropEnum, v)
	}
}

// property enum
func (m *ApplianceSystemInfo) validateOperationalStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applianceSystemInfoTypeOperationalStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplianceSystemInfo) validateOperationalStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.OperationalStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperationalStatusEnum("OperationalStatus", "body", m.OperationalStatus); err != nil {
		return err
	}

	return nil
}

var applianceSystemInfoTypeTimeZonePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pacific/Niue","Pacific/Pago_Pago","Pacific/Honolulu","Pacific/Rarotonga","Pacific/Tahiti","Pacific/Marquesas","America/Anchorage","Pacific/Gambier","America/Los_Angeles","America/Tijuana","America/Vancouver","America/Whitehorse","Pacific/Pitcairn","America/Dawson_Creek","America/Denver","America/Edmonton","America/Hermosillo","America/Mazatlan","America/Phoenix","America/Yellowknife","America/Belize","America/Chicago","America/Costa_Rica","America/El_Salvador","America/Guatemala","America/Managua","America/Mexico_City","America/Regina","America/Tegucigalpa","America/Winnipeg","Pacific/Galapagos","America/Bogota","America/Cancun","America/Cayman","America/Guayaquil","America/Havana","America/Iqaluit","America/Jamaica","America/Lima","America/Nassau","America/New_York","America/Panama","America/Port-au-Prince","America/Rio_Branco","America/Toronto","Pacific/Easter","America/Caracas","America/Asuncion","America/Barbados","America/Boa_Vista","America/Campo_Grande","America/Cuiaba","America/Curacao","America/Grand_Turk","America/Guyana","America/Halifax","America/La_Paz","America/Manaus","America/Martinique","America/Port_of_Spain","America/Porto_Velho","America/Puerto_Rico","America/Santo_Domingo","America/Thule","Atlantic/Bermuda","America/St_Johns","America/Araguaina","America/Argentina/Buenos_Aires","America/Bahia","America/Belem","America/Cayenne","America/Fortaleza","America/Godthab","America/Maceio","America/Miquelon","America/Montevideo","America/Paramaribo","America/Recife","America/Santiago","America/Sao_Paulo","Antarctica/Palmer","Antarctica/Rothera","Atlantic/Stanley","America/Noronha","Atlantic/South_Georgia","America/Scoresbysund","Atlantic/Azores","Atlantic/Cape_Verde","Africa/Abidjan","Africa/Accra","Africa/Bissau","Africa/Casablanca","Africa/El_Aaiun","Africa/Monrovia","America/Danmarkshavn","Atlantic/Canary","Atlantic/Faroe","Atlantic/Reykjavik","Etc/GMT","Europe/Dublin","Europe/Lisbon","Europe/London","Africa/Algiers","Africa/Ceuta","Africa/Lagos","Africa/Ndjamena","Africa/Tunis","Africa/Windhoek","Europe/Amsterdam","Europe/Andorra","Europe/Belgrade","Europe/Berlin","Europe/Brussels","Europe/Budapest","Europe/Copenhagen","Europe/Gibraltar","Europe/Luxembourg","Europe/Madrid","Europe/Malta","Europe/Monaco","Europe/Oslo","Europe/Paris","Europe/Prague","Europe/Rome","Europe/Stockholm","Europe/Tirane","Europe/Vienna","Europe/Warsaw","Europe/Zurich","Africa/Cairo","Africa/Johannesburg","Africa/Maputo","Africa/Tripoli","Asia/Amman","Asia/Beirut","Asia/Damascus","Asia/Gaza","Asia/Jerusalem","Asia/Nicosia","Europe/Athens","Europe/Bucharest","Europe/Chisinau","Europe/Helsinki","Europe/Istanbul","Europe/Kaliningrad","Europe/Kiev","Europe/Riga","Europe/Sofia","Europe/Tallinn","Europe/Vilnius","Africa/Khartoum","Africa/Nairobi","Antarctica/Syowa","Asia/Baghdad","Asia/Qatar","Asia/Riyadh","Europe/Minsk","Europe/Moscow","Asia/Tehran","Asia/Baku","Asia/Dubai","Asia/Tbilisi","Asia/Yerevan","Europe/Samara","Indian/Mahe","Indian/Mauritius","Indian/Reunion","Asia/Kabul","Antarctica/Mawson","Asia/Aqtau","Asia/Aqtobe","Asia/Ashgabat","Asia/Dushanbe","Asia/Karachi","Asia/Tashkent","Asia/Yekaterinburg","Indian/Kerguelen","Indian/Maldives","Asia/Calcutta","Asia/Kolkata","Asia/Colombo","Asia/Katmandu","Antarctica/Vostok","Asia/Almaty","Asia/Bishkek","Asia/Dhaka","Asia/Omsk","Asia/Thimphu","Indian/Chagos","Asia/Rangoon","Indian/Cocos","Antarctica/Davis","Asia/Bangkok","Asia/Hovd","Asia/Jakarta","Asia/Krasnoyarsk","Asia/Saigon","Indian/Christmas","Antarctica/Casey","Asia/Brunei","Asia/Choibalsan","Asia/Hong_Kong","Asia/Irkutsk","Asia/Kuala_Lumpur","Asia/Macau","Asia/Makassar","Asia/Manila","Asia/Shanghai","Asia/Singapore","Asia/Taipei","Asia/Ulaanbaatar","Australia/Perth","Asia/Pyongyang","Asia/Dili","Asia/Jayapura","Asia/Seoul","Asia/Tokyo","Asia/Yakutsk","Pacific/Palau","Australia/Adelaide","Australia/Darwin","Antarctica/DumontDUrville","Asia/Magadan","Asia/Vladivostok","Australia/Brisbane","Australia/Hobart","Australia/Sydney","Pacific/Chuuk","Pacific/Guam","Pacific/Port_Moresby","Pacific/Efate","Pacific/Guadalcanal","Pacific/Kosrae","Pacific/Norfolk","Pacific/Noumea","Pacific/Pohnpei","Asia/Kamchatka","Pacific/Auckland","Pacific/Fiji","Pacific/Funafuti","Pacific/Kwajalein","Pacific/Majuro","Pacific/Nauru","Pacific/Tarawa","Pacific/Wake","Pacific/Wallis","Pacific/Apia","Pacific/Enderbury","Pacific/Fakaofo","Pacific/Tongatapu","Pacific/Kiritimati"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applianceSystemInfoTypeTimeZonePropEnum = append(applianceSystemInfoTypeTimeZonePropEnum, v)
	}
}

// property enum
func (m *ApplianceSystemInfo) validateTimeZoneEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applianceSystemInfoTypeTimeZonePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplianceSystemInfo) validateTimeZone(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeZone) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeZoneEnum("TimeZone", "body", *m.TimeZone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceSystemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceSystemInfo) UnmarshalBinary(b []byte) error {
	var res ApplianceSystemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
