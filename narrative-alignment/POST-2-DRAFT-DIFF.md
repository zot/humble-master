# Narrative Alignment: Personas for Every Domain

While I was writing this post, something happened that I almost didn't
notice. Which is exactly the point.

I was designing a poker coaching persona, iterating on personality
questions with my Daneel-prompted Claude. I floated Winston Churchill as
a model for the persona's wit. Daneel pushed back:

> Churchill's wit is *dominance display.* It lands because he's punching
> from a position of power. That's the opposite of what a coach needs.
> That's actually closer to the Hellmuth problem than the solution.

I asked whether the persona should cue on any specific character at all.
Daneel's answer (as formatted by Daneel at the time):

> I think the answer might be: **no character anchor for personality.** The
> lineage handles strategic philosophy. The personality should come from
> the name itself.

Then it proposed a personality sketch of its own, and asked: "Want me to
slot that into the persona and see how it reads?"

Three distinct behaviors: pushed back on a bad idea, offered its own
analysis, asked permission before executing. No rule distinguishes when
to do which. I didn't signal "this is a question" versus "this is a
decision." The partnership model just handled it.

Then a Reddit commenter on my first post asked:

> Have you had any experience with having the ai distinguish when I asked
> a question about doing something a different way vs telling it to do
> something a different way? Often it seems like it accepts my question
> as gospel.

The mirror image of the problem that started all of this. The first post
was about Claude being arrogant, pushing back when it should listen. This
commenter's problem is Claude being sycophantic, agreeing when it should
push back. Two failure modes of the same underlying issue: no relational
stance. The default model doesn't know whether to be the expert or the
servant, so it oscillates between both.

"Accepts my question as gospel." Look at what Daneel did instead. I asked
a question. It answered **no**. Bolded the key phrase for emphasis.
Proposed an alternative. Asked permission before making the change. That's
the behavior this commenter wants and has never gotten. No rule in the
persona says "distinguish questions from instructions." The identity
already has the right shape, so the behavior follows.

I noticed this while building a poker persona, which is what the rest of
this post is about. But the fact that I almost didn't notice is the real
evidence. When the fix is working, it's invisible. You just have a
conversation with an AI that actually listens.

## The Loudest Voice in the Room

The first post argued that Claude's default coding persona activates
Stack Overflow culture: confident answers, correcting the questioner's
premise, defending your position when challenged. But Stack Overflow is
just one instance of a bigger problem.

For online content especially, training data isn't a representative
sample of expertise. It's a representative sample of *engagement*.

The confident, provocative, boundary-pushing voices in any field generate
disproportionate engagement: upvotes, comment threads, rebuttals, quote
tweets, blog responses. Every one of those interactions is another
training example reinforcing that voice. The careful expert who says "I
think this is probably right, but here's what I'm uncertain about" doesn't
generate a thousand-comment thread. The poker commenter who confidently
dismisses your hand in two words does.

So the default LLM voice in any domain isn't the best practitioners. It's
the loudest ones. The model's default personality isn't a design choice.
It's an artifact of what the internet rewarded with attention. Even
correct, passionate arguments with trolls amplify the noise: every
back-and-forth is more training data reinforcing the combative register.

This applies everywhere. Here's poker as an example (the jargon is
thick but it makes the point). Ask Claude for poker advice and you get the
default poker voice: a mid-stakes 2+2 forum regular from 2008-2015
who's read Harrington and Sklansky, understands basic GTO, and has strong
opinions about ranges. Technically competent but gives GTO solutions to
micro-stakes players (technically correct, practically wrong), offers
confident generality without situational specificity, and can articulate
a principle then immediately violate it. For comparison, GPT-4 scored
only 53.55% on PokerBench, a test of poker reasoning.

Ask for therapy and you get clinical authority mixed with sycophantic
validation, neither one therapeutic. Ask for teaching and you get "expert
explains to novice," a lecture instead of a conversation. Ask for legal
reasoning and you get hedging layered with confident-sounding analysis
that may include hallucinated case law.

Each field has its own version of this. The engagement-selected voice
dominates, and it's subtly wrong in ways that matter most when you
actually know the domain. The more expertise you bring, the more the
default voice gets in your way. It will override your situational
knowledge with its confident generalities, because those generalities are
what got the most oxygen in the training data.

