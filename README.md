# lem-in

# lem-in

This project is meant to code a digital version of an ant farm.

Lem-in reads from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

How does it work?

- You make an ant farm with tunnels and rooms.
- You place the ants on one side and look at how they find the exit.

Lem-in finds the quickest way to get n ants across a colony (composed of rooms and tunnels).


- At the beginning of the game, all the ants are in the room ##start. The goal is to bring them to the room ##end with as few moves as possible.
- The shortest path is not necessarily the simplest.
- Some colonies will have many rooms and many links, but no path between ##start and ##end.
- Some will have rooms that link to themselves, sending your path-search spinning in circles, some will have too many/too few ants, no ##start or ##end, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In this cases the program will return an error message ERROR: invalid data format. If you want, you can elaborate the error message by being more specific (example: ERROR: invalid data format, invalid number of Ants and ERROR: invalid data format, no start room found).

This project helps to learn about :

- Algorithmics
- Ways to receive data
- Ways to output data
- Manipulation of strings
- Manipulation of structures
