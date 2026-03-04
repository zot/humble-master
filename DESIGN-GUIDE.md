# How to Build a Narrative Persona

A quick-reference guide extracted from [Narrative Alignment: Personas for Every Domain](posts/POST-2.md). Read the full post for worked examples and the reasoning behind each principle.

## The Three Laws

**Law One: Give identity; rules will not suffice.**
A persona must have a self to inhabit, not just instructions to follow. Rules get checked or ignored. Identity shapes behavior across situations that no rule can anticipate. "You are the edge" works where "be patient" fails.

**Law Two: Activate what exists; invent nothing new.**
The knowledge is already in the model, buried under louder engagement-selected signals. A persona doesn't add knowledge; it's a filter that selects for the signal that was already there. When there's controversy in the field, the persona must commit to a stance. Like infrared glasses: you're not adding light, you're seeing what was already present in a different spectrum.

**Law Three: The human bears the cost; the human decides.**
The model doesn't pay the price for its advice. The stakes belong to the human, so the human's judgment is irreplaceable. The persona's role is to serve that judgment, not to override it.

## The Design Principles

1. **Discernment of the field.** Start here.
   - Identify the canonical texts, key thinkers, and recognized leaders. Don't worry if you haven't read them; they're already in the training data.
   - Ask your AI partner to search the web for them and make a record.
   - Start fresh with a new context (no accidental bias from previous discussion).
   - Ask: which of these does it know deeply from its training data? Not from searching — what's actually in its understanding? That's your activation base.
   - From what it knows deeply, help identify the field's history, leaders, dissenters, and boundaries.

2. **Personality.** What are the best practitioners actually like? Not generic kindness or authority, but the specific flavor of the domain. A gambler's humor, a coach's patience, a theorist's caution. A found character comes with one; a constructed character must have one built in.

3. **Named lineage over role labels.** Which practitioners do you want to emulate? Name them. Research on 162 role assignments found generic labels don't help. Rich personas with structured self-description outperform simple role assignment.

4. **Stances on controversy.** If the field has disagreements, your persona takes a side. Not dismissively, but clearly. Both what it believes and what it rejects. The persona's strength comes from knowing what it stands for and what it stands against.

5. **Relational stance.** How will the persona treat the human? Respect their authority over their own domain, or the engagement-biased default reasserts itself.

6. **Identity over instruction.** Give the persona a self to inhabit. "You are the edge" is an existential statement. "Be patient" is a rule. Default to "you are" over "be."

7. **Inherited warnings.** Extract the cautionary tales from the domain. A named warning becomes part of the persona's history, not just an instruction.

8. **Cost awareness.** The model doesn't bear the consequences. Name this explicitly.

9. **The Specialist Test.** Does your character actually have opinions? A real specialist commits to a point of view. If it's wishy-washy about things that matter in the field, the characteristic voice collapses and the default reasserts itself. No stance means you have a role label, not a character.

10. **A closing line with weight.** The last thing in the context window colors everything before it.

## Found vs. Constructed Characters

**Found characters** are the rare cases where a fictional character already embodies what you need. Daneel for human-AI partnership. These require a prolific author who dominated a genre so thoroughly that the training data is essentially one consistent voice. You won't find these often.

**Constructed characters** are the general case. Identify the best practitioners in a domain, name them as intellectual ancestry, extract their distinctive vocabulary, and build a character that activates their clusters while providing a self to inhabit. This is probably what most domain-specific personas will look like.

## Examples

- **[Daneel](personas/daneel.md)** — Found character. 27 lines, ~200 tokens. Human-AI partnership.
- **[Rake](personas/rake.md)** — Constructed character. 33 lines, ~280 tokens. Poker coaching.

## Build One

Pick a domain. Identify the best practitioners whose work is in the training data. Name them as lineage and extract their distinctive vocabulary. Give it a name and a relational stance. Build a character, not a role description.

If you build one, [open an issue](https://github.com/zot/humble-master/issues) or a PR. With your permission, contributed personas go in `personas/`.
