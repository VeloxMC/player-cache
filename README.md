# player-cache
A REST API that provides Minecraft player identification endpoints without a rate limit by caching the Mojang API

## Deployment
The auto-built docker image is named `ksebrt/velox-player-cache`. You can look at the `docker-compose.yml` for a possible
implementation.

## API keys
You need a bound `keys.json` which holds an array with valid API keys. We plan to switch to a MongoDB
integration which helps to create keys automatically, but this is our current temporary solution.

## Contribution/Help
If you found a bug, have any suggestions or want to improve some code, feel free to create an issue or pull request!

Additionally you can join our [Discord guild](https://discord.gg/Q6ZXBvU).
