# TODO: add parameters for all skipped endpoints

PARAMETERS = {
    "league": {
        "get": {"league_id": 4403},
        "get_points_systems": {"league_id": 4403},
        "seasons": {"league_id": 4403},
        "season_standings": {"league_id": 4403, "season_id": 111405},
        "season_sessions": {"league_id": 4403, "season_id": 111405},
    },
    "results": {
        "get": {
            "subsession_id": 67468746,
            "include_licenses": False,  # TODO: create the schema in case the parameter is True
        },
        "lap_data": {
            "subsession_id": 72820991,
            "simsession_number": 0,
            "cust_id": 911231,
            # "team_id": 0, # TODO
        },
    },
    "stats": {},
    "team": {
        "get": {"team_id": 381567},
    },
}
