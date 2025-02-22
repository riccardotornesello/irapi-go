# Add here some examples of the parameters that are used in the API calls
# This is used to generate the schemas for the API calls
# It is also possible to add some schema adjustments here, to convert objects to maps, and to add some request overrides
OVERRIDES = {
    "car": {
        "assets": {
            "schema": {"type": "object", "patternProperties": {r"^\d+$": None}},
        },
    },
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
    "hosted": {
        "combined_sessions": {
            "schema": {
                "type": "object",
                "properties": {
                    "sessions": {
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "count_by_car_id": {
                                    "type": "object",
                                    "patternProperties": {r"^\d+$": None},
                                },
                                "count_by_car_class_id": {
                                    "type": "object",
                                    "patternProperties": {r"^\d+$": None},
                                },
                            },
                        },
                    }
                },
            },
        },
    },
    "league": {
        "cust_league_sessions": {
            "schema": {
                "type": "object",
                "properties": {
                    "sessions": {
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "count_by_car_id": {
                                    "type": "object",
                                    "patternProperties": {r"^\d+$": None},
                                },
                                "count_by_car_class_id": {
                                    "type": "object",
                                    "patternProperties": {r"^\d+$": None},
                                },
                            },
                        },
                    },
                },
            },
        },
        "get": {
            "params": {
                "league_id": 4403,
            },
        },
        "get_points_systems": {
            "params": {
                "league_id": 4403,
            },
        },
        "seasons": {
            "params": {
                "league_id": 4403,
            },
        },
        "roster": {
            "params": {
                "league_id": 4403,
                "include_licenses": False,  # TODO: create the schema in case the parameter is True
            },
            "s3_cache": False,
        },
        "season_standings": {
            "params": {
                "league_id": 4403,
                "season_id": 111405,
            },
        },
        "season_sessions": {
            "params": {
                "league_id": 4403,
                "season_id": 111405,
            },
        },
    },
    "lookup": {
        "club_history": {
            "params": {
                "season_year": 2024,
                "season_quarter": 1,
            },
        },
    },
    "member": {
        "awards": {
            "s3_cache": False,
        },
        "award_instances": {
            "s3_cache": False,
            "params": {
                "cust_id": 911231,
                "award_id": 47780700,
            },
        },
        "chart_data": {
            "params": {
                "cust_id": 911231,
                "category_id": 5,
                "chart_type": 1,
            },
        },
        "get": {
            "params": {
                "cust_ids": [911231, 408068],
                "include_licenses": False,  # TODO: create the schema in case the parameter is True
            },
        },
    },
    "results": {
        "get": {
            "params": {
                "subsession_id": 67468746,
                "include_licenses": False,  # TODO: create the schema in case the parameter is True
            },
        },
        "lap_data": {
            "params": {
                "subsession_id": 72820991,
                "simsession_number": 0,
                "cust_id": 911231,
                # "team_id": 0, # TODO
            },
        },
    },
    "series": {
        "assets": {
            "schema": {"type": "object", "patternProperties": {r"^\d+$": None}},
        },
        "seasons": {
            "schema": {
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
            },
        },
    },
    "stats": {
        "params": {
            "member_division": {
                "season_id": 5244,
                "event_type": 5,
            },
        },
    },
    "team": {
        "get": {
            "params": {
                "team_id": 381567,
            },
        },
    },
    "track": {
        "assets": {
            "schema": {"type": "object", "patternProperties": {r"^\d+$": None}},
        },
    },
}
