// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirClaimResponse struct {
	models.ResourceBase
	// The creation date
	// https://hl7.org/fhir/r4/search.html#date
	Created time.Time `gorm:"column:created;type:datetime" json:"created,omitempty"`
	// The contents of the disposition message
	// https://hl7.org/fhir/r4/search.html#string
	Disposition string `gorm:"column:disposition;type:text" json:"disposition,omitempty"`
	// The identity of the ClaimResponse
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// The organization which generated this resource
	// https://hl7.org/fhir/r4/search.html#reference
	Insurer datatypes.JSON `gorm:"column:insurer;type:text;serializer:json" json:"insurer,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// The processing outcome
	// https://hl7.org/fhir/r4/search.html#token
	Outcome datatypes.JSON `gorm:"column:outcome;type:text;serializer:json" json:"outcome,omitempty"`
	// The expected payment date
	// https://hl7.org/fhir/r4/search.html#date
	PaymentDate time.Time `gorm:"column:paymentDate;type:datetime" json:"paymentDate,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// The claim reference
	// https://hl7.org/fhir/r4/search.html#reference
	Request datatypes.JSON `gorm:"column:request;type:text;serializer:json" json:"request,omitempty"`
	// The Provider of the claim
	// https://hl7.org/fhir/r4/search.html#reference
	Requestor datatypes.JSON `gorm:"column:requestor;type:text;serializer:json" json:"requestor,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// The status of the ClaimResponse
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
	// The type of claim
	// https://hl7.org/fhir/r4/search.html#token
	Use datatypes.JSON `gorm:"column:use;type:text;serializer:json" json:"use,omitempty"`
}

func (s *FhirClaimResponse) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"created":     "date",
		"disposition": "string",
		"identifier":  "token",
		"insurer":     "reference",
		"language":    "token",
		"lastUpdated": "date",
		"outcome":     "token",
		"paymentDate": "date",
		"profile":     "reference",
		"request":     "reference",
		"requestor":   "reference",
		"sourceUri":   "uri",
		"status":      "token",
		"tag":         "token",
		"text":        "string",
		"type":        "special",
		"use":         "token",
	}
	return searchParameters
}
func (s *FhirClaimResponse) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Created
	createdResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'ClaimResponse.created')[0]")
	if err == nil && createdResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, createdResult.String())
		if err == nil {
			s.Created = t
		}
	}
	// extracting Disposition
	dispositionResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'ClaimResponse.disposition')[0]")
	if err == nil && dispositionResult.String() != "undefined" {
		s.Disposition = dispositionResult.String()
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'ClaimResponse.identifier')
							IdentifierProcessed = IdentifierResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(IdentifierProcessed)
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Insurer
	insurerResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'ClaimResponse.insurer'))")
	if err == nil && insurerResult.String() != "undefined" {
		s.Insurer = []byte(insurerResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'Resource.language')
							LanguageProcessed = LanguageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(LanguageProcessed)
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = t
		}
	}
	// extracting Outcome
	outcomeResult, err := vm.RunString(` 
							OutcomeResult = window.fhirpath.evaluate(fhirResource, 'ClaimResponse.outcome')
							OutcomeProcessed = OutcomeResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(OutcomeProcessed)
						 `)
	if err == nil && outcomeResult.String() != "undefined" {
		s.Outcome = []byte(outcomeResult.String())
	}
	// extracting PaymentDate
	paymentDateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'ClaimResponse.payment.date')[0]")
	if err == nil && paymentDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, paymentDateResult.String())
		if err == nil {
			s.PaymentDate = t
		}
	}
	// extracting Profile
	profileResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.meta.profile'))")
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting Request
	requestResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'ClaimResponse.request'))")
	if err == nil && requestResult.String() != "undefined" {
		s.Request = []byte(requestResult.String())
	}
	// extracting Requestor
	requestorResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'ClaimResponse.requestor'))")
	if err == nil && requestorResult.String() != "undefined" {
		s.Requestor = []byte(requestorResult.String())
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'ClaimResponse.status')
							StatusProcessed = StatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(StatusProcessed)
						 `)
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'Resource.meta.tag')
							TagProcessed = TagResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(TagProcessed)
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	// extracting Use
	useResult, err := vm.RunString(` 
							UseResult = window.fhirpath.evaluate(fhirResource, 'ClaimResponse.use')
							UseProcessed = UseResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(UseProcessed)
						 `)
	if err == nil && useResult.String() != "undefined" {
		s.Use = []byte(useResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirClaimResponse) TableName() string {
	return "fhir_claim_response"
}
