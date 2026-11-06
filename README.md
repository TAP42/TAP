A README.md file must be provided at the root of your Git repository. Its purpose is
to allow anyone unfamiliar with the project (peers, staff, recruiters, etc.) to quickly
understand what the project is about, how to run it, and where to find more information
on the topic.
The README.md must include at least:
• The very first line must be italicized and read: This project has been created as part
of the 42 curriculum by <login1>[, <login2>[, <login3>[...]]].
• A “Description” section that clearly presents the project, including its goal and a
brief overview.
• An “Instructions” section containing any relevant information about compilation,
installation, and/or execution.
• A “Resources” section listing classic references related to the topic (documen-
tation, articles, tutorials, etc.), as well as a description of how AI was used —
specifying for which tasks and which parts of the project.
➠ Additional sections may be required depending on the project (e.g., usage
examples, feature list, technical choices, etc.).
Any required additions will be explicitly listed below.
For this project, the following additional sections are required in your README.md:
• A "Architecture" section explaining your server design choices (dispatcher/router
vs inline handling, concurrency model, etc.).
• A "Protocol Implementation" section documenting any deviations from RFC
42TAP and justifying your choices.
• A "Combat System" section describing your turn-based combat mechanics, dam-
age formulas, initiative order, and additional combat commands (DEFEND, FLEE,
etc.).
• A "Quest System" section explaining your quest progression mechanics, comple-
tion validation, and reward systems.
13
TAP A shared-world retro text adventure
• A "World Design" section describing your world layout, room connections, NPC
roles, and item distribution.
• A "Server Logging" section documenting your logging implementation, including
log format, event types, output destinations, and how to monitor server behavior
and detect abuse patterns.
• A "Group Contributions" section clearly indicating each team member’s respon-
sibilities and contributions to different components (server, CLI client, GUI client,
world design, etc.).
• A "Building and Running" section with detailed instructions for your chosen
building tool and how to run each component (server, CLI client, GUI client).
• A "Testing" section explaining how to test the multiplayer functionality, combat
system, and quest mechanics.
