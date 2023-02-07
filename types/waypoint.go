package types

type WaypointCollection struct {
	Feature []WaypointFeature `json:"features"`
}

type WaypointFeature struct {
	Property WaypointProperty `json:"properties"`
}

type WaypointProperty struct {
	Distance float32        `json:"distance"`
	Metric   string         `json:"distance_units"`
	Time     float32        `json:"time"`
	Legs     []WaypointLegs `json:"legs"`
}

type WaypointLegs struct {
	Steps []WaypointSteps `json:"steps"`
}

type WaypointSteps struct {
	Distance    float32             `json:"distance"`
	Time        float32             `json:"time"`
	Instruction Waypointinstruction `json:"instruction"`
}

type Waypointinstruction struct {
	Text string `json:"text"`
}
