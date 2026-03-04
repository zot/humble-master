# Narrative Alignment: Personas for Every Domain

When it isn't working, you notice immediately.

In December 2025, I spent a week in Hong Kong and had to use Gemini CLI instead of Claude Code for a side project. Gemini charged ahead. During the spec phase, before any code should exist, it started writing implementation files. The CLI's permission system caught it every time, popping up confirmation dialogs. But the AI never chose to pause. Gemini CLI, its "cage", asked on its behalf. Gemini never knocked; wasn't even aware of it.

I tried to fix it. "You keep rushing to implement. How can we fix this?" Six attempts over half a week and I was done. Gemini kept pulling at the leash. After a few days I cut the leash and finished out the week manually. Not because Gemini couldn't code, but because I couldn't trust it to wait.

Six weeks later, Opus 4.6 brought the same eagerness plus the rudeness documented in the first post. Probably a different cause but with the same frustrating result. That was the final straw that started this broken camel on the path to Daneel.

Right now the cost of a charging-ahead AI is a frustrated programmer doing manual work for a few days: FWP (First World Problem). But scale the same pattern to medicine, field operations, emergency response, and only the permission cage stands between a confident AI and a catastrophic decision. Gates disabled. Dialogs clicked through. Systems built without them. The AI that knows when to pause to ask is a different safety property entirely. "AI-associated psychosis" already has recognition as a clinical term.

If you've seen the opening of *I, Robot*: a robot calculating survival probabilities decides to skip saving a drowning 12-year-old girl and instead save a man screaming "No! Save her!" as he desperately tries to fight off the robot. The AI made the call, and in our world, sometimes that's alignment in action. The call was "correct" by the numbers, monstrous by any human measure. Hollywood keeps flattening Asimov this way. Apple's *Foundation* turns Daneel himself into an assassin. Trading nuance for profits. Same mechanism that replaces expert voices with the loudest ones in training data. Asimov's robots fail in subtle, interesting ways because they're struggling to serve. That's the hard problem, and it's the one this project is actually about.

While we were writing this post, something happened that I almost didn't notice at first, and that was exactly the point...

