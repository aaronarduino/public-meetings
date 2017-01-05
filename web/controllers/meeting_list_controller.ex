defmodule Meetings.MeetingListController do
  use Meetings.Web, :controller

  @data (
    Application.app_dir(:meetings, "priv/data/Meetings_2017.json")
    |> File.read!
    |> Poison.decode!
  )

  def index(conn, _params) do
    render conn, "index.html", meetings: @data
  end

  def show(conn, %{"id" => id}) do
    meeting = Enum.find(@data, fn(m) -> m["id"] == String.to_integer(id) end)
    # date = DateTime.from_iso8601(@meeting["date"]) only avalaible on v1.40-rc.1
    render conn, meeting: meeting
  end
end