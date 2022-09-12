# Seeding data

When developing Fleet, it may be useful to create seed data that includes users and teams.

Check out this Loom demo video that walks through creating teams seed data:
https://www.loom.com/share/1c41a1540e8f41328a7a6cfc56ad0a01

For a text-based walkthrough, check out the following steps:

First, create an `env` file with the following contents:

```
export SERVER_URL=https://localhost:8080 # your Fleet server URL and port
export CURL_FLAGS='-k -s' # set insecure flag
export TOKEN=eyJhbGciOi... # your login token
```

Next, set the `FLEET_ENV_PATH` to point to the `env` file. This will let the scripts in the `fleet/` folder source the env file.

```
export FLEET_ENV_PATH=/Users/victor/fleet_env
```

Finally, run one of the bash scripts located in the [/tools/api](../../tools/api/README.md) directory.

The `fleet/create_free` script will generate an environment to roughly reflect an installation of Fleet Free. The script creates three users with different roles.

```
./tools/api/fleet/teams/create_free
```

The `fleet/create_premium` script will generate an environment to roughly reflect an installation of Fleet Premium. The script will create two teams of four users with different roles.

```
./tools/api/fleet/teams/create_premium
```

The `fleet/create_figma` script will generate an environment to reflect the mockups in the Fleet EE (current) Figma file. The script creates three teams and twelve users with different roles.

```
./tools/api/fleet/teams/create_figma
```

Each user-generated by the script has its password set to `user123#`.

In order to run scripts that make use of premium features, make sure you have started the server with the correct flags as described in [Testing](./Testing-and-local-development.md#license-key).

<meta name="pageOrderInSection" value="600">