"""
Applet: Fortnite Stats
Summary: View your Fortnite stats
Description: Displays your Fortnite stats.
Author: naominori
"""

load("render.star", "render")
load("schema.star", "schema")

def main(config):
    return render.Root(
        child = render.Box(height = 1, width = 1),
    )

def get_schema():
    return schema.Schema(
        version = "1",
        fields = [],
    )