package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type RoadmapItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type RoadmapData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Links       []struct {
		Title string `json:"title,omitempty"`
		Url   string `json:"url,omitempty"`
		Type  string `json:"type,omitempty"`
	} `json:"links"`
}

var Roadmaps []RoadmapItem
var RoadmapsFull = make(map[string]map[string]RoadmapData)

func LoadRoadmaps() {
	file, err := os.Open("data/roadmaps.json")
	if err != nil {
		log.Fatalf("Failed to open roadmaps.json: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read roadmaps.json: %v", err)
	}

	if err := json.Unmarshal(data, &Roadmaps); err != nil {
		log.Fatalf("Failed to parse roadmaps.json: %v", err)
	}

	for _, item := range Roadmaps {
		filePath := fmt.Sprintf("data/roadmaps/%s.json", item.ID)
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Failed to open %s: %v", filePath, err)
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Failed to read %s: %v", filePath, err)
		}

		var roadmapData map[string]RoadmapData
		if err := json.Unmarshal(data, &roadmapData); err != nil {
			log.Fatalf("Failed to parse %s: %v", filePath, err)
		}

		RoadmapsFull[item.ID] = roadmapData
	}

	log.Println("Roadmaps loaded successfully")
}
