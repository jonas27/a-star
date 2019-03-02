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
