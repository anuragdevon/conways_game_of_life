Problem Statement: Conways game of life
    - Requirement 2-D grid
    - Start with Random initial state -> 2-D grid
    - Using the current state of grid we need to move to next state
    - Point -> Print to console(Using a methods are class, but do not use lingering prints)
    - Game Rules:
         - cell[i][j] < 2 => dies - underpopulation
         - cell[i][j] >= 2 => continues to live on next state
         - cell[i][j] == 3 => dead cell comes to life
         - cell[i][j] > 3 => dies - overpopulation


testgridcheck:
[1 1 0]
[0 0 1]
[1 0 0]
1 - alive
0 - dead