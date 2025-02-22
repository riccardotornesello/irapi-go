# Add here some examples of the parameters that are used in the API calls
# This is used to generate the schemas for the API calls
PARAMETERS = {
    "league": {
        "get": {"league_id": 4403},
        "get_points_systems": {"league_id": 4403},
        "seasons": {"league_id": 4403},
        "roster": {
            "league_id": 4403,
            "include_licenses": False,  # TODO: create the schema in case the parameter is True
        },
        "season_standings": {"league_id": 4403, "season_id": 111405},
        "season_sessions": {"league_id": 4403, "season_id": 111405},
    },
    "lookup": {
        "club_history": {
            "season_year": 2024,
            "season_quarter": 1,
        },
    },
    "member": {
        "award_instances": {
            "cust_id": 911231,
            "award_id": 47780700,
        },
        "chart_data": {
            "cust_id": 911231,
            "category_id": 5,
            "chart_type": 1,
        },
        "get": {
            "cust_ids": [911231, 408068],
            "include_licenses": False,  # TODO: create the schema in case the parameter is True
        },
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
    "stats": {
        "member_division": {
            "season_id": 5244,
            "event_type": 5,
        },
    },
    "team": {
        "get": {"team_id": 381567},
    },
}

# Add here the schema adjustments that are not automatically generated, to convert objects to maps
SCHEMA_ADJUSTMENTS = {
    "car": {
        "assets": {"type": "object", "patternProperties": {r"^\d+$": None}},
    },
    "series": {
        "seasons": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "allowed_season_members": {
                        "type": "object",
                        "patternProperties": {r"^\d+$": None},
                    }
                },
            },
        }
    },
}

REQUEST_OVERRIDES = {
    "constants": {
        "categories": {
            "s3_cache": False,
        },
        "divisions": {
            "s3_cache": False,
        },
        "event_types": {
            "s3_cache": False,
        },
    },
    "driver_stats_by_category": {
        "dirt_oval": {"format": "csv"},
        "dirt_road": {"format": "csv"},
        "formula_car": {"format": "csv"},
        "oval": {"format": "csv"},
        "road": {"format": "csv"},
        "sports_car": {"format": "csv"},
    },
    "league": {
        "roster": {
            "s3_cache": False,
        },
    },
    "member": {
        "awards": {
            "s3_cache": False,
        },
        "award_instances": {
            "s3_cache": False,
        },
    },
}
