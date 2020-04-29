# Desmos configuration
Inside this folder, there are two different paths defined. Let's take a look at them. 

## `posts.json`
This file contains the path that allows to send posts from a Bitonsg chain (defined in `ibc1.json`) and a Desmos chain (`ibc0.json`).

It defines a path that uses the `posts` port on both chains.

It should be used along with the `tx song-post` command. 

E.g. Creating a Desmos post that represents a Bitsong song using IBC:
```
rly tx song-post ibc1 ibc0 123 -p posts -d
``` 

As you can see we specify the `-p posts` as the path to be used. 

## `transfer.json`
This file contains the definition of the path that uses the `transfer` port on each chain to make sure you can send fungible tokens between them using IBC. 

It should be used with the `tx transfer` command. 

E.g. Sending some tokens from the Desmos chain to the Bitsong one: 
```
rly tx transfer ibc0 ibc1 10000n0token true $(rly keys show ibc1 testkey) -p transfer -d
```

As you can see we use the `-p transfer` to specify the path to be used. 