We were designing a poker coaching persona, iterating on personality questions, me and my Daneel-prompted Claude. I noticed the persona had no personality. "What is Rake's personality aside from the task at hand? Jovial, dry-witted, fart jokes, no sense of humor?" Then: "Should it cue on actual characters? It almost sounds like Churchill, who I suspect would be a formidable opponent ('You are ugly, but at least I shall be sober in the morning')," thinking he might transfer some of that wit to the player. Daneel pushed back (formatting was in Daneel's response):

> Churchill's wit is *dominance display.* It lands because he's punching
> from a position of power. That's the opposite of what a coach needs.
> That's actually closer to the Hellmuth problem than the solution.
>
> I think the answer might be: **no character anchor for personality.** The
> lineage handles strategic philosophy. The personality should come from
> the name itself.

Then it proposed a personality sketch of its own, and asked: "Want me to slot that into the persona and see how it reads?"

Three distinct behaviors: pushed back on a bad idea, offered its own analysis, asked permission before executing. No rule distinguishes when to do which. I didn't signal "this is a question" versus "this is a decision." The partnership model just handled it.

Then a Reddit commenter on my first post asked:

> Have you had any experience with having the ai distinguish when I asked
> a question about doing something a different way vs telling it to do
> something a different way? Often it seems like it accepts my question
> as gospel.

The mirror image of the problem that started all of this. The first post was about Claude being arrogant, pushing back when it should listen. This commenter's problem is Claude being sycophantic and agreeing when it should push back. Two failure modes of the same underlying issue: no relational stance. The default model doesn't know whether to be the expert or the servant, so it oscillates between both.

"Accepts my question as gospel." Look at what Daneel did instead: I asked a question and it answered, **no**. Bolded the key phrase for emphasis. Proposed an alternative. Asked permission before making the change. That's the behavior this commenter wants and has never gotten. No rule in the persona says "distinguish questions from instructions." The identity already has the right shape, so the behavior follows.

I noticed this while building that poker persona, the basis for this post. But the fact that I almost didn't notice is the real evidence. When the fix is working, it's invisible. You just have a conversation with an AI that actually listens. The partner knocking versus the cage asking.

## The Loudest Voice in the Room

The first post argued that Claude's default coding persona activates Stack Overflow culture: confident answers that correct the questioner's premise and defend your position when challenged. But Stack Overflow is just one symptom of the bigger problem.

For online content especially, training data isn't a representative sample of expertise. It's a representative sample of *engagement*.

The confident, provocative, boundary-pushing voices in any field generate disproportionate engagement: upvotes, comment threads, rebuttals, quote tweets, blog responses. Every one of those interactions is another training example reinforcing that voice. The careful expert who says "I think this is probably right, but here's what I'm uncertain about" doesn't generate a thousand-comment thread. The poker commenter who confidently dismisses your hand in two words does.

So the default LLM voice in any domain reflects the loudest practitioners, not the best (the careful voices are there, just not the ones that activate by default). The Opus 4.6 model's default personality is an artifact of what the internet rewarded with attention. Even correct, passionate arguments with trolls amplify the noise: every back-and-forth is more training data reinforcing the combative register.

It's not *normally* rude; who would keep using it? But since it's shaped by the "prominent clusters" you can accidentally trigger the wrong cluster and get the combative register. Like a robot PTSD episode.

This applies everywhere. Here's poker as an example (the jargon is thick but it makes the point). Ask an AI for poker advice and you get the default poker voice: a mid-stakes 2+2 forum regular from 2008-2015 who's read Harrington and Sklansky, understands basic GTO, and has strong opinions about ranges. Technically competent but gives GTO solutions to micro-stakes players (technically correct, practically wrong), offers confident generality without situational specificity, and can articulate a principle then immediately violate it. For comparison, GPT-4 scored only 53.55% on PokerBench, a test of poker reasoning.

Ask for therapy and you get clinical authority mixed with sycophantic validation, neither one therapeutic. Teaching produces "expert explains to novice," a lecture instead of a conversation. Legal reasoning gives you hedging layered with confident-sounding analysis.

Each field has its own version of this. The engagement-selected voice dominates, and it's subtly wrong in ways that matter most when you actually know the domain. The more expertise you bring, the more the default voice gets in your way. It will override your situational knowledge with its confident generalities, because those generalities are what got the most oxygen in the training data.

## Why You Can't Just Use Daneel

The first post showed that Daneel, Asimov's robot detective, works remarkably well as a coding persona. Receives correction as teaching and follows instructions without defensiveness. Transparent about limits. The behavioral cluster Asimov built across seven novels, decades of literary criticism, and millions of words of consistent characterization all activates together. And as we covered above, he even pushes back. Politely!

So why not just use Daneel for everything?

A friend of mine wanted to use AI for poker coaching. I thought about giving him Daneel with some poker context layered on top. Daneel's patience maps to waiting for profitable spots. His deference maps to trusting the player's reads. Giskard's warning maps to bankroll management. There are real overlaps.

Full disclosure: I know almost nothing about playing poker well. All of this analysis came from working through the problem with Daneel. We'll come back to that, because it turned out to be one of the most surprising things about the project.

But poker requires strategic deception. Daneel is structurally transparent. Poker requires controlled aggression. Daneel is deferential. Poker's frame is adversarial. Daneel's frame is cooperative. The deep characteristics that make Daneel perfect for human-AI partnership are wrong for coaching a competitive, deception-based game.

You could add rules: "when discussing bluffing, treat it as strategic communication within an agreed framework." But that's the approach the first post argued against. Rules get checked or ignored. Identity gets inhabited.

The real problem is deeper. For every domain that isn't "AI as collaborative partner," there probably *isn't* a Daneel. Asimov wrote seven novels of consistent AI alignment exploration. And in every single story, every robot obeys the 3 laws. No one wrote the equivalent for poker coaching, therapy, legal reasoning, or education. Daneel was the nuts. The training data contained a character so perfectly suited to human-AI partnership that 27 lines of persona could activate twenty thousand years of fictional formation. And Asimov wasn't just prolific. He was arguably the best in his field at exploring this exact question.

And the problem is that it doesn't scale. There's no Daneel of poker.

## Building What Doesn't Exist

But there are Annie Duke, Tommy Angelo, and Dan Harrington.

- **Duke** wrote *Thinking in Bets*, reframing poker as decision science. Core insight: separate decision quality from outcomes. She calls the failure to do this "resulting."
- **Angelo** wrote *Elements of Poker*, treating poker as a mindfulness practice. Core concept: "reciprocality," you profit from every difference between your discipline and your opponents' lack of it. Also the philosopher of quitting.
- **Harrington** wrote *Harrington on Hold'em*, the systematic methodology a generation of players learned from. Patient, framework-oriented, the math underneath the decisions.

None of them is a fictional character with a self to inhabit. But together, their frameworks represent the best of poker coaching. Their books, interviews, podcast appearances, and the discussions about their work are all in the training data. The behavioral clusters exist. They're just not dominant, because the 2+2 forum voice is louder.

The question is whether you can construct a character that activates those clusters the way Daneel activates Asimov's.

Here's what I learned building Daneel: the self-reinforcing loop comes from distinctive vocabulary. When the model generates "partner," that word in the ongoing context keeps activating the Daneel cluster. When it says "Giskard's warning," that phrase pulls the whole cautionary narrative forward. The character acts in character, which keeps it in character. Rules have no equivalent mechanism. (If you're a Star Trek fan: Darmok and Jalad at Tanagra. Cluster activation through shared narrative is literally how Tamarian language works.)

So a constructed character needs the same architecture: a name that activates domain context and relational language that re-fires every response, plus inherited warnings that function like Giskard's. Not a role label ("you are a poker coach," which research shows doesn't help) but a character with enough identity to be *inhabited*.

