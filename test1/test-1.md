# Test Number 1

============================ 
1. Simple Database querying 
============================ 
Terdapat sebuah table "USER" yg memiliki 3 kolom: ID, UserName, Parent. Di mana: Kolom ID adalah Primary Key 
Kolom UserName adalah Nama User 
Kolom Parent adalah ID dari User yang menjadi Creator untuk User tertentu. 
eg. 
—————————————————————————— 
| ID | UserName | Parent | 
—————————————————————————— 
| 1 | Ali | 2 | 
| 2 | Budi | 0 | 
| 3 | Cecep | 1 | 
—————————————————————————- 
Tuliskan SQL Query untuk mendapatkan data berisi: 
—————————————————————————————————— 
| ID | UserName | ParentUserName | 
—————————————————————————————————— 
| 1 | Ali | Budi | 
| 2 | Budi | NULL | 
| 3 | Cecep | Ali | 
—————————————————————————————————— 
*Kolom ParentUserName adalah UserName berdasarkan value Parent 
================================================= 

```text
Answer : SELECT child.ID, child.UserName, parent.UserName as ParentUserName FROM users AS child LEFT JOIN users as parent ON child.Parent = parent.ID
```