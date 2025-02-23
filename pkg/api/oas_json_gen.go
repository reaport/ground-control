// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *AirplaneGetParkingSpotOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AirplaneGetParkingSpotOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("nodeId")
		e.Str(s.NodeId)
	}
}

var jsonFieldsNameOfAirplaneGetParkingSpotOK = [1]string{
	0: "nodeId",
}

// Decode decodes AirplaneGetParkingSpotOK from json.
func (s *AirplaneGetParkingSpotOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AirplaneGetParkingSpotOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "nodeId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.NodeId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nodeId\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AirplaneGetParkingSpotOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAirplaneGetParkingSpotOK) {
					name = jsonFieldsNameOfAirplaneGetParkingSpotOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AirplaneGetParkingSpotOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AirplaneGetParkingSpotOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *AirplaneIDServiceTypeGetOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AirplaneIDServiceTypeGetOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("nodeId")
		e.Str(s.NodeId)
	}
}

var jsonFieldsNameOfAirplaneIDServiceTypeGetOK = [1]string{
	0: "nodeId",
}

// Decode decodes AirplaneIDServiceTypeGetOK from json.
func (s *AirplaneIDServiceTypeGetOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AirplaneIDServiceTypeGetOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "nodeId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.NodeId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nodeId\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AirplaneIDServiceTypeGetOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAirplaneIDServiceTypeGetOK) {
					name = jsonFieldsNameOfAirplaneIDServiceTypeGetOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AirplaneIDServiceTypeGetOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AirplaneIDServiceTypeGetOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *AirportMap) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AirportMap) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("nodes")
		e.ArrStart()
		for _, elem := range s.Nodes {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("edges")
		e.ArrStart()
		for _, elem := range s.Edges {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfAirportMap = [2]string{
	0: "nodes",
	1: "edges",
}

// Decode decodes AirportMap from json.
func (s *AirportMap) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AirportMap to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "nodes":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Nodes = make([]Node, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Node
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Nodes = append(s.Nodes, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nodes\"")
			}
		case "edges":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				s.Edges = make([]Edge, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Edge
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Edges = append(s.Edges, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"edges\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AirportMap")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAirportMap) {
					name = jsonFieldsNameOfAirportMap[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AirportMap) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AirportMap) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Edge) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Edge) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("from")
		e.Str(s.From)
	}
	{
		e.FieldStart("to")
		e.Str(s.To)
	}
	{
		e.FieldStart("distance")
		e.Float64(s.Distance)
	}
}

var jsonFieldsNameOfEdge = [3]string{
	0: "from",
	1: "to",
	2: "distance",
}

// Decode decodes Edge from json.
func (s *Edge) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Edge to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "from":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.From = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from\"")
			}
		case "to":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.To = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"to\"")
			}
		case "distance":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Float64()
				s.Distance = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"distance\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Edge")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfEdge) {
					name = jsonFieldsNameOfEdge[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Edge) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Edge) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes MovingGetRouteOKApplicationJSON as json.
func (s MovingGetRouteOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []string(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		e.Str(elem)
	}
	e.ArrEnd()
}

// Decode decodes MovingGetRouteOKApplicationJSON from json.
func (s *MovingGetRouteOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MovingGetRouteOKApplicationJSON to nil")
	}
	var unwrapped []string
	if err := func() error {
		unwrapped = make([]string, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem string
			v, err := d.Str()
			elem = string(v)
			if err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = MovingGetRouteOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s MovingGetRouteOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MovingGetRouteOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MovingGetRouteReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MovingGetRouteReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("from")
		e.Str(s.From)
	}
	{
		e.FieldStart("to")
		e.Str(s.To)
	}
	{
		e.FieldStart("type")
		s.Type.Encode(e)
	}
}

var jsonFieldsNameOfMovingGetRouteReq = [3]string{
	0: "from",
	1: "to",
	2: "type",
}

// Decode decodes MovingGetRouteReq from json.
func (s *MovingGetRouteReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MovingGetRouteReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "from":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.From = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from\"")
			}
		case "to":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.To = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"to\"")
			}
		case "type":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MovingGetRouteReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMovingGetRouteReq) {
					name = jsonFieldsNameOfMovingGetRouteReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MovingGetRouteReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MovingGetRouteReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MovingNotifyArrivalReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MovingNotifyArrivalReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("vehicleId")
		e.Str(s.VehicleId)
	}
	{
		e.FieldStart("vehicleType")
		s.VehicleType.Encode(e)
	}
	{
		e.FieldStart("nodeId")
		e.Str(s.NodeId)
	}
}

var jsonFieldsNameOfMovingNotifyArrivalReq = [3]string{
	0: "vehicleId",
	1: "vehicleType",
	2: "nodeId",
}

