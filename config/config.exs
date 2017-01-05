# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.
use Mix.Config

# General application configuration
config :meetings,
  ecto_repos: [Meetings.Repo]

# Configures the endpoint
config :meetings, Meetings.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "HCWK0CvoO5N+pt+rYdGuKJWNIB4yg2aTq4c/DxV7pgCCZ0ME6HqPSVGR5HCJFmpl",
  render_errors: [view: Meetings.ErrorView, accepts: ~w(html json)],
  pubsub: [name: Meetings.PubSub,
           adapter: Phoenix.PubSub.PG2]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env}.exs"
