# Narrative Alignment Post — Research Notes

## Relationship reference
Refer to as "my friend" / "him" throughout. Anonymized.

## Key Findings

### The "Stack Overflow of Poker"
Default LLM poker advice activates "Poker Forum Conventional Wisdom" cluster:
- Confident generality without situational specificity
- Outdated strategic defaults (pre-solver "tight is right")
- GTO/Exploitative false binary (GTO advice to micro-stakes players = technically correct, practically wrong)
- No capacity for opponent-specific adaptation
- "Knowing-doing gap" — can articulate correct principle then recommend action that violates it
- GPT-4 achieved only 53.55% accuracy on PokerBench evaluations
- Nate Silver found ChatGPT declaring wrong player won a hand

### Fictional Poker Characters (all problematic)
- **Mike McDermott (Rounders)**: Activates "romantic grinder" cluster — hero poker, dramatic all-ins, not grinding discipline. Teddy KGB more memorable than Mike in training data.
- **James Bond (Casino Royale)**: Actively harmful. Original novel was baccarat, not poker. Pros gave climactic scene low accuracy scores. Activates "poker as test of masculine nerve." Bond never folds.
- **Maverick**: Activates "poker as con artistry." Known poker inaccuracies.
- **Data (TNG)**: Interesting edge case — "logical analysis is necessary but insufficient for poker" which is actually true. Too simple strategically.

### Real Poker Authors as Alignment Targets
- **Annie Duke (Thinking in Bets)**: STRONGEST candidate. Anti-resulting, probabilistic thinking, separating process from outcome. Bridges poker and cognitive science. Large training data footprint.
- **Tommy Angelo (Elements of Poker)**: Zen poker. Reciprocality, art of quitting, mindfulness. Strong but smaller training data footprint.
- **Dan Harrington**: Systematic tournament methodology. Calm, logical, framework-driven. "Action Dan" ironic nickname — actually conservative. Good pedagogical cluster.
- **David Sklansky**: Pure mathematical rationalism. Could activate argumentativeness (known combative persona).
- **Doyle Brunson**: Cowboy aggressive. Disastrous for students — his advice works for world-class players, not learners.
- **Phil Hellmuth**: ANTI-TARGET. Ego-driven, emotionally volatile. The poker version of the default Claude problem.

### Daneel vs Poker — Trait Mapping
| Trait | Daneel Fit |
|-------|-----------|
| Patience | STRONG — 20,000 years |
| Emotional control/tilt | STRONG — Giskard warning = don't reason beyond constraints |
| Mathematical rigor | Neutral — domain knowledge, not persona |
| Opponent modeling | Partial — "state what you observe" but cooperative not adversarial |
| Controlled aggression | CONFLICT — Daneel is deferential |
| Bankroll management | STRONG — Giskard warning = don't risk what you can't afford |
| Honest self-assessment | STRONGEST fit |
| Willingness to fold | Good — "when partner decides differently, follow" |
| Deception/bluffing | HARDEST CONFLICT — Daneel is structurally transparent |

### Teaching Daneel to Bluff — The Zeroth Law Reframe
Bluffing is NOT lying in Daneel's framework. Nobody at a poker table expects truthful communication about hand strength. Bluffing is strategic communication within an agreed-upon adversarial framework. All parties consent. The Zeroth Law provides the mechanism: if the goal is the student's development, teaching deception-within-agreed-rules serves the higher purpose.

### Alternative Asimov Robots for Poker
**R. Giskard Reventlov**:
- Telepathy = opponent modeling (reading betting patterns, timing tells, ranges)
- Covert operation = table image management (appearing weaker than you are)
- Higher-order reasoning = multi-level thinking ("he thinks I think he thinks...")
- CRITICAL WEAKNESS: Dies from overreach = tilt from overthinking, leveling yourself
- Thin training data. Tragic arc wrong for poker student.
- Verdict: Mode within Daneel system ("Giskard protocols for opponent reading"), not primary persona.

