defmodule Meetings.Repo.Migrations.CreateMeeting do
  use Ecto.Migration

  def change do
    create table(:meetings) do
      add :type, :string
      add :subtype, :string
      add :title, :string
      add :description, :string
      add :location, :string
      add :hour, :integer
      add :minute, :integer
      add :duration, :integer
      add :email, :string
      add :agenda, :string

      timestamps()
    end

  end
end
