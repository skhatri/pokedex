### pokedex

The purpose of this project is to implement a basic example in a few languages.

The api will touch upon the following:

1. using list/map/struct/class
2. making http calls
3. serialising/deserialising json
4. unit tests
5. local integration tests against the mock server

### Configuration

Generate BUILD files in each module

```
#create BUILD files
bazel run //:gazelle 

#generate repository configuration using go.sum files
./update-repos.sh
```

### Server

Start the server on 0.0.0.0:6100

```
bazel run //server:server
```

The following endpoints will be available:

|Api|Sample URL|
|---|---|
|Find by Type|http://localhost:6100/search/by-type?q=flying,water|
|Find by Name|http://localhost:6100/search/by-name?q=Charizard|
|Find Effective Pokemon Against Type|http://localhost:6100/search/effective-against/by-type?q=dragon|
|Find Effective Pokemon Against Pokemon|http://localhost:6100/search/effective-against/by-name?q=Charizard|

### Tests

Following Questions will be answered and asserted in the tests we write for each client.

1. What is the ability of Mimikyu?
2. How many pokemons are both flying and water type?
3. What is the name of the pokemon that is of type bug and ghost and has the lowest HP?
4. What is the type of pokemon is Zeraora?

### Clients
The following clients are implemented:
* [java](#java)
* [golang](#golang)


#### java
External Jvm Rules have been imported to handle maven dependencies. 

Query Dependency Name with
```
bazel query @maven//:all --output=build
```


#### golang
the go-client is implemented in go-client

Run tests using
```
bazel test //go-client/api:api_test
```

```
bazel coverage  --combined_report=lcov -- //...:all
```

#### rust

```
bazel run //rust:rust-client
```
