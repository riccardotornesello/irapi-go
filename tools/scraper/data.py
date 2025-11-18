"""
Configuration overrides for API endpoint behavior.

This module contains the OVERRIDES dictionary which specifies custom behavior
for specific API endpoints, including:
- Response format (json/csv)
- S3 caching behavior
- Required parameters and sample values for testing

The overrides are organized hierarchically by category and endpoint name.
"""

OVERRIDES = {
    # Configuration overrides by category and endpoint
    "constants": {
        "categories": {
            "s3_cache": False,  # Don't follow S3 links for this endpoint
        },
        "divisions": {
            "s3_cache": False,
        },
        "event_types": {
            "s3_cache": False,
        },
    },
    "driver_stats_by_category": {
        # CSV format endpoints
        "dirt_oval": {"format": "csv"},
        "dirt_road": {"format": "csv"},
        "formula_car": {"format": "csv"},
        "oval": {"format": "csv"},
        "road": {"format": "csv"},
        "sports_car": {"format": "csv"},
    },
    "league": {
        "get": {
            # Sample parameters for endpoints that require them
            "params": [
                {
                    "league_id": 4403,
                }
            ],
        },
        "get_points_systems": {
            "params": [
                {
                    "league_id": 4403,
                }
            ],
        },
        "roster": {
            # Multiple parameter sets to test different response types
            "params": [
                {
                    "league_id": 4403,
                    "include_licenses": False,
                },
                {
                    "league_id": 4403,
                    "include_licenses": True,
                },
            ],
            "s3_cache": False,
        },
        "seasons": {
            "params": [
                {
                    "league_id": 4403,
                }
            ],
        },
        "season_standings": {
            "params": [
                {
                    "league_id": 4403,
                    "season_id": 111405,
                }
            ],
        },
        "season_sessions": {
            "params": [
                {
                    "league_id": 4403,
                    "season_id": 111405,
                }
            ],
        },
    },
    "lookup": {
        "drivers": {
            "params": [
                {
                    "search_term": "911231",
                }
            ],
        },
        "club_history": {
            "params": [
                {
                    "season_year": 2024,
                    "season_quarter": 1,
                }
            ],
        },
    },
    "member": {
        "awards": {
            "s3_cache": False,
            "params": [
                {
                    "cust_id": 911231,
                }
            ],
        },
        "award_instances": {
            "s3_cache": False,
            "params": [
                {
                    "cust_id": 911231,
                    "award_id": 47780700,
                }
            ],
        },
        "chart_data": {
            "params": [
                {
                    "cust_id": 911231,
                    "category_id": 5,
                    "chart_type": 1,
                }
            ],
        },
        "get": {
            "params": [
                {
                    "cust_ids": [911231, 408068],
                    "include_licenses": False,
                },
                {
                    "cust_ids": [911231, 408068],
                    "include_licenses": True,
                },
            ],
        },
    },
    "results": {
        "get": {
            "params": [
                {
                    "subsession_id": 67468746,
                    "include_licenses": False,
                },
                {
                    "subsession_id": 67468746,
                    "include_licenses": True,
                },
            ],
        },
        "lap_data": {
            "params": [
                {
                    "subsession_id": 72820991,
                    "simsession_number": 0,
                    "cust_id": 911231,
                    # "team_id": 0, # TODO
                }
            ],
        },
    },
    "stats": {
        "params": [
            {
                "member_division": {
                    "season_id": 5244,
                    "event_type": 5,
                },
            }
        ],
    },
    "team": {
        "get": {
            "params": [
                {
                    "team_id": 381567,
                }
            ],
        },
    },
}
