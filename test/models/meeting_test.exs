defmodule Meetings.MeetingTest do
  use Meetings.ModelCase

  alias Meetings.Meeting

  @valid_attrs %{agenda: "some content", description: "some content", duration: 42, email: "some content", hour: 42, location: "some content", minute: 42, subtype: "some content", title: "some content", type: "some content"}
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    changeset = Meeting.changeset(%Meeting{}, @valid_attrs)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Meeting.changeset(%Meeting{}, @invalid_attrs)
    refute changeset.valid?
  end
end