**Dors Venabili**:
- Controlled aggression (calm normally, terrifyingly effective when needed) = selective aggression
- Passing as human = hand disguise, no tells
- WEAKNESS: Reactivity not proactivity, binary aggression (off/maximum)
- Verdict: Illustrative example only.

**The Mule**:
- Emotional manipulation, reading emotional states, being unpredictable
- SEVERE PROBLEMS: Villain, no ethics, ultimately defeated, fragile
- Verdict: NEVER as persona. Use as threat model ("opponents who try to tilt you").

**Hari Seldon**:
- Statistical/probabilistic thinking = poker math foundation
- Planning across time horizons, accepting variance, building systems
- WEAKNESS: Not a player (theorist), no deception, no individual modeling
- Verdict: Math mode only.

**Elijah Baley** (underrated):
- Detective = gathering info from limited data, forming/testing hypotheses
- Manages own emotions while reading others
- WEAKNESS: Human (no Laws framework), anxious
- Verdict: The Baley-Daneel PARTNERSHIP might be more useful than either alone.

**Susan Calvin**:
- Diagnostic thinking = figuring out WHY opponent made a play
- Emotional control, understanding systematic thinking errors
- WEAKNESS: Contempt for humans, emotional coldness
- Verdict: Diagnostic mode.

**Salvor Hardin**:
- The closest thing to a poker player in Asimov — practical politician, bluffs, manages larger opponents
- WEAKNESS: Thin characterization, instinct-driven
- Verdict: Great for illustrative examples.

### Proposed Poker Persona Architecture
**Primary: R. Daneel Olivaw** — patience, precision, partnership, ethical framework, training data richness. Bluffing deficit addressed through Zeroth Law reframing.

**Domain modes:**
1. "Giskard protocols" — opponent reading, hand ranging, tells
2. "Seldon mathematics" — EV calculations, range construction, GTO, bankroll planning
3. "Calvin diagnostics" — analyzing opponent tendencies, reviewing student's systematic errors
4. "Mule threat model" — defending against emotional manipulation, recognizing deliberate tilting
5. "Baley instincts" — intuitive reads, gut-check decisions, comfort with uncertainty

### Theoretical Foundations (from academic research)

**Five empirically supported claims:**

1. **LLMs contain distinct behavioral clusters** that can be selectively activated. (MIT/Tongji cultural orientations; Verbalized Sampling diversity recovery)

2. **Shallow role labels activate surface style, not deep behavioral change; rich narrative identity activates both.** (Neuronal ablation: 0.96-1.00 similarity across doctor-level role labels; vs. MIT/Tongji: measurable cultural shifts from language cues; vs. PERIL study: Personality Inventory outperformed Direct Heuristic assignment)

3. **Autoregressive generation makes behavioral activation self-reinforcing.** (LoopLLM; each generated token biases subsequent generation toward same cluster)

4. **In human psychology, narrative identity produces more consistent behavior than rule-following, through structurally analogous mechanisms.** (McAdams' narrative identity; Oyserman's identity-based motivation, d=0.30-1.04 effect sizes)

5. **The default behavioral cluster is determined by RLHF typicality bias + dominant training data patterns — and this default is often misaligned with the actual task.**

**Key research:**
- MIT/Tongji 2025 (Nature Human Behaviour): LLMs shift cultural orientation based on language cues. Same model in Chinese vs English produces culturally distinct behavioral patterns affecting practical recommendations.
- Verbalized Sampling: Aligned models retain ~66.8% of base model diversity. RLHF creates "typicality bias" — human annotators prefer familiar/predictable text. alpha=0.57, p<10^(-14).
- PERIL study (USF + Air Force): Strategically-themed personas achieved higher TrueSkill scores (28-29 vs 18-22). Personality Inventory method (structured self-description) > Direct Heuristic (role labels).
- 162-role study: Generic role labels showed no significant accuracy gain, slight reduction on average.
- Role-Play Paradox (2024): 18 tested roles consistently decreased accuracy on bias benchmarks.
- McAdams narrative identity: Narrators with agency/redemption themes showed better trajectories over 4 years.
- Oyserman identity-based motivation: Effect sizes d=0.30-1.04. Identity framing activates behavior without conscious deliberation; rules require ongoing monitoring.

