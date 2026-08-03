package cli

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	geremmyas "github.com/woliveiras/geremmyas"
)

type gameRoutingCorpus struct {
	Version    int               `json:"version"`
	Method     string            `json:"method"`
	Limitation string            `json:"limitation"`
	Cases      []gameRoutingCase `json:"cases"`
}

type gameRoutingCase struct {
	Name          string   `json:"name"`
	Prompt        string   `json:"prompt"`
	ExpectedSkill string   `json:"expected_skill"`
	Signals       []string `json:"signals"`
}

func TestGameSkillRoutingDescriptionConformance(t *testing.T) {
	data, err := geremmyas.EmbeddedFiles.ReadFile("catalog/game-dev-routing.json")
	if err != nil {
		t.Fatalf("read routing corpus: %v", err)
	}
	var corpus gameRoutingCorpus
	if err := json.Unmarshal(data, &corpus); err != nil {
		t.Fatalf("decode routing corpus: %v", err)
	}
	if corpus.Method != "lexical-description-proxy" || corpus.Limitation == "" {
		t.Fatalf("routing corpus must identify its proxy method and limitation: %#v", corpus)
	}

	descriptions := map[string]string{}
	for _, name := range gameSkillNames() {
		description, err := embeddedSkillDescription(name)
		if err != nil {
			t.Fatal(err)
		}
		descriptions[name] = strings.ToLower(description)
	}

	covered := map[string]bool{}
	for _, testCase := range corpus.Cases {
		t.Run(testCase.Name, func(t *testing.T) {
			expectedDescription, ok := descriptions[testCase.ExpectedSkill]
			if !ok {
				t.Fatalf("unknown expected skill %q", testCase.ExpectedSkill)
			}
			if len(testCase.Signals) < 2 {
				t.Fatalf("case has %d signals, want at least 2", len(testCase.Signals))
			}
			prompt := strings.ToLower(testCase.Prompt)
			for _, signal := range testCase.Signals {
				if !strings.Contains(prompt, strings.ToLower(signal)) {
					t.Errorf("prompt does not contain declared signal %q", signal)
				}
			}

			winner, winnerScore, tied := "", -1, false
			for skill, description := range descriptions {
				score := 0
				for _, signal := range testCase.Signals {
					if strings.Contains(description, strings.ToLower(signal)) {
						score++
					}
				}
				if score > winnerScore {
					winner, winnerScore, tied = skill, score, false
				} else if score == winnerScore {
					tied = true
				}
			}
			if winnerScore < 2 || tied || winner != testCase.ExpectedSkill {
				t.Fatalf("lexical routing winner = %q score=%d tied=%t, want unique %q; expected description=%q", winner, winnerScore, tied, testCase.ExpectedSkill, expectedDescription)
			}
			covered[testCase.ExpectedSkill] = true
		})
	}
	for _, name := range gameSkillNames() {
		if !covered[name] {
			t.Errorf("routing corpus does not cover game skill %q", name)
		}
	}
}

func embeddedSkillDescription(name string) (string, error) {
	data, err := geremmyas.EmbeddedFiles.ReadFile("content/skills/" + name + "/SKILL.md")
	if err != nil {
		return "", fmt.Errorf("read %s skill: %w", name, err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "description:") {
			return strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "description:")), "\"'"), nil
		}
	}
	return "", fmt.Errorf("skill %s has no frontmatter description", name)
}

func gameSkillNames() []string {
	return []string{
		"game-ai-2d",
		"game-art-2d",
		"game-audio-2d",
		"game-build-and-release",
		"game-feel-2d",
		"game-performance-2d",
		"game-save-n-progress",
		"game-testing-2d",
		"game-ui-accessibility",
		"gameplay-programming-2d",
		"procedural-generation-2d",
	}
}
