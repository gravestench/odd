# About
ODD is an open source isometric action RPG game engine.

ODD is written on top of `director`, and `director` is an ECS scene abstraction written in `akara` and `raylib`.

# Concepts
Game implementations are to be written as either:
* pure golang, as director scenes
* mods, written as a collection of lua scripts for scenes and systems.

**The modding API is not yet defined, still working on it!**

# Project Structure
* **engine/components** - contains the declarations of ECS components, for `akara`.

* **engine/mpq_loader** - contains an mpq archive file loader implementation, for `director`.

* **engine/scenes** - contains graphical scenes _only related to the engine itself_.
This does not contain any specific game code.

* **engine/systems** - contains non-graphical systems _only related to the engine itself_.
This does not contain any specific game code.