package flightserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"net/url"
)

// GetFlightLogs gets all the flight logs for a specific user.
func (s *Service) GetFlightLogs(UserUUID uuid.UUID) (FlightLogs, dutil.Error) {
	// set path
	s.serv.URL.Path = "log"
	// set query params
	qs := url.Values{"userUUID": {UserUUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return FlightLogs{}, e
	}

	// response structure
	type Data struct {
		FlightLogs `json:"flightLogs"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return FlightLogs{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return FlightLogs{}, e
	}
	return res.Data.FlightLogs, nil
}

func (s *Service) GetFlightLog(userUUID, UUID uuid.UUID) (FlightLog, dutil.Error) {
	// set path
	s.serv.URL.Path = "/log/-"
	// set query params
	qs := url.Values{"userUUID": {userUUID.String()}, "UUID": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return FlightLog{}, e
	}

	// response structure
	type Data struct {
		FlightLog `json:"flightLog"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return FlightLog{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return FlightLog{}, e
	}
	return res.Data.FlightLog, nil
}

// CreateFlightLog creates a new FlightLog entry.
func (s *Service) CreateFlightLog(log FlightLog) (FlightLog, dutil.Error) {
	// set path
	s.serv.URL.Path = "/log"
	// parse body
	p, e := dutil.MarshalReader(log)
	if e != nil {
		return FlightLog{}, e
	}
	// do request
	r, e := s.serv.NewRequest("POST", s.serv.URL.String(), nil, p)
	if e != nil {
		return FlightLog{}, e
	}
	// response structure
	type Data struct {
		FlightLog `json:"flightLog"`
	}
	res := struct {
		Data         `json:"data"`
		dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return FlightLog{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return FlightLog{}, e
	}

	return res.Data.FlightLog, nil
}

// UpdateFlightLog updates a new FlightLog entry.
func (s *Service) UpdateFlightLog(log FlightLog) (FlightLog, dutil.Error) {
	// set path
	s.serv.URL.Path = "/log/-"
	// parse body
	p, e := dutil.MarshalReader(log)
	if e != nil {
		return FlightLog{}, e
	}
	// do request
	r, e := s.serv.NewRequest("PUT", s.serv.URL.String(), nil, p)
	if e != nil {
		return FlightLog{}, e
	}
	// response structure
	type Data struct {
		FlightLog `json:"flightLog"`
	}
	res := struct {
		Data         `json:"data"`
		dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return FlightLog{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return FlightLog{}, e
	}

	return res.Data.FlightLog, nil
}
