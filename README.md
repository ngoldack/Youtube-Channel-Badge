# Youtube-Channel-Badge

A shields.io API for your youtube channel to protect your api key.

## Why

You can't publicly access the google/youtube api without an api key. That's why there is no shields.io badge for youtube channels. This repo aims to be the solution to this, because you deploy your own public api which only fetches your channel.

## Features

- Cache time to prevent rate limiting
- Secure api key due env variable

## Examples

### Urls

| Name | Example | Description | URL |
| --- | --- | --- | --- |
| Subscriber count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fsubscriber) | Shows the amount of subscribers | /subscriber |
| View count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fviews) | Shows the amount of youtube views of all videos | /views |
| Video count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fvideos) | Shows the amount of videos on the channel | /videos |
| Comment count | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fcomments) | Shows the amount of youtube comments of the channel | /comments |

### Styles

| Name | Examples |
| --- | --- |
| plastic | ![Plastic badge](https://img.shields.io/endpoint?style=plastic&url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fsubscriber) |
| flat | ![Custom badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fsubscriber) |
| flat-square | ![Custom badge](https://img.shields.io/endpoint?style=flat-square&url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fsubscriber) |
| for-the-badge | ![Custom badge](https://img.shields.io/endpoint?style=for-the-badge&url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fsubscriber) |
| social | ![Custom badge](https://img.shields.io/endpoint?style=social&url=https%3A%2F%2Fyoutube-channel-stats.ntec-io.vercel.app%2Fapi%2Fsubscriber) |

## Configuration

| Variable Name | Description | Required |
| --- | --- | --- |
| API_KEY | Your google api key. See section below on how to get one | **TRUE**
| CHANNEL_ID | The youtube channel id of which the stats should be processed | **TRUE** |
| CACHE_TIME | Time in seconds on how long the last result is cached. Default is 300. | **FALSE** |

<details>
    <summary>Getting a google api key</summary>
    TODO
</details>
<details>
    <summary>Find your youtube channel id</summary>
    TODO
</details>

## Deployment with vercel

It is neccessary to deploy this repo by yourself to get a domain for the desired youtube channel.
An easy and free way to deploy this repo is by using vercel. Just use the button below.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/import/git?s=https%3A%2F%2Fgithub.com%2Fntec-io%2Fyoutube-channel-badge&env=API_KEY,CHANNEL_ID,CACHE_TIME&envDescription=Find%20information%20on%20how%20to%20get%20these%20in%20the%20readme&envLink=https%3A%2F%2Fgithub.com%2Fntec-io%2Fyoutube-channel-badge%23configuration)

<details>
    <summary>Step by step deployment</summary>
    TODO
</details>

## Usage

1. Deploy the repo and get your url from **domains**

    Example in vercel:

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/deployUrl.png?raw=true)

2. Go to shields.io/endpoint and scroll down

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/usage01.png?raw=true)

3. Add your vercel url to the **url** field
4. Add your wanted url suffix:
    - to get the subscriber count: add api/subscriber

    - to get the view count: add api/views

    - to get the comment count: add api/comments

    - to get the video count: add api/videos

    It should look like this:

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/usage02.png?raw=true)

5. Override the labels/colors/logo if you want (optional)
6. Click on the button to copy the badge url

    ![url image](https://github.com/ntec-io/youtube-channel-badge/blob/master/docs/images/usage03.png?raw=true)
