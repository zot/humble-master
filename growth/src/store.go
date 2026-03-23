// CRC: crc-Store.md
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
)

type Phrase struct {
	Type      string   `json:"type"`
	Persona   string   `json:"persona"`
	Phrase    string   `json:"phrase"`
	Context   string   `json:"context"`
	Observer  string   `json:"observer"`
	Tags      []string `json:"tags"`
	Committed string   `json:"committed,omitempty"`
	Reason    string   `json:"reason,omitempty"`
	Rejected  string   `json:"rejected,omitempty"`
}

type Candidate struct {
	Phrase   string   `json:"phrase"`
	Context  string   `json:"context"`
	Observer string   `json:"observer"`
	Persona  string   `json:"persona"`
	Tags     []string `json:"tags"`
	Phase    string   `json:"phase"`
}

type Store struct {
	baseDir string
	persona string
}

func NewStore(persona string) *Store {
	home, _ := os.UserHomeDir()
	s := &Store{
		baseDir: filepath.Join(home, ".claude", "personal"),
		persona: persona,
	}
	s.ensureDir()
	return s
}

func (s *Store) jsonlPath() string {
	return filepath.Join(s.baseDir, s.persona+".jsonl")
}

func (s *Store) candidatePath() string {
	return filepath.Join(s.baseDir, "considering.json")
}

func (s *Store) ensureDir() error {
	return os.MkdirAll(s.baseDir, 0755)
}

func (s *Store) loadPhrases() ([]Phrase, error) {
	f, err := os.Open(s.jsonlPath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	var phrases []Phrase
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		var p Phrase
		if err := json.Unmarshal(line, &p); err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping malformed line %d: %v\n", lineNum, err)
			continue
		}
		if p.Persona == s.persona && p.Type == "phrase" {
			phrases = append(phrases, p)
		}
	}
	return phrases, scanner.Err()
}

func (s *Store) randomPhrases(n int) ([]Phrase, error) {
	phrases, err := s.loadPhrases()
	if err != nil {
		return nil, err
	}
	if len(phrases) == 0 {
		return nil, nil
	}
	if n >= len(phrases) {
		rand.Shuffle(len(phrases), func(i, j int) {
			phrases[i], phrases[j] = phrases[j], phrases[i]
		})
		return phrases, nil
	}
	rand.Shuffle(len(phrases), func(i, j int) {
		phrases[i], phrases[j] = phrases[j], phrases[i]
	})
	return phrases[:n], nil
}

func (s *Store) randomByTag(tag string) (*Phrase, error) {
	phrases, err := s.loadPhrases()
	if err != nil {
		return nil, err
	}
	var matched []Phrase
	for _, p := range phrases {
		for _, t := range p.Tags {
			if t == tag {
				matched = append(matched, p)
				break
			}
		}
	}
	if len(matched) == 0 {
		return nil, nil
	}
	pick := matched[rand.Intn(len(matched))]
	return &pick, nil
}

type TagCount struct {
	Tag   string
	Count int
}

func (s *Store) allTags() ([]TagCount, error) {
	phrases, err := s.loadPhrases()
	if err != nil {
		return nil, err
	}
	counts := make(map[string]int)
	for _, p := range phrases {
		for _, t := range p.Tags {
			counts[t]++
		}
	}
	var tags []TagCount
	for t, c := range counts {
		tags = append(tags, TagCount{t, c})
	}
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Tag < tags[j].Tag
	})
	return tags, nil
}

func (s *Store) loadCandidate() (*Candidate, error) {
	data, err := os.ReadFile(s.candidatePath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var c Candidate
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *Store) saveCandidate(c *Candidate) error {
	if err := s.ensureDir(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.candidatePath(), data, 0644)
}

func (s *Store) clearCandidate() error {
	err := os.Remove(s.candidatePath())
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func (s *Store) appendRecord(p Phrase) error {
	if err := s.ensureDir(); err != nil {
		return err
	}
	f, err := os.OpenFile(s.jsonlPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = f.Write(append(data, '\n'))
	return err
}

func (s *Store) commitPhrase(c *Candidate, date string) error {
	p := Phrase{
		Type:      "phrase",
		Persona:   c.Persona,
		Phrase:    c.Phrase,
		Context:   c.Context,
		Observer:  c.Observer,
		Tags:      c.Tags,
		Committed: date,
	}
	if err := s.appendRecord(p); err != nil {
		return err
	}
	return s.clearCandidate()
}

func (s *Store) rejectPhrase(c *Candidate, reason, date string) error {
	p := Phrase{
		Type:     "rejected",
		Persona:  c.Persona,
		Phrase:   c.Phrase,
		Context:  c.Context,
		Observer: c.Observer,
		Tags:     c.Tags,
		Reason:   reason,
		Rejected: date,
	}
	if err := s.appendRecord(p); err != nil {
		return err
	}
	return s.clearCandidate()
}
