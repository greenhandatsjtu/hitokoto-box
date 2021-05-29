<p align="center">
  <img width="400" src="https://user-images.githubusercontent.com/40566803/114980457-11d44c00-9ebf-11eb-8feb-682587918420.png">
  <h3 align="center">hitokoto-box</h3>
  <p align="center">🌧Update a pinned gist to contain a random <a href="https://github.com/hitokoto-osc/hitokoto-api">hitokoto</a></p>
</p>

---

[中文](README.md) | English

> 📌✨ For more pinned-gist projects like this one, check out: https://github.com/matchai/awesome-pinned-gists

## Setup

### Prep work

1. Create a new public GitHub Gist with `🌧Hitokoto` filename (https://gist.github.com/)
1. Create a token with the `gist` scope and copy it. (https://github.com/settings/tokens/new)

### Project setup

1. Fork this repo
2. Open the **Actions** tab of your fork and enable the **Update gist with new hitokoto** workflow
3. Edit the [environment variable](https://github.com/greenhandatsjtu/hitokoto-box/blob/main/.github/workflows/update.yml#L13-L15) in `.github/workflows/update.yml`:

   - **GIST_ID:** The ID portion from your gist url: `https://gist.github.com/greenhandatsjtu/`**`8c2cd8ea266a45c1aa9cb5f78d066b60`**.
   - **CATEGORY**: Preferred categories of hitokoto to be fetched. You can find code of all available categories [here](https://developer.hitokoto.cn/sentence/#%E5%8F%A5%E5%AD%90%E7%B1%BB%E5%9E%8B-%E5%8F%82%E6%95%B0). Default value is `abh`, which stands for anime, manga and film. Leave it blank if you want to get hitokoto of all categories.
4. By default, hitokoto-box will fetch a new hitokoto **every hour**, you can change the frequency by setting [cron](https://github.com/greenhandatsjtu/hitokoto-box/blob/8e94d65a0193555978a20229f80a72155c4410d9/.github/workflows/update.yml#L7) in `.github/workflows/update.yml`
5. Go to the repo **Settings > Secrets**
6. Click **New repository secret** and add the following repository secrets:
   - **GH_TOKEN:** The GitHub token generated above.