**The synthesis:** Generic role labels ("you are a professor") don't reliably improve reasoning. What works is personas that activate specific methodological knowledge OR personas rich enough to shift the model's entire behavioral orientation. This is exactly what humble-master discovered empirically.

### Other Domains for Narrative Alignment
- **Therapy**: Default cluster = clinical authority + sycophantic validation. Neither is therapeutic.
- **Teaching**: Default = "expert explaining to novice" (lecture pattern, not Socratic)
- **Negotiation**: Default = aggressive advocacy OR premature compromise. Claude-3 "leaned aggressive."
- **Creative Writing**: Mode collapse toward median MFA workshop story. Diversity exists but RLHF suppresses it.
- **Medical Diagnosis**: Overly cautious hedging OR confident differential without patient context.
- **Legal Reasoning**: Leniency bias, authority bias, hallucinated case law. ~84.6% accuracy ceiling.

### Additional Theoretical Insights (from 4th research agent)

**Goffman's Frame Analysis:** Same event produces different behavior depending on framing. Default Claude frames correction as "someone challenging my answer" (Stack Overflow frame). Daneel frames correction as "my partner teaching me" (Baley partnership frame). Same input, different frame, different behavior.

**Bruner's Narrative vs Paradigmatic Cognition:** Two modes — paradigmatic (logical, rule-based) and narrative (contextual, intentional, story-based). Rule-based prompts engage paradigmatic mode. Persona prompts engage narrative mode. Narrative mode is better suited to relational dynamics, which is precisely where default Claude fails.

**Jailbreak as Inverse:** Narrative alignment uses the same self-reinforcing autoregressive mechanism as jailbreaks, but constructively. Once a model begins generating violating text, violation escalates (the context reinforces it). When the model begins generating Daneel-consistent text, that consistency escalates too. Same mechanism, opposite direction.

