// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MarketChangeMessage market change message
//
// swagger:model MarketChangeMessage
type MarketChangeMessage struct {
	idField int32

	// Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect)
	Clk string `json:"clk,omitempty"`

	// Conflate Milliseconds - the conflation rate (may differ from that requested if subscription is delayed)
	ConflateMs int64 `json:"conflateMs,omitempty"`

	// Change Type - set to indicate the type of change - if null this is a delta)
	// Enum: [SUB_IMAGE RESUB_DELTA HEARTBEAT]
	Ct string `json:"ct,omitempty"`

	// Heartbeat Milliseconds - the heartbeat rate (may differ from requested: bounds are 500 to 30000)
	HeartbeatMs int64 `json:"heartbeatMs,omitempty"`

	// Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect)
	InitialClk string `json:"initialClk,omitempty"`

	// MarketChanges - the modifications to markets (will be null on a heartbeat
	Mc []*MarketChange `json:"mc"`

	// Publish Time (in millis since epoch) that the changes were generated
	Pt int64 `json:"pt,omitempty"`

	// Segment Type - if the change is split into multiple segments, this denotes the beginning and end of a change, and segments in between. Will be null if data is not segmented
	// Enum: [SEG_START SEG SEG_END]
	SegmentType string `json:"segmentType,omitempty"`

	// Stream status: set to null if the exchange stream data is up to date and 503 if the downstream services are experiencing latencies
	Status int32 `json:"status,omitempty"`
}

// ID gets the id of this subtype
func (m *MarketChangeMessage) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *MarketChangeMessage) SetID(val int32) {
	m.idField = val
}

// Op gets the op of this subtype
func (m *MarketChangeMessage) Op() string {
	return "mcm"
}

// SetOp sets the op of this subtype
func (m *MarketChangeMessage) SetOp(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MarketChangeMessage) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect)
		Clk string `json:"clk,omitempty"`

		// Conflate Milliseconds - the conflation rate (may differ from that requested if subscription is delayed)
		ConflateMs int64 `json:"conflateMs,omitempty"`

		// Change Type - set to indicate the type of change - if null this is a delta)
		// Enum: [SUB_IMAGE RESUB_DELTA HEARTBEAT]
		Ct string `json:"ct,omitempty"`

		// Heartbeat Milliseconds - the heartbeat rate (may differ from requested: bounds are 500 to 30000)
		HeartbeatMs int64 `json:"heartbeatMs,omitempty"`

		// Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect)
		InitialClk string `json:"initialClk,omitempty"`

		// MarketChanges - the modifications to markets (will be null on a heartbeat
		Mc []*MarketChange `json:"mc"`

		// Publish Time (in millis since epoch) that the changes were generated
		Pt int64 `json:"pt,omitempty"`

		// Segment Type - if the change is split into multiple segments, this denotes the beginning and end of a change, and segments in between. Will be null if data is not segmented
		// Enum: [SEG_START SEG SEG_END]
		SegmentType string `json:"segmentType,omitempty"`

		// Stream status: set to null if the exchange stream data is up to date and 503 if the downstream services are experiencing latencies
		Status int32 `json:"status,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ID int32 `json:"id,omitempty"`

		Op string `json:"op,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result MarketChangeMessage

	result.idField = base.ID

	if base.Op != result.Op() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid op value: %q", base.Op)
	}

	result.Clk = data.Clk
	result.ConflateMs = data.ConflateMs
	result.Ct = data.Ct
	result.HeartbeatMs = data.HeartbeatMs
	result.InitialClk = data.InitialClk
	result.Mc = data.Mc
	result.Pt = data.Pt
	result.SegmentType = data.SegmentType
	result.Status = data.Status

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MarketChangeMessage) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect)
		Clk string `json:"clk,omitempty"`

		// Conflate Milliseconds - the conflation rate (may differ from that requested if subscription is delayed)
		ConflateMs int64 `json:"conflateMs,omitempty"`

		// Change Type - set to indicate the type of change - if null this is a delta)
		// Enum: [SUB_IMAGE RESUB_DELTA HEARTBEAT]
		Ct string `json:"ct,omitempty"`

		// Heartbeat Milliseconds - the heartbeat rate (may differ from requested: bounds are 500 to 30000)
		HeartbeatMs int64 `json:"heartbeatMs,omitempty"`

		// Token value (non-null) should be stored and passed in a MarketSubscriptionMessage to resume subscription (in case of disconnect)
		InitialClk string `json:"initialClk,omitempty"`

		// MarketChanges - the modifications to markets (will be null on a heartbeat
		Mc []*MarketChange `json:"mc"`

		// Publish Time (in millis since epoch) that the changes were generated
		Pt int64 `json:"pt,omitempty"`

		// Segment Type - if the change is split into multiple segments, this denotes the beginning and end of a change, and segments in between. Will be null if data is not segmented
		// Enum: [SEG_START SEG SEG_END]
		SegmentType string `json:"segmentType,omitempty"`

		// Stream status: set to null if the exchange stream data is up to date and 503 if the downstream services are experiencing latencies
		Status int32 `json:"status,omitempty"`
	}{

		Clk: m.Clk,

		ConflateMs: m.ConflateMs,

		Ct: m.Ct,

		HeartbeatMs: m.HeartbeatMs,

		InitialClk: m.InitialClk,

		Mc: m.Mc,

		Pt: m.Pt,

		SegmentType: m.SegmentType,

		Status: m.Status,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ID int32 `json:"id,omitempty"`

		Op string `json:"op,omitempty"`
	}{

		ID: m.ID(),

		Op: m.Op(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this market change message
func (m *MarketChangeMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var marketChangeMessageTypeCtPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUB_IMAGE","RESUB_DELTA","HEARTBEAT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketChangeMessageTypeCtPropEnum = append(marketChangeMessageTypeCtPropEnum, v)
	}
}

// property enum
func (m *MarketChangeMessage) validateCtEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, marketChangeMessageTypeCtPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MarketChangeMessage) validateCt(formats strfmt.Registry) error {

	if swag.IsZero(m.Ct) { // not required
		return nil
	}

	// value enum
	if err := m.validateCtEnum("ct", "body", m.Ct); err != nil {
		return err
	}

	return nil
}

func (m *MarketChangeMessage) validateMc(formats strfmt.Registry) error {

	if swag.IsZero(m.Mc) { // not required
		return nil
	}

	for i := 0; i < len(m.Mc); i++ {
		if swag.IsZero(m.Mc[i]) { // not required
			continue
		}

		if m.Mc[i] != nil {
			if err := m.Mc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var marketChangeMessageTypeSegmentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEG_START","SEG","SEG_END"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketChangeMessageTypeSegmentTypePropEnum = append(marketChangeMessageTypeSegmentTypePropEnum, v)
	}
}

// property enum
func (m *MarketChangeMessage) validateSegmentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, marketChangeMessageTypeSegmentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MarketChangeMessage) validateSegmentType(formats strfmt.Registry) error {

	if swag.IsZero(m.SegmentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSegmentTypeEnum("segmentType", "body", m.SegmentType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this market change message based on the context it is used
func (m *MarketChangeMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketChangeMessage) contextValidateMc(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mc); i++ {

		if m.Mc[i] != nil {
			if err := m.Mc[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketChangeMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketChangeMessage) UnmarshalBinary(b []byte) error {
	var res MarketChangeMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
