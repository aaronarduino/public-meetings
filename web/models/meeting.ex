defmodule Meetings.Meeting do
  use Meetings.Web, :model

  schema "meetings" do
    field :type, :string
    field :subtype, :string
    field :title, :string
    field :description, :string
    field :location, :string
    field :hour, :integer
    field :minute, :integer
    field :duration, :integer
    field :email, :string
    field :agenda, :string

    timestamps()
  end

  @doc """
  Builds a changeset based on the `struct` and `params`.
  """
  def changeset(struct, params \\ %{}) do
    struct
    |> cast(params, [:type, :subtype, :title, :description, :location, :hour, :minute, :duration, :email, :agenda])
    |> validate_required([:type, :subtype, :title, :description, :location, :hour, :minute, :duration, :email, :agenda])
  end
end
