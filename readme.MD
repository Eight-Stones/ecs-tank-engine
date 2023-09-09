## Game tanks engine
### Game interface
#### Rotate
##### Aim
Turns object in select direction.
##### Params
- id (string): identificator of objects.
- direction (int): code identificator of direction.  
  - Left = 1
  - Right = 2
  - Up = 3
  - Down = 4
##### Return
- Bit mask of actions.
#### Move
##### Aim
Turn and moves an object in its direction.
##### Params
- id (string): identificator of objects.
- direction (int): code identificator of direction.
    - Left = 1
    - Right = 2
    - Up = 3
    - Down = 4
##### Return
- Bit mask of actions.
#### Shoot
##### Aim
Shoots in object direction.
##### Params
- id (string): identificator of objects.
##### Return
- Bit mask of actions.
#### Vision
##### Aim
Return small area around object.
##### Params
- id (string): identificator of objects.
##### Return
- Bit mask of actions.
- List with data (tank/bullets/walls/etc), area around object.
#### Radar
##### Aim
Return big area around object, has long recharge.
##### Params
- id (string): identificator of objects.
##### Return
- Bit mask of actions.
- List with data (tank/bullets/walls/etc), area around object.
### Game Config
### Info statistic interface
Engine return statistic in map of maps of interfaces
#### List of actions 
| Value                    | Bit                     |
|--------------------------|-------------------------|
| 	Fail                    | 0b1                     |
| 	Ok                      | 0b10                    |
| 	Ban                     | 0b100                   |
| 	OkBorder                | 0b1000                  |
| 	FailBorder              | 0b10000                 |
| 	OkStep                  | 0b100000                |
| 	FailStep                | 0b1000000               |
| 	OkRotate                | 0b10000000              |
| 	FailRotate              | 0b100000000             |
| 	OkCollision             | 0b1000000000            |
| 	NoCollision             | 0b10000000000           |
| 	NotInterruptOkCollision | 0b1000000000000         |
| 	Damaged                 | 0b10000000000000        |
| 	BothDamaged             | 0b100000000000000       |
| 	OkShot                  | 0b10000000000000000     |
| 	FailShot                | 0b100000000000000000    |
| 	Disappear               | 0b1000000000000000000   |
| 	Found                   | 0b10000000000000000000  |
| 	NotFound                | 0b100000000000000000000 |

#### Methods
##### CurrentState
Return info about alive objects
##### ResultState
Return info about objects
#### Format of result
map1[key1]map2[key2]value
- map1 - summary map with all objects
- key - id of object
- map2 summary about concrete object
- key2 - one of keys from table below

| Key                       | Value      | Meaning                                               |
|---------------------------|------------|-------------------------------------------------------|
| KeyStatObjectKind         | kind       | type of object, can be tank or bullet etc             |
| KeyObjectKindTank         | tank       | object is tank                                        |
| KeyObjectKindBullet       | bullet     | object is bullet                                      |
| KeyStatPositionCoordinate | coordinate | coordinate of object, list of two el (x and y values) |
| KeyStatMovementDirection  | direction  | direction of object                                   |
| KeyStatHitPoints          | hit-points | current level of hit point of objects                 |
| KeyStatActions            | actions    | list codes of action of player                        |