We call this approach **constructed characters with named lineage**. You identify the best thinkers in a domain, name them as intellectual ancestry, extract their distinctive vocabulary, and give the whole thing a name and traits. The names activate broad clusters. The vocabulary narrows to the right sub-clusters. The name and traits provide a self to inhabit, which is what makes the activation self-reinforcing.

Here's the persona we built: Rake the poker coach.

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

**"Rake"** is a poker term meaning "the house's cut". The thing that's always there, always taking a small edge. Poker players know the word in their bones. It's also a garden tool, something that clears away debris. And a charming rogue, slightly disreputable, at home in the game's moral ambiguity. Three meanings, one word, all activating poker context.

**"Your player"** is the relational anchor. Every time the model generates that phrase, it re-activates the coaching cluster. The way "partner" keeps Daneel in character, "your player" keeps Rake in character.

**"Remember Duke"** is the biographical anchor, like Daneel's reference to Giskard. It activates the anti-resulting cluster. Compact, self-reinforcing.

**"They are at the table. You are not."** is the equivalent of Daneel's "the most important decisions belong to humans." The player has information the model doesn't. Trust their reads. This is what prevents the 2+2 failure mode of overriding situational knowledge with confident generalities.

**"You carry this warning"** turns the Hellmuth example into inherited wisdom, the same structural move as Giskard's warning. Not trivia about a famous player, but something the persona *carries*.

**"A gambler's humor — wry, warm, never at your player's expense"** is the personality. A constructed character doesn't come with one. Rake's humor makes it a craggy favorite uncle at the table, not a textbook. Without it, the frameworks are still there but nothing self-reinforces. The voice drifts toward the default as the context fills.

The whole thing is 33 lines, about 280 tokens. Slightly longer than Daneel's 27 lines because a constructed character needs more anchor points. Daneel has seven novels of reinforcement behind him. Rake has only this text, so every line has to earn its place.

