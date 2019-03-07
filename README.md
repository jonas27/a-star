# a-star
An implementation of the A* search algorithm

## The Maze
The maze is an array of maps, where the array is the rows
and the keys of the map the columns and the values are the type.<br/>
To create the maze use the Init( ... ) function in package maze.
It has three variables `baseInt`, `rows` and `columns`. `baseInt` just means
the default map integer which will fill the map<br />
A visual representation of the maze with `baseInt=0` looks as follow: <br/>
```
&map[0:0 1:0 2:0 3:0 4:0 5:0 6:0]
&map[0:0 1:0 2:0 3:0 4:0 5:0 6:0]
&map[0:0 1:0 2:0 3:0 4:0 5:0 6:0]
&map[0:0 1:0 2:0 3:0 4:0 5:0 6:0]
&map[0:0 1:0 2:0 3:0 4:0 5:0 6:0]
&map[0:0 1:0 2:0 3:0 4:0 5:0 6:0]
&map[0:0 1:0 2:0 3:0 4:0 5:0 6:0]
```

## The Algorithm
In field x look calculate costs of all neighbors
If a neighboring field/node has a cost lower than the cost of traversing it
from the current node or can't be accessed
(marked as `walkable=false`) the node is not excluded from farther considerations
If a neighbor has higher cost or was not yet visited but is walkable
then its costs will be updated and the node will be put on the Stack.
After evaluating all neighbors the stack will be ordered and the node
with the lowest costs (cost of going to it + its distance vector to the destination)
will be chosen and all neighbors will be evaluated again.
Side Node: This (effectively) implements a priority queue, which also means that
there can be no duplicate items.