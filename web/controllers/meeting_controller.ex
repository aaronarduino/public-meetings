defmodule Meetings.MeetingController do
  use Meetings.Web, :controller

  alias Meetings.Meeting

  def index(conn, _params) do
    meetings = Repo.all(Meeting)
    render(conn, "index.html", meetings: meetings)
  end

  def new(conn, _params) do
    changeset = Meeting.changeset(%Meeting{})
    render(conn, "new.html", changeset: changeset)
  end

  def create(conn, %{"meeting" => meeting_params}) do
    changeset = Meeting.changeset(%Meeting{}, meeting_params)

    case Repo.insert(changeset) do
      {:ok, _meeting} ->
        conn
        |> put_flash(:info, "Meeting created successfully.")
        |> redirect(to: meeting_path(conn, :index))
      {:error, changeset} ->
        render(conn, "new.html", changeset: changeset)
    end
  end

  def show(conn, %{"id" => id}) do
    meeting = Repo.get!(Meeting, id)
    render(conn, "show.html", meeting: meeting)
  end

  def edit(conn, %{"id" => id}) do
    meeting = Repo.get!(Meeting, id)
    changeset = Meeting.changeset(meeting)
    render(conn, "edit.html", meeting: meeting, changeset: changeset)
  end

  def update(conn, %{"id" => id, "meeting" => meeting_params}) do
    meeting = Repo.get!(Meeting, id)
    changeset = Meeting.changeset(meeting, meeting_params)

    case Repo.update(changeset) do
      {:ok, meeting} ->
        conn
        |> put_flash(:info, "Meeting updated successfully.")
        |> redirect(to: meeting_path(conn, :show, meeting))
      {:error, changeset} ->
        render(conn, "edit.html", meeting: meeting, changeset: changeset)
    end
  end

  def delete(conn, %{"id" => id}) do
    meeting = Repo.get!(Meeting, id)

    # Here we use delete! (with a bang) because we expect
    # it to always work (and if it does not, it will raise).
    Repo.delete!(meeting)

    conn
    |> put_flash(:info, "Meeting deleted successfully.")
    |> redirect(to: meeting_path(conn, :index))
  end
end
