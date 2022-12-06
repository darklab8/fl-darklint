# Releases

[download here](https://github.com/darklab8/darklab_freelancer_darktool/releases)

currently built for:

- linux amd64
- windows amd64

# Description

- Set of dev tools for Freelancer Discovery community, in order to help them in their effort
- Main goal is to create classic linter for game files, whichi can autofix config files
- Posisbly creating graphical web interface that utilizes ORM for automaticly validated configuration management

# Features:

- Processes
  - `market_commodities.ini`
  - `market_misc.ini`
  - `market_ships.ini`
  - `universe_ini`
  - all system files like `universe/systems/**/br01.ini`
- For processed files brings to lower case allowed set of keys, like `base = GA06_03_base` to `base = ga06_03_base`
- to `market_*.ini` files it adds to bases human readable name extracted from infocard.txt
- to `market_*.ini` reports if base is recycle_candidate, by checking missmatch in its set system and pressence in files + if system is `fp7` or `ga13`, example:
  - `;%is_recycle_candidate = DARK_ERR_0001 base_good.base=ga06_03_base not in universe.ini->Base.system->System.file->systems\ga13\ga13.ini | universe.ini->Base.system=ga13 in [[ga13 fp7]]`
  - see picture example below

![](assets/diff_example.png)

# Dev Requirements

- cobra generator https://github.com/spf13/cobra-cli/blob/main/README.md
- cobra guide https://github.com/spf13/cobra/blob/main/user_guide.md
- godoc
- add binary discovery for cobra-cli, godoc detection
  - export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"
- Git hooks of conventional commits
    - https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13
    - https://www.npmjs.com/package/git-conventional-commits