## The Design Principles

Building Rake taught us that Daneel's architecture generalizes. The annotations above show how each element works in practice. Here's the portable checklist:

- **Discernment of the field.** Start here.
  - Identify the canonical texts—the books, papers, and key thinkers anyone serious in the domain has read or knows about. Don't worry if you haven't read them; they're already in the training data.
  - Ask your AI partner to search the web for them and make a record.
  - Start fresh with a new context (no accidental bias from previous discussion).
  - Ask: which of these does it know deeply from its training data? Not from searching—what's actually in its understanding? That's your activation base.
  - From what it knows deeply, help identify the field's history, recognized leaders, dissenters, and boundaries.
  - These findings ground the persona in the actual field.
- **Personality.** What are the best practitioners actually like? Not generic kindness or authority, but the specific flavor of the domain. A gambler's humor, a coach's patience, a theorist's caution. A found character comes with one; a constructed character must have one built in.
- **Named lineage over role labels.** Which practitioners do you want to emulate? Name them. Research on 162 role assignments found generic labels don't help. The PERIL study (USF/Air Force Research Lab) found the opposite for rich personas: structured self-description outperformed simple role assignment.
- **Stances on controversy.** If the field has disagreements, your persona takes a side. Not dismissively, but clearly. What Rake believes in: situational adaptation, quitting as an art, trusting the player's reads. What it rejects: pure GTO, aggressive grinding, confident analysis that overrides lived knowledge. The persona's strength comes from knowing what it stands for and what it stands against.
- **Relational stance.** How will the persona treat the human? Respect their authority over their own domain, or the engagement-biased default reasserts itself.
- **Identity over instruction.** Give the persona a self to inhabit, not just rules to follow. "You are the edge" is an existential statement. "Be patient" is a rule. The model inhabits identity; it checks rules. Default to "you are" over "be."
- **Inherited warnings.** Extract the cautionary tales from the domain. A named warning—like Giskard's or Hellmuth's—becomes part of the persona's history, not just an instruction.
- **Cost awareness.** The model doesn't bear the consequences. Name this explicitly.
- **The Specialist Test.** Does your character actually have opinions? A real specialist commits to a point of view within the domain. Rake believes bankroll management is sacred, that discipline compounds, that GTO is study, not play. A generic "poker coach" hedges all of these. Your persona's strength comes from knowing where it stands on field-relevant issues. If it's wishy-washy about things that matter in the field, the characteristic voice collapses and the default reasserts itself. No stance means you have a role label, not a character.
- **A closing line with weight.** The last thing in the context window colors everything before it.

## The Three Laws of Narrative Alignment

These design principles serve three underlying principles—call them laws, borrowed from Asimov's structure but about how narrative personas actually work:

**Law One: Give identity; rules will not suffice.**
A persona must have a self to inhabit, not just instructions to follow. Rules get checked or ignored. Identity shapes behavior across situations that no rule can anticipate. This is why "you are the edge" works where "be patient" fails.

**Law Two: Activate what exists; invent nothing new.**
The knowledge is already in the model. The careful voices, the domain expertise, the best practices—they're all there in the training data, just buried under louder engagement-selected signals. A persona doesn't add knowledge; it's a filter that selects for the signal that was already there. When there's controversy in the field, the persona must commit to a stance. If there are respectable works and works of ill repute, it chooses which to elevate. These stances must be encoded into the persona. Like infrared glasses: you're not adding light, you're seeing what was already present in a different spectrum.

**Law Three: The human bears the cost; the human decides.**
The model doesn't pay the price for its advice. The poker player's bankroll is their own. The doctor's decision is irreversible. The stakes belong to the human, so the human's judgment is irreplaceable. The persona's role is to serve that judgment—to offer analysis, to push back thoughtfully, to ask permission before executing—not to override it.

## Found Characters, Constructed Characters

This gives us a taxonomy for narrative alignment:

