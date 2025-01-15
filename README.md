# AniLody

[**AniLody**](https://github.com/PyroWilDx/AniLody/) is a tool for downloading opening and ending songs based on your [AniList](https://anilist.co/) anime list.

AniLody utilizes [AniList API](https://docs.anilist.co/) and [AnimeThemes API](https://api-docs.animethemes.moe/) to retrieve songs.

## App Set-Up

The configuration file is located at `config/Config.txt`. Below is a list of available configurations and their usage.

- `userSite` &ndash; Website from which your anime list will be fetched.

> [!NOTE]
> Currently, only [AniList](https://anilist.co/) is supported.

- `userName` &ndash; Your [AniList](https://anilist.co/) user-name.

- `outPath` &ndash; Directory where the downloaded songs will be saved.

- `musicNameFormat` &ndash; Format for the downloaded song file names. Below are the available variables.
  - `#AnimeTitle` &ndash; Title of the anime.
  - `#Slug` &ndash; Song type and number (e.g. Op1 for the first opening, Ed1 for the first ending).
  - `#SongTitle` &ndash; Title of the song.

> [!NOTE]
> **Example**
>
> If you download **Unravel** from **Tokyo Ghoul**, and use the format `#AnimeTitle - #Slug (#SongTitle)`, the resulting file name will be `Tokyo Ghoul - Op1 (Unravel)`.

- `capWords` &ndash; Capitalize the first letter of each word in song file names.

- `lowWords` &ndash; Lowercase all letters (except the first one) of each word in song file names.

- `addImage` &ndash; Attach key visual image to the song files.

- `incOp` &ndash; Enable downloading of opening theme songs.

- `incEd` &ndash; Enable downloading of ending theme songs.

- `minScore` &ndash; Minimum score an anime must have on your list to be included.

- `maxScore` &ndash; Maximum score an anime can have on your list to be included.

- `statusList` &ndash; Filter anime by their status in your anime list. Below are the possible status options (use `|` to separate multiple options).
    - `Completed` - Anime you have finished.
    - `Current` - Anime you are currently watching.
    - `Repeating` - Anime you are re-watching.
    - `Paused` - Anime you have paused.
    - `Dropped` - Anime you have dropped.
    - `Planning` - Anime you plan to watch.

## Download

<div align="center">

| [<img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/windows8/windows8-original.svg" width="60"/>](https://www.mediafire.com/file/rsr6mdbm6wlxasm/AniLody.zip/) |
|---|

</div>

## Development Set-Up

<div align="center">

| [<img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" width="60"/>](https://go.dev/) | [<img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/goland/goland-original.svg" width="60"/>](https://www.jetbrains.com/go/) | [<img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/windows8/windows8-original.svg" width="60"/>](https://www.microsoft.com/windows/) |
|---|---|---|

</div>

### How To Use

- Run w/ Go.

---

<div align="center">
  Copyright &#169; 2024 PyroWilDx. All Rights Reserved.
</div>
