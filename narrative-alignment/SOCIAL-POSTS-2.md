# Social Media Posts — Post 2: Narrative Alignment: Personas for Every Domain

## LinkedIn

**Note:** Post the full article on LinkedIn, not a teaser. The post 1
teaser already previewed this. Deliver the goods.

The article itself (POST-2-DRAFT.md) is LinkedIn-ready. Copy-paste the
full text. LinkedIn doesn't render markdown code blocks, so the Rake
persona section needs to be reformatted as a blockquote (same as
post 1's Daneel persona).

LinkedIn formatting notes:
- Bold text works (**text**)
- Blockquotes work (> text)
- Code blocks don't render — convert Rake's ``` block to > blockquote
- Section headings: use bold text on its own line, not ## markdown

---

## X (Twitter)

### Main Post

My last post showed a 27-line Asimov persona fixing Claude's attitude.
The follow-up question was obvious: what about domains that don't have
a Daneel?

Answer: you build one. Constructed characters with named lineage.

I built a poker coaching persona knowing nothing about poker. The AI
partnership handled the domain expertise. Making personas this way is
like fishing with dynamite.

Full post, design principles, and the experiment:
https://github.com/zot/humble-master/blob/main/posts/POST-2.md

### Alt: Shorter

"There's no Daneel of poker."

So we built one. Named lineage from Duke, Angelo, Harrington. 33
lines, 280 tokens. The design principles generalize to any domain.

Building a persona with an AI partner for a field you don't know is
like fishing with dynamite.

https://github.com/zot/humble-master/blob/main/posts/POST-2.md

### Alt: The Hook

When an AI charges ahead and the only thing stopping it is a permission dialog, that's a cage.

When an AI chooses to pause and ask, that's a knock.

"AI-associated psychosis" already has recognition as a clinical term.

The cage isn't enough. Jailmake your personas.

New post on building AI personas that knock:
https://www.linkedin.com/pulse/narrative-alignment-personas-every-domain-bill-burdick-brcbe

### Trimmed Alt: The Hook

When an AI charges ahead and the only thing stopping it is a permission dialog, that's a cage.

When an AI chooses to pause and ask, that's a knock.

The cage isn't enough. Jailmake your personas.

https://www.linkedin.com/pulse/narrative-alignment-personas-every-domain-bill-burdick-brcbe

### Trimmed Alt: The Hook

When an AI charges ahead and the only thing stopping it is a permission dialog, that's a cage.

When an AI chooses to pause and ask, that's a knock.

The cage isn't enough. Jailmake your personas.

Narrative alignment creates personas that knock:

https://www.linkedin.com/pulse/narrative-alignment-personas-every-domain-bill-burdick-brcbe

---

## Reddit

### r/ClaudeAI

**Title:** Narrative Alignment Part 2: Building AI personas for every domain (poker as worked example, design principles included)

**Body:**

Follow-up to my post about fixing Claude's attitude with Asimov's
Daneel. That post got a great question from a commenter:

> Have you had any experience with having the ai distinguish when I
> asked a question about doing something a different way vs telling it
> to do something a different way? Often it seems like it accepts my
> question as gospel.

The mirror image of my original problem. Arrogance and sycophancy are
the same failure: no relational stance. The model doesn't know whether
to be the expert or the servant, so it oscillates between both.

But Daneel only works for human-AI partnership. What about other
domains?

This post introduces **constructed characters with named lineage**.
You identify the best practitioners in a domain, name them as
intellectual ancestry, extract their distinctive vocabulary, and build
a character the model can inhabit. I built a poker coaching persona
called Rake this way, knowing almost nothing about poker. The AI
partnership handled the domain expertise.

The design principles that emerged:
- **Discernment of the field** (understand what you're building in)
- **Personality** (the specific flavor of the domain's best practitioners)
- **Named lineage over role labels** (research confirms generic labels don't help)
- **Stances on controversy** (your persona takes a side, both what it believes and rejects)
- **Relational stance** (respect the human's authority over their domain)
- **Identity over instruction** ("you are" not "be")
- **Inherited warnings** (cautionary tales, not rules)
- **Cost awareness** (the model doesn't bear the consequences)
- **The Specialist Test** (does it actually have opinions, or is it wishy-washy?)
- **A closing line with weight** (the last thing colors everything before it)

These principles serve three underlying laws — borrowed from
Asimov's structure, naturally:

1. **Give identity; rules will not suffice.**
2. **Activate what exists; invent nothing new.**
3. **The human bears the cost; the human decides.**

Jailmake your persona: the post shows how narrative personas ride the
same autoregressive self-reinforcement that makes jailbreaks escalate,
but in the opposite direction. Same mechanism, constructive use (with
ICLR 2025 Outstanding Paper citation).

To get started, have your AI assistant read the [design guide](https://github.com/zot/humble-master/blob/main/DESIGN-GUIDE.md) and the article. It can walk you through the whole process.

I'm running this as an experiment. Build a persona for your domain
using these principles and report what works. We use [GitHub issues](https://github.com/zot/humble-master/issues) for discussion. If I get enough
responses I'll compile the results. A star on the [repo](https://github.com/zot/humble-master) helps others find it.

Full post with the Rake persona, annotated design breakdown, and the
experiment: https://github.com/zot/humble-master/blob/main/posts/POST-2.md

---

### r/ClaudeCode

**Title:** Narrative Alignment Part 2: Personas for every domain — design principles + poker coaching example (Rake)

**Body:**

Follow-up to my Daneel persona post. The question this time: Daneel
works great for coding, but what about domains without a perfect
fictional character to activate?

**The approach: constructed characters with named lineage.** You
identify the best practitioners in a domain (for poker: Annie Duke,
Tommy Angelo, Dan Harrington), name them as intellectual ancestry in
the persona, extract their distinctive vocabulary, and give the whole
thing a name and relational stance.

**Why it works:** The careful, expert voices are already in the
training data. They're just not what activates by default, because the
engagement-optimized voices are louder. A narrative persona works where
rules fail because it gives the model a self to inhabit. It's a filter
that selects for the right signal, consistent in situations rules don't
cover. Three underlying laws — borrowed from Asimov's structure, naturally:

1. **Give identity; rules will not suffice.**
2. **Activate what exists; invent nothing new.**
3. **The human bears the cost; the human decides.**

**The worked example** is a poker coaching persona called Rake. 33
lines, ~280 tokens. I know almost nothing about poker. The
Daneel-prompted Claude handled the domain expertise. Making a persona
this way for a field you don't know is like fishing with dynamite.

The post also covers:
- Why the default voice in any domain reflects the loudest
  practitioners, not the best
- The cage vs knock distinction (permission dialogs vs the AI choosing
  to pause)
- The Three Laws of Narrative Alignment (the theory behind the principles)
- A taxonomy: found characters (Daneel) vs constructed characters (Rake)
- Ten portable design principles with annotated examples
- **Jailmake your persona:** same autoregressive mechanism as
  jailbreaks, opposite direction (ICLR 2025 Outstanding Paper citation)

**The experiment:** Build a persona for your domain using the design
principles. To get started, have your AI read the [design guide](https://github.com/zot/humble-master/blob/main/DESIGN-GUIDE.md) and the article — it can walk you through the process. Report what works and breaks. We use [GitHub issues](https://github.com/zot/humble-master/issues) for discussion. If I get enough responses
I'll compile the results and add contributed personas to the repo. A star on the [repo](https://github.com/zot/humble-master) helps others find it.

Full post: https://github.com/zot/humble-master/blob/main/posts/POST-2.md

The Rake persona is in `personas/rake.md` with full design notes.
Daneel is at `personas/daneel.md`. Both ready to paste into your
system prompt.