// Decode decodes MovingNotifyArrivalReq from json.
func (s *MovingNotifyArrivalReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MovingNotifyArrivalReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "vehicleId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.VehicleId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicleId\"")
			}
		case "vehicleType":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.VehicleType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicleType\"")
			}
		case "nodeId":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.NodeId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nodeId\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MovingNotifyArrivalReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMovingNotifyArrivalReq) {
					name = jsonFieldsNameOfMovingNotifyArrivalReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MovingNotifyArrivalReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MovingNotifyArrivalReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MovingRegisterVehicleOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MovingRegisterVehicleOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("nodeId")
		e.Str(s.NodeId)
	}
	{
		e.FieldStart("vehicleId")
		e.Str(s.VehicleId)
	}
}

var jsonFieldsNameOfMovingRegisterVehicleOK = [2]string{
	0: "nodeId",
	1: "vehicleId",
}

// Decode decodes MovingRegisterVehicleOK from json.
func (s *MovingRegisterVehicleOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MovingRegisterVehicleOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "nodeId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.NodeId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"nodeId\"")
			}
		case "vehicleId":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.VehicleId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicleId\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MovingRegisterVehicleOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMovingRegisterVehicleOK) {
					name = jsonFieldsNameOfMovingRegisterVehicleOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MovingRegisterVehicleOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MovingRegisterVehicleOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MovingRequestMoveOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MovingRequestMoveOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("distance")
		e.Float64(s.Distance)
	}
}

var jsonFieldsNameOfMovingRequestMoveOK = [1]string{
	0: "distance",
}

// Decode decodes MovingRequestMoveOK from json.
func (s *MovingRequestMoveOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MovingRequestMoveOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "distance":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Float64()
				s.Distance = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"distance\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MovingRequestMoveOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMovingRequestMoveOK) {
					name = jsonFieldsNameOfMovingRequestMoveOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MovingRequestMoveOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MovingRequestMoveOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *MovingRequestMoveReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *MovingRequestMoveReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("vehicleId")
		e.Str(s.VehicleId)
	}
	{
		e.FieldStart("vehicleType")
		s.VehicleType.Encode(e)
	}
	{
		e.FieldStart("from")
		e.Str(s.From)
	}
	{
		e.FieldStart("to")
		e.Str(s.To)
	}
}

var jsonFieldsNameOfMovingRequestMoveReq = [4]string{
	0: "vehicleId",
	1: "vehicleType",
	2: "from",
	3: "to",
}

// Decode decodes MovingRequestMoveReq from json.
func (s *MovingRequestMoveReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode MovingRequestMoveReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "vehicleId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.VehicleId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicleId\"")
			}
		case "vehicleType":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.VehicleType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicleType\"")
			}
		case "from":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.From = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"from\"")
			}
		case "to":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Str()
				s.To = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"to\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode MovingRequestMoveReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfMovingRequestMoveReq) {
					name = jsonFieldsNameOfMovingRequestMoveReq[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *MovingRequestMoveReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *MovingRequestMoveReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Node) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Node) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Str(s.ID)
	}
	{
		e.FieldStart("types")
		e.ArrStart()
		for _, elem := range s.Types {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		if s.Vehicles != nil {
			e.FieldStart("vehicles")
			e.ArrStart()
			for _, elem := range s.Vehicles {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfNode = [3]string{
	0: "id",
	1: "types",
	2: "vehicles",
}

// Decode decodes Node from json.
func (s *Node) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Node to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "types":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				s.Types = make([]VehicleType, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem VehicleType
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Types = append(s.Types, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"types\"")
			}
		case "vehicles":
			if err := func() error {
				s.Vehicles = make([]Vehicle, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Vehicle
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Vehicles = append(s.Vehicles, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vehicles\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Node")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfNode) {
					name = jsonFieldsNameOfNode[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Node) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Node) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Vehicle) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Vehicle) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Str(s.ID)
	}
	{
		e.FieldStart("type")
		s.Type.Encode(e)
	}
}

var jsonFieldsNameOfVehicle = [2]string{
	0: "id",
	1: "type",
}

// Decode decodes Vehicle from json.
func (s *Vehicle) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Vehicle to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "type":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Vehicle")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfVehicle) {
					name = jsonFieldsNameOfVehicle[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Vehicle) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Vehicle) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VehicleType as json.
func (s VehicleType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes VehicleType from json.
func (s *VehicleType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VehicleType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch VehicleType(v) {
	case VehicleTypeAirplane:
		*s = VehicleTypeAirplane
	case VehicleTypeCatering:
		*s = VehicleTypeCatering
	case VehicleTypeRefueling:
		*s = VehicleTypeRefueling
	case VehicleTypeCleaning:
		*s = VehicleTypeCleaning
	case VehicleTypeBaggage:
		*s = VehicleTypeBaggage
	case VehicleTypeFollowMe:
		*s = VehicleTypeFollowMe
	case VehicleTypeCharging:
		*s = VehicleTypeCharging
	case VehicleTypeBus:
		*s = VehicleTypeBus
	case VehicleTypeRamp:
		*s = VehicleTypeRamp
	default:
		*s = VehicleType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s VehicleType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VehicleType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
