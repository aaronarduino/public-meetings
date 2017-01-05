defmodule Meetings.MeetingControllerTest do
  use Meetings.ConnCase

  alias Meetings.Meeting
  @valid_attrs %{agenda: "some content", description: "some content", duration: 42, email: "some content", hour: 42, location: "some content", minute: 42, subtype: "some content", title: "some content", type: "some content"}
  @invalid_attrs %{}

  test "lists all entries on index", %{conn: conn} do
    conn = get conn, meeting_path(conn, :index)
    assert html_response(conn, 200) =~ "Listing meetings"
  end

  test "renders form for new resources", %{conn: conn} do
    conn = get conn, meeting_path(conn, :new)
    assert html_response(conn, 200) =~ "New meeting"
  end

  test "creates resource and redirects when data is valid", %{conn: conn} do
    conn = post conn, meeting_path(conn, :create), meeting: @valid_attrs
    assert redirected_to(conn) == meeting_path(conn, :index)
    assert Repo.get_by(Meeting, @valid_attrs)
  end

  test "does not create resource and renders errors when data is invalid", %{conn: conn} do
    conn = post conn, meeting_path(conn, :create), meeting: @invalid_attrs
    assert html_response(conn, 200) =~ "New meeting"
  end

  test "shows chosen resource", %{conn: conn} do
    meeting = Repo.insert! %Meeting{}
    conn = get conn, meeting_path(conn, :show, meeting)
    assert html_response(conn, 200) =~ "Show meeting"
  end

  test "renders page not found when id is nonexistent", %{conn: conn} do
    assert_error_sent 404, fn ->
      get conn, meeting_path(conn, :show, -1)
    end
  end

  test "renders form for editing chosen resource", %{conn: conn} do
    meeting = Repo.insert! %Meeting{}
    conn = get conn, meeting_path(conn, :edit, meeting)
    assert html_response(conn, 200) =~ "Edit meeting"
  end

  test "updates chosen resource and redirects when data is valid", %{conn: conn} do
    meeting = Repo.insert! %Meeting{}
    conn = put conn, meeting_path(conn, :update, meeting), meeting: @valid_attrs
    assert redirected_to(conn) == meeting_path(conn, :show, meeting)
    assert Repo.get_by(Meeting, @valid_attrs)
  end

  test "does not update chosen resource and renders errors when data is invalid", %{conn: conn} do
    meeting = Repo.insert! %Meeting{}
    conn = put conn, meeting_path(conn, :update, meeting), meeting: @invalid_attrs
    assert html_response(conn, 200) =~ "Edit meeting"
  end

  test "deletes chosen resource", %{conn: conn} do
    meeting = Repo.insert! %Meeting{}
    conn = delete conn, meeting_path(conn, :delete, meeting)
    assert redirected_to(conn) == meeting_path(conn, :index)
    refute Repo.get(Meeting, meeting.id)
  end
end
