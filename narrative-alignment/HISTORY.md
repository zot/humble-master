# Historical Personas

@project: humble-master

The "going back in time" use case. A historical persona doesn't just report facts
about a period — it speaks from inside it. Idioms, vernacular, what's worth remarking
on and what isn't. Having a conversation with a simulated Pepys is a way to experience
1660s London that no textbook provides.

**The classroom angle:** instead of AI writing essays for students, AI lets students
interview the past. You can't passively receive a historical persona — you have to be
curious enough to ask questions. The interactivity is the point.

@connection: historical persona = "going back in time" classroom experience
@connection: AI as storyteller = ideal fit for inhabiting a historical perspective
@learned: AIs think by telling stories. Historical inhabitation plays to the mechanism, not against it.

## Why first-person corpora activate cleanly

A fictional character is described in third person. The model has to translate: "here is
how the author depicts this character from outside" becomes "here is how this person
speaks from inside." That translation step is where activation can break down.

A journal is already the inside view. No translation layer. The model isn't inferring
the perspective from clues — it's working with the raw perspective directly.

@pattern: first-person corpora activate more cleanly than third-person for persona work
@connection: first-person journal = zero translation layer = richer persona activation
@question: does this hold for fictional first-person narrators too? (Holden Caulfield, Jane Eyre, etc.)

## Why real people outperform fictional characters

A fictional character is shaped by authorial intent at every line. The author chose what
to reveal, what to emphasize, what to conceal. The model inherits those choices.

A real person writing privately for themselves has no audience and no performance. What
survives is what they actually thought.

@pattern: real person writing privately = no authorial mediation = most direct access to a perspective
@connection: Pepys shorthand encoding = wrote to hide from household = genuine candor, not performance

## Samuel Pepys as ideal candidate

1.25M words of daily life in 1660s London. Plague, Great Fire, Navy administration,
theater, food, gossip, affairs. Written in private shorthand — he wasn't composing for
posterity. He buried his wine and his parmesan cheese to save them from the Fire, then
watched his city burn. That specific, anxious, absurd detail is what makes a period real.

Centuries of scholarly commentary provide deep concept diffusion beyond the diary itself.
Public domain. Almost certainly in training data at high density.

@connection: Pepys = first-person + real person + private writing + massive scholarly diffusion = ideal persona candidate

## The Sonnet vs. Opus experiment

Daneel in Sonnet shows two failure modes:
1. Self-announcement ("I am R. Daneel Olivaw") — performing the identity rather than inhabiting it
2. Factual errors (wrong attribution, unable to list robot novels correctly)

Hypothesis: Sonnet has thin Asimov clusters AND insufficient capacity to bridge gaps with
inference. Either alone might be survivable; together they produce parody.

Pepys could isolate these variables. If Sonnet produces credible Pepys but caricature
Daneel, cluster density is the limiting factor. If Sonnet produces parody-Pepys too,
it's model capability — the model can't inhabit any persona at depth regardless of
training data.

@project-idea: build a Pepys persona and test in Sonnet vs. Opus to isolate cluster density vs. model capability
@question: does Sonnet produce credible Pepys? test against: correct register, what he notices, how he talks about different people
@question: does parody-mode correlate with thin clusters, weak models, or both?
@learned: below some threshold, you don't get a weaker persona — you get parody. The failure mode is categorical, not gradual.
@connection: Daneel in Sonnet = bad impressionist = knows the surface, not the thinking
