<p align="center">
  <img width="400" src="https://user-images.githubusercontent.com/40566803/114980457-11d44c00-9ebf-11eb-8feb-682587918420.png">
  <h3 align="center">hitokoto-box</h3>
  <p align="center">🌧将一条随机<a href="https://github.com/hitokoto-osc/hitokoto-api">一言</a>显示在 pinned gist</p>
</p>



---

中文 | [English](README-en.md)

> 📌✨ 若想了解更多类似我的 pinned gist 项目，请查看：https://github.com/matchai/awesome-pinned-gists

## 使用

### 准备工作

1. 创建一个公开的 GitHub Gist，文件名为`🌧Hitokoto` (https://gist.github.com/)
1. 创建一个GitHub token，需要勾选 gist 权限，复制刚才创建的 token (https://github.com/settings/tokens/new)

### 安装

1. Fork 我
2. 进入你 fork 的仓库的 **Actions** 标签页，启用 **Update gist with new hitokoto** workflow
3. 在  `.github/workflows/update.yml` 中编辑[环境变量](https://github.com/greenhandatsjtu/hitokoto-box/blob/main/.github/workflows/update.yml#L13-L15)：

   - **GIST_ID:** 你的 gist url 的ID部分（后缀） `https://gist.github.com/greenhandatsjtu/`**`8c2cd8ea266a45c1aa9cb5f78d066b60`**.
   - **CATEGORY**: 你希望`hitokoto-box`获取的一言的类别，所有类别的编码在[这里](https://developer.hitokoto.cn/sentence/#%E5%8F%A5%E5%AD%90%E7%B1%BB%E5%9E%8B-%E5%8F%82%E6%95%B0)。默认值为 `abh`，代表动画、漫画、影视。若希望获取所有类别的一言，将本环境变量置空即可。
4.  `hitokoto-box`默认**每小时**更新一次一言，如果你想更改更新频率，需要在 `.github/workflows/update.yml` 中编辑 [cron](https://github.com/greenhandatsjtu/hitokoto-box/blob/8e94d65a0193555978a20229f80a72155c4410d9/.github/workflows/update.yml#L7)
5. 前往仓库的 **Settings > Secrets**
6. 点击 **New repository secret** 并添加如下仓库秘密 (repository secrets) ：
   - **GH_TOKEN:** 刚才复制的 GitHub token