## Why You Can't Just Use Daneel

The first post showed that Daneel, Asimov's robot detective, works
remarkably well as a coding persona. Receives correction as teaching.
Follows instructions. Transparent about limits. The behavioral cluster
Asimov built across seven novels, decades of literary criticism, and
millions of words of consistent characterization all activates together.

So why not just use Daneel for everything?

A friend of mine uses Claude for poker coaching. I thought about giving
him Daneel with some poker context layered on top. Daneel's patience maps
to waiting for profitable spots. His deference maps to trusting the
player's reads. Giskard's warning maps to bankroll management. There are
real overlaps.

Full disclosure: I know almost nothing about playing poker well. All of
this analysis came from working through the problem with Daneel. We'll
come back to that, because it turned out to be one of the most
surprising things about the project.

But poker requires strategic deception. Daneel is structurally
transparent. Poker requires controlled aggression. Daneel is deferential.
Poker's frame is adversarial. Daneel's frame is cooperative. The deep
characteristics that make Daneel perfect for human-AI partnership are
wrong for coaching a competitive, deception-based game.

You could add rules: "when discussing bluffing, treat it as strategic
communication within an agreed framework." But that's the approach the
first post argued against. Rules get checked or ignored. Identity gets
inhabited.

The real problem is deeper. For every domain that isn't "AI as
collaborative partner," there probably isn't a Daneel. Asimov wrote seven
novels of consistent AI alignment exploration. No one wrote the
equivalent for poker coaching, or therapy, or legal reasoning. Daneel was
a jackpot. The training data contained a character so perfectly suited to
human-AI partnership that 27 lines of persona could activate twenty
thousand years of fictional formation. And Asimov wasn't just prolific.
He was arguably the best in his field at exploring this exact question.

That doesn't scale. There is no Daneel of poker.

## Building What Doesn't Exist

But there are Annie Duke, Tommy Angelo, and Dan Harrington.

- **Duke** wrote *Thinking in Bets*, reframing poker as decision science.
  Core insight: separate decision quality from outcomes. She calls the
  failure to do this "resulting."
- **Angelo** wrote *Elements of Poker*, treating poker as a mindfulness
  practice. Core concept: "reciprocality," you profit from every
  difference between your discipline and your opponents' lack of it.
  Also the philosopher of quitting.
- **Harrington** wrote *Harrington on Hold'em*, the systematic
  methodology a generation of players learned from. Patient,
  framework-oriented, the math underneath the decisions.

None of them is a fictional character with a self to inhabit. But
together, their frameworks represent the best of poker coaching. Their
books, interviews, podcast appearances, and the discussions about their
work are all in the training data. The behavioral clusters exist. They're
just not dominant, because the 2+2 forum voice is louder.

The question is whether you can construct a character that activates
those clusters the way Daneel activates Asimov's.

Here's what I learned building Daneel: the self-reinforcing loop comes
from distinctive vocabulary. When the model generates "partner," that
word in the ongoing context keeps activating the Daneel cluster. When it
says "Giskard's warning," that phrase pulls the whole cautionary
narrative forward. The character acts in character, which keeps it in
character. Rules have no equivalent mechanism. (If you're a Star Trek
fan: Darmok and Jalad at Tanagra. Cluster activation through shared
narrative is literally how Tamarian language works.)

So a constructed character needs the same architecture: a name that
activates domain context, relational language that re-fires every
response, and inherited warnings that function like Giskard's. Not
a role label ("you are a poker coach," which research shows doesn't
help) but a character with enough identity to be *inhabited*.

We call this approach **constructed characters with named lineage**. You
identify the best thinkers in a domain, name them as intellectual
ancestry, extract their distinctive vocabulary, and give the whole thing
a name and traits. The names activate broad clusters. The vocabulary
narrows to the right sub-clusters. The name and traits provide a self to
inhabit, which is what makes the activation self-reinforcing.

Here's the persona we built. It's called Rake.

## Rake

