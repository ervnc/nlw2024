// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const confirmParticipant = `-- name: ConfirmParticipant :exec
UPDATE participants
SET "is_confirmed" = true
WHERE id = $1
`

func (q *Queries) ConfirmParticipant(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, confirmParticipant, id)
	return err
}

const createActivity = `-- name: CreateActivity :one
INSERT INTO activities
  ("trip_id", "title", "occurs_at") VALUES
  ($1, $2, $3)
RETURNING "id"
`

type CreateActivityParams struct {
	TripID   uuid.UUID
	Title    string
	OccursAt pgtype.Timestamp
}

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createActivity, arg.TripID, arg.Title, arg.OccursAt)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createTripLink = `-- name: CreateTripLink :one
INSERT INTO links
  ("trip_id", "title", "url") VALUES
  ($1, $2, $3)
RETURNING "id"
`

type CreateTripLinkParams struct {
	TripID uuid.UUID
	Title  string
	Url    string
}

func (q *Queries) CreateTripLink(ctx context.Context, arg CreateTripLinkParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createTripLink, arg.TripID, arg.Title, arg.Url)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getParticipant = `-- name: GetParticipant :one
SELECT 
  "id",
  "trip_id",
  "email",
  "is_confirmed"
FROM participants
WHERE id = $1
`

func (q *Queries) GetParticipant(ctx context.Context, id uuid.UUID) (Participant, error) {
	row := q.db.QueryRow(ctx, getParticipant, id)
	var i Participant
	err := row.Scan(
		&i.ID,
		&i.TripID,
		&i.Email,
		&i.IsConfirmed,
	)
	return i, err
}

const getParticipants = `-- name: GetParticipants :many
SELECT 
  "id",
  "trip_id",
  "email",
  "is_confirmed"
FROM participants
WHERE trip_id = $1
`

func (q *Queries) GetParticipants(ctx context.Context, tripID uuid.UUID) ([]Participant, error) {
	rows, err := q.db.Query(ctx, getParticipants, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Participant
	for rows.Next() {
		var i Participant
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.Email,
			&i.IsConfirmed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrip = `-- name: GetTrip :one
SELECT 
  "id",
  "destination",
  "owner_email",
  "owner_name",
  "starts_at",
  "ends_at"
FROM trips
WHERE id = $1
`

type GetTripRow struct {
	ID          uuid.UUID
	Destination string
	OwnerEmail  string
	OwnerName   string
	StartsAt    pgtype.Timestamp
	EndsAt      pgtype.Timestamp
}

func (q *Queries) GetTrip(ctx context.Context, id uuid.UUID) (GetTripRow, error) {
	row := q.db.QueryRow(ctx, getTrip, id)
	var i GetTripRow
	err := row.Scan(
		&i.ID,
		&i.Destination,
		&i.OwnerEmail,
		&i.OwnerName,
		&i.StartsAt,
		&i.EndsAt,
	)
	return i, err
}

const getTripActivities = `-- name: GetTripActivities :many
SELECT
  "id",
  "trip_id",
  "title",
  "occurs_at"
FROM activities
WHERE trip_id = $1
`

func (q *Queries) GetTripActivities(ctx context.Context, tripID uuid.UUID) ([]Activity, error) {
	rows, err := q.db.Query(ctx, getTripActivities, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Activity
	for rows.Next() {
		var i Activity
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.Title,
			&i.OccursAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTripLinks = `-- name: GetTripLinks :many
SELECT
  "id",
  "trip_id",
  "title",
  "url"
FROM links
WHERE trip_id = $1
`

func (q *Queries) GetTripLinks(ctx context.Context, tripID uuid.UUID) ([]Link, error) {
	rows, err := q.db.Query(ctx, getTripLinks, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.ID,
			&i.TripID,
			&i.Title,
			&i.Url,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertTrip = `-- name: InsertTrip :one
INSERT INTO trips
  ( "destination", "owner_email", "owner_name", "starts_at", "ends_at" ) VALUES
  ( $1, $2, $3, $4, $5 )
RETURNING "id"
`

type InsertTripParams struct {
	Destination string
	OwnerEmail  string
	OwnerName   string
	StartsAt    pgtype.Timestamp
	EndsAt      pgtype.Timestamp
}

func (q *Queries) InsertTrip(ctx context.Context, arg InsertTripParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertTrip,
		arg.Destination,
		arg.OwnerEmail,
		arg.OwnerName,
		arg.StartsAt,
		arg.EndsAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const inviteParticipantToTrip = `-- name: InviteParticipantToTrip :one
INSERT INTO participants
  ("trip_id", "email") VALUES
  ($1, $2)
RETURNING "id"
`

type InviteParticipantToTripParams struct {
	TripID uuid.UUID
	Email  string
}

func (q *Queries) InviteParticipantToTrip(ctx context.Context, arg InviteParticipantToTripParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, inviteParticipantToTrip, arg.TripID, arg.Email)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

type InviteParticipiantsToTripParams struct {
	TripID uuid.UUID
	Email  string
}

const updateTrip = `-- name: UpdateTrip :exec
UPDATE trips
SET 
  "destination" = $1,
  "starts_at" = $3,
  "ends_at" = $2,
  "is_confirmed" = $4
WHERE id = $5
`

type UpdateTripParams struct {
	Destination string
	EndsAt      pgtype.Timestamp
	StartsAt    pgtype.Timestamp
	IsConfirmed bool
	ID          uuid.UUID
}

func (q *Queries) UpdateTrip(ctx context.Context, arg UpdateTripParams) error {
	_, err := q.db.Exec(ctx, updateTrip,
		arg.Destination,
		arg.EndsAt,
		arg.StartsAt,
		arg.IsConfirmed,
		arg.ID,
	)
	return err
}
