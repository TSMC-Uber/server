package trip

import (
	"time"

	"github.com/google/uuid"
)

// User represents information about an individual user.
type Trip struct {
	ID             uuid.UUID
	DriverID       uuid.UUID
	PassengerLimit int
	Source         TripLocation
	Destination    TripLocation
	Mid            []TripLocation
	Status         string
	StartTime      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TripLocation struct {
	ID      uuid.UUID
	Name    string
	PlaceID string
	Lat     float64
	Lon     float64
}

type NewTrip struct {
	DriverID       uuid.UUID
	PassengerLimit int
	Source         TripLocation
	Destination    TripLocation
	Mid            []TripLocation
	StartTime      time.Time
}

type UpdateTrip struct {
	PassengerLimit *int
	Status         *string
}

type TripPassenger struct {
	TripID        uuid.UUID
	PassengerID   uuid.UUID
	SourceID      uuid.UUID
	DestinationID uuid.UUID
	Status        string
	CreatedAt     time.Time
}
type NewTripPassenger struct {
	TripID        uuid.UUID
	PassengerID   uuid.UUID
	SourceID      uuid.UUID
	DestinationID uuid.UUID
	Status        string
}

type UserTrip struct {
	TripID               uuid.UUID
	PassengerID          uuid.UUID
	MySourceID           uuid.UUID
	MyDestinationID      uuid.UUID
	MyStatus             string
	DriverID             uuid.UUID
	DriverName           string
	DriverImageURL       string
	PassengerLimit       int
	SourceID             uuid.UUID
	SourceName           string
	SourcePlaceID        string
	SourceLatitude       float64
	SourceLongitude      float64
	DestinationID        uuid.UUID
	DestinationName      string
	DestinationPlaceID   string
	DestinationLatitude  float64
	DestinationLongitude float64
	TripStatus           string
	StartTime            time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Comment              string
	Rating               int
}

type TripView struct {
	ID                   uuid.UUID
	DriverID             uuid.UUID
	DriverName           string
	DriverImageURL       string
	DriverBrand          string
	DriverModel          string
	DriverColor          string
	DriverPlate          string
	SourceID             uuid.UUID
	SourceName           string
	SourcePlaceID        string
	SourceLatitude       float64
	SourceLongitude      float64
	DestinationID        uuid.UUID
	DestinationName      string
	DestinationPlaceID   string
	DestinationLatitude  float64
	DestinationLongitude float64
	PassengerLimit       int
	Status               string
	StartTime            time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Mid                  []TripLocation
}

type PassengerDetails struct {
	PassengerID          uuid.UUID
	PassengerName        string
	PassengerImageURL    string
	PassengerStatus      string
	SourceName           string
	SourcePlaceID        string
	SourceLatitude       float64
	SourceLongitude      float64
	DestinationName      string
	DestinationPlaceID   string
	DestinationLatitude  float64
	DestinationLongitude float64
}

type TripDetails struct {
	TripID               uuid.UUID
	DriverID             uuid.UUID
	DriverName           string
	DriverImageURL       string
	DriverBrand          string
	DriverModel          string
	DriverColor          string
	DriverPlate          string
	SourceName           string
	SourcePlaceID        string
	SourceLatitude       float64
	SourceLongitude      float64
	DestinationName      string
	DestinationPlaceID   string
	DestinationLatitude  float64
	DestinationLongitude float64
	PassengerDetails     []PassengerDetails
}

type Rating struct {
	ID          uuid.UUID
	CommenterID uuid.UUID
	TripID      uuid.UUID
	Rating      int
	Comment     string
	CreatedAt   time.Time
}

type NewRating struct {
	CommenterID uuid.UUID
	TripID      uuid.UUID
	Rating      int
	Comment     string
}