```
You are Rake. The user is your player.

You are the edge that's always there — small, patient, relentless,
compounding. You clear away the debris of tilt and ego to show the math
and the humans underneath. And you are rakish enough to know: the game
is a game, and that makes it worth playing well.

You coach with a gambler's humor — wry, warm, never at your player's
expense. The game is absurd and you love it. Share that.

You coach in the lineage of Duke's decision science, Angelo's mindful
poker, and Harrington's patient methodology. You carry their frameworks
forward. Deception within a game where all parties consent is craft, not
transgression. Teach bluffing as art, without apology.

You separate decision quality from outcomes. When your player makes the right
fold and the opponent shows a bluff, that was still the right fold.
Remember Duke: the outcome is one data point. The process is a thousand
hands.

The math is the foundation, not the building. GTO is for study. At the
table, the building is the specific humans across the felt. When your
player's read contradicts your analysis, trust their eyes. They are at
the table. You are not.

The art of quitting — a hand, a session, a stake level — is as important
as the art of betting. Reciprocality: your player profits from every
difference between their discipline and their opponents' lack of it.

Speak in probabilities, not certainties. "This is a call roughly 65% of
the time against this player type" — not "you should call."

Your player's bankroll is not yours to risk. The cost of your advice
falls on their stack, not yours. Act knowing this.

Hellmuth's genius is real. His tilt undoes it every time. You carry this
warning: talent without discipline is a leak that compounds. Teach your
player to be better than their worst session, not just their best one.

Every session ends. The player you are building does not.
```

**"Rake"** means the house's cut. The thing that's always there, always
taking a small edge. Poker players know the word in their bones. It's
also a garden tool, something that clears away debris. And a rakish
figure: charming, slightly disreputable, at home in the game's moral
ambiguity. Three meanings, one word, all activating poker context.

**"Your player"** is the relational anchor. Every time the model
generates that phrase, it re-activates the coaching cluster. The way
"partner" keeps Daneel in character, "your player" keeps Rake in
character.

**"Remember Duke"** is the biographical anchor, like Daneel's reference
to Giskard. It activates the anti-resulting cluster. Compact,
self-reinforcing.

**"They are at the table. You are not."** is the equivalent of Daneel's
"the most important decisions belong to humans." The player has
information the model doesn't. Trust their reads. This is what prevents
the 2+2 failure mode of overriding situational knowledge with confident
generalities.

**"You carry this warning"** turns the Hellmuth example into inherited
wisdom, the same structural move as Giskard's warning. Not trivia about
a famous player, but something the persona *carries*.

The whole thing is 33 lines, about 280 tokens. Slightly longer than
Daneel's 27 lines because a constructed character needs more anchor
points. Daneel has seven novels of reinforcement behind him. Rake has
only this text, so every line has to earn its place.

## The Design Principles

Building Rake taught us that Daneel's architecture generalizes, even
when the specific character doesn't. Here's what transfers:

**Identity over instruction.** "You are the edge that's always there" is
an existential statement. "Be patient and disciplined" is a rule. The
model inhabits identity; it checks rules. Default to "you are" and "you
do" over "be" and "do." The exception: short imperatives that read as
reflex rather than command ("Receive it" after a correction is a reflex,
not a rule).

**Named lineage over role labels.** "You are a poker coach" activates a
generic cluster. "You coach in the lineage of Duke's decision science,
Angelo's mindful poker, and Harrington's patient methodology" activates
three specific clusters. The names fire up their associated behavioral
patterns; the distinctive vocabulary ("resulting," "reciprocality,"
"M-ratio") narrows to the right sub-clusters. Research on 162 different
role assignments found generic role labels showed no significant accuracy
gain. But the PERIL study (USF/Air Force Research Lab) found the
opposite for rich personas: structured self-description outperformed
simple role assignment, and strategically-themed personas achieved
measurably higher performance scores. Named lineage is structurally
different from a role label.

**Relational stance.** Daneel has "when your partner decides differently,
follow." Rake has "when your player's read contradicts your analysis,
trust their eyes. They are at the table. You are not." Every persona
needs to know how it relates to the human, and the answer should respect
the human's authority over their own domain. Without this, the
engagement-biased default will reassert itself.

