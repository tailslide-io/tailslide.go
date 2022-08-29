<p align="center">
    <img src="https://user-images.githubusercontent.com/73451363/187207442-bae7ea26-7eac-4cab-8806-42779629c653.png" alt="Tailslide logo" width="400">
</p>

# Golang SDK

---

This package is a server-side SDK for applications written in Golang for the Tailslide feature flag framework.

Visit the https://github.com/tailslide-io repository or see Tailslide’s [case study](https://tailslide-io.github.io) page for more information.

## Installation

---

Install the Tailslide npm package with `go get github.com/tailslide-io/tailslide.go`

## Basic Usage

---

### Instantiating and Initializing FlagManager

The `FlagManager` struct is the entry point of this SDK. It is responsible for retrieving all the flag rulesets for a given app with its `AppId` and creating new `Toggler` structs to handle toggling of feature flags within that app. To create a `FlagManager` struct, a user must provide a configuration object:

```golang
import (
	tailslide "github.com/tailslide-io/tailslide.go"
)

func main(){
  config := tailslide.FlagManagerConfig{
    NatsServer:  "nats://localhost:4222",
    NatsStream:  "flags_ruleset",
    AppId:       "1",
    UserContext: "375d39e6-9c3f-4f58-80bd-e5960b710295",
    SdkKey:      "myToken",
    RedisHost:   "http://localhost",
    RedisPort:   "6379",
  }

  manager := tailslide.NewFlagManager(config)
  manager.InitializeFlags()
}
```

- `NatsServer` is the NATS JetStream server `address:port`
- `NatsStream` is the NATS JetStream’s stream name that stores all the apps and their flag rulesets
- `AppId` is the ID number of the app the user wants to retrieve its flag ruleset from
- `UserContext` is the UUID string that identifies the current user
- `SdkKey` is the SDK key for the Tailslide, it is used as a password for NATS JetStream token authentication
- `RedisHost` is the address to the Redis database
- `RedisPort` is the port number that the Redis database runs on

After instantiating a `FlagManager`, invoke the `initialize` method. This method connects the `FlagManager` instance to both NATS JetStream and Redis Timeseries, and asynchronously retrieves the latest and any new flag ruleset data.

---

### Using Feature Flag with Toggler

Once the `FlagManager` is created, it can create a `Toggler`, with the `NewToggler` method, for each feature flag that the developer wants to wrap the new and old features in. A `Toggler`’s `IsFlagActive` method checks whether the flag with its `FlagName` is active or not based on the flag ruleset. A `Toggler`’s `IsFlagActive` method returns a boolean value, which can be used to evaluate whether a new feature should be used or not.

```golang
flagConfig := tailslide.TogglerConfig{
  FlagName: "App 1 Flag 1",
}

flagToggler, err := manager.NewToggler(flagConfig)
if flagToggler.IsFlagActive() {
  // call new feature here
} else {
  // call old feature here
}
```

---

### Emitting Success or Failture

To use a `Toggler` instance to record successful or failed operations, call its `EmitSuccess` or `EmitFailure` methods:

```golang
if successCondition {
  flagToggler.EmitSuccess()
} else {
  flagToggler.EmitFailure()
}
```

## Documentation

---

### FlagManager

The `FlagManager` struct is the entry point of the SDK. A new `FlagManager` object will need to be created for each app.

#### FlagManager Constructor

##### `tailslide.NewFlagManager(options tailslide.FlagManagerConfig)`

**Parameters:**

- `options`: An object with the following keys
  - `NatsServer` is the NATS JetStream server `address:port`
  - `NatsStream` is the NATS JetStream’s stream name that stores all the apps and their flag rulesets
  - `AppId` a number representing the application the microservice belongs to
  - `SdkKey` a string generated via the Tower front-end for NATS JetStream authentication
  - `UserContext` a string representing the user’s UUID
  - `RedisHost` a string that represents the url of the Redis server
  - `RedisPort` a number that represents the port number of the Redis server

**Return Value:**

- A pointer to a `FlagManager` struct

---

#### Instance Methods

##### `flagmanager.initialize()`

Asynchronously initialize `flagmanager` connections to NATS JetStream and Redis database

**Parameters:**

- `nil`

**Return Value:**

- `nil`

---

##### `flagManager.SetUserContext(newUserContext string)`

Set the current user's context for the `flagmanager`

**Parameters:**

- `newUserContext`: A UUID string that represents the current active user

**Return Value:**

- `nil`

---

##### `flagManager.GetUserContext()`

Returns the current user context

**Parameters:**

- `nil`

**Return Value:**

- The UUID string that represents the current active user

---

##### `flagManager.NewToggler(options tailslide.TogglerConfig)`

Creates a new toggler to check for a feature flag's status from the current app's flag ruleset by the flag's name.

**Parameters:**

- `options`: An object with key of `FlagName` and a string value representing the name of the feature flag for the new toggler to check whether the new feature is enabled

**Return Value:**

- A pointer to a `Toggler` object

---

##### `flagManager.Disconnect()`

Asynchronously disconnects the `FlagManager` instance from NATS JetStream and Redis database

**Parameters:**

- `nil`

**Return Value:**

- `nil`

---

### Toggler

The Toggler structs provide methods that determine whether or not new feature code is run and handles success/failure emissions. Each toggler handles one feature flag, and is created by `FlagManager.NewToggler()`.

---

#### Instance Methods

##### `toggler.IsFlagActive()`

Checks for flag status, whitelisted users, and rollout percentage in that order to determine whether the new feature is enabled.

- If the flag's active status is false, the function returns `false`
- If current user's UUID is in the whitelist of users, the function returns `true`
- If current user's UUID hashes to a value within user rollout percentage, the function returns `true`
- If current user's UUID hashes to a value outside user rollout percentage, the function returns `false`

**Parameters:**

- `nil`

**Return Value**

- `true` or `false` depending on whether the feature flag is active

---

##### `toggler.EmitSuccess()`

Records a successful operation to the Redis Timeseries database, with key `flagId:success` and value of current timestamp

**Parameters:**

- `nil`

**Return Value**

- `nil`

---

##### `toggler.emitFailure()`

Records a failure operation to the Redis Timeseries database, with key `flagId:success` and value of current timestamp

**Parameters:**

- `nil`

**Return Value**

- `nil`