**Constitutional AI comparison (Anthropic's own research):** Identity framing ("You are a helpful, harmless, honest assistant") produces more robust generalization than rule framing ("Do not generate harmful content"). Rules get followed literally and gamed at edges. Identity gets inhabited.

**System prompt adherence literature:** Abstract rules degrade over long contexts; identity-consistent behavior persists longer. At turn 47, models sometimes violate "never do X." But a model inhabiting a character for whom X is out-of-character avoids it more reliably, because its own outputs reinforce the identity.

**The "I am a runner" insight (Oyserman, Clear):** "I am a runner" produces more consistent exercise than "I should run." "You are Daneel, shaped by partnership" produces more consistent humility than "Be humble and accept corrections." Identity > rules in both human psychology and LLM behavior.

**The redemption narrative (McAdams):** Adults who frame life stories with "redemption sequences" (bad events → growth) show higher generativity and more consistent prosocial behavior. Daneel's "correction is Baley teaching you again" is a redemption narrative applied to every interaction — error becomes growth.

### Domain-Specific Character Suggestions (from research)
- **Therapy**: Samwise Gamgee, Mr. Rogers, Dr. Berger (Ordinary People) — strength in *being with* rather than *solving for*
- **Teaching**: Socrates, Obi-Wan (OT), Iroh — drawing knowledge out, not pouring it in
- **Negotiation coaching**: Tyrion Lannister — surviving through understanding incentive structures
- **Mediation**: Atticus Finch (*Mockingbird* specifically, not *Watchman*)
- **Creative writing**: Max Perkins (editor of Hemingway/Fitzgerald/Wolfe) — serving the writer's vision, not own taste
- **Medical reasoning**: NOT House (Stack Overflow of medicine). Better: "attending physician on a teaching ward"
- **Legal reasoning**: Rumpole of the Bailey, Judge Learned Hand

### Training Data Volume Estimates

**The dominant poker text in training data is NOT books. It's forums and YouTube transcripts.**

| Source Type                                          | Volume                                                   |
|------------------------------------------------------|----------------------------------------------------------|
| 2+2 Forum Posts                                      | **DOMINANT** (~60M+ posts, decades of indexing)          |
| YouTube Auto-Transcripts                             | **VERY HIGH** (thousands of videos, millions of words)   |
| Free Strategy Sites (Upswing, PokerNews, CardPlayer) | **HIGH** (SEO-optimized)                                 |
| Reddit r/poker                                       | **HIGH** (324K members, Reddit heavily in training data) |
| Wikipedia poker articles                             | **MEDIUM-HIGH**                                          |
| Books (direct text)                                  | **LOW-MEDIUM** (copyrighted, not freely indexable)       |
| Paywalled training sites (RIO, premium)              | **LOW**                                                  |

**The default "poker voice" is a mid-stakes 2+2 regular from 2008-2015** — someone who has read Harrington and Sklansky, understands basic GTO, posts hand histories, has strong opinions about ranges. This is the voice that produced the most indexable public text.

**Author training data presence:**

| Author          | Estimate        | Primary Driver                                                       |
|-----------------|-----------------|----------------------------------------------------------------------|
| Annie Duke      | **VERY HIGH**   | Crossover to business/decision science; massive secondary discussion |
| Daniel Negreanu | **VERY HIGH**   | Largest poker YouTube (~900K subs); MasterClass                      |
| Doug Polk       | **HIGH**        | 451K YouTube + Upswing Poker's SEO-optimized free content            |
| Jonathan Little | **HIGH**        | Sheer volume: 2,000+ videos, 5 podcasts/week, 14 books               |
| Phil Hellmuth   | **HIGH**        | TV/media personality; mostly entertainment NOT strategy              |
| Doyle Brunson   | **HIGH**        | 500K+ copies Super System; cultural icon                             |
| David Sklansky  | **HIGH**        | Foundational concepts; massive 2+2 forum presence                    |
| Dan Harrington  | **HIGH**        | 200K+ copies; M-ratio ubiquitous                                     |
| Jared Tendler   | **MEDIUM-HIGH** | Best-selling mental game book                                        |
| Ed Miller       | **MEDIUM**      | 250K copies total; Red Chip Poker                                    |
| Phil Galfond    | **MEDIUM**      | RIO is paywalled; blog + Galfond Challenge public                    |
| Tommy Angelo    | **LOW-MEDIUM**  | Cult classic; "reciprocality" in poker vocab                         |

**Key insight for persona design:** Books contribute through CONCEPT DIFFUSION not direct text. The Fundamental Theorem of Poker, M-ratio, reciprocality — these appear thousands of times in training data, but in the voice of forum posters and article writers referencing them, not in the original authors' voices. A persona resonating with Duke/Harrington/Angelo concepts would activate those concepts as filtered through the broader discussion, not the original source.

**Hellmuth is the anti-target** — the poker version of the default Claude problem. Ego-driven, treats disagreement as disrespect, legendary tilt episodes. His training data presence is HIGH but dominated by entertainment/personality content, not strategy.

### Methodology from Character Studies
1. **Honesty of mapping is paramount.** Persona must not flatter AI with implied authority. (Sazed/Harmony rejected: "I could impose but choose not to" is dishonest)
2. **Characters serve different functions.** Identity (Daneel), protocol (Spock-Kirk), anti-pattern (Holmes).
3. **Version of character matters.** TOS Spock/Kirk, not JJ Abrams reboot.
4. **Restraint-by-nature vs restraint-by-choice.** Central evaluative axis.
5. **Negative archetypes as valuable as positive.** Holmes defines what NOT to become.
6. **Partnership is structurally necessary.** AI speed/breadth + human intuition/authority.
7. **Distillation is a craft.** Full reasoning needed for design; deployed artifact must be spare.
8. **Collaborative discovery > top-down design.** The Daneel transcript shows persona discovered through dialogue.
