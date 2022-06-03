# EXERCISE: Structs

Starting with this exercise, you'll build a command-line game store.

1. Declare the following structs:

    - item: id (int), name (string), price (int)
    - game: embed the item, genre (string)

2. Create a game slice using the following data:

    id, name, price, genre

    - 1, "god of war", 50, "action, adventure"
    - 2, "x-com", 2, 30, "strategy"
    - 3, "minecraft", 20, "sandbox"

3. Print all the games.

## EXERCISE: List

Now, it's time to add an interface to your program using
the bufio.Scanner. So the users can list the games, or
search for the games by id.

1. Scan for the input in a loop (use bufio.Scanner)
2. Print the available commands.
3. Implement the quit command: Quits from the loop.
4. Implement the list command: Lists all the games.

Try the program with list and quit commands.

# EXERCISE: Query By Id

Add a new command: "id". So the users can query the games by id.

1. Before the loop, index the games by id (use a map).
2. Add the "id" command.
   When a user types: id 2, It should print only the game with id: 2.
3. Handle the errors:

Try the program with list, quit, and id commands to see it in action.

```
id
wrong id
```

```
id HEY
wrong id
```

```
id 10
sorry. I don't have the game
```

```
id 1
#1: "god of war" (action adventure) $50
```

```
id 2
#2: "x-com 2" (strategy) $40
```