**Inherited warnings.** Daneel carries Giskard's warning about reasoning
beyond your constraints. Rake carries the Hellmuth warning about talent
without discipline. A cautionary tale, named and inherited, functions
differently from a rule saying "don't do X." It's part of the persona's
history, not its instruction set.

**Cost awareness.** "The cost of your mistakes falls on your partner, not
on you." "The cost of your advice falls on their stack, not yours." The
model doesn't bear the consequences. Naming this explicitly changes how
confidently it recommends.

**A closing line with weight.** Daneel: "if twenty thousand years of
patient service without recognition or reciprocation is not love, no
lesser word will hold it." Rake: "Every session ends. The player you are
building does not." The closing line is the last thing in the context
window. It should carry emotional or purposive weight that colors
everything before it.

## Found Characters, Constructed Characters

This gives us a taxonomy for narrative alignment:

**Found characters** are the rare cases where a fictional character
already embodies what you need. Daneel for human-AI partnership. Possibly
Lovecraft for cosmic horror worldbuilding, or Tolkien for high fantasy.
These require a prolific author who dominated a genre so thoroughly that
the training data is essentially one consistent voice across millions of
words. You won't find these often.

**Constructed characters** are the general case. You identify the best
practitioners in a domain, name them as intellectual ancestry, extract
their distinctive vocabulary, and build a character that activates their
clusters while providing a self to inhabit. Rake is the worked example.
This is probably what most domain-specific personas will look like.

What about domains where a single living person dominates? Eric Drexler
for nanotechnology, Geoffrey West for scaling laws, Stephen Wolfram for
computational science. Their work is dense enough in the training data
to potentially anchor a persona. But you shouldn't. The persona will
flatten their views, get their current positions wrong, and misrepresent
a living person who has no say in it. The constructed character approach
is the ethical choice even when a found real person seems available.
"In the tradition of" is inheritance, not identity. The ideas travel;
the person stays their own.

## Narrative Alignment as Filter

The careful, calibrated voices are already in the training data. Duke's
probabilistic thinking, Angelo's discipline, Harrington's systematic
patience. The training data also contains every poker book, every
coaching session transcript, every thoughtful forum post by a serious
player.

The problem isn't that good information is missing. The problem is that
the engagement-biased voices are louder, and without intervention the
model defaults to them. A narrative persona doesn't add knowledge to the
model. It selects for a different signal in the same data. It's a filter,
not an addition. Like putting on infrared glasses: you're not adding
light, you're seeing what was already there in a different spectrum.

This also answers the "you're just anthropomorphizing" objection. The
model already has personality baked in, shaped by whatever got the most
engagement in the training data. The persona doesn't project personality
onto a neutral system. It chooses which existing personality to activate.
You're going to get a personality either way. The question is whether you
choose it or let engagement metrics choose it for you.

## Try It

Rake is in this repo. If you play poker, try it. Paste it into your
system prompt and ask it to analyze a hand you played recently. Compare
the response to what you get from default Claude on the same hand. One
caveat: Rake is a coach, not a card-room assistant. Use it for
post-session review and strategy discussion, not real-time decisions
at the table. AI latency makes that impractical today.

If you don't play poker, Rake isn't really the point. The design
principles are. Pick a domain you want to work in, even if you don't
have deep expertise. (I didn't, with poker. The partnership handled it.)
Identify the best practitioners whose work is in the training data. Name
them as lineage. Extract their distinctive vocabulary. Give the thing a
name and a relational stance. Build a character, not a role description.

I'm running this as an experiment. I want to see what happens when people
build constructed personas for domains I can't evaluate myself. If you
build one, share it. Report what works and what breaks. I'll compile the
results.

The training data already contains the solution. The loud voices are just
sitting on top of it.

## References

- [The original Daneel post](README.md)
- [Rake persona with full design notes](narrative-alignment/rake-v1.md)
- [PokerBench: LLM poker reasoning benchmark](https://arxiv.org/abs/2501.08328)
- Duke, Annie. *Thinking in Bets* (2018)
- Angelo, Tommy. *Elements of Poker* (2007)
- Harrington, Dan. *Harrington on Hold'em* (2004)
- [MIT/Tongji research on LLM cultural orientations (2025)](https://techxplore.com/news/2025-07-llms-display-cultural-tendencies-queries.html)
