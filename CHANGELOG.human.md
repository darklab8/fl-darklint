# Changelog

## v0.1.0

- Added parsing and validating feature to darktool on 'darktool validate'
- Parses market_ships.ini, universe.ini and infocards.txt
- Base names from infocards forwards to market_ships.ini as a commented out parameter
- Rewrites market_ships.ini as a completely enforced in style file
- Supports having comments at the file beginning and as commented out parameters on top of ini syntax
- Implements unique INI reader capable to parse Freelancer INI configs and write them back