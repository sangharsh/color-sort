# color-sort

### Basics
1. Colors
1. Testtubes
1. Level

### GamePlay
1. Pour
1. Undo
1. Win

### Under the hood
1. Generate levels
1. Ensure solvability
1. Define & assess difficulty
1. Track user competence

### Client
1. Render
1. Game play
1. Gather user performance

### Features
1. Show whether game is solvable without undo

### Experience
1. Ideally game should entirely be on client side
1. Use cookie / localstorage to track progress across levels and within level

## API
1. Get game level
    - int GameLevel
1. Make a move
    - src, dst - status, updateLevel, moves?
1. Show win
1. Go to another level
1. Undo

## My Wishlist
1. Protobuf and gRPC
1. CI/CD
    1. Bazel?
1. Solver and Difficulty
1. Monorepo - Client, server, infra

## Extras
1. Add a testtube
1. Hidden color (?) gameplay
1. Animations


Q:
Trust client or not?
- Don't trust