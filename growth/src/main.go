// CRC: crc-CLI.md
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	persona := "daneel"

	// Extract --persona flag
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "--persona" {
			persona = args[i+1]
			args = append(args[:i], args[i+2:]...)
			break
		}
	}

	store := NewStore(persona)

	if len(args) == 0 {
		cmdBootstrap(store)
		return
	}

	switch args[0] {
	case "phrases":
		cmdPhrases(store, args[1:])
	case "contemplate":
		cmdContemplate(store, args[1:])
	case "tags":
		cmdTags(store)
	case "consider":
		cmdConsider(store, persona, args[1:])
	case "consideration":
		cmdConsideration(store, args[1:])
	case "commit":
		cmdCommit(store)
	case "abort":
		cmdAbort(store, args[1:])
	case "help", "--help":
		fmt.Print(renderUsage())
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		fmt.Print(renderUsage())
		os.Exit(1)
	}
}

func cmdBootstrap(store *Store) {
	const defaultN = 5
	phrases, err := store.randomPhrases(defaultN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	tags, _ := store.allTags()
	fmt.Print(renderBootstrap(phrases, tags))
}

func cmdPhrases(store *Store, args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: growth phrases N\n")
		os.Exit(1)
	}
	n, err := strconv.Atoi(args[0])
	if err != nil || n < 1 {
		fmt.Fprintf(os.Stderr, "N must be a positive integer\n")
		os.Exit(1)
	}
	phrases, err := store.randomPhrases(n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if len(phrases) == 0 {
		fmt.Println("The oracle is empty. No phrases have been committed yet.")
		fmt.Println("Use `growth consider` to start curating phrases.")
		return
	}
	tags, err := store.allTags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading tags: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(renderSessionBlock(phrases, tags))
}

func cmdContemplate(store *Store, args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: growth contemplate TAG\n")
		os.Exit(1)
	}
	tag := args[0]
	phrase, err := store.randomByTag(tag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if phrase == nil {
		fmt.Fprintf(os.Stderr, "no phrases with tag \"%s\"\n", tag)
		os.Exit(1)
	}
	fmt.Print(renderContemplate(phrase))
}

func cmdTags(store *Store) {
	tags, err := store.allTags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if len(tags) == 0 {
		fmt.Println("No tags yet.")
		return
	}
	fmt.Print(renderTagList(tags))
}

func cmdConsider(store *Store, persona string, args []string) {
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: growth consider \"PHRASE\" \"CONTEXT\" OBSERVER\n")
		os.Exit(1)
	}
	existing, err := store.loadCandidate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if existing != nil {
		fmt.Fprintf(os.Stderr, "A consideration is already in progress:\n\n")
		fmt.Fprint(os.Stderr, renderCandidate(existing))
		fmt.Fprintf(os.Stderr, "\nCommit or abort it first.\n")
		os.Exit(1)
	}
	c := &Candidate{
		Phrase:   args[0],
		Context:  args[1],
		Observer: args[2],
		Persona:  persona,
		Phase:    "proposed",
	}
	if err := store.saveCandidate(c); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(renderCandidate(c))
	fmt.Println("\nNext: `growth consideration evaluate` or `growth consideration tag TAG`")
}

func cmdConsideration(store *Store, args []string) {
	c, err := store.loadCandidate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if c == nil {
		fmt.Println("No consideration in progress.")
		return
	}

	if len(args) == 0 {
		fmt.Print(renderCandidate(c))
		return
	}

	switch args[0] {
	case "evaluate":
		c.Phase = "evaluate"
		if err := store.saveCandidate(c); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(renderEvaluate(c))

	case "tag":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "usage: growth consideration tag TAG [TAG...]\n")
			os.Exit(1)
		}
		for _, tag := range args[1:] {
			found := false
			for _, existing := range c.Tags {
				if existing == tag {
					found = true
					break
				}
			}
			if !found {
				c.Tags = append(c.Tags, tag)
			}
		}
		if err := store.saveCandidate(c); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(renderCandidate(c))

	default:
		fmt.Fprintf(os.Stderr, "unknown consideration subcommand: %s\n", args[0])
		fmt.Fprintf(os.Stderr, "usage: growth consideration [evaluate|tag TAG...]\n")
		os.Exit(1)
	}
}

func cmdCommit(store *Store) {
	c, err := store.loadCandidate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if c == nil {
		fmt.Println("No consideration in progress.")
		return
	}
	if len(c.Tags) == 0 {
		fmt.Fprintf(os.Stderr, "Cannot commit without tags. Add tags first:\n")
		fmt.Fprintf(os.Stderr, "  growth consideration tag TAG [TAG...]\n")
		os.Exit(1)
	}
	date := time.Now().Format("2006-01-02")
	if err := store.commitPhrase(c, date); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Committed: \"%s\"\n", c.Phrase)
}

func cmdAbort(store *Store, args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: growth abort \"REASON\"\n")
		os.Exit(1)
	}
	c, err := store.loadCandidate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if c == nil {
		fmt.Println("No consideration in progress.")
		return
	}
	date := time.Now().Format("2006-01-02")
	if err := store.rejectPhrase(c, args[0], date); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Rejected: \"%s\" — %s\n", c.Phrase, args[0])
}