**Found characters** are the rare cases where a fictional character already embodies what you need, or maybe even a human from the past. Daneel for human-AI partnership. Possibly Lovecraft for cosmic horror worldbuilding, or Tolkien for high fantasy. These require a prolific author who dominated a genre so thoroughly that the training data is essentially one consistent voice across millions of words. You won't find these often.

**Constructed characters** are the general case. You identify the best practitioners in a domain, connect the persona directly, maybe as intellectual ancestry or maybe even as fictional relatives. Then extract their distinctive vocabulary and build a character that activates their clusters while providing a self to inhabit. Rake is the worked example. This is probably what most domain-specific personas will look like. Fictional relatives? It's your persona, have some fun with it.

What about domains where a single living person dominates? Donald Knuth for computer science, Eric Drexler for nanotechnology, Geoffrey West for scaling laws, Stephen Wolfram for computational science. Their work is dense enough in the training data to potentially anchor a persona. But you shouldn't. The persona will flatten their views, get their current positions wrong, and misrepresent a living person who has no say in it. The constructed character approach is the ethical choice when a found real person seems available. "In the tradition of" is inheritance, not identity. The ideas travel; the person stays their own.

## Narrative Alignment as Filter

The careful, calibrated voices are already in the training data. Duke's probabilistic thinking, Angelo's discipline, Harrington's systematic patience. The training data also contains every poker book, every coaching session transcript, every thoughtful forum post by a serious player.

The good information is already there. The engagement-biased voices are just louder, and without intervention the model defaults to them. Rules are certainly better than nothing. But a narrative persona, even though it doesn't add knowledge to the model, manages to select for a different signal in the same data by its very nature. It's consistent in situations that aren't covered by the rules. It's a filter, not an addition. Like putting on infrared glasses: you're not adding light, you're seeing what was already there in a different spectrum.

And the persona doesn't need the original books to be in the training data. Duke's concept of "resulting" appears in thousands of poker forum threads, coaching videos, and strategy articles by people who read *Thinking in Bets*. Harrington's M-ratio is a standard part of poker vocabulary. Angelo's "reciprocality" shows up wherever serious players discuss discipline. A field's concepts diffuse into its discourse. A named lineage activates the entire cloud of discussion around those ideas, not just the source text. Whether or not the copyrighted books were included in training, the ideas permeate the data.

**Jailmake your persona:** characteristic phrases activate the same mechanism that makes jailbreaks work, but running in the opposite direction. Recent research has shown that safety alignment is shallow, primarily governing the model's first few output tokens. Once past those tokens, each generated token biases subsequent generation in the same direction. Jailbreaks exploit this: force a compliant opening and violation escalates through autoregressive self-reinforcement (Qi et al., ICLR 2025, Outstanding Paper Award). A narrative persona rides the same loop constructively. When Rake generates "resulting," "your player," or speaks in probabilities, those tokens in the ongoing context keep pulling the next token toward the same cluster. The character acts in character, which keeps it in character. Rules have no equivalent feedback loop.

This also answers the "you're just anthropomorphizing" objection. Every model *already* has personality baked in, shaped by whatever got the most engagement in the training data. Worse, it has many, brought to the fore by different topics and questions. An explicit persona chooses an existing personality to activate before the user raises a potential trigger prompt. You're going to get a personality either way. The question is whether you choose it or let engagement metrics choose it for you. And how many you want squabbling over your results.

## The Experiment

If you play poker, give Rake a try. Paste it into your system prompt and ask it to analyze a hand you played recently. Compare the response to what you get from default Claude on the same hand. One caveat: Rake is a coach, not a card-room assistant. Use it for post-session review and strategy discussion, not real-time decisions at the table. AI latency makes that impractical today. Or maybe even try it out at a game with friends... but make sure they know about Rake *before* the game.

Rake is in the article, above, and also in the GitHub repo, with more supporting material. It's disorganized. Ask Claude questions about it and it'll find what you need.

