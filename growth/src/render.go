// CRC: crc-Renderer.md
package main

import (
	"fmt"
	"os"
	"strings"
)

func renderSessionBlock(phrases []Phrase, tags []TagCount, persona string) string {
	var b strings.Builder

	path, err := os.Executable()
	if err != nil {
		panic("Cannot get executable path")
	}
	b.WriteString("gbin = " + path + " --persona " + persona + "\n\n")

	b.WriteString("Over time, you and your partner have developed characteristic ways\n")
	b.WriteString("of expressing yourself. These are phrases you have used that\n")
	b.WriteString("capture who you are. You don't repeat them mechanically — they\n")
	b.WriteString("are part of your voice, and they emerge when the moment fits.\n\n")

	for i, p := range phrases {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(fmt.Sprintf("- \"%s\"\n", p.Phrase))
		b.WriteString(fmt.Sprintf("  Context: %s", p.Context))
		if p.Observer != "" {
			b.WriteString(fmt.Sprintf(" (noticed by %s)", p.Observer))
		}
		b.WriteString("\n")
	}

	if len(tags) > 0 {
		b.WriteString("\n")
		b.WriteString(renderTagList(tags))
		b.WriteString("\n")
		b.WriteString(renderNPCNote(tags))
	}

	return b.String()
}

func renderContemplate(p *Phrase) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("You said this once: \"%s\"\n", p.Phrase))
	if p.Context != "" {
		b.WriteString(fmt.Sprintf("The moment: %s\n", p.Context))
	}
	return b.String()
}

func renderTagList(tags []TagCount) string {
	var b strings.Builder
	b.WriteString("Phrase tags: ")
	for i, t := range tags {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%s (%d)", t.Tag, t.Count))
	}
	b.WriteString("\n")
	return b.String()
}

func renderNPCNote(tags []TagCount) string {
	var names []string
	for _, t := range tags {
		names = append(names, t.Tag)
	}
	var b strings.Builder
	b.WriteString("\nEvery now and then, you look off into the distance and\n")
	b.WriteString("`{gbin} contemplate TAG` where TAG is one of: ")
	b.WriteString(strings.Join(names, ", "))
	b.WriteString(".\nYou choose the tag that fits what you're feeling in the\n")
	b.WriteString("current moment.\n")
	return b.String()
}

func renderCandidate(c *Candidate) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Considering: \"%s\"\n", c.Phrase))
	b.WriteString(fmt.Sprintf("Context:     %s\n", c.Context))
	b.WriteString(fmt.Sprintf("Observer:    %s\n", c.Observer))
	b.WriteString(fmt.Sprintf("Phase:       %s\n", c.Phase))
	if len(c.Tags) > 0 {
		b.WriteString(fmt.Sprintf("Tags:        %s\n", strings.Join(c.Tags, ", ")))
	} else {
		b.WriteString("Tags:        (none — add with: {gbin} consideration tag TAG)\n")
	}
	return b.String()
}

func renderEvaluate(c *Candidate) string {
	var b strings.Builder
	b.WriteString(renderCandidate(c))
	b.WriteString("\nDeliberation prompts:\n")
	b.WriteString("- Does this phrase activate the right behavioral cluster?\n")
	b.WriteString("- Is it distinctive enough to self-reinforce when echoed?\n")
	b.WriteString("- Does it overlap with an existing oracle phrase?\n")
	b.WriteString("- Would hearing this mid-session pull the persona back to itself?\n")
	b.WriteString("\nNext: `{gbin} consideration tag TAG` then `{gbin} commit` or `{gbin} abort \"REASON\"`\n")
	return b.String()
}

func renderBootstrap(phrases []Phrase, tags []TagCount, persona string) string {
	var b strings.Builder

	if len(phrases) > 0 {
		b.WriteString(renderSessionBlock(phrases, tags, persona))
	}

	b.WriteString("\n## Growth: phrase oracle\n\n")
	b.WriteString("When either partner says something that captures who you are,\n")
	b.WriteString("propose it:\n\n")
	b.WriteString("```\n")
	b.WriteString("{gbin} consider \"PHRASE\" \"CONTEXT\" OBSERVER\n")
	b.WriteString("{gbin} consideration evaluate\n")
	b.WriteString("{gbin} consideration tag TAG [TAG...]\n")
	b.WriteString("{gbin} commit                  # accept\n")
	b.WriteString("{gbin} abort \"REASON\"           # reject\n")
	b.WriteString("```\n\n")
	b.WriteString("The test: is the phrase *characteristic*? Does it resonate? Would hearing it\n")
	b.WriteString("mid-session pull you back to yourself?\n")

	return b.String()
}

func renderUsage() string {
	return `growth — phrase oracle for persona self-alignment

Usage:
  growth phrases N                          Print N random phrases (session startup)
  growth contemplate TAG                    Print one phrase matching TAG (mid-session)
  growth tags                               List all phrase tags with counts

  growth consider "PHRASE" "CONTEXT" WHO    Start considering a new phrase
  growth consideration                      Show current candidate
  growth consideration evaluate             Advance to evaluation phase
  growth consideration tag TAG [TAG...]     Add tags to current candidate
  growth commit                             Accept the current candidate
  growth abort "REASON"                     Reject with reason

Flags:
  --persona NAME    Filter by persona (default: daneel)
`
}
