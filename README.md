# Chords generator

This is a program that computes the notes for triad and seventh chords. I use an automaton $\text{Chords computer}$ to acomplish that:

$$
\text{Chords computer} = S \cup C
$$

where $C$ is the automaton that computes chord notes and $S$ is the automaton that splits the symbols of the chord description $w$ as letters $x$ for the automaton $C$

## TODO

- Support seventh chords
- Create the automaton $S$ for recognizing valid words for the language of the automaton $C$