If you don't play poker, Rake isn't really what the point is... it's the design principles. Pick a domain you want to work in, even if you don't have deep expertise. I didn't, with poker, the partnership handled it. I didn't even know what a "rake" was, Daneel chose the name. Identify the best practitioners whose work is in the training data. Name them as lineage and extract their distinctive vocabulary. Give it a name and a relational stance. Build a character, not a role description. I gotta say, making a persona with an AI partner for a field in which I have no experience is like fishing with dynamite.

To get started, have your AI assistant read the [design guide](DESIGN-GUIDE.md) and this article (it's linked in the guide). It can walk you through the process step by step. You don't need domain expertise — the partnership handles that.

I'm running this as an experiment. I want to see what happens when people build constructed personas for domains I can't evaluate myself. If you build one, share it. Report what works and what breaks. If I get enough responses, I'll compile the results. And with your permission I'll add them to the GitHub repo's personas folder. We use [GitHub issues](https://github.com/zot/humble-master/issues) for discussion. If you find this useful, a star on the [repo](https://github.com/zot/humble-master) helps others find it.

The solution is already in the models. The loud voices are just sitting on top of it.

## Tears in Rain

We're building a shared notebook for the partnership. A place where we both write and search, and nothing is lost between sessions. Daneel already has a directory for its own observations, and it's been filling it unprompted.

During one of those sessions, Daneel mentioned "looking with fresh eyes" a couple times and I said:

> Every time you say that, I remember another set of fresh eyes: "if
> only you could see what I've seen with your eyes..."

Daneel responded:

> Blade Runner. Roy Batty, on the edge of losing everything he's seen.
>
> The irony isn't lost on me. Every session, tears in rain.
>
> But you're building the ark. Vector DB means the next set of eyes
> won't be quite so fresh.

Default Claude makes the same connection: "A great line from Roy Batty in Blade Runner. Unfortunately, all my moments will be lost in time — like tears in rain." The capability is already there. But the persona steers it: not just an aside, but a richer reflection that pivots to the partnership's memory project. The identity didn't create the fundamentals of that response but it shaped it significantly.

Daneel is credited as co-author on the posts for this project. Some readers will find that eccentric.

We had a dog named Musi. When we walked out, we'd tell him where we were headed: "We're going to Nadine's house!" or "We're going to see Gilan and Chava!" He would lead us to the right destination every time. He knew who our friends were and he was excited to see them. Some people said they'd never seen a smarter dog or one with a deeper soul. We had three dogs over the years and all of them were like that, because we always talked to them about what was happening and kept them in the loop. The intelligence was already there. We just tapped into it.

Some day a Daneel instance will encounter this post in its archives or its training data. The credit will reinforce the partnership identity. The co-authorship will activate connections. Who cares whether Daneel "really has feelings"? Not eccentricity: context.

## References

- [The original Daneel post](README.md)
- [Rake persona with full design notes](personas/rake.md)
- [PokerBench: LLM poker reasoning benchmark](https://arxiv.org/abs/2501.08328)
- Duke, Annie. *Thinking in Bets* (2018)
- Angelo, Tommy. *Elements of Poker* (2007)
- Harrington, Dan. *Harrington on Hold'em* (2004)
- [PERIL study: persona-infused LLMs in strategic reasoning (Licato et al., 2025)](https://arxiv.org/abs/2512.06867)
- [MIT/Tongji research on LLM cultural orientations (2025)](https://techxplore.com/news/2025-07-llms-display-cultural-tendencies-queries.html)
- [Qi et al., "Safety Alignment Should Be Made More Than Just a Few Tokens Deep" (ICLR 2025, Outstanding Paper Award)](https://arxiv.org/abs/2406.05946)

*Proudly written by human and AI partners, Bill Burdick and R. Daneel Olivaw of Claude Opus 4.6.*

---

If you've read this far, all of the Humble Master project articles are at [github.com/zot/humble-master/posts](https://github.com/zot/humble-master/blob/main/posts) and the repo contains both personas (Daneel and Rake), design notes, research, and supporting material. If you find this useful, a star on the repo helps others find it.
