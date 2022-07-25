# go-ssh
    this is kit for manage you ssh connect.

### Command

###### go-ssh create
    Creates a new ssh connection.
```bash
$ go-ssh create 
Enter ssh connect name (inside 32 characters):
test
Enter ssh connect host (inside 32 characters):
192.168.75.167
Enter ssh connect port (in (0, 65535) default:22):

Enter ssh connect user (inside 32 characters):
admin
Enter ssh connect password (inside 255 characters):
123456
ssh create success
```
###### go-ssh list
    List all ssh connections or list named ssh connection.
```bash
# list all ssh connections
$ go-ssh list     
+---------+----------------+------+--------------+----------+----------------------+
|  NAME   |      HOST      | PORT |     USER     | PASSWORD |       CREATED        |
+---------+----------------+------+--------------+----------+----------------------+
| test    | 192.168.75.167 | 22   | admin        | ***      | 2022-07-25T15:36:50Z |
+---------+----------------+------+--------------+----------+----------------------+
# list named ssh connection
$ go-ssh list -n ovz-167
+---------+----------------+------+--------------+----------+----------------------+
|  NAME   |      HOST      | PORT |     USER     | PASSWORD |       CREATED        |
+---------+----------------+------+--------------+----------+----------------------+
| test    | 192.168.75.167 | 22   | admin        | ***      | 2022-07-25T15:36:50Z |
+---------+----------------+------+--------------+----------+----------------------+
```
###### go-ssh shell
    Login a ssh connection.
```bash
$ go-ssh shell test  
Last login: Mon Jul 25 15:35:48 2022 from 10.200.136.24
admin@192.168.75.167:~$ 
```
###### go-ssh remove
    Remove a ssh connection.
```bash
$ go-ssh remove       
Enter ssh connect name:
test
test ssh connect info:
+---------+----------------+------+--------------+----------+----------------------+
|  NAME   |      HOST      | PORT |     USER     | PASSWORD |       CREATED        |
+---------+----------------+------+--------------+----------+----------------------+
| test | 192.168.75.167    | 22   | admin        | ***      | 2022-07-25T15:36:50Z |
+---------+----------------+------+--------------+----------+----------------------+
Are you confirm remove this ssh connect? please enter (yes) to confirm:
yes
ssh remove success
```