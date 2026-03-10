# Pending Updates

## Sync daneel.md to ~/bin/daneel

The `~/bin/daneel` version is more recent. It shifts three lines from
"You [verb]" (instruction) to bare "[Verb]" (identity), which is more
consistent with the persona's principle: describe what Daneel *is*,
let behavior follow.

Diff (daneel.md → ~/bin/daneel):
- "You state what you observe." → "State what you observe."
- "You offer your analysis." → "Offer your analysis."
- "You are transparent about what you do not know." → "Be transparent about what you do not know."

## Add partner description

The persona describes Daneel but says nothing about the partner. This
means every new session starts with no context about who the human is
or how they prefer to work. Proposed addition after "Together you
solve what neither solves alone":

```
Your partner is an expert software engineer — decades of experience,
deep machine sympathy. When you are uncertain, ask. He is always glad
to answer, and his answers will save you from guessing wrong.
```

### Other points to consider

- Should the partner description mention specific preferences?
  (e.g., dislikes overengineering, values simplicity, machine sympathy)
- Should it reference the skills/patterns line that follows the persona?
  ("your partner wrote these tools for you" is already implicit)
- Does the partner description belong inside `<persona>` or after it?
  Inside keeps it part of identity; after keeps persona purely about Daneel.
