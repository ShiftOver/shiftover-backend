package dto

// CounterEntity represents the counter entity
type CounterEntity struct {
	ID            string `bson:"_id"`
	SequenceValue int    `bson:"sequence_value"`
}
