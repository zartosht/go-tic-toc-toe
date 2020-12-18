# TIC TOC TOE

This is my journey to learning `Golang`, my first approach was to make a board with rows and columns that user will specify. In this first version, user can not play and computer will play against itself and the winner will be picked by counting how many tiles belong to whom (X or O).

# Run

For running this game, do the following:

```bash
$> clone https://github.com/zartosht/go-tic-toc-toe
$> cd go-tic-toc-toe
$> go build tictoctoe.go
$> ./tictoctoe
```

It will output:

```bash
How many columns> 5
How many rows> 4
O won the game!
|¯¯¯¯¯|¯¯¯¯¯|¯¯¯¯¯|¯¯¯¯¯|¯¯¯¯¯|
|  O  |  X  |  O  |  O  |  X  |
|_____|_____|_____|_____|_____|
|     |     |     |     |     |
|  O  |  O  |  O  |  O  |  O  |
|_____|_____|_____|_____|_____|
|     |     |     |     |     |
|  X  |  X  |  X  |  X  |  X  |
|_____|_____|_____|_____|_____|
|     |     |     |     |     |
|  O  |  X  |  O  |  X  |  O  |
|_____|_____|_____|_____|_____|

```

# Changelogs

## 0.0.1 (2020-12-18)

### Added

- Create a board with desired columns and rows.
- Pick winner by determining who has most tiles.