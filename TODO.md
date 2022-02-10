## Update snakes.go
    - checking for larger / smaller snakes ought to check all directions around the target cell, not just up/down/left/right,
    because it is possible for a snake to be in the kitty-corner position and presently we are unaware of that.

## Update wrapped logic
    - find a way to count empty cells in a given direction, UP TO finding any snake body
    - must "wrap" around game board and not stop at edges
